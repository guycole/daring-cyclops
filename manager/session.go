package main

import (
	"encoding/json"
	"log"

	redis "github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type Session struct {
	Email string
	Id    string
	//	CreatedAt time.Time
}

func getSession(rc *redis.Client, id string) (session *Session, err error) {
	rawJson, err := rc.Get(redisCtx, id).Result()
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	var result Session
	err = json.Unmarshal([]byte(rawJson), &result)
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	return &result, nil
}

func setSession(rc *redis.Client, email string) (session *Session, err error) {
	sessionId := uuid.NewString()

	result := Session{Email: email, Id: sessionId}

	payload, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
	}

	err = rc.Set(redisCtx, sessionId, payload, 0).Err()
	if err != nil {
		log.Println(err)
	}

	return &result, nil
}
