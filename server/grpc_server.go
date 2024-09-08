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

type cyclopsServiceServer struct {
	ft *FacadeType

	cyclopsv1connect.UnimplementedCyclopsServiceHandler
}

func (css *cyclopsServiceServer) GameCatalog(ctx context.Context, req *connect.Request[v1.GameCatalogRequest]) (*connect.Response[v1.GameCatalogResponse], error) {
	css.ft.sugarLog.Debug("gameCatalog:", req.Msg.GetGameKey())

	var retCode uint32 = shared.RetCodeSuccess
	results := []*v1.GameSummary{}

	gat := css.ft.gameCatalog()
	for _, gt := range gat {
		if gt == nil {
			continue
		}

		temp := v1.GameSummary{Age: gt.age, Key: gt.key.key}
		results = append(results, &temp)
	}

	return connect.NewResponse(&v1.GameCatalogResponse{GameSummary: results, Retcode: retCode}), nil
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
