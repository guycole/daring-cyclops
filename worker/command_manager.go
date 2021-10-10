// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import (
	"context"
	"encoding/json"
	"log"

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

// newCommand convenience function to populate struct
func newCommand(name, id string, size int, commands commandArrayType) *CommandType {
	result := CommandType{Name: name, RequestId: id, CommandSize: size, Commands: commands}
	return &result
}

func commandFromManager(channelName string, commandQueue *commandQueueType) {
	log.Println("commandFromManager entry")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "bigSekret",
		DB:       0, // use default DB
	})

	topic := rdb.Subscribe(context.Background(), channelName)

	for {
		// blocking read
		msg, err := topic.ReceiveMessage(context.Background())
		if err != nil {
			log.Println("err err err err")
			log.Println(err)
			continue
		}

		var ct CommandType
		err = json.Unmarshal([]byte(msg.Payload), &ct)
		if err != nil {
			log.Println("err err err err 222")
			log.Println(err)
			continue
		}

		commandQueue.enqueue(&ct)
	}

	log.Println("commandFromManager exit")
}

func okResponse(channelName string) {
	log.Println("OK response")
}

func responseToManager(channelName string, ct *CommandType) {
	log.Println("responseToManager entry")
	log.Println(channelName)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "bigSekret",
		DB:       0, // use default DB
	})

	payload, err := json.Marshal(ct)
	if err != nil {
		log.Println(err)
	}

	err = rdb.Publish(context.Background(), channelName, payload).Err()
	if err != nil {
		log.Fatal(err)
	}
}
