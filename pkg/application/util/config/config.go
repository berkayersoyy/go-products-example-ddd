package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	MysqlDSN      string        `mapstructure:"MYSQL_DSN"`
	RedisHost     string        `mapstructure:"REDIS_HOST"`
	AccessSecret  string        `mapstructure:"ACCESS_SECRET"`
	RefreshSecret string        `mapstructure:"REFRESH_SECRET"`
	Timeout       time.Duration `mapstructure:"TIMEOUT"`
	EndpointUrl   string        `mapstructure:"EndpointUrl"`
	Region        string        `mapstructure:"Region"`
	ID            string        `mapstructure:"ID"`
	SecretKey     string        `mapstructure:"SecretKey"`
	AccessKey     string        `mapstructure:"AccessKey"`
	Profile       string        `mapstructure:"Profile"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
