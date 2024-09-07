// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/guycole/daring-cyclops/grpc/gen/cyclops/v1/cyclopsv1connect"

	shared "github.com/guycole/daring-cyclops/shared"
)

type AppType struct {
	Ft           *FacadeType
	FeatureFlags uint32 // control run time features
	GrpcPort     int
	SugarLog     *zap.SugaredLogger
}

func (at *AppType) Initialize(featureFlags string) {
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

	at.Ft, err = newFacade(at.FeatureFlags, at.SugarLog)
	if err != nil {
		at.SugarLog.Fatal("newFacade failure")
	}
}

// Run pacifier
func (at *AppType) Run(address string) {
	at.SugarLog.Info("run run run")

	mux := http.NewServeMux()
	mux.Handle(cyclopsv1connect.NewCyclopsServiceHandler(&cyclopsServiceServer{Ft: at.Ft}))
	err := http.ListenAndServe(address, h2c.NewHandler(mux, &http2.Server{}))
	at.SugarLog.Fatalf("listen failure: %v", err)
}
