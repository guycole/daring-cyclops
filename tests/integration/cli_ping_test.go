package integration_test

import (
	"net"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	pingv1 "github.com/guycole/daring-cyclops/gen/proto/ping/v1"
	pingserver "github.com/guycole/daring-cyclops/internal/server/ping"
	"google.golang.org/grpc"
)

func TestCLIPingSuccess(t *testing.T) {
	address := startTCPServer(t, pingserver.NewService("2.3.4", func() time.Time {
		return time.Date(2026, time.June, 15, 12, 30, 0, 0, time.UTC)
	}))

	output, err := runCLI(t, "ping", "-server", address, "-timeout", "2s")
	if err != nil {
		t.Fatalf("CLI ping returned error: %v\noutput:\n%s", err, output)
	}

	if !strings.Contains(output, "Server Version: 2.3.4") {
		t.Fatalf("success output missing version: %q", output)
	}

	if !strings.Contains(output, "Server Time: 2026-06-15T12:30:00Z") {
		t.Fatalf("success output missing server time: %q", output)
	}
}

func TestCLIPingFailure(t *testing.T) {
	address := closedAddress(t)

	output, err := runCLI(t, "ping", "-server", address, "-timeout", "500ms")
	if err == nil {
		t.Fatalf("CLI ping unexpectedly succeeded\noutput:\n%s", output)
	}

	if !strings.Contains(output, "ping failed:") {
		t.Fatalf("failure output missing ping error prefix: %q", output)
	}

	if strings.Contains(output, "Server Version:") || strings.Contains(output, "Server Time:") {
		t.Fatalf("failure output should not contain success fields: %q", output)
	}
}

func startTCPServer(t *testing.T, service pingv1.PingServiceServer) string {
	t.Helper()

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Listen failed: %v", err)
	}

	grpcServer := grpc.NewServer()
	pingv1.RegisterPingServiceServer(grpcServer, service)

	go func() {
		_ = grpcServer.Serve(listener)
	}()

	t.Cleanup(func() {
		grpcServer.Stop()
		_ = listener.Close()
	})

	return listener.Addr().String()
}

func closedAddress(t *testing.T) string {
	t.Helper()

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Listen failed: %v", err)
	}
	address := listener.Addr().String()
	_ = listener.Close()
	return address
}

func runCLI(t *testing.T, args ...string) (string, error) {
	t.Helper()

	repoRoot, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatalf("Abs failed: %v", err)
	}

	command := exec.Command("go", append([]string{"run", "./cmd/cyclops"}, args...)...)
	command.Dir = repoRoot
	output, err := command.CombinedOutput()
	return string(output), err
}
