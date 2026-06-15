package ping

import (
	"strings"
	"testing"
	"time"
)

func TestFormatIncludesLabelsAndValues(t *testing.T) {
	result := NewResult("1.2.3", time.Date(2026, time.June, 15, 12, 0, 0, 0, time.UTC))

	output := Format(result)

	if !strings.Contains(output, "Server Version: 1.2.3") {
		t.Fatalf("output missing version label: %q", output)
	}

	if !strings.Contains(output, "Server Time: 2026-06-15T12:00:00Z") {
		t.Fatalf("output missing time label: %q", output)
	}
	if !strings.HasSuffix(output, "\n") {
		t.Fatalf("output should end with newline: %q", output)
	}
}
