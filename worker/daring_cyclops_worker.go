// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import (
	"flag"
	"log"
	"time"
	//	"flag"
	//	"log"
	//	"net/http"
	//	"time"
	//	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// banner splash message
const banner = "Daring Cyclops Worker V0.0"

func testClient(gameId string) {
	log.Printf("test client mode %s", gameId)
}

var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

func main3() {
	log.Println(banner)

	/*
		flag.Parse()
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(*addr, nil))

		log.Println("fell through")
	*/

	/*
		gameId := os.Getenv("gameId")
		if len(gameId) > 0 {
			log.Printf("gameId:%s", gameId)
		} else {
			log.Fatalf("missing gameId")
		}
	*/

	var gameId = "testGame0"

	game := newGame(gameId, emptyBoard)

	go requestFromManager(game.inboundQueue, game.requestQueue)

	for {
		if game.shutDownFlag {
			log.Println("main loop break")
			break
		}

		start := time.Now()

		game.turnManager()

		elapsed := time.Since(start)
		log.Printf("turn %d took %s", game.turnCounter, elapsed)

		time.Sleep(5 * time.Second)
	}

	/*
		for ndx := 0; ndx < 13; ndx++ {
			start := time.Now()

			game.turnManager()

			elapsed := time.Since(start)
			log.Printf("turn %d took %s", ndx, elapsed)

			time.Sleep(1 * time.Second)
		}
	*/

	game.requestQueue.dump()
}
