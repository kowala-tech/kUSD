pragma solidity 0.4.21;

import "github.com/kowala-tech/kcoin/contracts/lifecycle/contracts/Pausable.sol" as pausable;
import "github.com/kowala-tech/kcoin/contracts/token/contracts/TokenReceiver.sol" as receiver;

contract ValidatorManager is pausable.Pausable, receiver.TokenReceiver {
    uint public baseDeposit;       
    uint public maxValidators;
    // period in days
    uint public unbondingPeriod;
    address public genesisValidator;

    // validatorsChecksum is a representation of the current set of validators
    bytes32 public validatorsChecksum;

    // Deposit represents the collateral - staked tokens
    struct Deposit {
        uint amount;
        uint availableAt;
    }

    // Validator represents a consensus validator      
    struct Validator {
        uint index;
        bool isValidator;

        // @NOTE (rgeraldes) - users can have more than one deposit
        // Example: user leaves and re-enters the election. At this point
        // the initial deposit will have a release date and the validator 
        // will have a new deposit for the current election.
        Deposit[] deposits; 
    }
    
    mapping (address => Validator) private validatorRegistry;
    
    // validatorPool contains the validator code ordered by the biggest deposit to
    // the smallest deposit.
    address[] validatorPool;

    // onlyWithMinDeposit requires a minimum deposit to proceed
    modifier onlyWithMinDeposit {
        require(msg.value >= getMinimumDeposit());
        _;
    }

    // onlyValidator requires the sender to be a validator
    modifier onlyValidator {
        require(isValidator(msg.sender));
        _;
    }

    // onlyNewCandidate required the sender to be a new candidate
    modifier onlyNewCandidate {
        require(!isValidator(msg.sender));
        _;
    }

    function ValidatorManager(uint _baseDeposit, uint _maxValidators, uint _unbondingPeriod, address _genesis) public {
        require(_maxValidators >= 1);

        baseDeposit = _baseDeposit * 1 ether;
        maxValidators = _maxValidators;
        unbondingPeriod = _unbondingPeriod * 1 days;
        genesisValidator = _genesis;
    
        _insertValidator(_genesis, baseDeposit);
    }

    function isGenesisValidator(address code) public view returns (bool isIndeed) {
        return code == genesisValidator;
    }

    function isValidator(address code) public view returns (bool isIndeed) {
        return validatorRegistry[code].isValidator;
    }

    function getValidatorCount() public view returns (uint count) {
        return validatorPool.length;
    }

    function getValidatorAtIndex(uint index) public view returns (address code, uint deposit) {
        code = validatorPool[index];
        Validator validator = validatorRegistry[code];
        deposit = validator.deposits[validator.deposits.length - 1].amount;
    }

    function _hasAvailability() public view returns (bool available) {
        return (maxValidators - validatorPool.length) > 0;
    }

    // getMinimumDeposit returns the base deposit if there are positions available or
    // the current smallest deposit required if there aren't positions availabe.
    function getMinimumDeposit() public view returns (uint deposit) {
        // there are positions for validator available
        if (_hasAvailability()) {
            return baseDeposit;
        } else {
            Validator smallestBidder = validatorRegistry[validatorPool[validatorPool.length - 1]];               
            return smallestBidder.deposits[smallestBidder.deposits.length - 1].amount + 1;
        }
    }

    function _updateChecksum() private {
        validatorsChecksum = keccak256(validatorPool);
    }

    function _insertValidator(address code, uint deposit) private {
        Validator sender = validatorRegistry[code];
        sender.index = validatorPool.push(code) - 1;
        sender.isValidator = true;
        sender.deposits.push(Deposit({amount:deposit, availableAt: 0}));

        for (uint index = sender.index; index > 0; index--) {
            Validator target = validatorRegistry[validatorPool[index - 1]];
            Deposit collateral = target.deposits[target.deposits.length - 1];
            if (deposit <= collateral.amount) {
                break;
            }
            validatorPool[index] = validatorPool[index - 1];
            validatorPool[index - 1] = code; 
            // update indexes
            target.index = index;
            sender.index = index - 1;
        }

        _updateChecksum();
    }

    function setBaseDeposit(uint deposit) public onlyOwner {
        baseDeposit = deposit;
    }

    function _deleteValidator(address account) private {
        Validator validator = validatorRegistry[account];
        for (uint index = validator.index; index < validatorPool.length - 1; index++) {
            validatorPool[index] = validatorPool[index + 1];
        }
        validatorPool.length--;

        validator.isValidator = false;
        validator.deposits[validator.deposits.length - 1].availableAt = now + unbondingPeriod;

        _updateChecksum();
    }

    // _deleteSmallestBidder removes the validator with the smallest deposit
    function _deleteSmallestBidder() private {
        _deleteValidator(validatorPool[validatorPool.length - 1]);
    }

    function setMaxValidators(uint max) public onlyOwner { 
        if (max < validatorPool.length) {
            uint toRemove = validatorPool.length - max;
            for (uint i = 0; i < toRemove; i++) {
                _deleteSmallestBidder();
            }
        }
        maxValidators = max;   
    }

    function getDepositCount() public view returns (uint count) {
        return validatorRegistry[msg.sender].deposits.length; 
    }

    function getDepositAtIndex(uint index) public view returns (uint amount, uint availableAt) {
        Deposit deposit = validatorRegistry[msg.sender].deposits[index];
        return (deposit.amount / 1 ether, deposit.availableAt);
    }

    function registerValidator() public tokenPayable whenNotPaused onlyNewCandidate onlyWithMinDeposit {
        if (!_hasAvailability()) {
            _deleteSmallestBidder();
        }
        _insertValidator(tkn.sender, tkn.value);
    }

    function deregisterValidator() public whenNotPaused onlyValidator {
        _deleteValidator(msg.sender);
    }

    function _removeDeposits(address code, uint index) private {
        if (index == 0) return;

        Validator validator = validatorRegistry[code];
        uint lo = 0;
        uint hi = index;
        while (hi < validator.deposits.length) {
            validator.deposits[lo] = validator.deposits[hi];
            lo++;
            hi++;
        }
        validator.deposits.length = lo;
    }

    // releaseDeposits transfers locked deposit(s) back the user account if they
    // are past the unbonding period
    function releaseDeposits() public whenNotPaused {
        uint refund = 0;
        uint i = 0;
        Deposit[] deposits = validatorRegistry[msg.sender].deposits;
        
        for (; i < deposits.length && deposits[i].availableAt != 0; i++) {
            if (now < deposits[i].availableAt) {
                // @NOTE (rgeraldes) - no need to iterate further since the 
                // release date (if is different than 0) of the following deposits
                // will always be past than the current one.
                break;
            }
            refund += deposits[i].amount;
        }
        
        _removeDeposits(msg.sender, i);

        if (refund > 0) {
            msg.sender.transfer(refund);
        }
    }

    // ERC223 - mUSD support
    struct TKN {
        address sender;
        uint value;
        bytes data;
        bytes4 sig;
    }

    TKN tkn;

    bool isTokenFallback = true;

    modifier tokenPayable {
        if (!isTokenFallback) revert();
        _;
    }

    function tokenFallback(address _from, uint _value, bytes _data) public {}

    // @NOTE (rgeraldes) - limitation - this is necessary for now since the golang bindings do not support
    // method overloading and the choosen method for the bindings will be the method with the
    // biggest signature which requires a custom fallback for ERC233 token transfer.
    function customFallback(address _from, uint _value, bytes _data) public {
        uint32 u = uint32(_data[3]) + (uint32(_data[2]) << 8) + (uint32(_data[1]) << 16) + (uint32(_data[0]) << 24);
        tkn = TKN(_from, _value, _data, bytes4(u));
        if (!address(this).delegatecall(_data)) revert();
    }
}