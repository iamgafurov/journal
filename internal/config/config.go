package config

import (
	jsoniter "github.com/json-iterator/go"
	"os"
)

type Config struct {
	ServerPrefix          string `json:"restPrefix"`
	ServerPort            string `json:"serverPort"`
	PostgresConnStr       string `json:"postgresConnStr"`
	MSSQLConnStr          string `json:"mssqlConnStr"`
	MasterKey             string `json:"masterKey"`
	SentryDSN             string `json:"sentryDSN"`
	TokensDurationInHours int    `json:"tokensDurationInHours"`
	Debug                 bool   `json:"debug"`
}

func New() (cfg *Config, err error) {
	bt, err := os.ReadFile("./config.json")
	if err != nil {
		return
	}

	err = jsoniter.Unmarshal(bt, &cfg)
	if err != nil {
		return
	}

	if cfg.TokensDurationInHours == 0 {
		cfg.TokensDurationInHours = 24 * 30
	}

	return cfg, nil
}
