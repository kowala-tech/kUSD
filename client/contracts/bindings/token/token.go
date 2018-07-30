package token

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/accounts"
	"github.com/kowala-tech/kcoin/client/common"
)

type Token interface {
	Transfer(walletAccount accounts.WalletAccount, to common.Address, value *big.Int, data []byte, customFallback string) (common.Hash, error)
	BalanceOf(target common.Address) (*big.Int, error)
	Name() (string, error)
	Cap() (*big.Int, error)
	TotalSupply() (*big.Int, error)
}
