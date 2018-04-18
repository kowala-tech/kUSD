package tests

import (
	"math/big"

	"github.com/DATA-DOG/godog"
	"github.com/kowala-tech/kcoin/tests/features"
)

var (
	chainID = big.NewInt(519374298533)
)

func FeatureContext(s *godog.Suite) {
	context := features.NewTestContext(chainID)

	s.BeforeSuite(func() {
		if err := context.PrepareCluster(); err != nil {
			panic(err)
		}
	})

	s.AfterSuite(func() {
		if err := context.DeleteCluster(); err != nil {
			panic(err)
		}
	})

	s.BeforeScenario(func(interface{}) {
		context.Reset()
	})
	s.Step(`^I have the following accounts:$`, context.IHaveTheFollowingAccounts)
	s.Step(`^I transfer (\d+) kcoins? from (\w+) to (\w+)$`, context.ITransferKUSD)
	s.Step(`^I try to transfer (\d+) kcoins? from (\w+) to (\w+)$`, context.ITryTransferKUSD)
	s.Step(`^the balance of (\w+) should be (\d+) kcoins?$`, context.TheBalanceIsExactly)
	s.Step(`^the balance of (\w+) should be around (\d+) kcoins?$`, context.TheBalanceIsAround)
	s.Step(`^the transaction should fail$`, context.LastTransactionFailed)
}
