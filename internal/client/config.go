package client

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	DefaultServerAddress = "127.0.0.1:50051"
	DefaultTimeout       = 5 * time.Second
)

type Config struct {
	ServerAddress string
	Timeout       time.Duration
}

type Connection interface {
	grpc.ClientConnInterface
	io.Closer
}

type Dialer func(context.Context, Config) (Connection, Config, error)

func (config Config) WithDefaults() (Config, error) {
	if strings.TrimSpace(config.ServerAddress) == "" {
		config.ServerAddress = DefaultServerAddress
	}

	switch {
	case config.Timeout < 0:
		return Config{}, fmt.Errorf("timeout must be positive")
	case config.Timeout == 0:
		config.Timeout = DefaultTimeout
	}

	return config, nil
}

func Dial(ctx context.Context, config Config) (Connection, Config, error) {
	normalized, err := config.WithDefaults()
	if err != nil {
		return nil, Config{}, err
	}

	dialContext, cancel := context.WithTimeout(ctx, normalized.Timeout)
	defer cancel()

	connection, err := grpc.DialContext(
		dialContext,
		normalized.ServerAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, Config{}, err
	}

	return connection, normalized, nil
}
