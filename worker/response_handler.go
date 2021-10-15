// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import (
	"context"
	"encoding/json"
	"log"

	redis "github.com/go-redis/redis/v8"
)

type ResponseType struct {
	RequestId    string
	Response     int
	ArgumentSize int
	Arguments    argumentArrayType
}

// newResponse convenience function to populate struct
func newResponse(response responseEnum, id string, size int, arguments argumentArrayType) *ResponseType {
	result := ResponseType{RequestId: id, ArgumentSize: size, Arguments: arguments}
	result.Response = int(response)
	return &result
}

func okArgument() (int, argumentArrayType) {
	log.Println("OK argument")

	var aat argumentArrayType
	aat[0] = "ok"
	return 1, aat
}

func unknownArgument() (int, argumentArrayType) {
	log.Println("unknown argument")

	var aat argumentArrayType
	aat[0] = "unknown"
	return 1, aat
}

func responseToManager(channelName string, rt *ResponseType) {
	log.Println("-x-x-x-x-x-x-x-x-x-x-x-x-x-")
	log.Println(rt)
	log.Println("-x-x-x-x-x-x-x-x-x-x-x-x-x-")

	// TODO get these arguments from secrets
	rdb := redis.NewClient(&redis.Options{
		Addr:     "cyclops-redis-master:6379",
		Password: "bigSekret",
		DB:       0, // use default DB
	})

	payload, err := json.Marshal(rt)
	if err != nil {
		log.Println(err)
	}

	err = rdb.Publish(context.Background(), channelName, payload).Err()
	if err != nil {
		log.Fatal(err)
	}
}
