package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	clientcfg "github.com/guycole/daring-cyclops/internal/client"
	pingcmd "github.com/guycole/daring-cyclops/internal/client/ping"
	sessioncmd "github.com/guycole/daring-cyclops/internal/client/session"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printUsage(stderr)
		return 1
	}

	switch args[0] {
	case "ping":
		return runPing(args[1:], stdout, stderr)
	case "session":
		return runSession(args[1:], stdout, stderr)
	default:
		fmt.Fprintf(stderr, "unknown command %q\n", args[0])
		printUsage(stderr)
		return 1
	}
}

// ---------------------------------------------------------------------------
// ping command
// ---------------------------------------------------------------------------

func runPing(args []string, stdout, stderr io.Writer) int {
	flagSet := flag.NewFlagSet("ping", flag.ContinueOnError)
	flagSet.SetOutput(stderr)

	serverAddress := clientcfg.DefaultServerAddress
	timeout := clientcfg.DefaultTimeout

	flagSet.StringVar(&serverAddress, "server", serverAddress, "gRPC server address")
	flagSet.DurationVar(&timeout, "timeout", timeout, "ping timeout")

	if err := flagSet.Parse(args); err != nil {
		return 1
	}

	if flagSet.NArg() != 0 {
		fmt.Fprintln(stderr, "ping does not accept positional arguments")
		return 1
	}

	result, err := pingcmd.Run(context.Background(), clientcfg.Config{
		ServerAddress: serverAddress,
		Timeout:       timeout,
	})
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}

	fmt.Fprint(stdout, pingcmd.Format(result))
	return 0
}

// ---------------------------------------------------------------------------
// session command dispatcher
// ---------------------------------------------------------------------------

func runSession(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printSessionUsage(stderr)
		return 1
	}

	switch args[0] {
	case "create":
		return runSessionCreate(args[1:], stdout, stderr)
	case "list":
		return runSessionList(args[1:], stdout, stderr)
	case "join":
		return runSessionJoin(args[1:], stdout, stderr)
	case "lobby":
		return runSessionLobby(args[1:], stdout, stderr)
	case "activate":
		return runSessionActivate(args[1:], stdout, stderr)
	case "kill-queue":
		return runSessionKillQueue(args[1:], stdout, stderr)
	case "status":
		return runSessionStatus(args[1:], stdout, stderr)
	default:
		fmt.Fprintf(stderr, "unknown session sub-command %q\n", args[0])
		printSessionUsage(stderr)
		return 1
	}
}

func commonSessionFlags(fs *flag.FlagSet) *clientcfg.Config {
	cfg := &clientcfg.Config{
		ServerAddress: clientcfg.DefaultServerAddress,
		Timeout:       clientcfg.DefaultTimeout,
	}
	fs.StringVar(&cfg.ServerAddress, "server", cfg.ServerAddress, "gRPC server address")
	fs.DurationVar(&cfg.Timeout, "timeout", cfg.Timeout, "RPC timeout")
	return cfg
}

// ---------------------------------------------------------------------------
// session create
// ---------------------------------------------------------------------------

func runSessionCreate(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("session create", flag.ContinueOnError)
	fs.SetOutput(stderr)
	cfg := commonSessionFlags(fs)

	player := fs.String("player", "", "player name (required)")
	mode := fs.String("mode", "regular", "game mode: regular or tournament")
	tournament := fs.String("tournament", "", "tournament name (required for tournament mode)")
	romulans := fs.Bool("romulans", true, "enable Romulan NPC")
	blackholes := fs.Bool("blackholes", false, "enable black holes")
	respawn := fs.Int("respawn", 120, "respawn wait ticks")

	if err := fs.Parse(args); err != nil {
		return 1
	}
	if *player == "" {
		fmt.Fprintln(stderr, "error: -player is required")
		return 1
	}

	resp, err := sessioncmd.RunCreate(context.Background(), *cfg, sessioncmd.CreateOptions{
		PlayerName:        *player,
		GameMode:          *mode,
		TournamentName:    *tournament,
		RomulansEnabled:   *romulans,
		BlackHolesEnabled: *blackholes,
		RespawnWaitTicks:  int32(*respawn),
	})
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	fmt.Fprint(stdout, sessioncmd.FormatCreate(resp))
	return 0
}

// ---------------------------------------------------------------------------
// session list
// ---------------------------------------------------------------------------

func runSessionList(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("session list", flag.ContinueOnError)
	fs.SetOutput(stderr)
	cfg := commonSessionFlags(fs)
	includeEnded := fs.Bool("ended", false, "include ended sessions")

	if err := fs.Parse(args); err != nil {
		return 1
	}

	resp, err := sessioncmd.RunList(context.Background(), *cfg, sessioncmd.ListOptions{
		IncludeEnded: *includeEnded,
	})
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	fmt.Fprint(stdout, sessioncmd.FormatList(resp))
	return 0
}

// ---------------------------------------------------------------------------
// session join
// ---------------------------------------------------------------------------

func runSessionJoin(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("session join", flag.ContinueOnError)
	fs.SetOutput(stderr)
	cfg := commonSessionFlags(fs)

	sessionID := fs.String("session", "", "session ID (required)")
	player := fs.String("player", "", "player name (required)")
	side := fs.String("side", "federation", "preferred side: federation or empire")

	if err := fs.Parse(args); err != nil {
		return 1
	}
	if *sessionID == "" {
		fmt.Fprintln(stderr, "error: -session is required")
		return 1
	}
	if *player == "" {
		fmt.Fprintln(stderr, "error: -player is required")
		return 1
	}

	resp, err := sessioncmd.RunJoin(context.Background(), *cfg, sessioncmd.JoinOptions{
		SessionID:     *sessionID,
		PlayerName:    *player,
		PreferredSide: *side,
	})
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	fmt.Fprint(stdout, sessioncmd.FormatJoin(resp))
	return 0
}

// ---------------------------------------------------------------------------
// session lobby
// ---------------------------------------------------------------------------

func runSessionLobby(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("session lobby", flag.ContinueOnError)
	fs.SetOutput(stderr)
	cfg := commonSessionFlags(fs)

	sessionID := fs.String("session", "", "session ID (required)")
	player := fs.String("player", "", "player name (required)")
	command := fs.String("cmd", "", "lobby command: help, news, users, time, points, summary, gripe, quit (required)")
	argument := fs.String("arg", "", "optional argument (e.g., gripe text)")

	if err := fs.Parse(args); err != nil {
		return 1
	}
	if *sessionID == "" {
		fmt.Fprintln(stderr, "error: -session is required")
		return 1
	}
	if *player == "" {
		fmt.Fprintln(stderr, "error: -player is required")
		return 1
	}
	if *command == "" {
		fmt.Fprintln(stderr, "error: -cmd is required")
		return 1
	}

	resp, err := sessioncmd.RunLobby(context.Background(), *cfg, sessioncmd.LobbyOptions{
		SessionID:  *sessionID,
		PlayerName: *player,
		Command:    strings.ToLower(*command),
		Argument:   *argument,
	})
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	fmt.Fprint(stdout, sessioncmd.FormatLobby(resp))
	return 0
}

// ---------------------------------------------------------------------------
// session activate
// ---------------------------------------------------------------------------

func runSessionActivate(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("session activate", flag.ContinueOnError)
	fs.SetOutput(stderr)
	cfg := commonSessionFlags(fs)

	sessionID := fs.String("session", "", "session ID (required)")
	player := fs.String("player", "", "player name (required)")
	ship := fs.String("ship", "", "preferred ship name (optional)")

	if err := fs.Parse(args); err != nil {
		return 1
	}
	if *sessionID == "" {
		fmt.Fprintln(stderr, "error: -session is required")
		return 1
	}
	if *player == "" {
		fmt.Fprintln(stderr, "error: -player is required")
		return 1
	}

	resp, err := sessioncmd.RunActivate(context.Background(), *cfg, sessioncmd.ActivateOptions{
		SessionID:         *sessionID,
		PlayerName:        *player,
		PreferredShipName: *ship,
	})
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	fmt.Fprint(stdout, sessioncmd.FormatActivate(resp))
	return 0
}

// ---------------------------------------------------------------------------
// session kill-queue
// ---------------------------------------------------------------------------

func runSessionKillQueue(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("session kill-queue", flag.ContinueOnError)
	fs.SetOutput(stderr)
	cfg := commonSessionFlags(fs)

	sessionID := fs.String("session", "", "session ID (required)")
	player := fs.String("player", "", "player name (required)")

	if err := fs.Parse(args); err != nil {
		return 1
	}
	if *sessionID == "" {
		fmt.Fprintln(stderr, "error: -session is required")
		return 1
	}
	if *player == "" {
		fmt.Fprintln(stderr, "error: -player is required")
		return 1
	}

	resp, err := sessioncmd.RunKillQueue(context.Background(), *cfg, sessioncmd.KillQueueOptions{
		SessionID:  *sessionID,
		PlayerName: *player,
	})
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	fmt.Fprint(stdout, sessioncmd.FormatKillQueue(resp))
	return 0
}

// ---------------------------------------------------------------------------
// session status
// ---------------------------------------------------------------------------

func runSessionStatus(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("session status", flag.ContinueOnError)
	fs.SetOutput(stderr)
	cfg := commonSessionFlags(fs)

	sessionID := fs.String("session", "", "session ID (required)")

	if err := fs.Parse(args); err != nil {
		return 1
	}
	if *sessionID == "" {
		fmt.Fprintln(stderr, "error: -session is required")
		return 1
	}

	resp, err := sessioncmd.RunStatus(context.Background(), *cfg, sessioncmd.StatusOptions{
		SessionID: *sessionID,
	})
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	fmt.Fprint(stdout, sessioncmd.FormatStatus(resp))
	return 0
}

// ---------------------------------------------------------------------------
// Usage
// ---------------------------------------------------------------------------

func printUsage(output io.Writer) {
	fmt.Fprintln(output, "usage: cyclops <command> [flags]")
	fmt.Fprintln(output, "commands: ping, session")
}

func printSessionUsage(output io.Writer) {
	fmt.Fprintln(output, "usage: cyclops session <sub-command> [flags]")
	fmt.Fprintln(output, "sub-commands: create, list, join, lobby, activate, kill-queue, status")
}
