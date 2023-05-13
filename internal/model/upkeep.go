package model

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Upkeep struct {
	Name           string
	EncryptedEmail []byte
	UpkeepContract ContractInfo
	GasLimit       uint32
	AdminAddress   common.Address
	CheckData      []byte
	Limit          uint64
	Offset         uint64
	OffchainConfig []byte
	Amount         *big.Int
	Sender         common.Address
}

type UpkeepTopup struct {
	UpkeepId big.Int
	Amount   big.Int
}
