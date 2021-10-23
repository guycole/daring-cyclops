package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// serialized for redis
type SessionTypeJson struct {
	Email string
	Id    string
	//	CreatedAt time.Time
}

func sessionCheck(id string) *SessionTypeJson {
	log.Println("sessionCheck")
	log.Println(id)

	rc := freshRedisConnection()

	rawJson, err := rc.Get(redisCtx, id).Result()
	if err != nil {
		log.Println(err)
		log.Println("session err1x")
		return nil
	}

	var result SessionTypeJson
	err = json.Unmarshal([]byte(rawJson), &result)
	if err != nil {
		log.Println(err)
		log.Println("session err2x")
		return nil
	}

	return &result
}

func sessionTest(request *http.Request, response http.ResponseWriter) *SessionTypeJson {
	//	func session(writer http.ResponseWriter, request *http.Request) (sess data.Session, err error) {

	log.Println("-x-x-x-x-x-x-x-x-")
	cookie, err := request.Cookie("_cookie")
	log.Println(cookie)
	log.Println(err)

	session := SessionTypeJson{}
	if err == nil {
		session = *sessionCheck(cookie.Value)
	}

	return &session
}

/*
func getSession(rc *redis.Client, id string) (session *Session, err error) {
	rawJson, err := rc.Get(redisCtx, id).Result()
	if err != nil {
		log.Println(err)
		log.Println("session err1")
		return nil, nil
	}

	var result Session
	err = json.Unmarshal([]byte(rawJson), &result)
	if err != nil {
		log.Println(err)
		log.Println("session err2")
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
*/
