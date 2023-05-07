package config

import (
	"github.com/cexles/automation-lb/internal/model"
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
	RegistryVersion         string          `yaml:"registry_version"`
	RegistryContracts       ContractsConfig `yaml:"registry_contracts"`
	CurrentRegistryContract *model.Contract
}

type ContractsConfig map[string]model.Contract

func NewConfig(filename string) (*Config, error) {
	if len(filename) == 0 {
		filename = DefaultConfigFile
	}
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Panic().Str("filename", filename).Msg(err.Error())
		return nil, err
	}

	cfg := &Config{}

	err = yaml.Unmarshal(f, cfg)
	if err != nil {
		log.Panic().Str("filename", filename).Msg(err.Error())
		return nil, err
	}

	for version, contract := range cfg.Chainlink.RegistryContracts {
		if version != cfg.Chainlink.RegistryVersion {
			continue
		}

		cfg.Chainlink.CurrentRegistryContract = &contract
		log.Info().Fields(map[string]any{
			"address": contract.Address,
			"version": contract.Version,
		}).Msg("Registry contract config found")
		break
	}

	if cfg.Chainlink.CurrentRegistryContract == nil {
		log.Panic().Str("registry_version", cfg.Chainlink.RegistryVersion).Msg("Registry Contract version unsupported.")
	}

	return cfg, err
}
