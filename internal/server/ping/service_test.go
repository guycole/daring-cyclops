package ping

import (
	"context"
	"testing"
	"time"
)

func TestServicePingReturnsVersionAndTime(t *testing.T) {
	fixedTime := time.Date(2026, time.June, 15, 12, 0, 0, 0, time.UTC)
	service := NewService("1.2.3", func() time.Time { return fixedTime })

	response, err := service.Ping(context.Background(), nil)
	if err != nil {
		t.Fatalf("Ping returned error: %v", err)
	}

	if response.GetServerVersion() != "1.2.3" {
		t.Fatalf("server version = %q, want %q", response.GetServerVersion(), "1.2.3")
	}

	if response.GetServerTime() == nil {
		t.Fatal("server time is nil")
	}

	if got := response.GetServerTime().AsTime(); !got.Equal(fixedTime) {
		t.Fatalf("server time = %s, want %s", got, fixedTime)
	}
}
