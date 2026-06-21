package session

import (
	"context"
	"testing"

	sessionv1 "github.com/guycole/daring-cyclops/gen/proto/session/v1"
)

// ---------------------------------------------------------------------------
// T038: Service unit tests (direct method calls, no gRPC transport)
// ---------------------------------------------------------------------------

func newTestService(t *testing.T) *Service {
	t.Helper()
	return NewService(NewRegistry())
}

func TestCreateSession_Regular(t *testing.T) {
	svc := newTestService(t)
	resp, err := svc.CreateSession(context.Background(), &sessionv1.CreateSessionRequest{
		PlayerName: "Kirk",
		Config: &sessionv1.SessionConfig{
			GameMode: sessionv1.GameMode_GAME_MODE_REGULAR,
		},
	})
	if err != nil {
		t.Fatalf("CreateSession: %v", err)
	}
	if resp.GetSessionId() == "" {
		t.Error("expected non-empty session_id")
	}
	if resp.GetConfig().GetGameMode() != sessionv1.GameMode_GAME_MODE_REGULAR {
		t.Errorf("game_mode = %v; want REGULAR", resp.GetConfig().GetGameMode())
	}
	if resp.GetConfig().GetRespawnWaitTicks() != 120 {
		t.Errorf("respawn_wait_ticks = %d; want 120 (default)", resp.GetConfig().GetRespawnWaitTicks())
	}
}

func TestCreateSession_Tournament(t *testing.T) {
	svc := newTestService(t)
	resp, err := svc.CreateSession(context.Background(), &sessionv1.CreateSessionRequest{
		PlayerName: "Kirk",
		Config: &sessionv1.SessionConfig{
			GameMode:       sessionv1.GameMode_GAME_MODE_TOURNAMENT,
			TournamentName: "Kobayashi Maru",
		},
	})
	if err != nil {
		t.Fatalf("CreateSession tournament: %v", err)
	}
	if resp.GetConfig().GetTournamentName() != "Kobayashi Maru" {
		t.Errorf("tournament_name = %q; want %q", resp.GetConfig().GetTournamentName(), "Kobayashi Maru")
	}
}

func TestCreateSession_MissingPlayerName(t *testing.T) {
	svc := newTestService(t)
	_, err := svc.CreateSession(context.Background(), &sessionv1.CreateSessionRequest{
		PlayerName: "",
	})
	if err == nil {
		t.Error("expected error for missing player_name")
	}
}

func TestCreateSession_TournamentMissingName(t *testing.T) {
	svc := newTestService(t)
	_, err := svc.CreateSession(context.Background(), &sessionv1.CreateSessionRequest{
		PlayerName: "Kirk",
		Config: &sessionv1.SessionConfig{
			GameMode: sessionv1.GameMode_GAME_MODE_TOURNAMENT,
		},
	})
	if err == nil {
		t.Error("expected error for tournament mode without tournament_name")
	}
}

func TestListSessions(t *testing.T) {
	svc := newTestService(t)

	// Create two sessions.
	_, err := svc.CreateSession(context.Background(), &sessionv1.CreateSessionRequest{
		PlayerName: "Kirk",
		Config:     &sessionv1.SessionConfig{GameMode: sessionv1.GameMode_GAME_MODE_REGULAR},
	})
	if err != nil {
		t.Fatalf("CreateSession 1: %v", err)
	}
	_, err = svc.CreateSession(context.Background(), &sessionv1.CreateSessionRequest{
		PlayerName: "Picard",
		Config:     &sessionv1.SessionConfig{GameMode: sessionv1.GameMode_GAME_MODE_REGULAR},
	})
	if err != nil {
		t.Fatalf("CreateSession 2: %v", err)
	}

	resp, err := svc.ListSessions(context.Background(), &sessionv1.ListSessionsRequest{})
	if err != nil {
		t.Fatalf("ListSessions: %v", err)
	}
	if len(resp.GetSessions()) != 2 {
		t.Errorf("sessions count = %d; want 2", len(resp.GetSessions()))
	}
}

func TestJoinSession_Success(t *testing.T) {
	svc := newTestService(t)
	created, _ := svc.CreateSession(context.Background(), &sessionv1.CreateSessionRequest{
		PlayerName: "Kirk",
		Config:     &sessionv1.SessionConfig{GameMode: sessionv1.GameMode_GAME_MODE_REGULAR},
	})

	resp, err := svc.JoinSession(context.Background(), &sessionv1.JoinSessionRequest{
		SessionId:     created.GetSessionId(),
		PlayerName:    "Spock",
		PreferredSide: sessionv1.Side_SIDE_FEDERATION,
	})
	if err != nil {
		t.Fatalf("JoinSession: %v", err)
	}
	if resp.GetAssignedSide() != sessionv1.Side_SIDE_FEDERATION {
		t.Errorf("assigned_side = %v; want FEDERATION", resp.GetAssignedSide())
	}
	if resp.GetSideRedirected() {
		t.Error("should not be redirected")
	}
}

func TestJoinSession_NotFound(t *testing.T) {
	svc := newTestService(t)
	_, err := svc.JoinSession(context.Background(), &sessionv1.JoinSessionRequest{
		SessionId:  "nonexistent",
		PlayerName: "Kirk",
	})
	if err == nil {
		t.Error("expected NotFound error")
	}
}

func TestJoinSession_EndedSession(t *testing.T) {
	svc := newTestService(t)
	created, _ := svc.CreateSession(context.Background(), &sessionv1.CreateSessionRequest{
		PlayerName: "Kirk",
		Config:     &sessionv1.SessionConfig{GameMode: sessionv1.GameMode_GAME_MODE_REGULAR},
	})

	// Manually end the session.
	gs, _ := svc.registry.get(created.GetSessionId())
	gs.mu.Lock()
	gs.State = SessionStateEnded
	gs.mu.Unlock()

	_, err := svc.JoinSession(context.Background(), &sessionv1.JoinSessionRequest{
		SessionId:  created.GetSessionId(),
		PlayerName: "Spock",
	})
	if err == nil {
		t.Error("expected error joining ended session")
	}
}

func TestLobbyCommand_Help(t *testing.T) {
	svc, sessionID := createSessionWithPlayer(t, "Kirk")

	resp, err := svc.LobbyCommand(context.Background(), &sessionv1.LobbyCommandRequest{
		SessionId:  sessionID,
		PlayerName: "Kirk",
		Command:    sessionv1.LobbyCommand_LOBBY_COMMAND_HELP,
	})
	if err != nil {
		t.Fatalf("LobbyCommand help: %v", err)
	}
	if resp.GetOutput() == "" {
		t.Error("expected non-empty help output")
	}
	if resp.GetSessionEnded() {
		t.Error("help should not end session")
	}
}

func TestLobbyCommand_Quit(t *testing.T) {
	svc, sessionID := createSessionWithPlayer(t, "Kirk")

	resp, err := svc.LobbyCommand(context.Background(), &sessionv1.LobbyCommandRequest{
		SessionId:  sessionID,
		PlayerName: "Kirk",
		Command:    sessionv1.LobbyCommand_LOBBY_COMMAND_QUIT,
	})
	if err != nil {
		t.Fatalf("LobbyCommand quit: %v", err)
	}
	if !resp.GetSessionEnded() {
		t.Error("quit should set session_ended = true")
	}
}

func TestLobbyCommand_Unspecified(t *testing.T) {
	svc, sessionID := createSessionWithPlayer(t, "Kirk")

	_, err := svc.LobbyCommand(context.Background(), &sessionv1.LobbyCommandRequest{
		SessionId:  sessionID,
		PlayerName: "Kirk",
		Command:    sessionv1.LobbyCommand_LOBBY_COMMAND_UNSPECIFIED,
	})
	if err == nil {
		t.Error("expected error for unspecified command")
	}
}

func TestActivateShip_Success(t *testing.T) {
	svc, sessionID := createSessionWithPlayer(t, "Kirk")

	resp, err := svc.ActivateShip(context.Background(), &sessionv1.ActivateShipRequest{
		SessionId:  sessionID,
		PlayerName: "Kirk",
	})
	if err != nil {
		t.Fatalf("ActivateShip: %v", err)
	}
	if resp.GetShipName() == "" {
		t.Error("expected non-empty ship name")
	}
	if resp.GetSessionState() != sessionv1.SessionState_SESSION_STATE_ACTIVE {
		t.Errorf("session_state = %v; want ACTIVE", resp.GetSessionState())
	}
}

func TestActivateShip_SessionTransition(t *testing.T) {
	svc, sessionID := createSessionWithPlayer(t, "Kirk")

	// Verify session starts in WAITING.
	gs, _ := svc.registry.get(sessionID)
	gs.mu.Lock()
	if gs.State != SessionStateWaiting {
		t.Errorf("initial state = %v; want WAITING", gs.State)
	}
	gs.mu.Unlock()

	_, err := svc.ActivateShip(context.Background(), &sessionv1.ActivateShipRequest{
		SessionId:  sessionID,
		PlayerName: "Kirk",
	})
	if err != nil {
		t.Fatalf("ActivateShip: %v", err)
	}

	gs.mu.Lock()
	if gs.State != SessionStateActive {
		t.Errorf("post-activate state = %v; want ACTIVE", gs.State)
	}
	gs.mu.Unlock()
}

func TestGetKillQueueStatus_NotInQueue(t *testing.T) {
	svc, sessionID := createSessionWithPlayer(t, "Kirk")

	resp, err := svc.GetKillQueueStatus(context.Background(), &sessionv1.GetKillQueueStatusRequest{
		SessionId:  sessionID,
		PlayerName: "Kirk",
	})
	if err != nil {
		t.Fatalf("GetKillQueueStatus: %v", err)
	}
	if resp.GetInKillQueue() {
		t.Error("player should not be in kill queue")
	}
}

func TestGetSessionStatus(t *testing.T) {
	svc, sessionID := createSessionWithPlayer(t, "Kirk")

	resp, err := svc.GetSessionStatus(context.Background(), &sessionv1.GetSessionStatusRequest{
		SessionId: sessionID,
	})
	if err != nil {
		t.Fatalf("GetSessionStatus: %v", err)
	}
	st := resp.GetStatus()
	if st == nil {
		t.Fatal("expected non-nil status")
	}
	if st.GetSessionId() != sessionID {
		t.Errorf("session_id = %q; want %q", st.GetSessionId(), sessionID)
	}
	if len(st.GetFederationSlots()) != 9 {
		t.Errorf("federation_slots len = %d; want 9", len(st.GetFederationSlots()))
	}
}

func TestGetSessionStatus_NotFound(t *testing.T) {
	svc := newTestService(t)
	_, err := svc.GetSessionStatus(context.Background(), &sessionv1.GetSessionStatusRequest{
		SessionId: "nonexistent",
	})
	if err == nil {
		t.Error("expected NotFound error")
	}
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func createSessionWithPlayer(t *testing.T, playerName string) (*Service, string) {
	t.Helper()
	svc := newTestService(t)
	resp, err := svc.CreateSession(context.Background(), &sessionv1.CreateSessionRequest{
		PlayerName: playerName,
		Config:     &sessionv1.SessionConfig{GameMode: sessionv1.GameMode_GAME_MODE_REGULAR},
	})
	if err != nil {
		t.Fatalf("CreateSession: %v", err)
	}
	return svc, resp.GetSessionId()
}
