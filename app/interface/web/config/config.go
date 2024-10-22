package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	HttpConf *HttpConfig
	LogConf  *LogConfig
	EtcdConf *EtcdConfig
}

type HttpConfig struct {
	Port int
}

type LogConfig struct {
	Name          string
	LogDir        string
	Level         int32
	OpenAccessLog bool
}

type EtcdConfig struct {
	Addrs   []string
	Timeout int
}

var _configInstance *Config

func ConfigInstance() *Config {
	return _configInstance
}

func Init() {
	if err := initConfig(); err != nil {
		panic(err)
	}

	_configInstance = &Config{
		HttpConf: &HttpConfig{
			Port: viper.GetInt("http.port"),
		},
		LogConf: &LogConfig{
			Name:          viper.GetString("log.name"),
			LogDir:        viper.GetString("log.logDir"),
			Level:         viper.GetInt32("log.level"),
			OpenAccessLog: viper.GetBool("log.openAccessLog"),
		},
		EtcdConf: &EtcdConfig{
			Addrs:   viper.GetStringSlice("etcd.addrs"),
			Timeout: viper.GetInt("etcd.timeout"),
		},
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
