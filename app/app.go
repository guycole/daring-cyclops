// Copyright 2023 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package app

import (
	"strconv"

	"go.uber.org/zap"
)

type AppType struct {
	FeatureFlags uint32             // control run time features
	GrpcPort     int                // gRPC port
	SugarLog     *zap.SugaredLogger // logging
}

func (at *AppType) Initialize(featureFlags, grpcPort string) {
	temp, err := strconv.Atoi(featureFlags)
	if err == nil {
		at.SugarLog.Infof("featureFlags: %x", temp)
		at.FeatureFlags = uint32(temp)
	} else {
		at.SugarLog.Fatal("bad featureFlags")
	}

	if isDevelopmentModeLogging(at.FeatureFlags) {
		at.SugarLog = ZapSetup(true)
		at.SugarLog.Debug("debug level log entry")
	}

}

// Run pacifier
func (at *AppType) Run() {
	at.SugarLog.Info("run run run")
}
