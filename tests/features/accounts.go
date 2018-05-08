package features

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	"github.com/DATA-DOG/godog/gherkin"
	"github.com/kowala-tech/kcoin/accounts"
)

var (
	NoAccount = accounts.Account{}

	unnamedAccountName = "no-name"
)

type AccountEntry struct {
	AccountName     string
	AccountPassword string
	Funds           int64
}

func parseAccountsDataTable(accountsDataTable *gherkin.DataTable) ([]*AccountEntry, error) {
	var fields []string
	head := accountsDataTable.Rows[0].Cells
	for _, cell := range head {
		fields = append(fields, cell.Value)
	}

	var accounts []*AccountEntry

	for i := 1; i < len(accountsDataTable.Rows); i++ {
		account := &AccountEntry{}
		for n, cell := range accountsDataTable.Rows[i].Cells {
			switch head[n].Value {
			case "account":
				account.AccountName = cell.Value
			case "password":
				account.AccountPassword = cell.Value
			case "funds":
				parsed, err := strconv.ParseInt(cell.Value, 10, 64)
				if err != nil {
					return nil, err
				}
				account.Funds = parsed
			default:
				return nil, fmt.Errorf("unexpected column name: %s", head[n].Value)
			}
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (ctx *Context) ICreatedAnAccountWithPassword(password string) error {
	_, err := ctx.createAccount(unnamedAccountName, password)
	return err
}

func (ctx *Context) IHaveTheFollowingAccounts(accountsDataTable *gherkin.DataTable) error {
	accountsData, err := parseAccountsDataTable(accountsDataTable)
	if err != nil {
		return err
	}

	for _, accountData := range accountsData {
		acct, err := ctx.createAccount(accountData.AccountName, accountData.AccountPassword)
		if err != nil {
			return err
		}
		if _, err := ctx.sendFundsAndWait(ctx.seederAccount, acct, accountData.Funds); err != nil {
			return err
		}
	}

	return nil
}

func (ctx *Context) ITryUnlockMyAccountWithPassword(password string) error {
	return ctx.ITryUnlockAccountWithPassword(unnamedAccountName, password)
}

func (ctx *Context) ITryUnlockAccountWithPassword(accountName, password string) error {
	ctx.lastUnlockErr = ctx.IUnlockAccountWithPassword(accountName, password)
	return nil
}

func (ctx *Context) IUnlockAccountWithPassword(accountName, password string) error {
	acct, found := ctx.accounts[accountName]
	if !found {
		return fmt.Errorf("account not created")
	}
	return ctx.AccountsStorage.Unlock(acct, password)
}

func (ctx *Context) IGotAccountUnlocked() error {
	return ctx.lastUnlockErr
}
func (ctx *Context) IGotErrorUnlocking() error {
	if ctx.lastUnlockErr == nil {
		return fmt.Errorf("unlocking expected to fail, but didn't fail.")
	}
	return nil
}

func (ctx *Context) TheBalanceIsExactly(accountName string, kcoin int64) error {
	expected := toWei(kcoin)

	account := ctx.accounts[accountName]
	balance, err := ctx.client.BalanceAt(context.Background(), account.Address, nil)
	if err != nil {
		return err
	}
	if balance.Cmp(expected) != 0 {
		return fmt.Errorf("Balance expected to be %v but is %v", expected, balance)
	}
	return nil
}

func (ctx *Context) TheBalanceIsAround(accountName string, kcoin int64) error {
	expected := toWei(kcoin)

	account := ctx.accounts[accountName]
	balance, err := ctx.client.BalanceAt(context.Background(), account.Address, nil)
	if err != nil {
		return err
	}
	diff := &big.Int{}
	diff.Sub(balance, expected)
	diff.Abs(diff)

	if diff.Cmp(big.NewInt(100000)) >= 0 {
		return fmt.Errorf("Balance expected to be around %v but is %v", expected, balance)
	}
	return nil
}

func (ctx *Context) createAccount(accountName string, password string) (accounts.Account, error) {
	if _, ok := ctx.accounts[accountName]; ok {
		return NoAccount, fmt.Errorf("an account with this name already exists: %s", accountName)
	}
	account, err := ctx.AccountsStorage.NewAccount(password)
	if err != nil {
		return NoAccount, err
	}

	ctx.accounts[accountName] = account
	return account, nil
}
