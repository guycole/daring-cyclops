// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found
// in the LICENSE file.
package main

import (
	"encoding/json"

	"context"
	"errors"
	"log"

	redis "github.com/go-redis/redis/v8"
)

// rankEnum
type rankEnum int

// must match order for legalRanks
const (
	unknownRank rankEnum = iota
	cadetRank
	lieutenantRank
	captainRank
	admiralRank
)

// must match order for rankEnum
var legalRanks = [...]string{
	"unknown",
	"cadet",
	"lieutenant",
	"captain",
	"admiral",
}

// must match order for rankEnum
func (re rankEnum) string() string {
	return [...]string{"unknown", "cadet", "lieutenant", "captain", "admiral"}[re]
}

func findRank(arg string) rankEnum {
	for ndx := 0; ndx < len(legalRanks); ndx++ {
		if legalRanks[ndx] == arg {
			return rankEnum(ndx)
		}
	}

	return rankEnum(unknownRank)
}

// teamEnum
type teamEnum int

// must match order for legalTeams
const (
	unknownTeam teamEnum = iota
	neutralTeam
	blueTeam
	redTeam
	acheronTeam
)

// must match order for teamEnum
var legalTeams = [...]string{
	"unknown",
	"neutral",
	"blue",
	"red",
	"acheron",
}

// must match order for teamEnum
func (te teamEnum) string() string {
	return [...]string{"unknown", "neutral", "blue", "red", "acheron"}[te]
}

func findTeam(arg string) teamEnum {
	for ndx := 0; ndx < len(legalTeams); ndx++ {
		if legalTeams[ndx] == arg {
			return teamEnum(ndx)
		}
	}

	return teamEnum(unknownTeam)
}

type playerType struct {
	email string   `json:"email"`
	name  string   `json:"name"`
	rank  rankEnum `json:"rank"`
	score int      `json:"score"`
	team  teamEnum `json:"team"`
}

const maxTeamPlayers = 5
const maxPlayers = maxTeamPlayers * 2

// playerArrayType contains all active players
type playerArrayType [maxPlayers]*playerType

// newPlayer convenience function to populate struct
func newPlayer(email, name string, rank string, team string) (*playerType, error) {
	if len(email) < 1 {
		return nil, errors.New("bad email")
	}

	if len(name) < 1 {
		return nil, errors.New("bad player name")
	}

	result := playerType{email: email, name: name}
	playerRank := findRank(rank)
	if playerRank == unknownRank {
		return nil, errors.New("unknown rank")
	}

	result.rank = playerRank

	playerTeam := findTeam(team)
	if playerTeam == unknownTeam {
		return nil, errors.New("unknown team")
	}

	result.team = playerTeam

	return &result, nil
}

const testPlayerEmail1 = "test1@bogus.com"
const testPlayerName1 = "testName1"

// testPlayer1 returns test player1
func testPlayer1() *playerType {
	np1, _ := newPlayer(testPlayerEmail1, testPlayerName1, "cadet", "blue")
	return np1
}

const testPlayerEmail2 = "test2@bogus.com"
const testPlayerName2 = "testName2"

// testPlayer2 returns test player2
func testPlayer2() *playerType {
	np2, _ := newPlayer(testPlayerEmail2, testPlayerName2, "admiral", "red")
	return np2
}

/////////////////////////

func (obj *playerType) marshalBinary() ([]byte, error) {
	return json.Marshal(obj)
}

// UnmarshalBinary -
func (obj *playerType) unmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &obj); err != nil {
		log.Println("choke choke")
		return err
	}

	return nil
}

/////////////////////////

func getPlayer(rdb *redis.Client, key string) *playerType {
	log.Println("getPlayer")
	log.Println(key)

	p := rdb.Get(context.Background(), key)
	log.Println("xoxoxoxoxoxo")
	log.Println(p)

	/*
		pong, err := rdb.Ping(context.Background()).Result()
		if err == nil {
			log.Println(pong)
		} else {
			log.Println(err)
		}
	*/

	return nil
}

func setPlayer(rdb *redis.Client, pt *playerType) {
	message4 := `{"email":"email2", "name":"name2", "rank":1, "score":123, "team":2}`

	var obj playerType
	err3 := obj.unmarshalBinary([]byte(message4))
	log.Println(err3)
	log.Println(obj)

	log.Println("-x--x-x-x-x-x-x")

	log.Println("setPlayer")
	log.Println(pt)

	jpt, err2 := pt.marshalBinary()
	log.Println(jpt)
	log.Println(err2)

	key := pt.name

	err := rdb.Set(context.Background(), key, jpt, 0).Err()
	if err != nil {
		log.Println(err)
		//		log.Error(err)
	}

	log.Println("select")

	p, _ := rdb.Get(context.Background(), key).Result()
	log.Println(p)

	/*
		cacheData, err := rdb.Get(ctx, ipaddress).Result()
		if err != nil {
			return nil, err
		}

		var obj IPStackResponseSuccess
		err = obj.UnmarshalBinary([]byte(cacheData))
		if err != nil {
			return nil, err
		}

		return &obj, nil
	*/

	log.Println("exit exit")
}

/*
func setPlayer(c *RedisClient, key string, value interface{}) error {
    p, err := json.Marshal(value)
    if err != nil {
       return err
    }
    return c.Set(key, p)
}

func get(c *RedisClient, key string, dest interface{}) error {
    p, err := c.Get(key)
    if err != nil {
       return err
    }
    return json.Unmarshal(p, dest)
}
*/
