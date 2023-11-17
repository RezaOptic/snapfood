package config

import (
	config "snapfood/utils/coreconfig"
)

// RedisConfig struct
type RedisConfig struct {
	config.TomlInterface
}

// Redis variable
var Redis RedisConfig

// Load load config
func (s RedisConfig) Load(env string) {
	config.LoadConfig("redis", &Redis)
}
