// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import (
	"context"
	"encoding/json"
	"log"

	redis "github.com/go-redis/redis/v8"
)

const maxRequestArguments2 = 5

type argumentArrayType2 [maxRequestArguments2]string

type ResponseType struct {
	Name         string
	RequestId    string
	ArgumentSize int
	Arguments    argumentArrayType2
}

// newResponse convenience function to populate struct
func newResponse(name, id string, size int, arguments requestArrayType) *ResponseType {
	result := ResponseType{Name: name, RequestId: id, ArgumentSize: size, Arguments: arguments}
	return &result
}

func okResponse(channelName string) {
	log.Println("OK response")
}

func responseToManager(channelName string, rt *ResponseType) {
	log.Println("responseToManager entry")
	log.Println(channelName)

	// TODO get these arguments from secrets
	rdb := redis.NewClient(&redis.Options{
		Addr:     "cyclops-redis-master:6379",
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
