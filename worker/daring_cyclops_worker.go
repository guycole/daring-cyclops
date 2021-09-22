// go mod init github.com/guycole/daring-cyclops/worker

package main

import (
	"log"
	"time"
)

// banner splash message
const banner = "Daring Cyclops Worker V0.0"

// redis begin
func connectToRedis() {
	/*
		rdb = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_URL"),
			Password: os.Getenv("REDIS_PASSWORD"), // no password set
			DB:       0,                           // use default DB
		})
	*/

	/*
		pong, err := rdb.Ping(ctx).Result()
		if err == nil {
			log.Println(pong, err)
		} else {
			log.Println(err)
		}
	*/
}

// redis end

func testClient2() {
	game := newGame("gameId", standardBoard)
	log.Println(game)

	for ndx := 0; ndx < 13; ndx++ {
		start := time.Now()

		turnManager(game)

		elapsed := time.Since(start)
		log.Printf("turn took %s", elapsed)

		time.Sleep(time.Second)
	}
}

func testClient(gameId string) {
	log.Printf("test client mode %s", gameId)
}

func main() {
	log.Println(banner)

	/*
		gameId := os.Getenv("gameId")
		if len(gameId) > 0 {
			log.Printf("gameId:%s", gameId)
		} else {
			log.Fatalf("missing gameId")
		}
	*/

	var gameId = "testGame0"

	//game := newGame(gameId, standardBoard)
	//log.Println(game)

	newGame(gameId, standardBoard)

	go commandFromManager(gameId + "m")

	for ndx := 0; ndx < 13; ndx++ {
		start := time.Now()

		//	turnManager(game)

		elapsed := time.Since(start)
		log.Printf("turn took %s", elapsed)

		time.Sleep(time.Second)
	}

	/*
		if len(os.Getenv("rabbit")) > 0 {
			rabbitClient(gameId)
		} else {
			testClient(gameId)
		}
	*/
}
