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

type PlayerType struct {
	Email string
	Name  string
	Rank  rankEnum
	Score int
	Team  teamEnum
}

// newPlayer convenience function to populate struct
func newPlayer(email, name string, rank string, team string) (*PlayerType, error) {
	if len(email) < 1 {
		return nil, errors.New("bad email")
	}

	if len(name) < 1 {
		return nil, errors.New("bad player name")
	}

	result := PlayerType{Email: email, Name: name}
	playerRank := findRank(rank)
	if playerRank == unknownRank {
		return nil, errors.New("unknown rank")
	}

	result.Rank = playerRank

	playerTeam := findTeam(team)
	if playerTeam == unknownTeam {
		return nil, errors.New("unknown team")
	}

	result.Team = playerTeam

	return &result, nil
}

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

const testPlayerEmail1 = "test1@bogus.com"
const testPlayerName1 = "testName1"

// testPlayer1 returns test player1
func testPlayer1() *PlayerType {
	np1, _ := newPlayer(testPlayerEmail1, testPlayerName1, "cadet", "blue")
	return np1
}

const testPlayerEmail2 = "test2@bogus.com"
const testPlayerName2 = "testName2"

// testPlayer2 returns test player2
func testPlayer2() *PlayerType {
	np2, _ := newPlayer(testPlayerEmail2, testPlayerName2, "admiral", "red")
	return np2
}

// getPlayer reads from redis
func getPlayer(rdb *redis.Client, key string) *PlayerType {
	log.Printf("getPlayer:%s", key)

	rawJson, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		log.Println(err)
		return nil
	}

	var pt PlayerType
	err = json.Unmarshal([]byte(rawJson), &pt)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &pt
}

// setPlayer writes to redis
func setPlayer(rdb *redis.Client, pt *PlayerType) {
	log.Printf("setPlayer:%s", pt.Name)

	payload, err := json.Marshal(pt)
	if err != nil {
		log.Println(err)
	}

	key := pt.Name

	err = rdb.Set(context.Background(), key, payload, 0).Err()
	if err != nil {
		log.Println(err)
	}
}
