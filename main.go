// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"flag"
	"math/rand"
	"os"
	"time"

	appx "github.com/guycole/daring-cyclops/app"
)

const banner = "daring-cyclops 0.0"

func main() {
	flag.Parse()

	sugarLog := appx.ZapSetup(false)
	sugarLog.Info(banner)

	rand.New(rand.NewSource(time.Now().UnixNano()))

	var featureFlags, grpcPort string

	envVars := [...]string{"FEATURE_FLAGS", "GRPC_PORT"}

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
		case "GRPC_PORT":
			grpcPort = temp
		default:
			sugarLog.Fatal("unknown environment var:", element)
		}
	}

	app := appx.AppType{SugarLog: sugarLog}
	app.Initialize(featureFlags, grpcPort)
	app.Run()
}
