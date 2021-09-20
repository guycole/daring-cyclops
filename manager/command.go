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

const maxCommandArguments = 5

type commandArrayType [maxCommandArguments]string

type CommandType struct {
	Name        string
	RequestId   string
	CommandSize int
	Commands    commandArrayType
}

func newCommand(name string, size int, commands commandArrayType) *CommandType {
	raw := CommandType{Name: name, CommandSize: size, Commands: commands}
	raw.RequestId = uuid.NewString()
	return &raw
}

func newPing(gameId, name string, rdb *redis.Client) {
	channel := gameId + "m"

	var commands commandArrayType
	commands[0] = "ping"

	ct := newCommand(name, 1, commands)
	log.Println(ct)

	payload, err := json.Marshal(ct)
	if err != nil {
		log.Println(err)
	}

	err = rdb.Publish(context.Background(), channel, payload).Err()
	if err != nil {
		log.Fatal(err)
	}
}
