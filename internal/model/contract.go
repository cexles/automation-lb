package model

import (
	"github.com/ethereum/go-ethereum/common"
)

type ContractInfo struct {
	Name           string         `yaml:"name"`
	ShortName      string         `yaml:"short_name"`
	Address        common.Address `yaml:"address"`
	Version        string         `yaml:"version"`
	GasLimit       uint32         `yaml:"gas_limit"`
	ExecutionLimit uint64         `yaml:"execution_limit"`
}
