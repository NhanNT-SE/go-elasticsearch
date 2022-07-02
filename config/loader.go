package config

import (
	"bytes"
	_ "embed"

	"github.com/spf13/viper"

	"go-elasticsearch/pkg/logger"
)

//go:embed config.yaml
var defaultConfig []byte

func MustLoad(cfg interface{}) {
	log := logger.New()

	// read from default config first
	v := viper.NewWithOptions(viper.KeyDelimiter("__"))
	v.SetConfigType("yaml")
	if err := v.ReadConfig(bytes.NewBuffer(defaultConfig)); err != nil {
		log.Fatal().Err(err).Msg("failed to read default config")
	}

	// try read from .env file
	// will overwrite previous config
	// v.SetConfigType("env")
	// v.SetConfigFile(".env")
	// if err := v.MergeInConfig(); err != nil {
	// 	log.Warn().Err(err).Msg("failed to read .env config")
	// }

	// read from environment variables
	// will overwrite previous config
	v.AutomaticEnv()

	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatal().Err(err).Msg("failed to unmarshal config")
	}

	return
}
