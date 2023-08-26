package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func NewConfig() *Config {
	path, err := filepath.Abs("config/config.yaml")
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic(err)
	}

	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	conf := &Config{}
	if err := v.UnmarshalKey("config", conf); err != nil {
		panic(err)
	}

	return conf
}
