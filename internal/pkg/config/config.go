package config

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
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

	// 设置当配置文件改变时的回调函数
	v.OnConfigChange(func(e fsnotify.Event) {
		slog.Info("Config file changed:", e.Name)
		if err := v.ReadInConfig(); err != nil {
			slog.Error("Error reading config file:", err)
		}
		if err := v.UnmarshalKey("config", conf); err != nil {
			slog.Error("Error unmarshalling config:", err)
		}
	})

	v.WatchConfig()

	return conf
}
