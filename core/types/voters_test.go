package types

import (
	"fmt"
	"github.com/kowala-tech/kUSD/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestValidator_Properties(t *testing.T) {
	address := common.Address{}
	deposit := uint64(100)
	weight := &big.Int{}
	voter := NewVoter(address, deposit, weight)

	assert.Equal(t, address, voter.Address())
	assert.Equal(t, deposit, voter.Deposit())
	assert.Equal(t, weight, voter.Weight())
}

func TestValidatorSet_EmptyReturnsError(t *testing.T) {
	voters, err := NewVoters(nil)

	require.Error(t, err)
	require.Nil(t, voters)
}

func TestValidatorSet_GetAtNegativeIndexReturnsNil(t *testing.T) {
	voter := makeValidator("0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed", 100, 100)
	voters, err := NewVoters([]*Voter{voter})
	require.NoError(t, err)

	voter = voters.At(-1)

	assert.Nil(t, voter)
}

func TestValidatorSet_GetAtOverLastReturnsNil(t *testing.T) {
	voter := makeValidator("0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed", 100, 100)
	voters, err := NewVoters([]*Voter{voter})
	require.NoError(t, err)

	voterAt := voters.At(0)
	assert.Equal(t, voter, voterAt)

	voterAt = voters.At(1)
	assert.Nil(t, voterAt)
}

func TestValidatorSet_One(t *testing.T) {
	address := common.HexToAddress("0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed")
	deposit := uint64(100)
	weight := &big.Int{}
	voter := NewVoter(address, deposit, weight)

	voters, err := NewVoters([]*Voter{voter})

	require.NoError(t, err)
	assert.Equal(t, 1, voters.Size())
	assert.Equal(t, voter, voters.At(0))
	assert.Equal(t, voter, voters.Get(address))
	assert.Equal(t, voter, voters.NextProposer())
}

func TestValidatorSet_UpdateWeightChangesProposer(t *testing.T) {
	voter := makeValidator("0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed", 100, 100)
	voter2 := makeValidator("0x6aaeb6053f3e94c9b9a09f33669435e7ef1beaed", 101, 101)
	voter3 := makeValidator("0x7aaeb6053f3e94c9b9a09f33669435e7ef1beaed", 99, 99)

	voters, err := NewVoters([]*Voter{voter, voter2, voter3})

	require.NoError(t, err)
	proposer := voters.NextProposer()
	assert.Equal(t, voter2, proposer)
	assert.Equal(t, big.NewInt(101), proposer.weight)
	assert.Equal(t, big.NewInt(200), voters.At(0).weight)
	assert.Equal(t, big.NewInt(101), voters.At(1).weight)
	assert.Equal(t, big.NewInt(198), voters.At(2).weight)
	assert.Equal(t, 3, voters.Size())
}

func TestValidatorSet_UpdateWeightChangesProposerElections(t *testing.T) {
	voter := makeValidator("0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed", 100, 100)
	voter2 := makeValidator("0x6aaeb6053f3e94c9b9a09f33669435e7ef1beaed", 101, 101)
	voter3 := makeValidator("0x7aaeb6053f3e94c9b9a09f33669435e7ef1beaed", 99, 99)

	voters, err := NewVoters([]*Voter{voter, voter2, voter3})
	require.NoError(t, err)
	require.Equal(t, 3, voters.Size())

	elections := []struct {
		proposerWeight *big.Int
		voter1weight   *big.Int
		voter2weight   *big.Int
		voter3weight   *big.Int
	}{
		{big.NewInt(101), big.NewInt(200), big.NewInt(101), big.NewInt(198)},
		{big.NewInt(200), big.NewInt(200), big.NewInt(202), big.NewInt(297)},
		{big.NewInt(297), big.NewInt(300), big.NewInt(303), big.NewInt(297)},
		{big.NewInt(303), big.NewInt(400), big.NewInt(303), big.NewInt(396)},
	}

	for round, tc := range elections {
		t.Run(fmt.Sprintf("round %d", round), func(t *testing.T) {
			proposer := voters.NextProposer()
			assert.Equal(t, tc.proposerWeight, proposer.weight)
			assert.Equal(t, tc.voter1weight, voters.At(0).weight)
			assert.Equal(t, tc.voter2weight, voters.At(1).weight)
			assert.Equal(t, tc.voter3weight, voters.At(2).weight)
		})
	}
}

func makeValidator(hexAddress string, deposit int, weight int64) *Voter {
	address := common.HexToAddress(hexAddress)
	return NewVoter(address, uint64(deposit), big.NewInt(weight))
}
