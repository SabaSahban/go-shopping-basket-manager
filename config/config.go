package config

import (
	"time"

	"github.com/spf13/viper"

	"github.com/sirupsen/logrus"
)

type (
	Config struct {
		Postgres Postgres `mapstructure:"postgres"`
		Server   Server   `mapstructure:"server"`
	}

	Postgres struct {
		Host               string        `mapstructure:"host"`
		Port               int           `mapstructure:"port"`
		Username           string        `mapstructure:"username"`
		Password           string        `mapstructure:"password"`
		DBName             string        `mapstructure:"db-name"`
		ConnectTimeout     time.Duration `mapstructure:"connect-timeout"`
		ConnectionLifetime time.Duration `mapstructure:"connection-lifetime"`
		MaxOpenConnections int           `mapstructure:"max-open-connections"`
		MaxIdleConnections int           `mapstructure:"max-idle-connections"`
		Debug              bool          `mapstructure:"debug"`
	}

	Server struct {
		Port int `mapstructure:"port"`
	}
)

func InitConfig() Config {
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatalf("failed to read config")
	}

	var conf Config
	err = viper.Unmarshal(&conf)
	if err != nil {
		logrus.Fatalf("failed to unmarshal config")
	}

	return conf
}
