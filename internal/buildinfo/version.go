package buildinfo

import "strings"

var Version = "dev"

func EffectiveVersion() string {
	version := strings.TrimSpace(Version)
	if version == "" {
		return "dev"
	}

	return version
}