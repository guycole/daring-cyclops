// go mod init github.com/guycole/daring-cyclops/manager

package main

import (
	"log"
	"os"
	"time"

	redis "github.com/go-redis/redis/v8"
)

var rdb *redis.Client

// banner splash message
const banner = "Daring Cyclops Manager V0.0"

// redis begin
func connectToRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
	})

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

func main() {
	log.Println(banner)

	for {
		log.Println("sleeping...")
		time.Sleep(8 * time.Second)
	}

}
