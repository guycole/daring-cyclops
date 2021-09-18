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

//var rdb *redis.Client

// banner splash message
const banner = "Daring Cyclops Manager V0.0"

// redis begin
func connectRedis(rdb *redis.Client) {
	// FIXME should be config map

	log.Println(ctx)

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	// game management

	/*
		rdb = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_URL"),
			Password: os.Getenv("REDIS_PASSWORD"), // no password set
			DB:       0,                           // use default DB
		})
	*/

	pong, err := rdb.Ping(ctx).Result()
	if err == nil {
		log.Println(pong)
	} else {
		log.Println(err)
	}

}

// redis end

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
	connectRedis(manager.rdb)

	for {
		log.Println("sleeping...")
		time.Sleep(8 * time.Second)
		connectRedis(manager.rdb)
	}
}
