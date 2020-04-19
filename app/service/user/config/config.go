package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	MysqlConf *MysqlConfig
	JwtConf   *JwtConfig
	TraceConf *TraceConfig
}

type MysqlConfig struct {
	Addr         string
	Db           string
	User         string
	Password     string
	MaxIdleConns int
	MaxOpenConns int
}

type JwtConfig struct {
	JwtKey  string
	JwtTime int
}

type TraceConfig struct {
	Addr         string
	TraceContext string
	ReqParam     string
}

func NewConfig() *Config {
	return &Config{
		MysqlConf: &MysqlConfig{
			Addr:         viper.GetString("mysql.addr"),
			Db:           viper.GetString("mysql.db"),
			User:         viper.GetString("mysql.user"),
			Password:     viper.GetString("mysql.password"),
			MaxIdleConns: viper.GetInt("mysql.maxIdleConns"),
			MaxOpenConns: viper.GetInt("mysql.maxOpenConns"),
		},
		JwtConf: &JwtConfig{
			JwtKey:  viper.GetString("jwt.key"),
			JwtTime: viper.GetInt("jwt.time"),
		},
		TraceConf: &TraceConfig{
			Addr:         viper.GetString("tracer.addr"),
			TraceContext: viper.GetString("tracer.context"),
			ReqParam:     viper.GetString("tracer.param"),
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
	viper.SetEnvPrefix("user-srv")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
