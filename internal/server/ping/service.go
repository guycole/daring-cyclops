package ping

import (
	"context"
	"time"

	pingv1 "github.com/guycole/daring-cyclops/gen/proto/ping/v1"
	"github.com/guycole/daring-cyclops/internal/buildinfo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	pingv1.UnimplementedPingServiceServer

	now     func() time.Time
	version string
}

func NewService(version string, now func() time.Time) *Service {
	if now == nil {
		now = time.Now
	}

	service := &Service{
		now:     now,
		version: version,
	}

	if service.version == "" {
		service.version = buildinfo.EffectiveVersion()
	}

	return service
}

func (service *Service) Ping(context.Context, *pingv1.PingRequest) (*pingv1.PingResponse, error) {
	response := &pingv1.PingResponse{
		ServerVersion: service.version,
		ServerTime:    timestamppb.New(service.now().UTC()),
	}

	if response.GetServerVersion() == "" {
		return nil, status.Error(codes.Internal, "ping response missing server version")
	}

	if response.GetServerTime() == nil || response.GetServerTime().CheckValid() != nil {
		return nil, status.Error(codes.Internal, "ping response missing valid server time")
	}

	return response, nil
}
