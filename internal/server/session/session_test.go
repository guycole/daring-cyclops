package session

import (
	"testing"
)

// ---------------------------------------------------------------------------
// T037: Domain model unit tests
// ---------------------------------------------------------------------------

func TestNewSession(t *testing.T) {
	cfg := SessionConfig{}
	cfg.ApplyDefaults()
	gs := NewSession("test-id", cfg)

	if gs.ID != "test-id" {
		t.Errorf("ID = %q; want %q", gs.ID, "test-id")
	}
	if gs.State != SessionStateWaiting {
		t.Errorf("State = %v; want WAITING", gs.State)
	}
	if len(gs.FederationSlots) != 9 {
		t.Errorf("FederationSlots len = %d; want 9", len(gs.FederationSlots))
	}
	if len(gs.EmpireSlots) != 9 {
		t.Errorf("EmpireSlots len = %d; want 9", len(gs.EmpireSlots))
	}
	// All slots should be UNOCCUPIED.
	for i, s := range gs.FederationSlots {
		if s.PlayerState != PlayerStateUnoccupied {
			t.Errorf("FederationSlots[%d].PlayerState = %v; want UNOCCUPIED", i, s.PlayerState)
		}
		if s.ShipName == "" {
			t.Errorf("FederationSlots[%d].ShipName is empty", i)
		}
	}
}

func TestApplyDefaults(t *testing.T) {
	cfg := SessionConfig{}
	cfg.ApplyDefaults()

	if cfg.GameMode != GameModeRegular {
		t.Errorf("GameMode = %v; want REGULAR", cfg.GameMode)
	}
	if !cfg.RomulansEnabled {
		t.Error("RomulansEnabled should be true by default")
	}
	if cfg.RespawnWaitTicks != 120 {
		t.Errorf("RespawnWaitTicks = %d; want 120", cfg.RespawnWaitTicks)
	}
}

func TestDeriveTournamentSeed(t *testing.T) {
	seed1 := DeriveTournamentSeed("Alpha")
	seed2 := DeriveTournamentSeed("Alpha")
	seed3 := DeriveTournamentSeed("Beta")

	if seed1 != seed2 {
		t.Error("same name should produce same seed")
	}
	if seed1 == seed3 {
		t.Error("different names should produce different seeds")
	}
	if seed1 == 0 {
		t.Error("seed should be non-zero")
	}
}

func TestAddPlayerToLobby(t *testing.T) {
	gs := newTestSession(t)

	if err := gs.AddPlayerToLobby("Kirk", SideFederation); err != nil {
		t.Fatalf("AddPlayerToLobby: %v", err)
	}

	if gs.CountLobby(SideFederation) != 1 {
		t.Errorf("CountLobby(Federation) = %d; want 1", gs.CountLobby(SideFederation))
	}

	// Duplicate join should fail.
	if err := gs.AddPlayerToLobby("Kirk", SideFederation); err == nil {
		t.Error("expected error on duplicate join")
	}
}

func TestAddPlayerToLobbyFull(t *testing.T) {
	gs := newTestSession(t)

	// Fill all 9 Federation slots.
	names := []string{"p1", "p2", "p3", "p4", "p5", "p6", "p7", "p8", "p9"}
	for _, n := range names {
		if err := gs.AddPlayerToLobby(n, SideFederation); err != nil {
			t.Fatalf("AddPlayerToLobby(%q): %v", n, err)
		}
	}

	// 10th player should fail.
	if err := gs.AddPlayerToLobby("p10", SideFederation); err == nil {
		t.Error("expected error when side is full")
	}
}

func TestAssignShip(t *testing.T) {
	gs := newTestSession(t)
	_ = gs.AddPlayerToLobby("Kirk", SideFederation)

	shipName, err := gs.AssignShip("Kirk", "")
	if err != nil {
		t.Fatalf("AssignShip: %v", err)
	}
	if shipName == "" {
		t.Error("expected non-empty ship name")
	}
	if gs.CountActive(SideFederation) != 1 {
		t.Errorf("CountActive(Federation) = %d; want 1", gs.CountActive(SideFederation))
	}
}

func TestAssignShipNotInSession(t *testing.T) {
	gs := newTestSession(t)

	if _, err := gs.AssignShip("Nobody", ""); err == nil {
		t.Error("expected error for player not in session")
	}
}

func TestKillPlayer(t *testing.T) {
	gs := newTestSession(t)
	_ = gs.AddPlayerToLobby("Kirk", SideFederation)
	_, _ = gs.AssignShip("Kirk", "")

	if err := gs.KillPlayer("Kirk"); err != nil {
		t.Fatalf("KillPlayer: %v", err)
	}

	if len(gs.KillQueue) != 1 {
		t.Errorf("KillQueue len = %d; want 1", len(gs.KillQueue))
	}
	if gs.KillQueue[0].PlayerName != "Kirk" {
		t.Errorf("KillQueue[0].PlayerName = %q; want Kirk", gs.KillQueue[0].PlayerName)
	}
}

func TestKillPlayerNotActive(t *testing.T) {
	gs := newTestSession(t)
	_ = gs.AddPlayerToLobby("Kirk", SideFederation)

	// Player is in LOBBY, not ACTIVE — should fail.
	if err := gs.KillPlayer("Kirk"); err == nil {
		t.Error("expected error when killing non-active player")
	}
}

func TestRespawnToLobby(t *testing.T) {
	gs := newTestSession(t)
	gs.Config.RespawnWaitTicks = 0 // instant respawn for test
	_ = gs.AddPlayerToLobby("Kirk", SideFederation)
	_, _ = gs.AssignShip("Kirk", "")
	_ = gs.KillPlayer("Kirk")

	if err := gs.RespawnToLobby("Kirk"); err != nil {
		t.Fatalf("RespawnToLobby: %v", err)
	}

	if len(gs.KillQueue) != 0 {
		t.Errorf("KillQueue len = %d; want 0 after respawn", len(gs.KillQueue))
	}
	if gs.CountLobby(SideFederation) != 1 {
		t.Errorf("CountLobby(Federation) = %d; want 1 after respawn", gs.CountLobby(SideFederation))
	}
}

func TestRespawnNotEligible(t *testing.T) {
	gs := newTestSession(t)
	gs.Config.RespawnWaitTicks = 120
	_ = gs.AddPlayerToLobby("Kirk", SideFederation)
	_, _ = gs.AssignShip("Kirk", "")
	_ = gs.KillPlayer("Kirk")

	// Current stardate is 0; eligible at 120 — should fail.
	if err := gs.RespawnToLobby("Kirk"); err == nil {
		t.Error("expected error when player not yet eligible to respawn")
	}
}

func TestCheckVictory(t *testing.T) {
	gs := newTestSession(t)
	gs.State = SessionStateActive

	// Empire wins: Federation has no active/lobby players, Empire has one active.
	_ = gs.AddPlayerToLobby("Kang", SideEmpire)
	_, _ = gs.AssignShip("Kang", "")

	won, side := gs.CheckVictory()
	if !won {
		t.Error("expected victory condition to be met")
	}
	if side != SideEmpire {
		t.Errorf("winning side = %v; want EMPIRE", side)
	}
}

func TestCheckVictoryNoWinner(t *testing.T) {
	gs := newTestSession(t)
	gs.State = SessionStateActive

	_ = gs.AddPlayerToLobby("Kirk", SideFederation)
	_, _ = gs.AssignShip("Kirk", "")
	_ = gs.AddPlayerToLobby("Kang", SideEmpire)
	_, _ = gs.AssignShip("Kang", "")

	won, _ := gs.CheckVictory()
	if won {
		t.Error("no victory expected when both sides have active players")
	}
}

func TestKillQueueMaxTen(t *testing.T) {
	gs := newTestSession(t)
	gs.Config.RespawnWaitTicks = 9999

	// Add and kill 11 players to verify the queue caps at 10.
	for i := 0; i < 11 && i < 9; i++ {
		name := []string{"p1", "p2", "p3", "p4", "p5", "p6", "p7", "p8", "p9"}[i]
		_ = gs.AddPlayerToLobby(name, SideFederation)
		_, _ = gs.AssignShip(name, "")
		_ = gs.KillPlayer(name)
	}
	// Add a kill manually to bypass slot limits for testing enqueueKill.
	for i := 9; i < 11; i++ {
		gs.enqueueKill(KillQueueEntry{
			PlayerName:         []string{"extra1", "extra2"}[i-9],
			Side:               SideFederation,
			ShipName:           "Lexington",
			KilledAtStardate:   float64(i),
			EligibleAtStardate: float64(i + 9999),
		})
	}

	if len(gs.KillQueue) > 10 {
		t.Errorf("KillQueue len = %d; want ≤ 10", len(gs.KillQueue))
	}
}

func TestLobbyCommands(t *testing.T) {
	gs := newTestSession(t)
	_ = gs.AddPlayerToLobby("Kirk", SideFederation)

	if out := gs.LobbyHelp(); out == "" {
		t.Error("LobbyHelp should return non-empty text")
	}
	if out := gs.LobbyUsers(); out == "" {
		t.Error("LobbyUsers should return non-empty text")
	}
	if out := gs.LobbyTime(); out == "" {
		t.Error("LobbyTime should return non-empty text")
	}
	if out := gs.LobbyPoints(); out == "" {
		t.Error("LobbyPoints should return non-empty text")
	}
	if out := gs.LobbyNews(); out == "" {
		t.Error("LobbyNews should return non-empty text")
	}
	if out := gs.LobbySummary(); out == "" {
		t.Error("LobbySummary should return non-empty text")
	}
	if out := gs.LobbyGripe("test complaint"); out == "" {
		t.Error("LobbyGripe should return non-empty text")
	}
	if out := gs.LobbyGripe(""); out == "" {
		t.Error("LobbyGripe with empty text should return an error message")
	}
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func newTestSession(t *testing.T) *GameSession {
	t.Helper()
	cfg := SessionConfig{RespawnWaitTicks: 0}
	cfg.ApplyDefaults()
	cfg.RespawnWaitTicks = 0
	return NewSession("test-session", cfg)
}
