package validator

import (
	"errors"
	"math/big"

	"github.com/kowala-tech/kUSD/accounts"
	"github.com/kowala-tech/kUSD/accounts/abi/bind"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/core/types"
)

// getTransactionOpts returns a new set of transaction options
func getTransactionOpts(walletAccount accounts.WalletAccount, value *big.Int, chainID *big.Int) *bind.TransactOpts {
	account := walletAccount.Account()
	opts := &bind.TransactOpts{
		From: account.Address,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			// @NOTE (rgeraldes) - ignore the proposed signer as by default it will be a unprotected signer.
			if address != account.Address {
				return nil, errors.New("not authorized to sign this account")
			}
			return walletAccount.SignTx(account, tx, chainID)
		},
	}
	if value != nil {
		opts.Value = value
	}

	return opts
}
