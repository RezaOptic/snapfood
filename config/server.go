package config

import (
	config "snapfood/utils/coreconfig"
)

// ServerConfig struct
type ServerConfig struct {
	config.TomlInterface
	ServiceName string `mapstructure:"service_name"`
	GRPCPort    string `mapstructure:"grpc_port"`
	HTTPPort    string `mapstructure:"http_port"`
	TimeZone    string `mapstructure:"time_zone"`
}

// Server variable
var Server ServerConfig

// Load method
func (u ServerConfig) Load(env string) {
	config.LoadConfig("server", &Server)
}
