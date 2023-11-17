package config

import (
	"flag"
	"path/filepath"
	"runtime"
	"snapfood/constants"
	config "snapfood/utils/coreconfig"
)

var (
	//EnvInp var
	EnvInp string
	//BaseConfigPathInp var
	BaseConfigPathInp string
	//Loaded var
	Loaded bool
)

// Init func
func Init(Config *string) {
	_, b, _, _ := runtime.Caller(0)
	BasePath := filepath.Dir(b)
	if Config != nil && *Config == constants.TestFlag {
		configPath := BasePath + "/tests"
		flag.StringVar(&EnvInp, "env", "devel", "Environment to run the application, could be devel, prod, or tests")
		flag.StringVar(&BaseConfigPathInp, "config-path", configPath, "Path to config directory")
	} else if !Loaded {
		// Register flags
		configPath := BasePath + "/files"
		if Config != nil && *Config != "" {
			configPath = *Config
		}
		flag.StringVar(&EnvInp, "env", "devel", "Environment to run the application, could be devel, prod, or tests")
		flag.StringVar(&BaseConfigPathInp, "config-path", configPath, "Path to config directory")

	}
	flag.Parse()
	config.Loader(EnvInp, BaseConfigPathInp, PsqlDB, Redis, Server, Secrets)
}
