// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package shared

const (
	developmentModeLoggingMask = 0x01
)

// enable development mode logging
func IsDevelopmentModeLogging(featureFlags uint32) bool {
	return (featureFlags & developmentModeLoggingMask) != 0
}
