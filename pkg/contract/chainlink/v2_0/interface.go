package v2_0

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

type Upkeeper interface {
	CheckUpkeep(opts *bind.CallOpts, checkData []byte) (struct {
		UpkeepNeeded bool
		PerformData  []byte
	}, error)
	PerformUpkeep(opts *bind.TransactOpts, performData []byte) (*types.Transaction, error)
}
