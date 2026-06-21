package integration_test

import (
	"context"
	"net"
	"testing"
	"time"

	pingv1 "github.com/guycole/daring-cyclops/gen/proto/ping/v1"
	clientcfg "github.com/guycole/daring-cyclops/internal/client"
	pingcmd "github.com/guycole/daring-cyclops/internal/client/ping"
	pingserver "github.com/guycole/daring-cyclops/internal/server/ping"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func TestRunWithDialerSucceedsAgainstBufconnServer(t *testing.T) {
	fixedTime := time.Date(2026, time.June, 15, 12, 0, 0, 0, time.UTC)
	dialer := newBufconnDialer(t, pingserver.NewService("9.9.9", func() time.Time { return fixedTime }))

	result, err := pingcmd.RunWithDialer(context.Background(), clientcfg.Config{Timeout: time.Second}, dialer)
	if err != nil {
		t.Fatalf("RunWithDialer returned error: %v", err)
	}

	if result.ServerVersion != "9.9.9" {
		t.Fatalf("server version = %q, want %q", result.ServerVersion, "9.9.9")
	}

	if !result.ServerTime.Equal(fixedTime) {
		t.Fatalf("server time = %s, want %s", result.ServerTime, fixedTime)
	}
}

func TestRunWithDialerRejectsInvalidResponse(t *testing.T) {
	dialer := newBufconnDialer(t, invalidPingService{})

	_, err := pingcmd.RunWithDialer(context.Background(), clientcfg.Config{Timeout: time.Second}, dialer)
	if err == nil {
		t.Fatal("RunWithDialer unexpectedly succeeded")
	}

	if got := err.Error(); got != "ping failed: response missing server version" {
		t.Fatalf("error = %q, want %q", got, "ping failed: response missing server version")
	}
}

type invalidPingService struct {
	pingv1.UnimplementedPingServiceServer
}

func (invalidPingService) Ping(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error) {
	return &pingv1.PingResponse{}, nil
}

func newBufconnDialer(t *testing.T, service pingv1.PingServiceServer) clientcfg.Dialer {
	t.Helper()

	listener := bufconn.Listen(1024 * 1024)
	grpcServer := grpc.NewServer()
	pingv1.RegisterPingServiceServer(grpcServer, service)

	go func() {
		_ = grpcServer.Serve(listener)
	}()

	t.Cleanup(func() {
		grpcServer.Stop()
		_ = listener.Close()
	})

	return func(ctx context.Context, config clientcfg.Config) (clientcfg.Connection, clientcfg.Config, error) {
		normalized, err := config.WithDefaults()
		if err != nil {
			return nil, clientcfg.Config{}, err
		}

		connection, err := grpc.DialContext(
			ctx,
			"bufnet",
			grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
				return listener.Dial()
			}),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			return nil, clientcfg.Config{}, err
		}

		return connection, normalized, nil
	}
}
