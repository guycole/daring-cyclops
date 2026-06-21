package integration_test

import (
	"net"
	"strings"
	"testing"

	sessionv1 "github.com/guycole/daring-cyclops/gen/proto/session/v1"
	sessionserver "github.com/guycole/daring-cyclops/internal/server/session"
	"google.golang.org/grpc"
)

// ---------------------------------------------------------------------------
// T040: CLI session integration tests
// ---------------------------------------------------------------------------

func TestCLISessionCreate(t *testing.T) {
	address := startSessionTCPServer(t)

	output, err := runCLI(t, "session", "create",
		"-server", address,
		"-player", "Kirk",
		"-timeout", "2s",
	)
	if err != nil {
		t.Fatalf("CLI session create returned error: %v\noutput:\n%s", err, output)
	}

	if !strings.Contains(output, "Session created:") {
		t.Fatalf("output missing 'Session created:': %q", output)
	}
	if !strings.Contains(output, "Mode: REGULAR") {
		t.Fatalf("output missing 'Mode: REGULAR': %q", output)
	}
}

func TestCLISessionList(t *testing.T) {
	address := startSessionTCPServer(t)

	// Create a session first.
	_, err := runCLI(t, "session", "create",
		"-server", address,
		"-player", "Kirk",
		"-timeout", "2s",
	)
	if err != nil {
		t.Fatalf("setup CreateSession failed: %v", err)
	}

	output, err := runCLI(t, "session", "list",
		"-server", address,
		"-timeout", "2s",
	)
	if err != nil {
		t.Fatalf("CLI session list returned error: %v\noutput:\n%s", err, output)
	}

	if !strings.Contains(output, "WAITING") {
		t.Fatalf("list output should contain WAITING state: %q", output)
	}
}

func TestCLISessionCreateMissingPlayer(t *testing.T) {
	address := startSessionTCPServer(t)

	_, err := runCLI(t, "session", "create",
		"-server", address,
		"-timeout", "2s",
	)
	if err == nil {
		t.Fatal("expected error when player name is missing")
	}
}

func TestCLISessionUnknownSubcommand(t *testing.T) {
	address := startSessionTCPServer(t)

	output, err := runCLI(t, "session", "badcmd",
		"-server", address,
	)
	if err == nil {
		t.Fatalf("expected error for unknown sub-command\noutput:\n%s", output)
	}
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func startSessionTCPServer(t *testing.T) string {
	t.Helper()

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Listen failed: %v", err)
	}

	grpcServer := grpc.NewServer()
	registry := sessionserver.NewRegistry()
	sessionv1.RegisterSessionServiceServer(grpcServer, sessionserver.NewService(registry))

	go func() {
		_ = grpcServer.Serve(listener)
	}()
	t.Cleanup(func() {
		grpcServer.Stop()
		_ = listener.Close()
	})

	return listener.Addr().String()
}
