package main

import (
	"context"
	"errors"
	"flag"
	"github.com/cexles/automation-lb/internal/config"
	"github.com/cexles/automation-lb/internal/model"
	"github.com/cexles/automation-lb/internal/service"
	chainlink2 "github.com/cexles/automation-lb/internal/service/chainlink"
	"github.com/cexles/automation-lb/internal/service/chainlink/v2_0"
	"github.com/cexles/automation-lb/pkg/contract/chainlink"
	v20contract "github.com/cexles/automation-lb/pkg/contract/chainlink/v2_0"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: zerolog.TimeFieldFormat})
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	cfgFilename := flag.String("config_file", config.DefaultConfigFile, "Config File")
	flag.Parse()

	cfg, err := config.NewConfig(*cfgFilename)
	if err != nil {
		log.Panic().Err(err).Msg("failed to initialize config")
	}

	zerolog.SetGlobalLevel(zerolog.Level(cfg.App.LogLevel))

	client, err := ethclient.DialContext(ctx, cfg.Network.RpcUrl)
	if err != nil {
		log.Panic().Str("rpc_url", cfg.Network.RpcUrl).Err(err).Msg("failed to connect to RPC")
	}
	defer client.Close()

	chainId, err := client.ChainID(ctx)
	if err != nil {
		log.Panic().Str("rpc_url", cfg.Network.RpcUrl).Err(err).Msg("failed to retrieve chain id")
	}

	if chainId.Cmp(cfg.Network.ChainId) != 0 {
		log.Panic().Fields(map[string]any{
			"rpc_url":      cfg.Network.RpcUrl,
			"rpc_chain_id": chainId.String(),
			"cfg_chain_id": cfg.Network.ChainId.String(),
		}).Err(err).Msg("RPC chain id is not equal to config value")
	} else {
		log.Trace().Str("chain_id", chainId.String()).Msg("RPC chain id is equal to config value")
	}

	account, err := model.NewAccount(cfg.Account.PrivateKey, chainId)
	if err != nil {
		log.Panic().Err(err).Msg("failed to initialize payer account")
	}

	accountService := service.NewAccountService(account)

	upkeepControllerContract, err := v20contract.NewUpkeepControllerContract(cfg.Chainlink.UpkeepControllerAddress, client)
	if err != nil {
		log.Panic().Err(err).Msg("failed to initialize upkeep controller contract")
	}

	linkTokenContract, err := chainlink.NewLinkTokenContract(cfg.Chainlink.TokenAddress, client)
	if err != nil {
		log.Panic().Err(err).Msg("failed to initialize link token contract")
	}

	sampleContract, ok := cfg.Contracts[cfg.App.ActiveContractKey] // todo mutlicontract
	if !ok {
		log.Panic().Err(errors.New("contract not found, check config")).Msg("failed to get logic contract")
	}

	log.Info().Fields(map[string]any{
		"name":       sampleContract.Name,
		"short_name": sampleContract.ShortName,
		"version":    sampleContract.Version,
	}).Msg("logic contract found")

	cdService := v2_0.NewCheckDataService()

	newUpkeepCh := make(chan model.Upkeep, 1)
	topUpCh := make(chan model.UpkeepTopup, 1)

	linkTokenService := chainlink2.NewLinkTokenService(ctx, cfg, accountService, linkTokenContract)
	err = linkTokenService.IncreaseAllowanceIfNotEnough(ctx, cfg.Chainlink.UpkeepControllerAddress, cfg.Chainlink.MinControllerApprove)
	if err != nil {
		log.Panic().Err(err).
			Str("upkeep_contract", cfg.Chainlink.UpkeepControllerAddress.String()).
			Msg("failed to increase allowance to upkeep controller")
	}

	upkeepService := v2_0.NewUpkeepService(ctx, cfg, client, cdService, accountService, linkTokenService, upkeepControllerContract, cfg.Chainlink.UpkeepControllerAddress, account.Address(), sampleContract, newUpkeepCh, topUpCh)

	subscribedToNewHeadFuncs := []func(chan *types.Header){
		upkeepService.StartBalanceWatcher, upkeepService.StartScaler, upkeepService.StartAllowanceWatcher,
	}
	chainService := service.NewChainService(ctx, client, subscribedToNewHeadFuncs)
	err = chainService.StartBlockListener()
	if err != nil {
		log.Panic().Err(err).
			Msg("failed to start block event listener")
	}

	go chainService.StartBlockBroadcaster()

	for k, subFunc := range subscribedToNewHeadFuncs {
		go subFunc(chainService.SubChannels[k])
	}

	go upkeepService.StartCreator()
	go upkeepService.StartPayer()

	for {
		select {
		case <-ctx.Done():
			break
		}
	}
}
