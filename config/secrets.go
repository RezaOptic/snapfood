package config

import (
	config "snapfood/utils/coreconfig"
)

// SecretsConfig struct
type SecretsConfig struct {
	config.TomlInterface
}

// Secrets variable
var Secrets SecretsConfig

// Load method
func (s SecretsConfig) Load(env string) {
	config.LoadConfig("secrets", &Secrets)
}
