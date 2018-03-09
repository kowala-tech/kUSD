package validator

import (
	"errors"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/kowala-tech/kUSD/accounts"
	"github.com/kowala-tech/kUSD/accounts/abi/bind"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/consensus"
	"github.com/kowala-tech/kUSD/contracts/network"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/state"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/core/vm"
	"github.com/kowala-tech/kUSD/event"
	"github.com/kowala-tech/kUSD/kusddb"
	"github.com/kowala-tech/kUSD/log"
	"github.com/kowala-tech/kUSD/params"
)

var (
	ErrCantStopNonStartedValidator       = errors.New("can't stop validator, not started")
	ErrCantVoteNotValidating             = errors.New("can't vote, not validating")
	ErrCantSetCoinbaseOnStartedValidator = errors.New("can't set coinbase, already started validating")
	ErrCantAddProposalNotValidating      = errors.New("can't add proposal, not validating")
	ErrCantAddBlockFragmentNotValidating = errors.New("can't add block fragment, not validating")
)

// Backend wraps all methods required for mining.
type Backend interface {
	AccountManager() *accounts.Manager
	BlockChain() *core.BlockChain
	TxPool() *core.TxPool
	ChainDb() kusddb.Database
}

type Validator interface {
	Start(coinbase common.Address, deposit uint64)
	Stop() error
	SetExtra(extra []byte) error
	Validating() bool
	SetCoinbase(addr common.Address) error
	SetDeposit(deposit uint64)
	Pending() (*types.Block, *state.StateDB)
	PendingBlock() *types.Block
	AddProposal(proposal *types.Proposal) error
	AddVote(vote *types.Vote) error
	AddBlockFragment(blockNumber *big.Int, round uint64, fragment *types.BlockFragment) error
}

// validator represents a consensus validator
type validator struct {
	Election           // consensus state
	maxTransitions int // max number of state transitions (tests) 0 - unlimited

	running    int32
	validating int32
	deposit    uint64

	signer types.Signer

	// blockchain
	backend  Backend
	chain    *core.BlockChain
	config   *params.ChainConfig
	engine   consensus.Engine
	vmConfig vm.Config

	network *network.NetworkContract // validators contract

	walletAccount accounts.WalletAccount

	// sync
	canStart    int32 // can start indicates whether we can start the validation operation
	shouldStart int32 // should start indicates whether we should start after sync

	// events
	eventMux *event.TypeMux

	wg sync.WaitGroup
}

// New returns a new consensus validator
func New(walletAccount accounts.WalletAccount, backend Backend, contract *network.NetworkContract, config *params.ChainConfig, eventMux *event.TypeMux, engine consensus.Engine, vmConfig vm.Config) *validator {
	validator := &validator{
		config:        config,
		backend:       backend,
		chain:         backend.BlockChain(),
		engine:        engine,
		network:       contract,
		eventMux:      eventMux,
		signer:        types.NewAndromedaSigner(config.ChainID),
		vmConfig:      vmConfig,
		canStart:      0,
		walletAccount: walletAccount,
	}

	go validator.sync()

	return validator
}

func (val *validator) sync() {
	if err := SyncWaiter(val.eventMux); err != nil {
		log.Warn("Failed to sync with network", "err", err)
	} else {
		val.finishedSync()
	}
}

func (val *validator) finishedSync() {
	start := atomic.LoadInt32(&val.shouldStart) == 1
	atomic.StoreInt32(&val.canStart, 1)
	atomic.StoreInt32(&val.shouldStart, 0)
	if start {
		val.Start(val.walletAccount.Account().Address, val.deposit)
	}
}

func (val *validator) Start(coinbase common.Address, deposit uint64) {
	if val.Validating() {
		log.Warn("Failed to start the validator - the state machine is already running")
		return
	}

	atomic.StoreInt32(&val.shouldStart, 1)

	newWalletAccount, _ := accounts.NewWalletAccount(val.walletAccount, accounts.Account{Address: coinbase})
	val.walletAccount = newWalletAccount
	val.deposit = deposit

	if atomic.LoadInt32(&val.canStart) == 0 {
		log.Info("Network syncing, will start validator afterwards")
		return
	}

	go val.run()
}

func (val *validator) run() {
	log.Info("Starting validation operation")
	val.wg.Add(1)
	atomic.StoreInt32(&val.running, 1)

	defer func() {
		val.wg.Done()
		atomic.StoreInt32(&val.running, 0)
	}()

	log.Info("Starting the consensus state machine")
	for state, numTransitions := val.notLoggedInState, 0; state != nil; numTransitions++ {
		state = state()
		if val.maxTransitions > 0 && numTransitions == val.maxTransitions {
			break
		}
	}
}

func (val *validator) Stop() error {
	if !val.Validating() {
		return ErrCantStopNonStartedValidator
	}
	log.Info("Stopping consensus validator")

	val.leave()
	val.wg.Wait() // waits until the validator is no longer registered as a voter.

	atomic.StoreInt32(&val.shouldStart, 0)
	log.Info("Consensus validator stopped")
	return nil
}

func (val *validator) SetExtra(extra []byte) error { return nil }

func (val *validator) Validating() bool {
	return atomic.LoadInt32(&val.validating) > 0
}

func (val *validator) SetCoinbase(address common.Address) error {
	if val.Validating() {
		return ErrCantSetCoinbaseOnStartedValidator
	}
	newWalletAccount, err := accounts.NewWalletAccount(val.walletAccount, accounts.Account{Address: address})
	if err != nil {
		return err
	}
	val.walletAccount = newWalletAccount
	return nil
}

func (val *validator) SetDeposit(deposit uint64) {
	val.deposit = deposit
}

// Pending returns the currently pending block and associated state.
func (val *validator) Pending() (*types.Block, *state.StateDB) {
	// @TODO (rgeraldes) - review
	// val.currentMu.Lock()
	// defer val.currentMu.Unlock()

	state, err := val.chain.State()
	if err != nil {
		log.Crit("Failed to fetch the latest state", "err", err)
	}

	return val.chain.CurrentBlock(), state
}

func (val *validator) PendingBlock() *types.Block {
	// @TODO (rgeraldes) - review
	// val.currentMu.Lock()
	// defer val.currentMu.Unlock()

	return val.chain.CurrentBlock()
}

func (val *validator) restoreLastCommit() {
	checksum, err := val.network.ValidatorsChecksum(&bind.CallOpts{})
	if err != nil {
		log.Crit("Failed to access the validators checksum", "err", err)
	}

	if err := val.updateValidators(checksum, true); err != nil {
		log.Crit("Failed to update the validator set", "err", err)
	}

	// @TODO (rgeraldes) - we need to request the validator vote weights

	currentBlock := val.chain.CurrentBlock()
	if currentBlock.Number().Cmp(big.NewInt(0)) == 0 {
		return
	}

	/*
		lastCommit := currentBlock.LastCommit()
		lastPreCommits := core.NewVotingTable(val.eventMux, val.signer, currentBlock.Number(), lastCommit.Round(), types.PreCommit, lastValidators)
		for _, preCommit := range lastCommit.Commits() {
			if preCommit == nil {
				continue
			}
			added, err := lastPreCommits.Add(preCommit, false)
			if !added || err != nil {
				// @TODO (rgeraldes) - this should not happen > complete
				log.Error("Failed to restore the latest commit")
			}
		}

		val.lastCommit = lastPreCommits
	*/
}

func (val *validator) init() error {
	parent := val.chain.CurrentBlock()

	checksum, err := val.network.ValidatorsChecksum(&bind.CallOpts{})
	if err != nil {
		log.Crit("Failed to access the validators checksum", "err", err)
	}

	if val.validatorsChecksum != checksum {
		if err := val.updateValidators(checksum, true); err != nil {
			log.Crit("Failed to update the validator set", "err", err)
		}
	}

	// @NOTE (rgeraldes) - start is not relevant for the first block as the first election will
	// wait until we have transactions
	start := time.Unix(parent.Time().Int64(), 0)
	val.start = start.Add(time.Duration(params.BlockTime) * time.Millisecond)
	val.blockNumber = parent.Number().Add(parent.Number(), big.NewInt(1))
	val.round = 0

	val.proposal = nil
	val.block = nil
	val.blockFragments = nil

	val.lockedRound = 0
	val.lockedBlock = nil
	val.commitRound = -1

	// voting system
	val.votingSystem = NewVotingSystem(val.eventMux, val.signer, val.blockNumber, val.validators)

	// @TODO (rgeraldes) - last validators
	// val.lastValidators

	// events
	val.blockCh = make(chan *types.Block)
	val.majority = val.eventMux.Subscribe(core.NewMajorityEvent{})

	// @TODO (rgeraldes) - review vs go-eth
	if err = val.makeCurrent(parent); err != nil {
		log.Error("Failed to create mining context", "err", err)
		return nil
	}

	return nil
}

func (val *validator) isProposer() bool {
	return val.validators.Proposer() == val.walletAccount.Account().Address
}

func (val *validator) AddProposal(proposal *types.Proposal) error {
	if !val.Validating() {
		return ErrCantAddProposalNotValidating
	}

	log.Info("Received Proposal")
	// @TODO (rgeraldes) - Add proposal validation

	/*
		// not relevant
		if proposal.BlockNumber != val.blockNumber && proposal.Round != val.Round {
			return
		}

		// proposer sent two proposals
		if val.proposal != nil {
			// @TODO (rgeraldes) - punish the proposer ?
			return
		}

		// if the proposal is already known, discard it
		hash := proposal.Hash()
		if val.all[hash] != nil {
			log.Trace("Discarding already known proposal", "hash", hash)
			return false, fmt.Errorf("known transaction: %x", hash)
		}

		// if the proposal fails validation, discard it
		if err := val.validateProposal(proposal); err != nil {
			log.Trace("Discarding invalid proposal", "hash", hash, "err", err)
		}


		return
	*/

	val.proposal = proposal
	val.blockFragments = types.NewDataSetFromMeta(proposal.BlockMetadata())

	return nil
}

func (val *validator) AddVote(vote *types.Vote) error {
	if !val.Validating() {
		return ErrCantVoteNotValidating
	}

	if err := val.addVote(vote); err != nil {
		switch err {
		}
	}
	return nil
}

func (val *validator) addVote(vote *types.Vote) error {
	// @NOTE (rgeraldes) - for now just pre-vote/pre-commit for the current block number
	added, err := val.votingSystem.Add(vote, false)
	if err != nil {
		// @TODO (rgeraldes)
	}

	if added {
		switch vote.Type {
		//case PreVote:
		//case PreCommit:
		}
	}

	return nil
}

func (val *validator) commitTransactions(mux *event.TypeMux, txs *types.TransactionsByPriceAndNonce, bc *core.BlockChain, coinbase common.Address) {
	gp := new(core.GasPool).AddGas(val.header.GasLimit)

	var coalescedLogs []*types.Log

	for {
		// Retrieve the next transaction and abort if all done
		tx := txs.Peek()
		if tx == nil {
			break
		}
		// Error may be ignored here. The error has already been checked
		// during transaction acceptance is the transaction pool.
		//
		// We use the eip155 signer regardless of the current hf.
		from, _ := types.TxSender(val.signer, tx)

		// Start executing the transaction
		val.state.Prepare(tx.Hash(), common.Hash{}, val.tcount)

		err, logs := val.commitTransaction(tx, bc, coinbase, gp)
		switch err {
		case core.ErrGasLimitReached:
			// Pop the current out-of-gas transaction without shifting in the next from the account
			log.Trace("Gas limit exceeded for current block", "sender", from)
			txs.Pop()

		case core.ErrNonceTooLow:
			// New head notification data race between the transaction pool and miner, shift
			log.Trace("Skipping transaction with low nonce", "sender", from, "nonce", tx.Nonce())
			txs.Shift()

		case core.ErrNonceTooHigh:
			// Reorg notification data race between the transaction pool and miner, skip account =
			log.Trace("Skipping account with hight nonce", "sender", from, "nonce", tx.Nonce())
			txs.Pop()

		case nil:
			// Everything ok, collect the logs and shift in the next transaction from the same account
			coalescedLogs = append(coalescedLogs, logs...)
			val.tcount++
			txs.Shift()

		default:
			// Strange error, discard the transaction and get the next in line (note, the
			// nonce-too-high clause will prevent us from executing in vain).
			log.Debug("Transaction failed, account skipped", "hash", tx.Hash(), "err", err)
			txs.Shift()
		}
	}

	if len(coalescedLogs) > 0 || val.tcount > 0 {
		// make a copy, the state caches the logs and these logs get "upgraded" from pending to mined
		// logs by filling in the block hash when the block was mined by the local miner. This can
		// cause a race condition if a log was "upgraded" before the PendingLogsEvent is processed.
		cpy := make([]*types.Log, len(coalescedLogs))
		for i, l := range coalescedLogs {
			cpy[i] = new(types.Log)
			*cpy[i] = *l
		}
		go func(logs []*types.Log, tcount int) {
			if len(logs) > 0 {
				mux.Post(core.PendingLogsEvent{Logs: logs})
			}
			if tcount > 0 {
				mux.Post(core.PendingStateEvent{})
			}
		}(cpy, val.tcount)
	}
}

func (val *validator) commitTransaction(tx *types.Transaction, bc *core.BlockChain, coinbase common.Address, gp *core.GasPool) (error, []*types.Log) {
	snap := val.state.Snapshot()

	receipt, _, err := core.ApplyTransaction(val.config, bc, &coinbase, gp, val.state, val.header, tx, val.header.GasUsed, vm.Config{})
	if err != nil {
		val.state.RevertToSnapshot(snap)
		return err, nil
	}
	val.txs = append(val.txs, tx)
	val.receipts = append(val.receipts, receipt)

	return nil, receipt.Logs
}

func (val *validator) makeDeposit() error {
	min, err := val.network.GetMinimumDeposit(&bind.CallOpts{})
	if err != nil {
		return err
	}

	var deposit big.Int
	if min.Cmp(deposit.SetUint64(val.deposit)) > 0 {
		log.Warn("Current deposit is not enough", "deposit", val.deposit, "minimum required", min)
		// @TODO (rgeraldes) - error handling?
		return fmt.Errorf("Current deposit - %d - is not enough. The minimum required is %d", val.deposit, min)
	}

	options := getTransactionOpts(val.walletAccount, deposit.SetUint64(val.deposit), val.config.ChainID)
	_, err = val.network.Deposit(options)
	if err != nil {
		return fmt.Errorf("Failed to transact the deposit: %x", err)
	}

	return nil
}

func (val *validator) leave() {
	options := getTransactionOpts(val.walletAccount, nil, val.config.ChainID)
	_, err := val.network.Leave(options)
	if err != nil {
		log.Error("Failed to leave the election", "err", err)
	}
}

func (val *validator) createProposalBlock() *types.Block {
	if val.lockedBlock != nil {
		log.Info("Picking a locked block")
		return val.lockedBlock
	}
	return val.createBlock()
}

func (val *validator) createBlock() *types.Block {
	log.Info("Creating a new block")
	// new block header
	parent := val.chain.CurrentBlock()
	blockNumber := parent.Number()
	tstart := time.Now()
	tstamp := tstart.Unix()
	if parent.Time().Cmp(new(big.Int).SetInt64(tstamp)) >= 0 {
		tstamp = parent.Time().Int64() + 1
	}
	header := &types.Header{
		ParentHash: parent.Hash(),
		Coinbase:   val.walletAccount.Account().Address,
		Number:     blockNumber.Add(blockNumber, common.Big1),
		GasLimit:   core.CalcGasLimit(parent),
		GasUsed:    new(big.Int),
		Time:       big.NewInt(tstamp),
	}
	val.header = header

	var commit *types.Commit

	// @NOTE (rgeraldes) - temporary
	first := types.NewVote(blockNumber, parent.Hash(), 0, types.PreCommit)

	if blockNumber.Cmp(big.NewInt(1)) == 0 {
		commit = &types.Commit{
			PreCommits:     types.Votes{first},
			FirstPreCommit: first,
		}
	} else {
		commit = &types.Commit{
			PreCommits:     types.Votes{first},
			FirstPreCommit: first,
		}
		//commit = val.lastCommit.Proof()
	}

	if err := val.engine.Prepare(val.chain, header); err != nil {
		log.Error("Failed to prepare header for mining", "err", err)
		// @TODO (rgeraldes) - review returning nil
		return nil
	}

	pending, err := val.backend.TxPool().Pending()
	if err != nil {
		log.Crit("Failed to fetch pending transactions", "err", err)
	}

	txs := types.NewTransactionsByPriceAndNonce(val.signer, pending)
	val.commitTransactions(val.eventMux, txs, val.chain, val.walletAccount.Account().Address)

	// Create the new block to seal with the consensus engine
	var block *types.Block
	if block, err = val.engine.Finalize(val.chain, header, val.state, val.txs, commit, val.receipts); err != nil {
		log.Crit("Failed to finalize block for sealing", "err", err)
	}

	return block
}

func (val *validator) propose() {
	block := val.createProposalBlock()

	//lockedRound, lockedBlock := val.votes.LockingInfo()
	lockedRound := 1
	lockedBlock := common.Hash{}

	// @TODO (rgeraldes) - review int/int64; address situation where validators size might be zero (no peers)
	// @NOTE (rgeraldes) - (for now size = block size) number of block fragments = number of validators - self
	fragments, err := block.AsFragments(int(block.Size().Int64()) /*/val.validators.Size() - 1 */)
	if err != nil {
		// @TODO(rgeraldes) - analyse consequences
		log.Crit("Failed to get the block as a set of fragments of information", "err", err)
	}

	proposal := types.NewProposal(val.blockNumber, val.round, fragments.Metadata(), lockedRound, lockedBlock)

	signedProposal, err := val.walletAccount.SignProposal(val.walletAccount.Account(), proposal, val.config.ChainID)
	if err != nil {
		log.Crit("Failed to sign the proposal", "err", err)
	}

	val.proposal = signedProposal
	val.block = block

	val.eventMux.Post(core.NewProposalEvent{Proposal: proposal})

	// post block segments events
	// @TODO(rgeraldes) - review types int/uint
	for i := uint(0); i < fragments.Size(); i++ {
		val.eventMux.Post(core.NewBlockFragmentEvent{
			BlockNumber: val.blockNumber,
			Round:       val.round,
			Data:        fragments.Get(int(i)),
		})
	}

}

func (val *validator) preVote() {
	var vote common.Hash
	switch {
	case val.lockedBlock != nil:
		log.Debug("Locked Block is not nil, voting for the locked block")
		vote = val.lockedBlock.Hash()
	case val.block == nil:
		log.Debug("Proposal's block is nil, voting nil")
		vote = common.Hash{}
	default:
		log.Debug("Voting for the proposal's block")
		vote = val.block.Hash()
	}

	val.vote(types.NewVote(val.blockNumber, vote, val.round, types.PreVote))
}

func (val *validator) preCommit() {
	var vote common.Hash
	// access prevotes
	winner := common.Hash{}
	switch {
	// no majority
	//case !val.hasPolka():
	// majority pre-voted nil
	case winner == common.Hash{}:
		log.Debug("Majority of validators pre-voted nil")
		// unlock locked block
		if val.lockedBlock != nil {
			val.lockedRound = 0
			val.lockedBlock = nil
		}
	case winner == val.lockedBlock.Hash():
		log.Debug("Majority of validators pre-voted the locked block")
		// update locked block round
		val.lockedRound = val.round
		// vote on the pre-vote election winner
		vote = winner
	case winner == val.block.Hash():
		log.Debug("Majority of validators pre-voted the proposed block")
		// lock block
		val.lockedRound = val.round
		val.lockedBlock = val.block
		// vote on the pre-vote election winner
		vote = winner
		// we don't have the current block (fetch)
		// @TODO (tendermint): in the future save the POL prevotes for justification.
	default:
		// fetch block, unlock, precommit
		// unlock locked block
		val.lockedRound = 0
		val.lockedBlock = nil
		//val.lockedBlockParts = nil
		//if !cs.ProposalBlockParts.HasHeader(blockID.PartsHeader) {
		val.block = nil
		//val.ProposalBlockParts = types.NewPartSetFromHeader(blockID.PartsHeader)
		//}
	}

	val.vote(types.NewVote(val.blockNumber, vote, val.round, types.PreCommit))
}

func (val *validator) vote(vote *types.Vote) {
	signedVote, err := val.walletAccount.SignVote(val.walletAccount.Account(), vote, val.config.ChainID)
	if err != nil {
		log.Crit("Failed to sign the vote", "err", err)
	}

	val.votingSystem.Add(signedVote, true)
}

func (val *validator) AddBlockFragment(blockNumber *big.Int, round uint64, fragment *types.BlockFragment) error {
	if !val.Validating() {
		return ErrCantAddBlockFragmentNotValidating
	}
	val.blockFragments.Add(fragment)

	// @NOTE (rgeraldes) - the whole section needs to be refactored
	if val.blockFragments.HasAll() {
		block, err := val.blockFragments.Assemble()
		if err != nil {
			log.Crit("Failed to assemble the block", "err", err)
		}

		// @TODO (rgeraldes) - refactor ; based on core/blockchain.go (InsertChain)
		// Start the parallel header verifier
		nBlocks := 1
		headers := make([]*types.Header, nBlocks)
		seals := make([]bool, nBlocks)
		headers[nBlocks-1] = block.Header()
		seals[nBlocks-1] = true

		abort, results := val.engine.VerifyHeaders(val.chain, headers, seals)
		defer close(abort)

		err = <-results
		if err == nil {
			err = val.chain.Validator().ValidateBody(block)
		}

		// @NOTE(rgeraldes) - ignore for now (assume that the block is ok)
		/*
			if err != nil {
				if err == ErrKnownBlock {
					stats.ignored++
					continue
				}

				if err == consensus.ErrFutureBlock {
					// Allow up to MaxFuture second in the future blocks. If this limit
					// is exceeded the chain is discarded and processed at a later time
					// if given.
					max := big.NewInt(time.Now().Unix() + maxTimeFutureBlocks)
					if block.Time().Cmp(max) > 0 {
						return i, fmt.Errorf("future block: %v > %v", block.Time(), max)
					}
					bc.futureBlocks.Add(block.Hash(), block)
					stats.queued++
					continue
				}

				if err == consensus.ErrUnknownAncestor && bc.futureBlocks.Contains(block.ParentHash()) {
					bc.futureBlocks.Add(block.Hash(), block)
					stats.queued++
					continue
				}

				bc.reportBlock(block, nil, err)
				return i, err
			}
		*/
		parent := val.chain.GetBlock(block.ParentHash(), block.NumberU64()-1)

		// Process block using the parent state as reference point.
		receipts, _, usedGas, err := val.chain.Processor().Process(block, val.state, val.vmConfig)
		if err != nil {
			log.Crit("Failed to process the block", "err", err)
			//bc.reportBlock(block, receipts, err)
			//return i, err
		}
		val.receipts = receipts

		// Validate the state using the default validator
		err = val.chain.Validator().ValidateState(block, parent, val.state, receipts, usedGas)
		if err != nil {
			log.Crit("Failed to validate the state", "err", err)
			//bc.reportBlock(block, receipts, err)
			//return i, err
		}

		val.block = block

		go func() { val.blockCh <- block }()
	}
	return nil
}

func (val *validator) makeCurrent(parent *types.Block) error {
	state, err := val.chain.StateAt(parent.Root())
	if err != nil {
		return err
	}
	work := &work{
		state: state,
	}

	// Keep track of transactions which return errors so they can be removed
	work.tcount = 0
	val.work = work
	return nil
}

func (val *validator) updateValidators(checksum [32]byte, genesis bool) error {
	count, err := val.network.GetValidatorCount(&bind.CallOpts{})
	if err != nil {
		return err
	}

	val.validatorsChecksum = checksum
	validators := make([]*types.Validator, count.Uint64())
	for i := int64(0); i < count.Int64(); i++ {
		validator, err := val.network.GetValidatorAtIndex(&bind.CallOpts{}, big.NewInt(i))
		if err != nil {
			return err
		}

		var weight *big.Int
		// @TODO (rgeraldes) - weight needs to be shared
		/*
			if !genesis && val.validators.Contains(validator.Addr) {
				// old validator
				old := val.validators.Get(validator.Addr)
				weight = old.Weight()
			} else {
				// new validator
				weight = big.NewInt(0)
			}
		*/
		// @TODO (rgeraldes) - remove this statement as soon as the previous one is sorted out
		weight = big.NewInt(0)

		validators[i] = types.NewValidator(validator.Code, validator.Deposit.Uint64(), weight)
	}
	val.validators = types.NewValidatorSet(validators)
	val.validatorsChecksum = checksum

	return nil
}
