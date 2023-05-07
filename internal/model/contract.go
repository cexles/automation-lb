package model

type Contract struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	Version string `yaml:"version"`
}
