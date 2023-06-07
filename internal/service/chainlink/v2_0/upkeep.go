package v2_0

import (
	"context"
	"errors"
	"fmt"
	"github.com/cexles/automation-lb/internal/config"
	"github.com/cexles/automation-lb/internal/model"
	"github.com/cexles/automation-lb/internal/service"
	"github.com/cexles/automation-lb/pkg/contract/chainlink/v2_0"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	"math/big"
	"time"
)

type UpkeepService struct {
	ctx                   context.Context
	cfg                   *config.Config
	client                *ethclient.Client
	checkDataService      *CheckDataService
	accountService        *service.AccountService
	upkeepContract        *v2_0.UpkeepControllerContract
	upkeepContractAddress common.Address
	adminAddress          common.Address
	logicContractInfo     model.ContractInfo
	newUpkeepCh           chan model.Upkeep
	topUpCh               chan model.UpkeepTopup
	moduleName            string
}

func NewUpkeepService(ctx context.Context, cfg *config.Config, client *ethclient.Client, checkDataService *CheckDataService, accountService *service.AccountService, upkeepContract *v2_0.UpkeepControllerContract, upkeepContractAddress common.Address, adminAddress common.Address, logicContractInfo model.ContractInfo, newUpkeepCh chan model.Upkeep, topUpCh chan model.UpkeepTopup) *UpkeepService {
	return &UpkeepService{
		ctx:                   ctx,
		cfg:                   cfg,
		client:                client,
		checkDataService:      checkDataService,
		accountService:        accountService,
		upkeepContract:        upkeepContract,
		upkeepContractAddress: upkeepContractAddress,
		adminAddress:          adminAddress,
		logicContractInfo:     logicContractInfo,
		newUpkeepCh:           newUpkeepCh,
		topUpCh:               topUpCh,
		moduleName:            "service.chainlink.v2_0.upkeepService",
	}
}

func (s *UpkeepService) StartBalanceWatcher(headerCh chan *types.Header) {
	for {
		select {
		case <-s.ctx.Done():
			return
		case h, ok := <-headerCh:
			if !ok {
				return
			}
			log.Debug().Fields(map[string]any{"height": h.Number}).Msg("executing balance watcher")
			upkeeps, err := s.GetUpkeeps()
			if err != nil {
				log.Err(err).Str("module", s.moduleName).
					Str("func", "StartBalanceWatcher").
					Str("upkeep_contract", s.upkeepContractAddress.String()).
					Msg("failed to get upkeeps")
				continue
			}
			if len(upkeeps) == 0 {
				// todo check paused state
				_ = s.CreateInitialUpkeep()
			}
			for _, upkeep := range upkeeps {
				log.Debug().Str("module", s.moduleName).
					Str("func", "StartBalanceWatcher").
					Fields(map[string]any{
						"upkeep_id": upkeep.Id.String(),
						"balance":   upkeep.Info.Balance.String(),
					}).Msg("balance check")
				if upkeep.Info.Balance.Cmp(upkeep.MinAmount) <= 0 {
					log.Warn().Str("module", s.moduleName).
						Str("func", "StartBalanceWatcher").
						Fields(map[string]any{
							"upkeep_id":   upkeep.Id.String(),
							"balance":     upkeep.Info.Balance.String(),
							"min_balance": upkeep.MinAmount.String(),
						}).Msg("upkeep balance is too low. Trying to top up")
					topup := model.UpkeepTopup{
						UpkeepId: *upkeep.Id,
						Amount:   *s.cfg.Chainlink.TopupAmount,
					}
					s.topUpCh <- topup
				}
			}
		}
	}
}

func (s *UpkeepService) StartPayer() {
	for {
		select {
		case <-s.ctx.Done():
			break
		case topup := <-s.topUpCh:
			(func() {
				ctx, cancel := context.WithDeadline(s.ctx, time.Now().Add(time.Second*30))
				defer cancel()
				opts, err := s.accountService.NewTxOpts(ctx)
				if err != nil {
					return
				}
				r, err := s.upkeepContract.AddFunds(opts, &topup.UpkeepId, &topup.Amount)
				if err != nil {
					log.Err(err).Str("module", s.moduleName).
						Str("func", "StartPayer").
						Fields(map[string]any{
							"upkeep_id": topup.UpkeepId.String(),
							"amount":    topup.Amount.String(),
						}).Msg("Failed to top up upkeep")
					return
				}
				log.Info().Str("module", s.moduleName).
					Str("func", "StartPayer").
					Fields(map[string]any{
						"upkeep_id": topup.UpkeepId.String(),
						"amount":    topup.Amount.String(),
						"txid":      r.Hash(),
					}).Msg("Funds were added to upkeep.")
			})()

		}
	}
}

func (s *UpkeepService) StartScaler(headerCh chan *types.Header) {
	for {
		select {
		case <-s.ctx.Done():
			return
		case h, ok := <-headerCh:
			if !ok {
				return
			}
			log.Debug().Fields(map[string]any{"height": h.Number}).Msg("executing scaler")
			r, err := s.upkeepContract.IsNewUpkeepNeeded(&bind.CallOpts{})
			if err != nil {
				log.Err(err).Str("module", s.moduleName).
					Str("func", "StartScaler").
					Fields(map[string]any{
						"contract_name":               s.logicContractInfo.Name,
						"logic_contract_address":      s.logicContractInfo.Address.String(),
						"controller_contract_address": s.upkeepContractAddress,
					}).Msg("failed to check IsNewUpkeepNeeded")
				continue
			}
			if !r.IsNeeded {
				continue
			}

			newUpkeep, err := s.NewUpkeep(s.logicContractInfo, s.adminAddress, r.NewLimit.Uint64(), r.NewOffset.Uint64())
			if err != nil {
				log.Err(err).Str("module", s.moduleName).
					Str("func", "StartScaler").
					Fields(map[string]any{
						"contract_name":               s.logicContractInfo.Name,
						"logic_contract_address":      s.logicContractInfo.Address.String(),
						"controller_contract_address": s.upkeepContractAddress,
						"new_limit":                   r.NewLimit.String(),
						"new_offset":                  r.NewOffset.String(),
					}).Msg("failed to create new scaled upkeep")
				continue
			}
			s.newUpkeepCh <- newUpkeep
		}
	}
}

func (s *UpkeepService) CreateInitialUpkeep() error {
	initialUpkeep, err := s.NewUpkeep(s.logicContractInfo, s.adminAddress, s.logicContractInfo.ExecutionLimit, 0)
	if err != nil {
		log.Err(err).Str("module", s.moduleName).
			Str("func", "CreateInitialUpkeep").
			Fields(map[string]any{
				"contract_name":               s.logicContractInfo.Name,
				"logic_contract_address":      s.logicContractInfo.Address.String(),
				"controller_contract_address": s.upkeepContractAddress,
			}).Msg("failed to create initial upkeep")
		return err
	}
	s.newUpkeepCh <- initialUpkeep
	log.Debug().Str("module", s.moduleName).
		Str("func", "CreateInitialUpkeep").
		Msg("trying to create initial upkeep")
	return nil
}

func (s *UpkeepService) StartCreator() {
	for {
		select {
		case <-s.ctx.Done():
			return
		case newUpkeep := <-s.newUpkeepCh:
			(func() {
				ctx, cancel := context.WithDeadline(s.ctx, time.Now().Add(time.Second*30))
				defer cancel()

				opts, err := s.accountService.NewTxOpts(ctx)
				if err != nil {
					return
				}

				tx, err := s.upkeepContract.RegisterAndPredictID(opts, v2_0.RegistrationParams{
					Name:           newUpkeep.Name,
					EncryptedEmail: newUpkeep.EncryptedEmail,
					UpkeepContract: newUpkeep.UpkeepContract.Address,
					GasLimit:       newUpkeep.GasLimit,
					AdminAddress:   newUpkeep.AdminAddress,
					CheckData:      newUpkeep.CheckData,
					OffchainConfig: newUpkeep.OffchainConfig,
					Amount:         newUpkeep.Amount,
				})

				if err != nil {
					log.Err(err).Str("module", s.moduleName).
						Str("func", "StartCreator").
						Fields(map[string]any{
							"contract_name":    newUpkeep.UpkeepContract.ShortName,
							"contract_address": newUpkeep.UpkeepContract.Address,
						}).Msg("failed to register upkeep")
					return
				}

				log.Info().Str("module", s.moduleName).
					Str("func", "StartCreator").
					Fields(map[string]any{
						"txid":             tx.Hash().String(),
						"contract_name":    newUpkeep.UpkeepContract.ShortName,
						"contract_address": newUpkeep.UpkeepContract.Address,
					}).Msg("upkeep was successfully registered")
			})()
		}
	}
}

func (s *UpkeepService) GetUpkeeps() ([]v2_0.UpkeepControllerExampleDetailedUpkeep, error) {
	return s.upkeepContract.GetDetailedUpkeeps(&bind.CallOpts{}, new(big.Int).SetUint64(0), new(big.Int).SetUint64(1000))
}

func (s *UpkeepService) NewUpkeep(contract model.ContractInfo, adminAddress common.Address, limit uint64, offset uint64) (model.Upkeep, error) {
	//todo contract automation compatibility prevalidation
	if limit == 0 {
		return model.Upkeep{}, errors.New("limit must be greater than 0")
	}

	checkData, err := s.checkDataService.NewCheckData(limit, offset)
	if err != nil {
		return model.Upkeep{}, err
	}
	return model.Upkeep{
		Name:           fmt.Sprintf("%s v%s L:%d O:%d", contract.ShortName, contract.Version, limit, offset),
		EncryptedEmail: []byte{0},
		UpkeepContract: contract,
		GasLimit:       contract.GasLimit,
		AdminAddress:   adminAddress,
		CheckData:      checkData,
		OffchainConfig: []byte{0},
		Amount:         s.cfg.Chainlink.TopupAmount,
		Sender:         adminAddress,
	}, nil
}
