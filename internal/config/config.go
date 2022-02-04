// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config stores application configuration
type Config struct {
	Server     ServerConfig
	Roc        RocConfig
	Prometheus PrometheusConfig
}

// ServerConfig -
type ServerConfig struct {
	Port int    `mapstructure:"PORT"`
	Env  string `mapstructure:"ENV"`
}

// RocConfig -
type RocConfig struct {
	Base   string `mapstructure:"BASE"`
	Target string `mapstructure:"TARGET"`
	Token  string `mapstructure:"TOKEN"`
}

// PrometheusConfig -
type PrometheusConfig struct {
	Acc string `mapstructure:"ACC"`
}

// New initializes a new Configuration from the ENV variables
func New() *Config {
	viper.SetConfigName("config")

	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error while reading config file: %v", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Configuration file changed")
	})

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode config file to struct, err: %v", err)
	}

	return &config
}
