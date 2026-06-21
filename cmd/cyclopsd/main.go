package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	pingv1 "github.com/guycole/daring-cyclops/gen/proto/ping/v1"
	"github.com/guycole/daring-cyclops/internal/buildinfo"
	pingserver "github.com/guycole/daring-cyclops/internal/server/ping"
	"google.golang.org/grpc"
)

const defaultListenAddress = "127.0.0.1:50051"

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	flagSet := flag.NewFlagSet("cyclopsd", flag.ContinueOnError)
	flagSet.SetOutput(os.Stderr)

	listenAddress := defaultListenAddress
	flagSet.StringVar(&listenAddress, "listen", listenAddress, "listen address")

	if err := flagSet.Parse(args); err != nil {
		return err
	}

	if flagSet.NArg() != 0 {
		return fmt.Errorf("unexpected positional arguments: %s", strings.Join(flagSet.Args(), " "))
	}

	listener, err := net.Listen("tcp", listenAddress)
	if err != nil {
		return fmt.Errorf("listen %s: %w", listenAddress, err)
	}

	grpcServer := grpc.NewServer()
	pingv1.RegisterPingServiceServer(grpcServer, pingserver.NewService(buildinfo.EffectiveVersion(), time.Now))

	if err := grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("serve grpc: %w", err)
	}

	return nil
}
