// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"log"
)

//const testPlayerName1 = "testName1"
//const testPlayerName2 = "testName2"

// banner splash message
const banner = "Daring Cyclops Manager V0.0"

func main() {
	log.Println(banner)

	log.Println(configuration)

	manager := newManager()
	log.Println(manager)

	rc := freshRedisConnection()
	log.Println(rc)
	pong, err := rc.Ping(redisCtx).Result()
	log.Println(pong, err)

	pt := playerType{email: "email@bogus.com", name: "testaroo", password: "password"}
	setPlayer(rc, &pt)

	getPlayer(rc, pt.email)

	//webPortal()

	/*
		var gameId = "testGame0"
		var counter int

		go responseFromWorker(gameId + "w")

		for {
			counter++

			start := time.Now()

			elapsed := time.Since(start)
			log.Printf("turn %d took %s", counter, elapsed)

			pingRequest(gameId, "pingTest"+strconv.Itoa(counter))
			playerCreateRequest(gameId, testPlayerName1, "cadet", "blue")
			playerCreateRequest(gameId, testPlayerName2, "admiral", "red")
			playerDeleteRequest(gameId, testPlayerName1)
			playerDeleteRequest(gameId, testPlayerName2)

			time.Sleep(5 * time.Second)

			counter += 1
		}

		shutDownRequest(gameId, "shutDown")
	*/
	/*
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

		newPlayer2(gwt.gameId, testPlayerName1, manager.rdb)
		newShip2(gwt.gameId, testPlayerName1, manager.rdb)

		newPing(gwt.gameId, testPlayerName1, manager.rdb)

		gamePlayerAdd(*(testPlayer1()), &gwt.blueTeam, &gwt.redTeam)
		gamePlayerAdd(*(testPlayer2()), &gwt.blueTeam, &gwt.redTeam)

		bluePopulation, redPopulation := gamePlayerCensus(*gwt)
		log.Printf("%d %d", bluePopulation, redPopulation)
	*/

	////////////
	// now start worker and write commands
	////////////

	//log.Println(manager)

	//gameDump(manager.games)

	/*
		for {
			log.Println("sleeping...")
			time.Sleep(8 * time.Second)
		}
	*/
}

func newPlayer2(gameId, name string) {
	channel := gameId + "m"
	log.Println(channel)
	/*
		var arguments argumentArrayType
		commands[0] = "playerCreate"
		commands[1] = "captain"
		commands[2] = "blue"

		nr := newRequest(name, 3, arguments)

		var commands commandArrayType


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
	*/
}
