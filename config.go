package main

import (
	"github.com/spf13/viper"
)

type Server struct {
	Host string
	Port uint16
}

type Config struct {
	Port    uint16
	Servers []Server
}

func readConfig() (*Config, error) {
	viper.SetConfigName("blncr")
	viper.SetConfigType("toml")

	// TODO: add config path for consuming from final bin file
	// Some alternatives could be `$HOME/.blncr`
	// or `/etc/blncr`
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := Config{}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
