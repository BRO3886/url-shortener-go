package config

import (
	"os"

	"github.com/spf13/viper"
)

var Config *Configuration

type Configuration struct {
	Redis struct {
		Addr     string `mapstructure:"addr"`
		UserName string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"redis"`

	Mongo struct {
		Addr     string `mapstructure:"addr"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Database string `mapstructure:"database"`
		Timeout  int    `mapstructure:"timeout"`
	} `mapstructure:"mongo"`
}

// reads the given config file and initalises the config
func NewConfig(configPath string) error {
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	parser := viper.New()
	parser.SetConfigType("yaml")
	parser.SetConfigFile(configPath)
	parser.AutomaticEnv()
	if err := parser.ReadInConfig(); err != nil {
		return err
	}
	if err := parser.Unmarshal(&Config); err != nil {
		return err
	}

	return nil
}
