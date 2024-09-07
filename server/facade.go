// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"go.uber.org/zap"
)

type FacadeType struct {
	FeatureFlags uint32
	SugarLog     *zap.SugaredLogger
}

func newFacade(featureFlags uint32, sugarLog *zap.SugaredLogger) (*FacadeType, error) {
	return &FacadeType{FeatureFlags: featureFlags, SugarLog: sugarLog}, nil
}
