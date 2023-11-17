package config

import (
	config "snapfood/utils/coreconfig"
)

// PsqlDBConfig struct
type PsqlDBConfig struct {
	config.TomlInterface
	// mongodb config
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Sslmode  string `mapstructure:"sslmode"`
}

// PsqlDB variable
var PsqlDB PsqlDBConfig

// Load method
func (s PsqlDBConfig) Load(env string) {
	config.LoadConfig("psql", &PsqlDB)
}
