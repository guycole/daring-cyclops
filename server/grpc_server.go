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
// buf curl --schema proto --data '{"name":"bufcurl"}' http://localhost:8080/cyclops.v1.CyclopsService/PlayerNew
// buf curl --schema proto --data '{"available":"bufcurl"}' http://localhost:8080/cyclops.v1.CyclopsService/GameSummary

type cyclopsServiceServer struct {
	ft *facadeType

	cyclopsv1connect.UnimplementedCyclopsServiceHandler
}

func (css *cyclopsServiceServer) GameSummary(ctx context.Context, req *connect.Request[v1.GameSummaryRequest]) (*connect.Response[v1.GameSummaryResponse], error) {
	css.ft.sugarLog.Debug("gameCatalog:", req.Msg.GetAvailable())

	var retCode uint32 = shared.RetCodeSuccess
	results := []*v1.GameSummary{}

	gsat := css.ft.gameSummary()
	for _, gst := range gsat {
		if gst == nil {
			continue
		}

		blueShips2 := uint32(gst.blueShips)
		redShips2 := uint32(gst.redShips)

		temp := v1.GameSummary{Age: uint64(gst.age), Key: gst.key.key, BlueScore: gst.blueScore, BlueShips: blueShips2, RedScore: gst.redScore, RedShips: redShips2}
		results = append(results, &temp)
	}

	// Note that values with zero value are not included in the response

	return connect.NewResponse(&v1.GameSummaryResponse{GameSummary: results, Retcode: retCode}), nil
}

func (css *cyclopsServiceServer) Ping(ctx context.Context, req *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error) {
	css.ft.sugarLog.Debug("ping:", req.Msg.GetSource())
	return connect.NewResponse(&v1.PingResponse{Retcode: shared.RetCodeSuccess}), nil
}

func (css *cyclopsServiceServer) PlayerNew(ctx context.Context, req *connect.Request[v1.PlayerNewRequest]) (*connect.Response[v1.PlayerNewResponse], error) {
	css.ft.sugarLog.Debug("playerNew:", req.Msg.GetName())

	//pt, err := css.ft.playerNew(req.Msg.GetName())

	return connect.NewResponse(&v1.PlayerNewResponse{Retcode: shared.RetCodeSuccess}), nil
}
