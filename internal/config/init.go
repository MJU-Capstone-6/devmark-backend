package config

import (
	"errors"
	"log"
	"sync"

	"github.com/spf13/viper"
)

var (
	config *Config
	once   sync.Once
)

func InitConfig() (*Config, error) {
	if config == nil {
		once.Do(func() {
			err := readConfig()
			if err != nil {
				log.Fatal("Something went wrong while read config")
			}
		})
	} else {
		return nil, errors.New("Config already configured")
	}
	return config, nil
}

func readConfig() error {
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return err
	}
	return nil
}
