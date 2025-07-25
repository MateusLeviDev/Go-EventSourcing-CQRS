package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeout  = 30 * time.Second
	maxConnIdleTime = 3 * time.Minute
	minPoolSize     = 20
	maxPoolSize     = 300
)

type Config struct {
	URI      string `mapstructure:"uri"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Db       string `mapstructure:"db"`
}

func NewMongoDBConn(ctx context.Context, cfg *Config) (*mongo.Client, error) {

	clientOpts := options.Client().
		ApplyURI(cfg.URI).
		SetAuth(options.Credential{
			Username: cfg.User,
			Password: cfg.Password,
		}).
		SetConnectTimeout(connectTimeout).
		SetMaxConnIdleTime(maxConnIdleTime).
		SetMinPoolSize(minPoolSize).
		SetMaxPoolSize(maxPoolSize)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}
