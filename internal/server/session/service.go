package session

import (
	"context"
	"crypto/rand"
	"fmt"
	"sync"

	sessionv1 "github.com/guycole/daring-cyclops/gen/proto/session/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ---------------------------------------------------------------------------
// Registry
// ---------------------------------------------------------------------------

// Registry is a thread-safe store of GameSession objects.
type Registry struct {
	mu       sync.RWMutex
	sessions map[string]*GameSession
}

// NewRegistry creates an empty Registry ready for use.
func NewRegistry() *Registry {
	return &Registry{sessions: make(map[string]*GameSession)}
}

// create adds a new session and returns it. The caller must NOT hold mu.
func (r *Registry) create(config SessionConfig) *GameSession {
	id := generateUUID()
	gs := NewSession(id, config)
	r.mu.Lock()
	r.sessions[id] = gs
	r.mu.Unlock()
	return gs
}

// get returns the session for the given ID under a read lock.
func (r *Registry) get(id string) (*GameSession, bool) {
	r.mu.RLock()
	gs, ok := r.sessions[id]
	r.mu.RUnlock()
	return gs, ok
}

// list returns all sessions, optionally including ENDED ones.
func (r *Registry) list(includeEnded bool) []*GameSession {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]*GameSession, 0, len(r.sessions))
	for _, gs := range r.sessions {
		if !includeEnded && gs.State == SessionStateEnded {
			continue
		}
		result = append(result, gs)
	}
	return result
}

// generateUUID returns a random v4 UUID using crypto/rand.
func generateUUID() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		panic(fmt.Sprintf("crypto/rand.Read: %v", err))
	}
	b[6] = (b[6] & 0x0f) | 0x40 // version 4
	b[8] = (b[8] & 0x3f) | 0x80 // variant bits
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

// ---------------------------------------------------------------------------
// Service — implements SessionServiceServer
// ---------------------------------------------------------------------------

// Service handles SessionService RPCs.
type Service struct {
	sessionv1.UnimplementedSessionServiceServer
	registry *Registry
}

// NewService creates a Service backed by the given Registry.
func NewService(registry *Registry) *Service {
	return &Service{registry: registry}
}

// ---------------------------------------------------------------------------
// US1 — CreateSession, ListSessions, GetSessionStatus
// ---------------------------------------------------------------------------

// CreateSession creates a new game session with the supplied configuration and
// places the creating player into the lobby.
func (s *Service) CreateSession(
	_ context.Context,
	req *sessionv1.CreateSessionRequest,
) (*sessionv1.CreateSessionResponse, error) {
	if req.GetPlayerName() == "" {
		return nil, status.Error(codes.InvalidArgument, "player_name is required")
	}

	config := protoToConfig(req.GetConfig())
	if config.GameMode == GameModeTournament && config.TournamentName == "" {
		return nil, status.Error(codes.InvalidArgument, "tournament_name is required for TOURNAMENT mode")
	}
	config.ApplyDefaults()

	gs := s.registry.create(config)

	gs.mu.Lock()
	defer gs.mu.Unlock()

	if err := gs.AddPlayerToLobby(req.GetPlayerName(), SideFederation); err != nil {
		return nil, status.Errorf(codes.Internal, "add player to lobby: %v", err)
	}

	return &sessionv1.CreateSessionResponse{
		SessionId: gs.ID,
		Config:    configToProto(gs.Config),
	}, nil
}

// ListSessions returns summaries for all active (and optionally ended) sessions.
func (s *Service) ListSessions(
	_ context.Context,
	req *sessionv1.ListSessionsRequest,
) (*sessionv1.ListSessionsResponse, error) {
	sessions := s.registry.list(req.GetIncludeEnded())
	summaries := make([]*sessionv1.SessionSummary, 0, len(sessions))
	for _, gs := range sessions {
		gs.mu.Lock()
		summary := &sessionv1.SessionSummary{
			SessionId:         gs.ID,
			State:             sessionStateToProto(gs.State),
			Config:            configToProto(gs.Config),
			FederationActive:  int32(gs.CountActive(SideFederation)),
			EmpireActive:      int32(gs.CountActive(SideEmpire)),
			FederationLobby:   int32(gs.CountLobby(SideFederation)),
			EmpireLobby:       int32(gs.CountLobby(SideEmpire)),
			CreatedAt:         timestamppb.New(gs.CreatedAt),
		}
		gs.mu.Unlock()
		summaries = append(summaries, summary)
	}
	return &sessionv1.ListSessionsResponse{Sessions: summaries}, nil
}

// GetSessionStatus returns the full state of a session including all ship slots
// and the kill queue.
func (s *Service) GetSessionStatus(
	_ context.Context,
	req *sessionv1.GetSessionStatusRequest,
) (*sessionv1.GetSessionStatusResponse, error) {
	if req.GetSessionId() == "" {
		return nil, status.Error(codes.InvalidArgument, "session_id is required")
	}
	gs, ok := s.registry.get(req.GetSessionId())
	if !ok {
		return nil, status.Errorf(codes.NotFound, "session %q not found", req.GetSessionId())
	}

	gs.mu.Lock()
	defer gs.mu.Unlock()

	return &sessionv1.GetSessionStatusResponse{
		Status: buildSessionStatus(gs),
	}, nil
}

// ---------------------------------------------------------------------------
// US2 — JoinSession
// ---------------------------------------------------------------------------

// JoinSession places an existing player (or new one) into the lobby of an
// active session. The server may redirect to the other side for balance.
func (s *Service) JoinSession(
	_ context.Context,
	req *sessionv1.JoinSessionRequest,
) (*sessionv1.JoinSessionResponse, error) {
	if req.GetSessionId() == "" {
		return nil, status.Error(codes.InvalidArgument, "session_id is required")
	}
	if req.GetPlayerName() == "" {
		return nil, status.Error(codes.InvalidArgument, "player_name is required")
	}

	gs, ok := s.registry.get(req.GetSessionId())
	if !ok {
		return nil, status.Errorf(codes.NotFound, "session %q not found", req.GetSessionId())
	}

	gs.mu.Lock()
	defer gs.mu.Unlock()

	if gs.State == SessionStateEnded {
		return nil, status.Errorf(codes.FailedPrecondition, "session %q has ended", req.GetSessionId())
	}

	preferredSide := protoToSide(req.GetPreferredSide())
	if preferredSide == SideUnspecified {
		preferredSide = SideFederation
	}

	// Check if returning player has a prior side preference.
	if prior, found := gs.LookupPriorSide(req.GetPlayerName()); found {
		preferredSide = prior
	}

	assignedSide, redirected, redirectReason := gs.chooseSide(preferredSide)

	if err := gs.AddPlayerToLobby(req.GetPlayerName(), assignedSide); err != nil {
		return nil, status.Errorf(codes.ResourceExhausted, "join failed: %v", err)
	}

	return &sessionv1.JoinSessionResponse{
		AssignedSide:     sideToProto(assignedSide),
		SideRedirected:   redirected,
		RedirectReason:   redirectReason,
		RomulansEnabled:  gs.Config.RomulansEnabled,
		BlackHolesEnabled: gs.Config.BlackHolesEnabled,
	}, nil
}

// chooseSide determines which side a player should be placed on, applying
// balance rules.  balance_threshold of 0 means no enforcement.
func (gs *GameSession) chooseSide(preferred Side) (assigned Side, redirected bool, reason string) {
	other := SideFederation
	if preferred == SideFederation {
		other = SideEmpire
	}

	preferredCount := gs.CountOccupied(preferred)
	otherCount := gs.CountOccupied(other)

	// Hard cap: preferred side full (9/9 slots).
	if preferredCount >= 9 {
		return other, true, fmt.Sprintf("%s is full", preferred)
	}

	// Balance check: redirect if preferred side has 2+ more players than the other.
	const balanceThreshold = 2
	if preferredCount-otherCount >= balanceThreshold && otherCount < 9 {
		return other, true, fmt.Sprintf("balance: %s has %d, %s has %d", preferred, preferredCount, other, otherCount)
	}

	return preferred, false, ""
}

// ---------------------------------------------------------------------------
// US3 — LobbyCommand
// ---------------------------------------------------------------------------

// LobbyCommand executes a lobby sub-command for a player in LOBBY state.
func (s *Service) LobbyCommand(
	_ context.Context,
	req *sessionv1.LobbyCommandRequest,
) (*sessionv1.LobbyCommandResponse, error) {
	if req.GetSessionId() == "" {
		return nil, status.Error(codes.InvalidArgument, "session_id is required")
	}
	if req.GetPlayerName() == "" {
		return nil, status.Error(codes.InvalidArgument, "player_name is required")
	}
	if req.GetCommand() == sessionv1.LobbyCommand_LOBBY_COMMAND_UNSPECIFIED {
		return nil, status.Error(codes.InvalidArgument, "command is required")
	}

	gs, ok := s.registry.get(req.GetSessionId())
	if !ok {
		return nil, status.Errorf(codes.NotFound, "session %q not found", req.GetSessionId())
	}

	gs.mu.Lock()
	defer gs.mu.Unlock()

	if gs.State == SessionStateEnded {
		return nil, status.Errorf(codes.FailedPrecondition, "session %q has ended", req.GetSessionId())
	}

	// Verify player is in LOBBY state (except QUIT which removes them).
	info, ok := gs.players[req.GetPlayerName()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "player %q not in session", req.GetPlayerName())
	}
	ss := gs.slots(info.side)
	if ss[info.shipSlotIndex].PlayerState != PlayerStateLobby &&
		req.GetCommand() != sessionv1.LobbyCommand_LOBBY_COMMAND_QUIT {
		return nil, status.Errorf(codes.FailedPrecondition,
			"player %q is not in lobby state", req.GetPlayerName())
	}

	var output string
	sessionEnded := false

	switch req.GetCommand() {
	case sessionv1.LobbyCommand_LOBBY_COMMAND_HELP:
		output = gs.LobbyHelp()
	case sessionv1.LobbyCommand_LOBBY_COMMAND_NEWS:
		output = gs.LobbyNews()
	case sessionv1.LobbyCommand_LOBBY_COMMAND_USERS:
		output = gs.LobbyUsers()
	case sessionv1.LobbyCommand_LOBBY_COMMAND_TIME:
		output = gs.LobbyTime()
	case sessionv1.LobbyCommand_LOBBY_COMMAND_POINTS:
		output = gs.LobbyPoints()
	case sessionv1.LobbyCommand_LOBBY_COMMAND_SUMMARY:
		output = gs.LobbySummary()
	case sessionv1.LobbyCommand_LOBBY_COMMAND_GRIPE:
		output = gs.LobbyGripe(req.GetArgument())
	case sessionv1.LobbyCommand_LOBBY_COMMAND_QUIT:
		gs.RemovePlayer(req.GetPlayerName())
		output = fmt.Sprintf("Player %q has left the session.", req.GetPlayerName())
		sessionEnded = true
	default:
		return nil, status.Errorf(codes.InvalidArgument, "unknown command: %v", req.GetCommand())
	}

	return &sessionv1.LobbyCommandResponse{
		Output:        output,
		SessionEnded:  sessionEnded,
	}, nil
}

// ---------------------------------------------------------------------------
// US4 — ActivateShip
// ---------------------------------------------------------------------------

// ActivateShip transitions a player from LOBBY to ACTIVE and, on the first
// activation, transitions the session from WAITING to ACTIVE.
func (s *Service) ActivateShip(
	_ context.Context,
	req *sessionv1.ActivateShipRequest,
) (*sessionv1.ActivateShipResponse, error) {
	if req.GetSessionId() == "" {
		return nil, status.Error(codes.InvalidArgument, "session_id is required")
	}
	if req.GetPlayerName() == "" {
		return nil, status.Error(codes.InvalidArgument, "player_name is required")
	}

	gs, ok := s.registry.get(req.GetSessionId())
	if !ok {
		return nil, status.Errorf(codes.NotFound, "session %q not found", req.GetSessionId())
	}

	gs.mu.Lock()
	defer gs.mu.Unlock()

	if gs.State == SessionStateEnded {
		return nil, status.Errorf(codes.FailedPrecondition, "session %q has ended", req.GetSessionId())
	}

	shipName, err := gs.AssignShip(req.GetPlayerName(), req.GetPreferredShipName())
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "activate ship: %v", err)
	}

	// First activation transitions session from WAITING → ACTIVE.
	if gs.State == SessionStateWaiting {
		gs.State = SessionStateActive
	}

	info := gs.players[req.GetPlayerName()]
	return &sessionv1.ActivateShipResponse{
		ShipName:     shipName,
		Side:         sideToProto(info.side),
		SessionState: sessionStateToProto(gs.State),
	}, nil
}

// ---------------------------------------------------------------------------
// US5 — GetKillQueueStatus
// ---------------------------------------------------------------------------

// GetKillQueueStatus returns the respawn countdown for a player in the kill queue.
func (s *Service) GetKillQueueStatus(
	_ context.Context,
	req *sessionv1.GetKillQueueStatusRequest,
) (*sessionv1.GetKillQueueStatusResponse, error) {
	if req.GetSessionId() == "" {
		return nil, status.Error(codes.InvalidArgument, "session_id is required")
	}
	if req.GetPlayerName() == "" {
		return nil, status.Error(codes.InvalidArgument, "player_name is required")
	}

	gs, ok := s.registry.get(req.GetSessionId())
	if !ok {
		return nil, status.Errorf(codes.NotFound, "session %q not found", req.GetSessionId())
	}

	gs.mu.Lock()
	defer gs.mu.Unlock()

	for _, e := range gs.KillQueue {
		if e.PlayerName == req.GetPlayerName() {
			remaining := e.EligibleAtStardate - gs.CurrentStardate
			if remaining < 0 {
				remaining = 0
			}
			return &sessionv1.GetKillQueueStatusResponse{
				InKillQueue:        true,
				CurrentStardate:    gs.CurrentStardate,
				EligibleAtStardate: e.EligibleAtStardate,
				TicksRemaining:     int32(remaining),
			}, nil
		}
	}

	return &sessionv1.GetKillQueueStatusResponse{
		InKillQueue:     false,
		CurrentStardate: gs.CurrentStardate,
	}, nil
}

// ---------------------------------------------------------------------------
// US6 — session end (CheckVictory called after state mutations)
// ---------------------------------------------------------------------------

// checkAndApplyVictory inspects the session state and, if a victory condition
// is met, transitions the session to ENDED.  Must be called with gs.mu held.
func checkAndApplyVictory(gs *GameSession) {
	if won, side := gs.CheckVictory(); won {
		gs.State = SessionStateEnded
		gs.WinningSide = side
	}
}

// ---------------------------------------------------------------------------
// Proto conversion helpers
// ---------------------------------------------------------------------------

func protoToConfig(pb *sessionv1.SessionConfig) SessionConfig {
	if pb == nil {
		return SessionConfig{}
	}
	gm := GameModeRegular
	switch pb.GetGameMode() {
	case sessionv1.GameMode_GAME_MODE_TOURNAMENT:
		gm = GameModeTournament
	}
	return SessionConfig{
		GameMode:          gm,
		TournamentName:    pb.GetTournamentName(),
		RomulansEnabled:   pb.GetRomulansEnabled(),
		BlackHolesEnabled: pb.GetBlackHolesEnabled(),
		RespawnWaitTicks:  pb.GetRespawnWaitTicks(),
	}
}

func configToProto(c SessionConfig) *sessionv1.SessionConfig {
	gm := sessionv1.GameMode_GAME_MODE_REGULAR
	if c.GameMode == GameModeTournament {
		gm = sessionv1.GameMode_GAME_MODE_TOURNAMENT
	}
	return &sessionv1.SessionConfig{
		GameMode:          gm,
		TournamentName:    c.TournamentName,
		RomulansEnabled:   c.RomulansEnabled,
		BlackHolesEnabled: c.BlackHolesEnabled,
		RespawnWaitTicks:  c.RespawnWaitTicks,
	}
}

func sessionStateToProto(s SessionState) sessionv1.SessionState {
	switch s {
	case SessionStateWaiting:
		return sessionv1.SessionState_SESSION_STATE_WAITING
	case SessionStateActive:
		return sessionv1.SessionState_SESSION_STATE_ACTIVE
	case SessionStateEnded:
		return sessionv1.SessionState_SESSION_STATE_ENDED
	default:
		return sessionv1.SessionState_SESSION_STATE_UNSPECIFIED
	}
}

func sideToProto(s Side) sessionv1.Side {
	switch s {
	case SideFederation:
		return sessionv1.Side_SIDE_FEDERATION
	case SideEmpire:
		return sessionv1.Side_SIDE_EMPIRE
	default:
		return sessionv1.Side_SIDE_UNSPECIFIED
	}
}

func protoToSide(pb sessionv1.Side) Side {
	switch pb {
	case sessionv1.Side_SIDE_FEDERATION:
		return SideFederation
	case sessionv1.Side_SIDE_EMPIRE:
		return SideEmpire
	default:
		return SideUnspecified
	}
}

func playerStateToProto(ps PlayerState) sessionv1.PlayerState {
	switch ps {
	case PlayerStateUnoccupied:
		return sessionv1.PlayerState_PLAYER_STATE_UNOCCUPIED
	case PlayerStateLobby:
		return sessionv1.PlayerState_PLAYER_STATE_LOBBY
	case PlayerStateActive:
		return sessionv1.PlayerState_PLAYER_STATE_ACTIVE
	case PlayerStateKillQueue:
		return sessionv1.PlayerState_PLAYER_STATE_KILL_QUEUE
	default:
		return sessionv1.PlayerState_PLAYER_STATE_UNSPECIFIED
	}
}

func slotToProto(s ShipSlot) *sessionv1.ShipSlot {
	return &sessionv1.ShipSlot{
		SlotIndex:   int32(s.SlotIndex),
		Side:        sideToProto(s.Side),
		ShipName:    s.ShipName,
		PlayerName:  s.PlayerName,
		PlayerState: playerStateToProto(s.PlayerState),
	}
}

func killQueueEntryToProto(e KillQueueEntry) *sessionv1.KillQueueEntry {
	return &sessionv1.KillQueueEntry{
		PlayerName:         e.PlayerName,
		Side:               sideToProto(e.Side),
		ShipName:           e.ShipName,
		KilledAtStardate:   e.KilledAtStardate,
		EligibleAtStardate: e.EligibleAtStardate,
	}
}

func buildSessionStatus(gs *GameSession) *sessionv1.SessionStatus {
	fedSlots := make([]*sessionv1.ShipSlot, 9)
	empSlots := make([]*sessionv1.ShipSlot, 9)
	for i := 0; i < 9; i++ {
		fedSlots[i] = slotToProto(gs.FederationSlots[i])
		empSlots[i] = slotToProto(gs.EmpireSlots[i])
	}
	kq := make([]*sessionv1.KillQueueEntry, len(gs.KillQueue))
	for i, e := range gs.KillQueue {
		kq[i] = killQueueEntryToProto(e)
	}
	return &sessionv1.SessionStatus{
		SessionId:       gs.ID,
		State:           sessionStateToProto(gs.State),
		Config:          configToProto(gs.Config),
		CurrentStardate: gs.CurrentStardate,
		FederationSlots: fedSlots,
		EmpireSlots:     empSlots,
		KillQueue:       kq,
		WinningSide:     sideToProto(gs.WinningSide),
	}
}


