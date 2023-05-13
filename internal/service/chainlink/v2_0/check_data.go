package v2_0

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"math/big"
)

type CheckDataService struct {
	args abi.Arguments
}

func NewCheckDataService() *CheckDataService {
	uint256Type, _ := abi.NewType("uint256", "", nil)
	return &CheckDataService{
		args: abi.Arguments{
			{Name: "performOffset", Type: uint256Type},
			{Name: "performLimit", Type: uint256Type},
		},
	}
}

func (s *CheckDataService) NewCheckData(offset uint64, limit uint64) ([]byte, error) {
	return abi.Arguments{s.args[0], s.args[1]}.Pack(
		new(big.Int).SetUint64(limit),
		new(big.Int).SetUint64(offset),
	)
}
