// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import (
	"encoding/json"

	"context"
	"log"

	"github.com/google/uuid"

	redis "github.com/go-redis/redis/v8"
)

const maxRequestArguments = 5

type argumentArrayType [maxRequestArguments]string

type RequestType struct {
	Name         string
	RequestId    string
	ArgumentSize int
	Arguments    argumentArrayType
}

func newRequest(name string, argSize int, arguments argumentArrayType) *RequestType {
	raw := RequestType{Name: name, ArgumentSize: argSize, Arguments: arguments}
	raw.RequestId = uuid.NewString()
	return &raw
}

func redisPublish(gameId string, rt *RequestType) {
	channel := gameId + "m"
	log.Println(channel)

	payload, err := json.Marshal(rt)
	if err != nil {
		log.Println(err)
	}

	log.Println(payload)

	// TODO get these arguments from secrets
	rdb := redis.NewClient(&redis.Options{
		Addr:     "cyclops-redis-master:6379",
		Password: "bigSekret",
		DB:       0, // use default DB
	})

	err = rdb.Publish(context.Background(), channel, payload).Err()
	if err != nil {
		log.Fatal(err)
	}
}

func playerCreateRequest(gameId, name, rank, team string) {
	var arguments argumentArrayType
	arguments[0] = "playerCreateRequest"
	arguments[1] = rank
	arguments[2] = team

	rt := newRequest(name, 3, arguments)
	redisPublish(gameId, rt)
}

func playerDeleteRequest(gameId, name string) {
	var arguments argumentArrayType
	arguments[0] = "playerDeleteRequest"

	rt := newRequest(name, 1, arguments)
	redisPublish(gameId, rt)
}

func pingRequest(gameId, name string) {
	var arguments argumentArrayType
	arguments[0] = "pingRequest"

	rt := newRequest(name, 1, arguments)
	redisPublish(gameId, rt)
}

func shutDownRequest(gameId, name string) {
	var arguments argumentArrayType
	arguments[0] = "shutDownRequest"

	rt := newRequest(name, 1, arguments)
	redisPublish(gameId, rt)
}
