// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import (
	"context"
	"log"

	redis "github.com/go-redis/redis/v8"
)

// run for game duration
func responseFromWorker(channelName string) {
	log.Println("requestFromManager entry")

	// TODO get these arguments from secrets
	rdb := redis.NewClient(&redis.Options{
		Addr:     "cyclops-redis-master:6379",
		Password: "bigSekret",
		DB:       0, // use default DB
	})

	topic := rdb.Subscribe(context.Background(), channelName)

	for {
		// blocking read
		message, err := topic.ReceiveMessage(context.Background())
		if err != nil {
			log.Println(err)
			log.Println("requestFromWorker skipping bad receive message")
			continue
		}

		log.Println(message)

		/*
			var rt RequestType
			err = json.Unmarshal([]byte(msg.Payload), &rt)
			if err != nil {
				log.Println(err)
				log.Println("requestFromManager skipping bad unmarshal")
				continue
			}

			requestQueue.enqueue(&rt)
		*/
	}

	log.Println("requestFromManager exit")
}
