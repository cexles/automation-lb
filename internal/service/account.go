package service

import (
	"context"
	"github.com/cexles/automation-lb/internal/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/rs/zerolog/log"
)

type AccountService struct {
	account    *model.Account
	moduleName string
}

func NewAccountService(account *model.Account) *AccountService {
	return &AccountService{account: account, moduleName: "service.AccountService"}
}

func (s *AccountService) GetAccount() *model.Account {
	return s.account
}

func (s *AccountService) NewTxOpts(ctx context.Context) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(s.account.PrivateKey(), s.account.ChainId())
	if err != nil {
		log.Err(err).Str("module", s.moduleName).
			Str("func", "NewTxOpts").
			Str("address", s.account.Address().String()).
			Msg("Failed to create tx options")
		return nil, err
	}
	opts.Context = ctx
	return opts, nil
}
