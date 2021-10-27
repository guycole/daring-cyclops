// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"encoding/json"

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

// serialized for redis
type PlayerTypeJson struct {
	//	CreatedAt time.Time `json:"name"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Rank     string `json:"rank"`
	Score    int    `json:"score"`
	Team     string `json:"team"`
}

type playerType struct {
	email    string
	name     string
	password string
	rank     rankEnum
	score    int
	team     teamEnum
}

// newPlayer convenience function to populate struct
func newPlayer(email, name string, rank rankEnum, team teamEnum) (*playerType, error) {
	if len(email) < 1 {
		return nil, errors.New("bad email")
	}

	if len(name) < 1 {
		return nil, errors.New("bad player name")
	}

	result := playerType{email: email, name: name, rank: rank, team: team}
	return &result, nil
}

const testPlayerEmail1 = "test1@bogus.com"
const testPlayerName1 = "testName1"

// testPlayer1 returns test player1
func testPlayer1() *playerType {
	np1, _ := newPlayer(testPlayerEmail1, testPlayerName1, cadetRank, blueTeam)
	return np1
}

const testPlayerEmail2 = "test2@bogus.com"
const testPlayerName2 = "testName2"

// testPlayer2 returns test player2
func testPlayer2() *playerType {
	np2, _ := newPlayer(testPlayerEmail2, testPlayerName2, admiralRank, redTeam)
	return np2
}

/*
func rankChange(pt *PlayerType, newRank rankEnum) {
	// TODO issue 16
	// test for consecutive rank
	pt.Rank = newRank
}

func rankPromotion(pt *PlayerType) {
	// TODO issue 16
	// test score for rank promotion
}

func teamChange(pt *PlayerType, newTeam teamEnum) {
	// TODO issue 17
	// ensure only red/blue team selection
	pt.Team = newTeam
}
*/

// getPlayer reads from redis
func getPlayer(rc *redis.Client, key string) *playerType {
	log.Printf("getPlayer:%s", key)

	rawJson, err := rc.Get(redisCtx, key).Result()
	if err != nil {
		log.Println(err)
		return nil
	}

	var ptj PlayerTypeJson
	err = json.Unmarshal([]byte(rawJson), &ptj)
	if err != nil {
		log.Println(err)
		return nil
	}

	pt := playerType{email: ptj.Email, name: ptj.Name, password: ptj.Password, score: ptj.Score}
	pt.rank = findRank(ptj.Rank)
	pt.team = findTeam(ptj.Team)
	log.Println(pt)

	return &pt
}

// setPlayer writes to redis
func setPlayer(rc *redis.Client, pt *playerType) {
	log.Printf("setPlayer:%s", pt.name)

	rank := pt.rank.string()
	team := pt.team.string()

	ptj := PlayerTypeJson{Email: pt.email, Name: pt.name, Password: pt.password, Rank: rank, Score: pt.score, Team: team}

	payload, err := json.Marshal(ptj)
	if err != nil {
		log.Println(err)
	}

	key := ptj.Email

	err = rc.Set(redisCtx, key, payload, 0).Err()
	if err != nil {
		log.Println(err)
	}
}

func newPlayerX(gameId, name string) {
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
