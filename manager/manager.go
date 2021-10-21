// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"log"
	"math/rand"
	"time"

	redis "github.com/go-redis/redis/v8"
)

// gameManagerType, only one instance
type gameManagerType struct {
	//	games gameArrayType
	rdb *redis.Client
}

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

/*
type gameMasterType struct {
	seqId     int
	shipPop   int
	blueScore int
	redScore  int
}
*/

/*
const maxGames = 5
type masterArrayType [maxGames]*gameMasterType
*/
