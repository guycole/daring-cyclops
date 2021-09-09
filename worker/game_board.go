// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found
// in the LICENSE file.
package main

import (
	"log"
	"math/rand"
)

// boardTypeEnum
type boardTypeEnum int

const (
	unknownBoardType boardTypeEnum = iota
	emptyBoard                     // no mines, planets, ships, stargates, voids
	standardBoard                  // standard game map w/stars, planets, stargates and voids
)

// must match order for boardTypeEnum
func (bte boardTypeEnum) string() string {
	return [...]string{"unknown", "empty", "standard"}[bte]
}

const maxBoardSideX = 75
const maxBoardSideY = 75

type boardArrayType [maxBoardSideX][maxBoardSideY]*boardCellType

// newBoard create fresh empty game board
func newBoard() boardArrayType {
	var bat boardArrayType

	for yy := 0; yy < maxBoardSideY; yy++ {
		for xx := 0; xx < maxBoardSideX; xx++ {
			bat[yy][xx] = newBoardCell()
		}
	}

	return bat
}

func boardDump(bat boardArrayType) {
	log.Println("=-=-=-= boardDump =-=-=-=")

	var buffer string

	for yy := maxBoardSideY - 1; yy >= 0; yy-- {
		buffer = ""

		for xx := 0; xx < maxBoardSideX; xx++ {
			if bat[yy][xx] == nil {
				buffer += "oo"
			} else {
				buffer += boardCellToken(*bat[yy][xx])
			}
		}

		log.Printf("%2.2d %s", yy+1, buffer)
	}

	log.Println("=-=-=-= boardDump =-=-=-=")
}

func boardGenerator(gt *gameType) {
	switch gt.boardType {
	case emptyBoard:
		log.Println("generating empty board")
	case standardBoard:
		log.Println("generating standard board")
		addStarGates(gt)
		addStars(gt)
		addPlanets(gt)
	default:
		log.Println("unsupported boardType in boardGenerator")
	}
}

// return a random location for stars and planets
func randomCelestialLocation(gt *gameType) *locationType {
	for ndx := 0; ndx < 100; ndx++ {
		position := randomLocation(maxBoardSideY, maxBoardSideX)

		// cannot have celestial objects adjacent to stargates
		_, locNdx := starGateAdjacent(position)
		if locNdx >= 0 {
			//log.Printf("stargate adjacent:%d %d %d", gateNdx, locNdx, ndx)
			continue
		}

		boardCell := gt.board[position.yy][position.xx]
		if !testForCelestial(*boardCell) {
			return position
		}
	}

	log.Println("unable to generate random celestial location")

	return nil
}

func addPlanets(gt *gameType) {
	quarter := int(maxPlanets / 4)
	planetPopulation := 3*quarter + rand.Intn(quarter)
	log.Printf("planetPopulation:%d", planetPopulation)
	for ndx := 0; ndx < planetPopulation; ndx++ {
		position := randomCelestialLocation(gt)
		if position == nil {
			log.Println("skipping nil position for planet")
		} else {
			planet := newPlanet(position)
			gt.planets[ndx] = planet
			setPlanet(gt.board[planet.position.yy][planet.position.xx], planet.uuid)
		}
	}
}

func addStars(gt *gameType) {
	quarter := int(maxStars / 4)
	starPopulation := 3*quarter + rand.Intn(quarter)
	log.Printf("starPopulation:%d", starPopulation)
	for ndx := 0; ndx < starPopulation; ndx++ {
		position := randomCelestialLocation(gt)
		if position == nil {
			log.Println("skipping nil position for star")
		} else {
			star := newStar(position)
			gt.stars[ndx] = star
			setStar(gt.board[star.position.yy][star.position.xx], star.uuid)
		}
	}
}

func addStarGates(gt *gameType) {
	for ndx := 0; ndx < maxStarGates; ndx++ {
		sg := newStarGate(ndx)
		gt.starGates[ndx] = sg
		setStarGate(gt.board[sg.position.yy][sg.position.xx], sg.uuid)
	}
}
