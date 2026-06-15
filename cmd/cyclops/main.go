package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	clientcfg "github.com/guycole/daring-cyclops/internal/client"
	pingcmd "github.com/guycole/daring-cyclops/internal/client/ping"
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
	default:
		fmt.Fprintf(stderr, "unknown command %q\n", args[0])
		printUsage(stderr)
		return 1
	}
}

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

func printUsage(output io.Writer) {
	fmt.Fprintln(output, "usage: cyclops ping [-server address] [-timeout duration]")
}
