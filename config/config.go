package config

import (
	"flag"

	"fmt"

	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/pkg/constants"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/pkg/eventstroredb"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/pkg/logger"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/pkg/mongodb"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/pkg/probes"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/pkg/tracing"
	"github.com/pkg/errors"

	"github.com/spf13/viper"

	"os"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "ES microservice config path")
}

type Config struct {
	ServiceName      string                         `mapstructure:"serviceName"`
	Logger           *logger.Config                 `mapstructure:"logger"`
	GRPC             GRPC                           `mapstructure:"grpc"`
	Mongo            *mongodb.Config                `mapstructure:"mongo"`
	MongoCollections MongoCollections               `mapstructure:"mongoCollections"`
	Probes           probes.Config                  `mapstructure:"probes"`
	Jaeger           *tracing.Config                `mapstructure:"jaeger"`
	EventStoreConfig eventstroredb.EventStoreConfig `mapstructure:"eventStoreConfig"`
}

type GRPC struct {
	Port string `mapstructure:"port"`

	Development bool `mapstructure:"development"`
}

type MongoCollections struct {
	Products string `mapstructure:"products"`
}

func InitConfig() (*Config, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv(constants.ConfigPath)
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			getwd, err := os.Getwd()
			if err != nil {
				return nil, errors.Wrap(err, "os.Getwd")
			}
			configPath = fmt.Sprintf("%s/config/config.yaml", getwd)
		}
	}

	cfg := &Config{}

	viper.SetConfigType(constants.Yaml)

	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	grpcPort := os.Getenv(constants.GrpcPort)

	if grpcPort != "" {
		cfg.GRPC.Port = grpcPort
	}

	mongoURI := os.Getenv(constants.MongoDbURI)

	if mongoURI != "" {
		//cfg.Mongo.URI = "mongodb://host.docker.internal:27017"
		cfg.Mongo.URI = mongoURI

	}

	jaegerAddr := os.Getenv(constants.JaegerHostPort)

	if jaegerAddr != "" {
		cfg.Jaeger.HostPort = jaegerAddr
	}

	return cfg, nil
}
