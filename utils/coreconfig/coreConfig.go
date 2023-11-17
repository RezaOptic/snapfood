package coreconfig

import (
	// Go Native Packages

	"snapfood/utils/logger"
	"time"

	"github.com/spf13/viper"
)

const (
	//EnvDevel const
	EnvDevel = "devel"
	//EnvProd const
	EnvProd = "prod"
	//EnvTest const
	EnvTest = "test"
)

var (
	//Env var
	Env string
	//BaseConfigPath var
	BaseConfigPath string
	//TopoAddr var
	TopoAddr string
)

type configBase struct{}

// TomlInterface interface
type TomlInterface interface {
	Load(env string)
}

// Loader config
func Loader(env, configpath string, structs ...TomlInterface) {

	// logger
	defer logger.ZSLogger.Sync()

	// set variables
	Env = env
	BaseConfigPath = configpath

	// check structs input
	if len(structs) == 0 {
		logger.ZSLogger.Panicw("config struct is empty")
	}

	// call interface method
	for _, s := range structs {
		s.Load(Env)
	}
}

// LoadConfig load config file
func LoadConfig(configPath string, configStruct interface{}) {

	// logger
	defer logger.ZSLogger.Sync()

	viperInstance := viper.New()
	fullConfigPath := ""

	switch {
	case Env == EnvTest:
		fallthrough
	case Env == EnvDevel:
		fullConfigPath = BaseConfigPath + "/" + configPath + ".toml"
		viperInstance.SetConfigFile(fullConfigPath)
		err := viperInstance.ReadInConfig()
		if err != nil {
			logger.ZSLogger.Panicw("failed to read from file", "config_path", fullConfigPath, "error", err)
		}
	case Env == EnvProd:
		BaseConfigPath = "/config"
		fullConfigPath = BaseConfigPath + "/" + configPath + ".toml"
		viperInstance.SetConfigFile(fullConfigPath)
		err := viperInstance.ReadInConfig()
		if err != nil {
			logger.ZSLogger.Panicw("failed to read from file", "config_path", fullConfigPath, "error", err)
		}

		go periodicRead(viperInstance, configPath, configStruct)
	default:
		logger.ZSLogger.Panicw("unsupported environment")
	}

	err := viperInstance.Unmarshal(configStruct)
	if err != nil {
		logger.ZSLogger.Panicw("failed to unmarshal", "config_path", fullConfigPath, "error", err)
	}

	logger.ZSLogger.Infow("config loaded successfully", "config_path", fullConfigPath)
}

// LoadTranslation load translation file
func LoadTranslation(configPath string, configStruct interface{}) {

	// logger
	defer logger.ZSLogger.Sync()

	viperInstance := viper.New()
	fullConfigPath := ""

	fullConfigPath = configPath + ".toml"
	viperInstance.SetConfigFile(fullConfigPath)
	err := viperInstance.ReadInConfig()
	if err != nil {
		logger.ZSLogger.Panicw("failed to read from file", "config_path", fullConfigPath, "error", err)
	}

	err = viperInstance.Unmarshal(configStruct)
	if err != nil {
		logger.ZSLogger.Panicw("failed to unmarshal", "config_path", fullConfigPath, "error", err)
	}

	logger.ZSLogger.Infow("config loaded successfully", "config_path", fullConfigPath)
}

func periodicRead(viperInstance *viper.Viper, configName string, config interface{}) {

	// logger
	defer logger.ZSLogger.Sync()

	for {
		time.Sleep(5 * time.Second)

		err := viperInstance.ReadInConfig()
		for err != nil {
			logger.ZSLogger.Errorw("failed to read from config map after tries", "error", err)
			continue
		}

		err = viperInstance.Unmarshal(config)
		if err != nil {
			logger.ZSLogger.Errorw("error while unmarshaling viper config", "error", err, "config_name", configName)
		}
	}
}
