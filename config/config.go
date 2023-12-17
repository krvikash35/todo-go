package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	AppHttpPort string
	AppGrpcPort string
	Database    Database
}

var Value Config

func Init() error {
	viper.AutomaticEnv()
	viper.SetConfigName("application")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../..")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../../")
	viper.SetConfigType("yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("config.Init: could not read config due to err - %w", err)
	}

	if err := viper.Unmarshal(&Value); err != nil {
		return fmt.Errorf("config.Init: could not unmarshal config due to err - %w", err)
	}

	return nil
}
