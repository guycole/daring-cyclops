// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"net/http"
	"strconv"
	"strings"

	"go.uber.org/zap"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/guycole/daring-cyclops/grpc/gen/cyclops/v1/cyclopsv1connect"

	shared "github.com/guycole/daring-cyclops/shared"
)

type AppType struct {
	Ft           *facadeType
	FeatureFlags uint32 // control run time features
	SugarLog     *zap.SugaredLogger
}

func (at *AppType) Initialize(featureFlags, maxGames string) {
	temp, err := strconv.Atoi(featureFlags)
	if err == nil {
		at.SugarLog.Infof("featureFlags: %x", temp)
		at.FeatureFlags = uint32(temp)
	} else {
		at.SugarLog.Fatal("bad featureFlags")
	}

	if shared.IsDevelopmentModeLogging(at.FeatureFlags) {
		at.SugarLog = shared.ZapSetup(true)
		at.SugarLog.Debug("debug level log entry")
	}

	var maxGameLimit uint16

	temp, err = strconv.Atoi(maxGames)
	if err == nil {
		maxGameLimit = uint16(temp)
	} else {
		at.SugarLog.Fatal("bad maxGames")
	}

	gameManager := newGameManager(maxGameLimit, at.SugarLog)

	at.Ft = newFacade(at.FeatureFlags, gameManager, at.SugarLog)
}

// Run pacifier
func (at *AppType) Run(grpcAddress, runMode string) {
	if strings.Compare(runMode, "server") == 0 {
		at.SugarLog.Info("starting server mode")
		mux := http.NewServeMux()
		mux.Handle(cyclopsv1connect.NewCyclopsServiceHandler(&cyclopsServiceServer{ft: at.Ft}))
		err := http.ListenAndServe(grpcAddress, h2c.NewHandler(mux, &http2.Server{}))
		at.SugarLog.Fatalf("listen failure: %v", err)
	} else {
		at.SugarLog.Fatal("unsupported run mode")
	}
}
