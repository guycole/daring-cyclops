package session

import (
	"fmt"
	"hash/fnv"
	"strings"
	"sync"
	"time"
)

// ---------------------------------------------------------------------------
// Domain enumerations — mirror proto enums, no proto imports.
// ---------------------------------------------------------------------------

// SessionState tracks the lifecycle phase of a game session.
type SessionState int

const (
	SessionStateUnspecified SessionState = 0
	SessionStateWaiting     SessionState = 1 // created, no ship activated yet
	SessionStateActive      SessionState = 2 // at least one ship is active
	SessionStateEnded       SessionState = 3 // victory condition met
)

func (s SessionState) String() string {
	switch s {
	case SessionStateWaiting:
		return "WAITING"
	case SessionStateActive:
		return "ACTIVE"
	case SessionStateEnded:
		return "ENDED"
	default:
		return "UNSPECIFIED"
	}
}

// GameMode determines whether a session uses a fixed seed.
type GameMode int

const (
	GameModeUnspecified GameMode = 0
	GameModeRegular     GameMode = 1 // randomised game
	GameModeTournament  GameMode = 2 // seeded via tournament_name
)

// Side identifies which team a player belongs to.
type Side int

const (
	SideUnspecified Side = 0
	SideFederation  Side = 1
	SideEmpire      Side = 2
)

func (s Side) String() string {
	switch s {
	case SideFederation:
		return "FEDERATION"
	case SideEmpire:
		return "EMPIRE"
	default:
		return "UNSPECIFIED"
	}
}

// PlayerState represents what a player occupying a ship slot is doing.
type PlayerState int

const (
	PlayerStateUnoccupied PlayerState = 1 // slot has no player
	PlayerStateLobby      PlayerState = 2 // player connected, not yet active
	PlayerStateActive     PlayerState = 3 // ship is live in the game
	PlayerStateKillQueue  PlayerState = 4 // ship was destroyed; waiting to respawn
)

// ---------------------------------------------------------------------------
// Ship name constants
// ---------------------------------------------------------------------------

var FederationShipNames = [9]string{
	"Lexington", "Intrepid", "Enterprise", "Hornet", "Constitution",
	"Saratoga", "Ranger", "Yorktown", "Wasp",
}

var EmpireShipNames = [9]string{
	"Manta", "Cobra", "Hawk", "Condor", "Dragon",
	"Scorpion", "Vulture", "Raptor", "Viper",
}

// ---------------------------------------------------------------------------
// Session configuration
// ---------------------------------------------------------------------------

// SessionConfig holds immutable game options set by the first player.
type SessionConfig struct {
	GameMode          GameMode
	TournamentName    string
	TournamentSeed    int64
	RomulansEnabled   bool
	BlackHolesEnabled bool
	RespawnWaitTicks  int32
}

// ApplyDefaults fills in server-side defaults for zero-value fields.
func (c *SessionConfig) ApplyDefaults() {
	if c.GameMode == GameModeUnspecified {
		c.GameMode = GameModeRegular
	}
	if !c.RomulansEnabled {
		c.RomulansEnabled = true
	}
	if c.RespawnWaitTicks == 0 {
		c.RespawnWaitTicks = 120
	}
	if c.GameMode == GameModeTournament && c.TournamentName != "" {
		c.TournamentSeed = DeriveTournamentSeed(c.TournamentName)
	}
}

// DeriveTournamentSeed produces a deterministic int64 seed from a name using
// FNV-64a hashing so the raw seed is never accepted from clients.
func DeriveTournamentSeed(name string) int64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(strings.ToLower(strings.TrimSpace(name))))
	return int64(h.Sum64())
}

// ---------------------------------------------------------------------------
// Slot and player types
// ---------------------------------------------------------------------------

// ShipSlot represents one of 9 ship positions on a side.
type ShipSlot struct {
	SlotIndex   int
	Side        Side
	ShipName    string
	PlayerName  string
	PlayerState PlayerState
}

// playerInfo tracks per-player metadata not visible in ShipSlot.
type playerInfo struct {
	side          Side
	priorSide     Side
	shipSlotIndex int
}

// KillQueueEntry records a destroyed player waiting to respawn.
type KillQueueEntry struct {
	PlayerName          string
	Side                Side
	ShipName            string
	KilledAtStardate    float64
	EligibleAtStardate  float64
}

// ---------------------------------------------------------------------------
// GameSession — central domain object
// ---------------------------------------------------------------------------

// GameSession is the mutable state of one running game.
// Methods on GameSession are NOT goroutine-safe; callers must hold mu.
type GameSession struct {
	mu sync.Mutex // guards all fields below

	ID              string
	State           SessionState
	Config          SessionConfig
	CreatedAt       time.Time
	CurrentStardate float64
	WinningSide     Side
	FederationSlots [9]ShipSlot
	EmpireSlots     [9]ShipSlot
	KillQueue       []KillQueueEntry
	Gripes          []string

	players map[string]*playerInfo // player_name → internal info
}

// NewSession creates a session in WAITING state with slots pre-populated.
func NewSession(id string, config SessionConfig) *GameSession {
	gs := &GameSession{
		ID:        id,
		State:     SessionStateWaiting,
		Config:    config,
		CreatedAt: time.Now(),
		players:   make(map[string]*playerInfo),
	}
	for i := 0; i < 9; i++ {
		gs.FederationSlots[i] = ShipSlot{
			SlotIndex:   i,
			Side:        SideFederation,
			ShipName:    FederationShipNames[i],
			PlayerState: PlayerStateUnoccupied,
		}
		gs.EmpireSlots[i] = ShipSlot{
			SlotIndex:   i,
			Side:        SideEmpire,
			ShipName:    EmpireShipNames[i],
			PlayerState: PlayerStateUnoccupied,
		}
	}
	return gs
}

// slots returns a pointer to the slot array for the given side.
func (gs *GameSession) slots(side Side) *[9]ShipSlot {
	if side == SideEmpire {
		return &gs.EmpireSlots
	}
	return &gs.FederationSlots
}

// CountActive returns the number of ACTIVE players on a side.
func (gs *GameSession) CountActive(side Side) int {
	n := 0
	for _, s := range *gs.slots(side) {
		if s.PlayerState == PlayerStateActive {
			n++
		}
	}
	return n
}

// CountLobby returns the number of LOBBY players on a side.
func (gs *GameSession) CountLobby(side Side) int {
	n := 0
	for _, s := range *gs.slots(side) {
		if s.PlayerState == PlayerStateLobby {
			n++
		}
	}
	return n
}

// CountOccupied returns the number of slots not in UNOCCUPIED state on a side.
func (gs *GameSession) CountOccupied(side Side) int {
	n := 0
	for _, s := range *gs.slots(side) {
		if s.PlayerState != PlayerStateUnoccupied {
			n++
		}
	}
	return n
}

// LookupPriorSide returns the side a returning player was on, if known.
func (gs *GameSession) LookupPriorSide(playerName string) (Side, bool) {
	if info, ok := gs.players[playerName]; ok {
		return info.priorSide, true
	}
	return SideUnspecified, false
}

// AddPlayerToLobby assigns the first available UNOCCUPIED slot on side to the
// player and transitions it to LOBBY.  Returns an error if the side is full.
func (gs *GameSession) AddPlayerToLobby(playerName string, side Side) error {
	if _, exists := gs.players[playerName]; exists {
		return fmt.Errorf("player %q is already in this session", playerName)
	}
	ss := gs.slots(side)
	for i := range ss {
		if ss[i].PlayerState == PlayerStateUnoccupied {
			ss[i].PlayerName = playerName
			ss[i].PlayerState = PlayerStateLobby
			gs.players[playerName] = &playerInfo{
				side:          side,
				priorSide:     side,
				shipSlotIndex: i,
			}
			return nil
		}
	}
	return fmt.Errorf("side %s is full (9/9 slots occupied)", side)
}

// RemovePlayer transitions the player's slot back to UNOCCUPIED and removes
// internal tracking.  It also removes any kill-queue entry for the player.
func (gs *GameSession) RemovePlayer(playerName string) {
	if info, ok := gs.players[playerName]; ok {
		ss := gs.slots(info.side)
		ss[info.shipSlotIndex] = ShipSlot{
			SlotIndex:   info.shipSlotIndex,
			Side:        info.side,
			ShipName:    gs.shipNameForSide(info.side, info.shipSlotIndex),
			PlayerState: PlayerStateUnoccupied,
		}
		delete(gs.players, playerName)
	}
	// Remove kill-queue entry if present.
	for i, e := range gs.KillQueue {
		if e.PlayerName == playerName {
			gs.KillQueue = append(gs.KillQueue[:i], gs.KillQueue[i+1:]...)
			break
		}
	}
}

func (gs *GameSession) shipNameForSide(side Side, index int) string {
	if side == SideEmpire {
		return EmpireShipNames[index]
	}
	return FederationShipNames[index]
}

// AssignShip transitions the player from LOBBY to ACTIVE.  If preferredName is
// non-empty and that ship slot is free (or belongs to the player already), they
// get that ship; otherwise they keep their current lobby slot.
func (gs *GameSession) AssignShip(playerName, preferredName string) (string, error) {
	info, ok := gs.players[playerName]
	if !ok {
		return "", fmt.Errorf("player %q not in session", playerName)
	}
	ss := gs.slots(info.side)
	if ss[info.shipSlotIndex].PlayerState != PlayerStateLobby {
		return "", fmt.Errorf("player %q is not in lobby state", playerName)
	}

	// If a preferred ship was requested and is available, move the player.
	if preferredName != "" {
		for i := range ss {
			if strings.EqualFold(ss[i].ShipName, preferredName) &&
				(ss[i].PlayerState == PlayerStateUnoccupied ||
					(ss[i].PlayerName == playerName && ss[i].PlayerState == PlayerStateLobby)) {
				if i != info.shipSlotIndex {
					// Vacate current slot.
					ss[info.shipSlotIndex] = ShipSlot{
						SlotIndex:   info.shipSlotIndex,
						Side:        info.side,
						ShipName:    gs.shipNameForSide(info.side, info.shipSlotIndex),
						PlayerState: PlayerStateUnoccupied,
					}
					info.shipSlotIndex = i
					ss[i].PlayerName = playerName
					ss[i].PlayerState = PlayerStateLobby
				}
				break
			}
		}
	}

	ss[info.shipSlotIndex].PlayerState = PlayerStateActive
	return ss[info.shipSlotIndex].ShipName, nil
}

// KillPlayer transitions the player from ACTIVE to KILL_QUEUE and records the
// entry.  Returns an error if the player is not in ACTIVE state.
func (gs *GameSession) KillPlayer(playerName string) error {
	info, ok := gs.players[playerName]
	if !ok {
		return fmt.Errorf("player %q not in session", playerName)
	}
	ss := gs.slots(info.side)
	if ss[info.shipSlotIndex].PlayerState != PlayerStateActive {
		return fmt.Errorf("player %q is not in active state", playerName)
	}
	ss[info.shipSlotIndex].PlayerState = PlayerStateKillQueue

	entry := KillQueueEntry{
		PlayerName:         playerName,
		Side:               info.side,
		ShipName:           ss[info.shipSlotIndex].ShipName,
		KilledAtStardate:   gs.CurrentStardate,
		EligibleAtStardate: gs.CurrentStardate + float64(gs.Config.RespawnWaitTicks),
	}
	gs.enqueueKill(entry)
	return nil
}

// enqueueKill adds an entry to the kill queue, capped at 10 entries by evicting
// the oldest entry when the queue is full.
func (gs *GameSession) enqueueKill(entry KillQueueEntry) {
	if len(gs.KillQueue) >= 10 {
		gs.KillQueue = gs.KillQueue[1:]
	}
	gs.KillQueue = append(gs.KillQueue, entry)
}

// RespawnToLobby moves a player from KILL_QUEUE back to LOBBY state.
// Returns an error if the player is not in the kill queue or is not yet eligible.
func (gs *GameSession) RespawnToLobby(playerName string) error {
	idx := -1
	for i, e := range gs.KillQueue {
		if e.PlayerName == playerName {
			idx = i
			break
		}
	}
	if idx < 0 {
		return fmt.Errorf("player %q is not in the kill queue", playerName)
	}
	entry := gs.KillQueue[idx]
	if entry.EligibleAtStardate > gs.CurrentStardate {
		remaining := entry.EligibleAtStardate - gs.CurrentStardate
		return fmt.Errorf("player %q not yet eligible: %.0f ticks remaining", playerName, remaining)
	}

	// Remove from kill queue.
	gs.KillQueue = append(gs.KillQueue[:idx], gs.KillQueue[idx+1:]...)

	// Transition the slot back to LOBBY.
	info := gs.players[playerName]
	ss := gs.slots(info.side)
	ss[info.shipSlotIndex].PlayerState = PlayerStateLobby
	return nil
}

// CheckVictory returns true and the winning side if all ships on a side are
// either unoccupied or in the kill queue (i.e., no active or lobby players
// remain on that side while the other side has at least one active player).
// Returns false if no victory condition is met.
func (gs *GameSession) CheckVictory() (bool, Side) {
	if gs.State != SessionStateActive {
		return false, SideUnspecified
	}
	fedActive := gs.CountActive(SideFederation) + gs.CountLobby(SideFederation)
	empActive := gs.CountActive(SideEmpire) + gs.CountLobby(SideEmpire)

	if fedActive == 0 && empActive > 0 {
		return true, SideEmpire
	}
	if empActive == 0 && fedActive > 0 {
		return true, SideFederation
	}
	return false, SideUnspecified
}

// ---------------------------------------------------------------------------
// Lobby commands
// ---------------------------------------------------------------------------

// LobbyHelp returns the list of lobby commands.
func (gs *GameSession) LobbyHelp() string {
	return `Available lobby commands:
  help    - this help text
  news    - game news
  users   - show all connected players
  time    - display current stardate
  points  - show scoring summary
  summary - session overview
  gripe   - register a complaint
  quit    - leave the session`
}

// LobbyUsers returns a formatted list of all players in the session.
func (gs *GameSession) LobbyUsers() string {
	var sb strings.Builder
	sb.WriteString("Federation ships:\n")
	for _, s := range gs.FederationSlots {
		if s.PlayerState != PlayerStateUnoccupied {
			sb.WriteString(fmt.Sprintf("  %-15s  %-12s  %s\n", s.ShipName, s.PlayerName, playerStateLabel(s.PlayerState)))
		}
	}
	sb.WriteString("Empire ships:\n")
	for _, s := range gs.EmpireSlots {
		if s.PlayerState != PlayerStateUnoccupied {
			sb.WriteString(fmt.Sprintf("  %-15s  %-12s  %s\n", s.ShipName, s.PlayerName, playerStateLabel(s.PlayerState)))
		}
	}
	return sb.String()
}

func playerStateLabel(ps PlayerState) string {
	switch ps {
	case PlayerStateLobby:
		return "lobby"
	case PlayerStateActive:
		return "active"
	case PlayerStateKillQueue:
		return "killed"
	default:
		return "unknown"
	}
}

// LobbyTime returns the current simulated stardate.
func (gs *GameSession) LobbyTime() string {
	return fmt.Sprintf("Current stardate: %.2f", gs.CurrentStardate)
}

// LobbyPoints returns a brief scoring summary.
func (gs *GameSession) LobbyPoints() string {
	var sb strings.Builder
	fedActive := gs.CountActive(SideFederation)
	empActive := gs.CountActive(SideEmpire)
	fedLobby := gs.CountLobby(SideFederation)
	empLobby := gs.CountLobby(SideEmpire)
	sb.WriteString(fmt.Sprintf("Federation: %d active, %d in lobby\n", fedActive, fedLobby))
	sb.WriteString(fmt.Sprintf("Empire:     %d active, %d in lobby\n", empActive, empLobby))
	sb.WriteString(fmt.Sprintf("Kills in queue: %d\n", len(gs.KillQueue)))
	return sb.String()
}

// LobbyNews returns current news (static for now).
func (gs *GameSession) LobbyNews() string {
	return "No news. The galaxy is quiet. For now."
}

// LobbySummary returns an overview of the session configuration.
func (gs *GameSession) LobbySummary() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Session: %s\n", gs.ID))
	sb.WriteString(fmt.Sprintf("State:   %s\n", gs.State))
	sb.WriteString(fmt.Sprintf("Mode:    %s\n", gameModeLabel(gs.Config.GameMode)))
	if gs.Config.GameMode == GameModeTournament {
		sb.WriteString(fmt.Sprintf("Tournament: %s\n", gs.Config.TournamentName))
	}
	sb.WriteString(fmt.Sprintf("Romulans: %v  Black holes: %v\n", gs.Config.RomulansEnabled, gs.Config.BlackHolesEnabled))
	sb.WriteString(fmt.Sprintf("Respawn wait: %d ticks\n", gs.Config.RespawnWaitTicks))
	return sb.String()
}

func gameModeLabel(gm GameMode) string {
	switch gm {
	case GameModeRegular:
		return "REGULAR"
	case GameModeTournament:
		return "TOURNAMENT"
	default:
		return "UNSPECIFIED"
	}
}

// LobbyGripe records a complaint and returns an acknowledgement.
func (gs *GameSession) LobbyGripe(text string) string {
	text = strings.TrimSpace(text)
	if text == "" {
		return "Gripe text is required."
	}
	gs.Gripes = append(gs.Gripes, text)
	return fmt.Sprintf("Gripe recorded: %q", text)
}
