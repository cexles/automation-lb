package chainlink

import (
	"context"
	"github.com/cexles/automation-lb/internal/config"
	"github.com/cexles/automation-lb/internal/service"
	"github.com/cexles/automation-lb/pkg/contract/chainlink"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"math/big"
)

type LinkTokenService struct {
	ctx            context.Context
	cfg            *config.Config
	accountService *service.AccountService
	contract       *chainlink.LinkTokenContract
	moduleName     string
}

func NewLinkTokenService(ctx context.Context, cfg *config.Config, accountService *service.AccountService, contract *chainlink.LinkTokenContract) *LinkTokenService {
	return &LinkTokenService{
		ctx:            ctx,
		cfg:            cfg,
		accountService: accountService,
		contract:       contract,
		moduleName:     "service.chainlink.linkTokenService",
	}
}

func (s *LinkTokenService) IncreaseAllowanceIfNotEnough(ctx context.Context, spender common.Address, minApprove *big.Int) error {
	approve, err := s.CheckAllowance(context.Background(), spender)
	if err != nil {
		return err
	}

	approveCmp := approve.Cmp(minApprove)
	if approveCmp < 1 {
		err = s.IncreaseAllowance(ctx, spender, minApprove)
		if err != nil {
			return err
		}
		log.Info().Str("module", s.moduleName).
			Str("func", "IncreaseAllowanceIfNotEnough").
			Fields(map[string]any{
				"spender":        spender.String(),
				"approved_value": minApprove.Mul(minApprove, big.NewInt(2)),
			}).Msg("Allowance was successfully increased")
		return nil
	}
	return nil
}

func (s *LinkTokenService) CheckAllowance(ctx context.Context, spender common.Address) (*big.Int, error) {
	allowance, err := s.contract.Allowance(&bind.CallOpts{Context: ctx}, s.accountService.GetAccount().Address(), spender)
	if err != nil {
		log.Err(err).Str("module", s.moduleName).
			Str("func", "CheckAllowance").
			Fields(map[string]any{
				"owner_address":   s.accountService.GetAccount().Address(),
				"spender_address": spender.String(),
			}).Msg("Failed to approve LINK token")
		return nil, err
	}
	return allowance, err
}

func (s *LinkTokenService) Approve(ctx context.Context, spender common.Address, amount *big.Int) error {
	opts, err := s.accountService.NewTxOpts(ctx)
	if err != nil {
		return err
	}

	_, err = s.contract.Approve(opts, spender, amount)
	if err != nil {
		log.Err(err).Str("module", s.moduleName).
			Str("func", "Approve").
			Fields(map[string]any{
				"owner_address":   s.accountService.GetAccount().Address(),
				"spender_address": spender.String(),
			}).Msg("Failed to approve LINK token")
	}
	return err
}

func (s *LinkTokenService) IncreaseAllowance(ctx context.Context, spender common.Address, amount *big.Int) error {
	opts, err := s.accountService.NewTxOpts(ctx)
	if err != nil {
		return err
	}

	_, err = s.contract.IncreaseApproval(opts, spender, amount)
	if err != nil {
		log.Err(err).Str("module", s.moduleName).
			Str("func", "IncreaseAllowance").
			Fields(map[string]any{
				"owner_address":   s.accountService.GetAccount().Address(),
				"spender_address": spender.String(),
				"increase_amount": amount.String(),
			}).Msg("Failed to increase allowance of LINK token")
	}
	return err
}
