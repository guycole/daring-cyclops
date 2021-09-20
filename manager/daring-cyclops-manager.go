// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found
// in the LICENSE file.

// go mod init github.com/guycole/daring-cyclops/manager

package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	redis "github.com/go-redis/redis/v8"
)

// gameManagerType, only one instance
type gameManagerType struct {
	games gameArrayType
	rdb   *redis.Client
}

var ctx = context.Background()

// banner splash message
const banner = "Daring Cyclops Manager V0.0"

func newManager() *gameManagerType {
	log.Println("new manager")

	rand.Seed(time.Now().UnixNano())

	gmt := gameManagerType{}

	gmt.rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &gmt
}

func main() {
	log.Println(banner)

	manager := newManager()
	log.Println(manager)

	// known to redis
	setPlayer(manager.rdb, testPlayer1())
	setPlayer(manager.rdb, testPlayer2())

	gwt := newGame(testGame0)
	ndx := gameAdd(gwt, &manager.games)
	if ndx < 0 {
		log.Fatalf("unable to add game")
	}

	newPing(gwt.gameId, testPlayerName1, manager.rdb)

	gamePlayerAdd(*(testPlayer1()), &gwt.blueTeam, &gwt.redTeam)
	gamePlayerAdd(*(testPlayer2()), &gwt.blueTeam, &gwt.redTeam)

	bluePopulation, redPopulation := gamePlayerCensus(*gwt)
	log.Printf("%d %d", bluePopulation, redPopulation)

	////////////
	// now start worker and write commands
	////////////

	//log.Println(manager)

	gameDump(manager.games)

	/*
		for {
			log.Println("sleeping...")
			time.Sleep(8 * time.Second)
		}
	*/
}
