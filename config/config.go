package config

import (
	"fmt"
	"simple_brave_signer/logger"

	"github.com/spf13/viper"
)

func LoadYamlConfig() (*viper.Viper, error) {
	localViper := viper.New()
	localViper.SetConfigName("config")
	localViper.SetConfigType("yaml")
	localViper.AddConfigPath(".")

	err := localViper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Info("Config file not found, using CLI params or default settings...")
		} else {
			return localViper, fmt.Errorf("found config file, but encountered an error: %v", err)
		}
	}
	return localViper, nil
}
