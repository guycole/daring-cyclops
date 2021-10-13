// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

/*
// run for game duration
func responseFromWorker(channelName string, requestQueue *requestQueueType) {
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
		msg, err := topic.ReceiveMessage(context.Background())
		if err != nil {
			log.Println(err)
			log.Println("requestFromManager skipping bad receive message")
			continue
		}

		var rt RequestType
		err = json.Unmarshal([]byte(msg.Payload), &rt)
		if err != nil {
			log.Println(err)
			log.Println("requestFromManager skipping bad unmarshal")
			continue
		}

		requestQueue.enqueue(&rt)
	}

	log.Println("requestFromManager exit")
}
*/
