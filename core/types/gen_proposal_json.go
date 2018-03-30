// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/common/hexutil"
)

var _ = (*proposaldataMarshalling)(nil)

func (p proposaldata) MarshalJSON() ([]byte, error) {
	type proposaldata struct {
		BlockNumber   *hexutil.Big   `json:"blockNumber"   gencodec:"required"`
		Round         hexutil.Uint64 `json:"round"         gencodec:"required"`
		LockedRound   hexutil.Uint64 `json:"lockedRound"   gencodec:"required"`
		LockedBlock   common.Hash    `json:"lockedBlock"   gencodec:"required"`
		BlockMetadata *Metadata      `json:"metadata"      gencodec:"required"`
		V             *hexutil.Big   `json:"v"      gencodec:"required"`
		R             *hexutil.Big   `json:"r"      gencodec:"required"`
		S             *hexutil.Big   `json:"s"      gencodec:"required"`
	}
	var enc proposaldata
	enc.BlockNumber = (*hexutil.Big)(p.BlockNumber)
	enc.Round = hexutil.Uint64(p.Round)
	enc.LockedRound = hexutil.Uint64(p.LockedRound)
	enc.LockedBlock = p.LockedBlock
	enc.BlockMetadata = p.BlockMetadata
	enc.V = (*hexutil.Big)(p.V)
	enc.R = (*hexutil.Big)(p.R)
	enc.S = (*hexutil.Big)(p.S)
	return json.Marshal(&enc)
}

func (p *proposaldata) UnmarshalJSON(input []byte) error {
	type proposaldata struct {
		BlockNumber   *hexutil.Big    `json:"blockNumber"   gencodec:"required"`
		Round         *hexutil.Uint64 `json:"round"         gencodec:"required"`
		LockedRound   *hexutil.Uint64 `json:"lockedRound"   gencodec:"required"`
		LockedBlock   *common.Hash    `json:"lockedBlock"   gencodec:"required"`
		BlockMetadata *Metadata       `json:"metadata"      gencodec:"required"`
		V             *hexutil.Big    `json:"v"      gencodec:"required"`
		R             *hexutil.Big    `json:"r"      gencodec:"required"`
		S             *hexutil.Big    `json:"s"      gencodec:"required"`
	}
	var dec proposaldata
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.BlockNumber == nil {
		return errors.New("missing required field 'blockNumber' for proposaldata")
	}
	p.BlockNumber = (*big.Int)(dec.BlockNumber)
	if dec.Round == nil {
		return errors.New("missing required field 'round' for proposaldata")
	}
	p.Round = uint64(*dec.Round)
	if dec.LockedRound == nil {
		return errors.New("missing required field 'lockedRound' for proposaldata")
	}
	p.LockedRound = uint64(*dec.LockedRound)
	if dec.LockedBlock == nil {
		return errors.New("missing required field 'lockedBlock' for proposaldata")
	}
	p.LockedBlock = *dec.LockedBlock
	if dec.BlockMetadata == nil {
		return errors.New("missing required field 'metadata' for proposaldata")
	}
	p.BlockMetadata = dec.BlockMetadata
	if dec.V == nil {
		return errors.New("missing required field 'v' for proposaldata")
	}
	p.V = (*big.Int)(dec.V)
	if dec.R == nil {
		return errors.New("missing required field 'r' for proposaldata")
	}
	p.R = (*big.Int)(dec.R)
	if dec.S == nil {
		return errors.New("missing required field 's' for proposaldata")
	}
	p.S = (*big.Int)(dec.S)
	return nil
}
