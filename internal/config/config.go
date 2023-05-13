package config

import (
	"github.com/cexles/automation-lb/internal/model"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"math/big"
	"os"
)

const DefaultConfigFile = "config.yaml"

type Config struct {
	Account   AccountConfig   `yaml:"account"`
	Network   NetworkConfig   `yaml:"network"`
	App       AppConfig       `yaml:"app"`
	Contracts ContractsConfig `yaml:"contracts"`
	Chainlink ChainlinkConfig `yaml:"chainlink"`
}

type AccountConfig struct {
	PrivateKey string `yaml:"private_key"`
}

type NetworkConfig struct {
	RpcUrl    string   `yaml:"rpc_url"`
	IsTestnet bool     `yaml:"is_testnet"`
	ChainId   *big.Int `yaml:"chain_id"`
}

type AppConfig struct {
	LogLevel int8 `yaml:"log_level"`
}

type ChainlinkConfig struct {
	UpkeepControllerAddress common.Address `yaml:"upkeep_controller_address"`
	MinControllerApprove    *big.Int       `yaml:"min_controller_approve"`
	TopupAmount             *big.Int       `yaml:"topup_amount"`
	TokenAddress            common.Address `yaml:"token_address"`
}

type ContractsConfig map[string]model.ContractInfo

func NewConfig(filename string) (*Config, error) {
	if len(filename) == 0 {
		filename = DefaultConfigFile
	}
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Panic().Err(err).Str("filename", filename).Msg("failed to read config file")
		return nil, err
	}

	cfg := &Config{}

	err = yaml.Unmarshal(f, cfg)
	if err != nil {
		log.Panic().Err(err).Str("filename", filename).Msg("failed to unmarshal yaml")
		return nil, err
	}

	log.Info().Fields(map[string]any{
		"address": cfg.Chainlink.UpkeepControllerAddress,
	}).Msg("upkeep controller found")

	return cfg, err
}
