// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"context"

	"connectrpc.com/connect"

	v1 "github.com/guycole/daring-cyclops/grpc/gen/cyclops/v1"
	"github.com/guycole/daring-cyclops/grpc/gen/cyclops/v1/cyclopsv1connect"

	shared "github.com/guycole/daring-cyclops/shared"
)

// buf curl --schema proto --data '{"source":"bufcurl"}' http://localhost:8080/cyclops.v1.CyclopsService/Ping
// curl -v --header "Content-Type: application/json" --data '{"source":"curl"}' http://localhost:8080/cyclops.v1.CyclopsService/Ping
// curl -v --header "Content-Type: application/json" --data '{"source":"curl"}' http://192.168.1.102:8080/cyclops.v1.CyclopsService/Ping

type cyclopsServiceServer struct {
	Ft *FacadeType

	cyclopsv1connect.UnimplementedCyclopsServiceHandler
}

func (css *cyclopsServiceServer) Ping(ctx context.Context, req *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error) {
	css.Ft.SugarLog.Debug("ping:", req.Msg.GetSource())
	return connect.NewResponse(&v1.PingResponse{Retcode: shared.RetCodeSuccess}), nil
}
