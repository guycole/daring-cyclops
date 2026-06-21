package session

import (
	"context"
	"fmt"

	sessionv1 "github.com/guycole/daring-cyclops/gen/proto/session/v1"
	clientcfg "github.com/guycole/daring-cyclops/internal/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ---------------------------------------------------------------------------
// Option types
// ---------------------------------------------------------------------------

// CreateOptions parameterises a CreateSession call.
type CreateOptions struct {
	PlayerName        string
	GameMode          string // "regular" or "tournament"
	TournamentName    string
	RomulansEnabled   bool
	BlackHolesEnabled bool
	RespawnWaitTicks  int32
}

// ListOptions parameterises a ListSessions call.
type ListOptions struct {
	IncludeEnded bool
}

// JoinOptions parameterises a JoinSession call.
type JoinOptions struct {
	SessionID     string
	PlayerName    string
	PreferredSide string // "federation" or "empire"
}

// LobbyOptions parameterises a LobbyCommand call.
type LobbyOptions struct {
	SessionID  string
	PlayerName string
	Command    string // "help", "news", "users", "time", "points", "summary", "gripe", "quit"
	Argument   string
}

// ActivateOptions parameterises an ActivateShip call.
type ActivateOptions struct {
	SessionID         string
	PlayerName        string
	PreferredShipName string
}

// KillQueueOptions parameterises a GetKillQueueStatus call.
type KillQueueOptions struct {
	SessionID  string
	PlayerName string
}

// StatusOptions parameterises a GetSessionStatus call.
type StatusOptions struct {
	SessionID string
}

// ---------------------------------------------------------------------------
// Dialer helper
// ---------------------------------------------------------------------------

func dial(ctx context.Context, cfg clientcfg.Config) (*grpc.ClientConn, context.Context, context.CancelFunc, error) {
	conn, err := grpc.NewClient(
		cfg.ServerAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("dial %s: %w", cfg.ServerAddress, err)
	}
	rpcCtx, cancel := context.WithTimeout(ctx, cfg.Timeout)
	return conn, rpcCtx, cancel, nil
}

// ---------------------------------------------------------------------------
// RPC functions
// ---------------------------------------------------------------------------

// RunCreate calls CreateSession.
func RunCreate(ctx context.Context, cfg clientcfg.Config, opts CreateOptions) (*sessionv1.CreateSessionResponse, error) {
	conn, rpcCtx, cancel, err := dial(ctx, cfg)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	defer cancel()

	client := sessionv1.NewSessionServiceClient(conn)
	return client.CreateSession(rpcCtx, &sessionv1.CreateSessionRequest{
		PlayerName: opts.PlayerName,
		Config:     buildProtoConfig(opts),
	})
}

// RunList calls ListSessions.
func RunList(ctx context.Context, cfg clientcfg.Config, opts ListOptions) (*sessionv1.ListSessionsResponse, error) {
	conn, rpcCtx, cancel, err := dial(ctx, cfg)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	defer cancel()

	client := sessionv1.NewSessionServiceClient(conn)
	return client.ListSessions(rpcCtx, &sessionv1.ListSessionsRequest{
		IncludeEnded: opts.IncludeEnded,
	})
}

// RunJoin calls JoinSession.
func RunJoin(ctx context.Context, cfg clientcfg.Config, opts JoinOptions) (*sessionv1.JoinSessionResponse, error) {
	conn, rpcCtx, cancel, err := dial(ctx, cfg)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	defer cancel()

	client := sessionv1.NewSessionServiceClient(conn)
	return client.JoinSession(rpcCtx, &sessionv1.JoinSessionRequest{
		SessionId:     opts.SessionID,
		PlayerName:    opts.PlayerName,
		PreferredSide: parseSide(opts.PreferredSide),
	})
}

// RunLobby calls LobbyCommand.
func RunLobby(ctx context.Context, cfg clientcfg.Config, opts LobbyOptions) (*sessionv1.LobbyCommandResponse, error) {
	conn, rpcCtx, cancel, err := dial(ctx, cfg)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	defer cancel()

	client := sessionv1.NewSessionServiceClient(conn)
	return client.LobbyCommand(rpcCtx, &sessionv1.LobbyCommandRequest{
		SessionId:  opts.SessionID,
		PlayerName: opts.PlayerName,
		Command:    parseLobbyCommand(opts.Command),
		Argument:   opts.Argument,
	})
}

// RunActivate calls ActivateShip.
func RunActivate(ctx context.Context, cfg clientcfg.Config, opts ActivateOptions) (*sessionv1.ActivateShipResponse, error) {
	conn, rpcCtx, cancel, err := dial(ctx, cfg)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	defer cancel()

	client := sessionv1.NewSessionServiceClient(conn)
	return client.ActivateShip(rpcCtx, &sessionv1.ActivateShipRequest{
		SessionId:         opts.SessionID,
		PlayerName:        opts.PlayerName,
		PreferredShipName: opts.PreferredShipName,
	})
}

// RunKillQueue calls GetKillQueueStatus.
func RunKillQueue(ctx context.Context, cfg clientcfg.Config, opts KillQueueOptions) (*sessionv1.GetKillQueueStatusResponse, error) {
	conn, rpcCtx, cancel, err := dial(ctx, cfg)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	defer cancel()

	client := sessionv1.NewSessionServiceClient(conn)
	return client.GetKillQueueStatus(rpcCtx, &sessionv1.GetKillQueueStatusRequest{
		SessionId:  opts.SessionID,
		PlayerName: opts.PlayerName,
	})
}

// RunStatus calls GetSessionStatus.
func RunStatus(ctx context.Context, cfg clientcfg.Config, opts StatusOptions) (*sessionv1.GetSessionStatusResponse, error) {
	conn, rpcCtx, cancel, err := dial(ctx, cfg)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	defer cancel()

	client := sessionv1.NewSessionServiceClient(conn)
	return client.GetSessionStatus(rpcCtx, &sessionv1.GetSessionStatusRequest{
		SessionId: opts.SessionID,
	})
}

// ---------------------------------------------------------------------------
// Conversion helpers
// ---------------------------------------------------------------------------

func buildProtoConfig(opts CreateOptions) *sessionv1.SessionConfig {
	gm := sessionv1.GameMode_GAME_MODE_REGULAR
	if opts.GameMode == "tournament" {
		gm = sessionv1.GameMode_GAME_MODE_TOURNAMENT
	}
	return &sessionv1.SessionConfig{
		GameMode:          gm,
		TournamentName:    opts.TournamentName,
		RomulansEnabled:   opts.RomulansEnabled,
		BlackHolesEnabled: opts.BlackHolesEnabled,
		RespawnWaitTicks:  opts.RespawnWaitTicks,
	}
}

func parseSide(s string) sessionv1.Side {
	switch s {
	case "empire":
		return sessionv1.Side_SIDE_EMPIRE
	default:
		return sessionv1.Side_SIDE_FEDERATION
	}
}

func parseLobbyCommand(s string) sessionv1.LobbyCommand {
	switch s {
	case "help":
		return sessionv1.LobbyCommand_LOBBY_COMMAND_HELP
	case "news":
		return sessionv1.LobbyCommand_LOBBY_COMMAND_NEWS
	case "users":
		return sessionv1.LobbyCommand_LOBBY_COMMAND_USERS
	case "time":
		return sessionv1.LobbyCommand_LOBBY_COMMAND_TIME
	case "points":
		return sessionv1.LobbyCommand_LOBBY_COMMAND_POINTS
	case "summary":
		return sessionv1.LobbyCommand_LOBBY_COMMAND_SUMMARY
	case "gripe":
		return sessionv1.LobbyCommand_LOBBY_COMMAND_GRIPE
	case "quit":
		return sessionv1.LobbyCommand_LOBBY_COMMAND_QUIT
	default:
		return sessionv1.LobbyCommand_LOBBY_COMMAND_UNSPECIFIED
	}
}
