package ping

import (
	"context"
	"fmt"
	"strings"
	"time"

	pingv1 "github.com/guycole/daring-cyclops/gen/proto/ping/v1"
	clientcfg "github.com/guycole/daring-cyclops/internal/client"
)

func Run(ctx context.Context, config clientcfg.Config) (Result, error) {
	return RunWithDialer(ctx, config, clientcfg.Dial)
}

func RunWithDialer(ctx context.Context, config clientcfg.Config, dialer clientcfg.Dialer) (Result, error) {
	connection, normalized, err := dialer(ctx, config)
	if err != nil {
		return Result{}, fmt.Errorf("ping failed: %w", err)
	}
	defer connection.Close()

	rpcContext, cancel := context.WithTimeout(ctx, normalized.Timeout)
	defer cancel()

	response, err := pingv1.NewPingServiceClient(connection).Ping(rpcContext, &pingv1.PingRequest{})
	if err != nil {
		return Result{}, fmt.Errorf("ping failed: %w", err)
	}

	result, err := mapResponse(response)
	if err != nil {
		return Result{}, fmt.Errorf("ping failed: %w", err)
	}

	return result, nil
}

func mapResponse(response *pingv1.PingResponse) (Result, error) {
	if response == nil {
		return Result{}, fmt.Errorf("response missing payload")
	}

	serverVersion := strings.TrimSpace(response.GetServerVersion())
	if serverVersion == "" {
		return Result{}, fmt.Errorf("response missing server version")
	}

	serverTime := response.GetServerTime()
	if serverTime == nil {
		return Result{}, fmt.Errorf("response missing server time")
	}

	if err := serverTime.CheckValid(); err != nil {
		return Result{}, fmt.Errorf("response has invalid server time: %w", err)
	}

	return Result{
		ServerVersion: serverVersion,
		ServerTime:    serverTime.AsTime().UTC(),
	}, nil
}

func NewResult(serverVersion string, serverTime time.Time) Result {
	return Result{ServerVersion: serverVersion, ServerTime: serverTime}
}
