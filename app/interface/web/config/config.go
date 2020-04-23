package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	HttpConf *HttpConfig
}

type HttpConfig struct {
	Port int
}

func NewConfig() *Config {
	return &Config{
		HttpConf: &HttpConfig{
			Port: viper.GetInt("http.port"),
		},
	}
}

func init() {
	if err := initConfig(); err != nil {
		panic(err)
	}
}

func initConfig() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("web-web")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
