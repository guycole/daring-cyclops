package integration_test

import (
	"context"
	"net"
	"testing"
	"time"

	sessionv1 "github.com/guycole/daring-cyclops/gen/proto/session/v1"
	sessionserver "github.com/guycole/daring-cyclops/internal/server/session"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

// ---------------------------------------------------------------------------
// T039: gRPC integration test — full session lifecycle over bufconn
// ---------------------------------------------------------------------------

func TestGRPCSessionFullLifecycle(t *testing.T) {
	client := newSessionBufconnClient(t)
	ctx := context.Background()

	// 1. Create a session.
	createResp, err := client.CreateSession(ctx, &sessionv1.CreateSessionRequest{
		PlayerName: "Kirk",
		Config: &sessionv1.SessionConfig{
			GameMode:         sessionv1.GameMode_GAME_MODE_REGULAR,
			RomulansEnabled:  true,
			RespawnWaitTicks: 0, // instant respawn for testing
		},
	})
	if err != nil {
		t.Fatalf("CreateSession: %v", err)
	}
	sessionID := createResp.GetSessionId()
	if sessionID == "" {
		t.Fatal("expected non-empty session_id")
	}

	// 2. List sessions — should find our new session.
	listResp, err := client.ListSessions(ctx, &sessionv1.ListSessionsRequest{})
	if err != nil {
		t.Fatalf("ListSessions: %v", err)
	}
	if len(listResp.GetSessions()) == 0 {
		t.Fatal("expected at least one session in list")
	}
	found := false
	for _, s := range listResp.GetSessions() {
		if s.GetSessionId() == sessionID {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("session %q not found in list", sessionID)
	}

	// 3. Join the session as a second player.
	joinResp, err := client.JoinSession(ctx, &sessionv1.JoinSessionRequest{
		SessionId:     sessionID,
		PlayerName:    "Spock",
		PreferredSide: sessionv1.Side_SIDE_FEDERATION,
	})
	if err != nil {
		t.Fatalf("JoinSession: %v", err)
	}
	if joinResp.GetAssignedSide() != sessionv1.Side_SIDE_FEDERATION {
		t.Errorf("assigned_side = %v; want FEDERATION", joinResp.GetAssignedSide())
	}

	// 4. Execute a lobby command (help).
	lobbyResp, err := client.LobbyCommand(ctx, &sessionv1.LobbyCommandRequest{
		SessionId:  sessionID,
		PlayerName: "Kirk",
		Command:    sessionv1.LobbyCommand_LOBBY_COMMAND_HELP,
	})
	if err != nil {
		t.Fatalf("LobbyCommand help: %v", err)
	}
	if lobbyResp.GetOutput() == "" {
		t.Error("expected non-empty help output")
	}

	// 5. Activate Kirk's ship — should transition session to ACTIVE.
	activateResp, err := client.ActivateShip(ctx, &sessionv1.ActivateShipRequest{
		SessionId:  sessionID,
		PlayerName: "Kirk",
	})
	if err != nil {
		t.Fatalf("ActivateShip: %v", err)
	}
	if activateResp.GetShipName() == "" {
		t.Error("expected non-empty ship name")
	}
	if activateResp.GetSessionState() != sessionv1.SessionState_SESSION_STATE_ACTIVE {
		t.Errorf("session_state after first activation = %v; want ACTIVE", activateResp.GetSessionState())
	}

	// 6. Check Kirk's kill queue status — should not be in queue.
	kqResp, err := client.GetKillQueueStatus(ctx, &sessionv1.GetKillQueueStatusRequest{
		SessionId:  sessionID,
		PlayerName: "Kirk",
	})
	if err != nil {
		t.Fatalf("GetKillQueueStatus: %v", err)
	}
	if kqResp.GetInKillQueue() {
		t.Error("Kirk should not be in kill queue after activation")
	}

	// 7. Get full session status.
	statusResp, err := client.GetSessionStatus(ctx, &sessionv1.GetSessionStatusRequest{
		SessionId: sessionID,
	})
	if err != nil {
		t.Fatalf("GetSessionStatus: %v", err)
	}
	st := statusResp.GetStatus()
	if st.GetState() != sessionv1.SessionState_SESSION_STATE_ACTIVE {
		t.Errorf("status.state = %v; want ACTIVE", st.GetState())
	}
	if len(st.GetFederationSlots()) != 9 {
		t.Errorf("federation_slots len = %d; want 9", len(st.GetFederationSlots()))
	}
}

// ---------------------------------------------------------------------------
// T041: gRPC integration — concurrent session creation
// ---------------------------------------------------------------------------

func TestGRPCSessionConcurrentCreate(t *testing.T) {
	client := newSessionBufconnClient(t)
	ctx := context.Background()

	const n = 5
	type result struct {
		id  string
		err error
	}
	results := make(chan result, n)

	for i := 0; i < n; i++ {
		go func(i int) {
			resp, err := client.CreateSession(ctx, &sessionv1.CreateSessionRequest{
				PlayerName: "player",
				Config:     &sessionv1.SessionConfig{GameMode: sessionv1.GameMode_GAME_MODE_REGULAR},
			})
			if err != nil {
				results <- result{err: err}
				return
			}
			results <- result{id: resp.GetSessionId()}
		}(i)
	}

	ids := make(map[string]bool)
	for i := 0; i < n; i++ {
		r := <-results
		if r.err != nil {
			t.Errorf("CreateSession goroutine error: %v", r.err)
			continue
		}
		if r.id == "" {
			t.Error("expected non-empty session_id")
			continue
		}
		if ids[r.id] {
			t.Errorf("duplicate session_id: %q", r.id)
		}
		ids[r.id] = true
	}
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func newSessionBufconnClient(t *testing.T) sessionv1.SessionServiceClient {
	t.Helper()

	listener := bufconn.Listen(1024 * 1024)
	registry := sessionserver.NewRegistry()
	grpcServer := grpc.NewServer()
	sessionv1.RegisterSessionServiceServer(grpcServer, sessionserver.NewService(registry))

	go func() {
		_ = grpcServer.Serve(listener)
	}()
	t.Cleanup(func() {
		grpcServer.Stop()
		_ = listener.Close()
	})

	conn, err := grpc.DialContext(
		context.Background(),
		"bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return listener.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("grpc.DialContext bufconn: %v", err)
	}
	t.Cleanup(func() { _ = conn.Close() })

	// Unused variable workaround: the time package is needed for test helpers.
	_ = time.Second

	return sessionv1.NewSessionServiceClient(conn)
}
