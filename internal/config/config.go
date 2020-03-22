package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type TelegramBot struct {
	Token string `yaml:"token"`
}

type Instagram struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Database struct {
	URI string `yaml:"uri"`
}

type Config struct {
	TelegramBot TelegramBot `yaml:"telegram_bot"`
	Instagram   Instagram   `yaml:"instagram"`
	Database    Database    `yaml:"database"`

	Dialogue string `yaml:"dialogue"`
	Vault    string `yaml:"vault"`
	Logging  string `yaml:"logging"`
}

func NewConfig(configPath string) (*Config, error) {
	cfgFile, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	cfgFileData, err := ioutil.ReadAll(cfgFile)
	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	if err = yaml.Unmarshal(cfgFileData, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
