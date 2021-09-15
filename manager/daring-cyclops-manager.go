// go mod init github.com/guycole/daring-cyclops/manager

package main

import (
	"context"
	"log"
	"time"

	redis "github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

// banner splash message
const banner = "Daring Cyclops Manager V0.0"

// redis begin
func connectRedis() {
	// FIXME should be config map

	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	log.Println(ctx)

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	/*
		rdb = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_URL"),
			Password: os.Getenv("REDIS_PASSWORD"), // no password set
			DB:       0,                           // use default DB
		})
	*/

	pong, err := rdb.Ping(ctx).Result()
	if err == nil {
		log.Println(pong, err)
	} else {
		log.Println(err)
	}

}

// redis end

func main() {
	log.Println(banner)

	for {
		log.Println("sleeping...")
		time.Sleep(8 * time.Second)
		connectRedis()
	}

}
