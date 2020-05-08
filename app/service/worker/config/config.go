package config

import (
	"log"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

type Config struct {
	MysqlConf *MysqlConfig
	LogConf   *LogConfig
	EtcdConf  *EtcdConfig
}

type LogConfig struct {
	Name          string
	LogDir        string
	Level         int32
	OpenAccessLog bool
}

type MysqlConfig struct {
	Addr         string
	Db           string
	User         string
	Password     string
	MaxIdleConns int
	MaxOpenConns int
}

type EtcdConfig struct {
	Addrs   []string
	Timeout int
}

var (
	onceDo          sync.Once
	_configInstance *Config
)

func GetInstance() *Config {
	return _configInstance
}

func init() {
	if err := initConfig(); err != nil {
		panic(err)
	}

	onceDo.Do(func() {
		_configInstance = &Config{
			MysqlConf: &MysqlConfig{
				Addr:         viper.GetString("mysql.addr"),
				Db:           viper.GetString("mysql.db"),
				User:         viper.GetString("mysql.user"),
				Password:     viper.GetString("mysql.password"),
				MaxIdleConns: viper.GetInt("mysql.maxIdleConns"),
				MaxOpenConns: viper.GetInt("mysql.maxOpenConns"),
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
	})
}

func initConfig() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("worker-srv")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	watchConfig()

	return nil
}

func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s\n", e.Name)
	})
}
