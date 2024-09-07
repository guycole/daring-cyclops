// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"flag"
	"math/rand"
	"os"
	"time"

	server "github.com/guycole/daring-cyclops/server"
	shared "github.com/guycole/daring-cyclops/shared"
)

const banner = "daring-cyclops 0.0"

func main() {
	flag.Parse()

	sugarLog := shared.ZapSetup(false)
	sugarLog.Info(banner)

	rand.New(rand.NewSource(time.Now().UnixNano()))

	var featureFlags, grpcAddress, runMode string

	envVars := [...]string{"FEATURE_FLAGS", "GRPC_ADDRESS", "RUN_MODE"}

	for index, element := range envVars {
		temp, err := os.LookupEnv(element)
		if err {
			sugarLog.Infof("%d:%s:%s", index, element, temp)
		} else {
			sugarLog.Fatal("missing:", element)
		}

		switch element {
		case "FEATURE_FLAGS":
			featureFlags = temp
		case "GRPC_ADDRESS":
			grpcAddress = temp
		case "RUN_MODE":
			runMode = temp
		default:
			sugarLog.Fatal("unknown environment var:", element)
		}
	}

	app := server.AppType{SugarLog: sugarLog}
	app.Initialize(featureFlags)
	app.Run(grpcAddress, runMode)
}
