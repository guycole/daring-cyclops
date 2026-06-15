package ping

import (
	"fmt"
	"time"
)

type Result struct {
	ServerVersion string
	ServerTime    time.Time
}

func Format(result Result) string {
	return fmt.Sprintf(
		"Server Version: %s\nServer Time: %s\n",
		result.ServerVersion,
		result.ServerTime.UTC().Format(time.RFC3339Nano),
	)
}
