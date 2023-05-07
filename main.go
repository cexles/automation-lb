package main

import (
	"context"
	"flag"
	"github.com/cexles/automation-lb/internal/config"
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
		log.Panic().Msg(err.Error())
	}

	zerolog.SetGlobalLevel(zerolog.Level(cfg.App.LogLevel))

	client, err := ethclient.DialContext(ctx, cfg.Network.RpcUrl)
	if err != nil {
		log.Panic().Str("rpc_url", cfg.Network.RpcUrl).Err(err).Msg("Failed to connect to RPC")
	}
	defer client.Close()

	chainId, err := client.ChainID(ctx)
	if err != nil {
		log.Panic().Str("rpc_url", cfg.Network.RpcUrl).Err(err).Msg("Failed to retrieve chain id")
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
}
