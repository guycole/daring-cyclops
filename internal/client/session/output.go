package session

import (
	"fmt"
	"strings"

	sessionv1 "github.com/guycole/daring-cyclops/gen/proto/session/v1"
)

// FormatCreate renders a CreateSessionResponse for terminal display.
func FormatCreate(resp *sessionv1.CreateSessionResponse) string {
	if resp == nil {
		return "no response\n"
	}
	cfg := resp.GetConfig()
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Session created: %s\n", resp.GetSessionId()))
	sb.WriteString(fmt.Sprintf("Mode: %s\n", gameModeLabel(cfg.GetGameMode())))
	if cfg.GetGameMode() == sessionv1.GameMode_GAME_MODE_TOURNAMENT {
		sb.WriteString(fmt.Sprintf("Tournament: %s\n", cfg.GetTournamentName()))
	}
	sb.WriteString(fmt.Sprintf("Romulans: %v  Black holes: %v\n", cfg.GetRomulansEnabled(), cfg.GetBlackHolesEnabled()))
	sb.WriteString(fmt.Sprintf("Respawn wait: %d ticks\n", cfg.GetRespawnWaitTicks()))
	return sb.String()
}

// FormatList renders a ListSessionsResponse for terminal display.
func FormatList(resp *sessionv1.ListSessionsResponse) string {
	if resp == nil {
		return "no response\n"
	}
	sessions := resp.GetSessions()
	if len(sessions) == 0 {
		return "No sessions found.\n"
	}
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%-36s  %-9s  %-4s  %-4s  %-4s  %-4s\n",
		"SESSION ID", "STATE", "FA", "EA", "FL", "EL"))
	sb.WriteString(strings.Repeat("-", 72) + "\n")
	for _, s := range sessions {
		sb.WriteString(fmt.Sprintf("%-36s  %-9s  %4d  %4d  %4d  %4d\n",
			s.GetSessionId(),
			sessionStateLabel(s.GetState()),
			s.GetFederationActive(),
			s.GetEmpireActive(),
			s.GetFederationLobby(),
			s.GetEmpireLobby(),
		))
	}
	return sb.String()
}

// FormatJoin renders a JoinSessionResponse for terminal display.
func FormatJoin(resp *sessionv1.JoinSessionResponse) string {
	if resp == nil {
		return "no response\n"
	}
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Joined session on the %s side.\n", sideLabel(resp.GetAssignedSide())))
	if resp.GetSideRedirected() {
		sb.WriteString(fmt.Sprintf("Note: you were redirected — %s\n", resp.GetRedirectReason()))
	}
	sb.WriteString(fmt.Sprintf("Romulans: %v  Black holes: %v\n", resp.GetRomulansEnabled(), resp.GetBlackHolesEnabled()))
	return sb.String()
}

// FormatLobby renders a LobbyCommandResponse for terminal display.
func FormatLobby(resp *sessionv1.LobbyCommandResponse) string {
	if resp == nil {
		return "no response\n"
	}
	out := resp.GetOutput()
	if !strings.HasSuffix(out, "\n") {
		out += "\n"
	}
	return out
}

// FormatActivate renders an ActivateShipResponse for terminal display.
func FormatActivate(resp *sessionv1.ActivateShipResponse) string {
	if resp == nil {
		return "no response\n"
	}
	return fmt.Sprintf("Ship activated: %s (%s side)  Session: %s\n",
		resp.GetShipName(),
		sideLabel(resp.GetSide()),
		sessionStateLabel(resp.GetSessionState()),
	)
}

// FormatKillQueue renders a GetKillQueueStatusResponse for terminal display.
func FormatKillQueue(resp *sessionv1.GetKillQueueStatusResponse) string {
	if resp == nil {
		return "no response\n"
	}
	if !resp.GetInKillQueue() {
		return fmt.Sprintf("Not in kill queue.  Current stardate: %.2f\n", resp.GetCurrentStardate())
	}
	return fmt.Sprintf("In kill queue.  Eligible at stardate %.2f  (%d ticks remaining)\n",
		resp.GetEligibleAtStardate(),
		resp.GetTicksRemaining(),
	)
}

// FormatStatus renders a GetSessionStatusResponse for terminal display.
func FormatStatus(resp *sessionv1.GetSessionStatusResponse) string {
	if resp == nil {
		return "no response\n"
	}
	st := resp.GetStatus()
	if st == nil {
		return "no status\n"
	}
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Session:  %s\n", st.GetSessionId()))
	sb.WriteString(fmt.Sprintf("State:    %s\n", sessionStateLabel(st.GetState())))
	sb.WriteString(fmt.Sprintf("Stardate: %.2f\n", st.GetCurrentStardate()))
	if st.GetState() == sessionv1.SessionState_SESSION_STATE_ENDED {
		sb.WriteString(fmt.Sprintf("Winner:   %s\n", sideLabel(st.GetWinningSide())))
	}
	sb.WriteString("\nFederation ships:\n")
	for _, slot := range st.GetFederationSlots() {
		if slot.GetPlayerState() != sessionv1.PlayerState_PLAYER_STATE_UNOCCUPIED {
			sb.WriteString(fmt.Sprintf("  %-15s  %-12s  %s\n",
				slot.GetShipName(), slot.GetPlayerName(), playerStateLabel(slot.GetPlayerState())))
		}
	}
	sb.WriteString("\nEmpire ships:\n")
	for _, slot := range st.GetEmpireSlots() {
		if slot.GetPlayerState() != sessionv1.PlayerState_PLAYER_STATE_UNOCCUPIED {
			sb.WriteString(fmt.Sprintf("  %-15s  %-12s  %s\n",
				slot.GetShipName(), slot.GetPlayerName(), playerStateLabel(slot.GetPlayerState())))
		}
	}
	if len(st.GetKillQueue()) > 0 {
		sb.WriteString("\nKill queue:\n")
		for _, e := range st.GetKillQueue() {
			sb.WriteString(fmt.Sprintf("  %-12s  %-15s  eligible at %.2f\n",
				e.GetPlayerName(), e.GetShipName(), e.GetEligibleAtStardate()))
		}
	}
	return sb.String()
}

// ---------------------------------------------------------------------------
// Label helpers
// ---------------------------------------------------------------------------

func gameModeLabel(gm sessionv1.GameMode) string {
	switch gm {
	case sessionv1.GameMode_GAME_MODE_TOURNAMENT:
		return "TOURNAMENT"
	default:
		return "REGULAR"
	}
}

func sessionStateLabel(s sessionv1.SessionState) string {
	switch s {
	case sessionv1.SessionState_SESSION_STATE_WAITING:
		return "WAITING"
	case sessionv1.SessionState_SESSION_STATE_ACTIVE:
		return "ACTIVE"
	case sessionv1.SessionState_SESSION_STATE_ENDED:
		return "ENDED"
	default:
		return "UNKNOWN"
	}
}

func sideLabel(s sessionv1.Side) string {
	switch s {
	case sessionv1.Side_SIDE_FEDERATION:
		return "FEDERATION"
	case sessionv1.Side_SIDE_EMPIRE:
		return "EMPIRE"
	default:
		return "UNKNOWN"
	}
}

func playerStateLabel(ps sessionv1.PlayerState) string {
	switch ps {
	case sessionv1.PlayerState_PLAYER_STATE_LOBBY:
		return "lobby"
	case sessionv1.PlayerState_PLAYER_STATE_ACTIVE:
		return "active"
	case sessionv1.PlayerState_PLAYER_STATE_KILL_QUEUE:
		return "killed"
	default:
		return "unknown"
	}
}
