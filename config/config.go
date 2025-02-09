package config

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
	echoserver "github.com/vmdt/notification-system/pkg/http/echo/server"
	"github.com/vmdt/notification-system/pkg/logger"
	"github.com/vmdt/notification-system/pkg/postgres"
	"github.com/vmdt/notification-system/pkg/rabbitmq"
)

var configPath string

type Config struct {
	Logger *logger.LoggerConfig `mapstructure:"logger"`
	RabbitMQ *rabbitmq.RabbitMQConfig `mapstructure:"rabbitmq"`
	Echo *echoserver.EchoConfig `mapstructure:"echo"`
	GormPostgres *postgres.PostgresConfig `mapstructure:"postgres"`
}

func InitConfig() (*Config, *logger.LoggerConfig, *rabbitmq.RabbitMQConfig, *postgres.PostgresConfig, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv("CONFIG_PATH")
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			//https://stackoverflow.com/questions/31873396/is-it-possible-to-get-the-current-root-of-package-structure-as-a-string-in-golan
			//https://stackoverflow.com/questions/18537257/how-to-get-the-directory-of-the-currently-running-file
			d, err := dirname()
			if err != nil {
				return nil, nil, nil, nil, err
			}

			configPath = d
		}
	}

	cfg := &Config{}

	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, nil, nil, nil, err
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, nil, nil, nil, err
	}

	return cfg, cfg.Logger, cfg.RabbitMQ, cfg.GormPostgres, nil
}

func filename() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("unable to get the current filename")
	}
	return filename, nil
}

func dirname() (string, error) {
	filename, err := filename()
	if err != nil {
		return "", err
	}
	return filepath.Dir(filename), nil
}