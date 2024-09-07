package app

const (
	developmentModeLoggingMask = 0x01
)

// enable development mode logging
func isDevelopmentModeLogging(featureFlags uint32) bool {
	return (featureFlags & developmentModeLoggingMask) != 0
}
