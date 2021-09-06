// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found
// in the LICENSE file.
package main

import (
	"log"
)

// boardTypeEnum
type boardTypeEnum int

const (
	unknownBoardType boardTypeEnum = iota
	emptyBoard01                   // no mines, planets, ships, stargates, voids
	standardBoard01                // standard game map w/planets, stargates and voids
	randomBoard01                  // standard game map w/randomly generated planets
)

// must match order for boardTypeEnum
func (bte boardTypeEnum) string() string {
	return [...]string{"unknown", "empty01", "standard01", "random01"}[bte]
}

const maxBoardSideX = 75
const maxBoardSideY = 75

type boardArrayType [maxBoardSideX][maxBoardSideY]boardCellType

type gameBoardType struct {
	boardType  boardTypeEnum
	boardArray boardArrayType
}

func newGameBoard(bt boardTypeEnum) gameBoardType {
	result := gameBoardType{boardType: bt}
	return result
}

func boardDump(gt gameType) {
	log.Println("=-=-=-= boardDump =-=-=-=")

	boardType := gt.gameBoard.boardType.string()
	log.Printf("boardType:%s", boardType)

	/*
		for yy := 0; yy < maxBoardSideY; yy++ {
			var buffer string

			for xx := 0; xx < maxBoardSideX; xx++ {
				buffer += " "
			}

			log.Println(buffer)
		}
	*/

	/*
		for ndx := 0; ndx < maxPlayers; ndx++ {
			rank := gt.players[ndx].rank.string()
			team := gt.players[ndx].team.string()

			log.Printf("%d %t %s %s %s %s", ndx, gt.players[ndx].active, gt.players[ndx].name, rank, team, gt.players[ndx].uuid)
		}
	*/

	log.Println("=-=-=-= boardDump =-=-=-=")
}

func randomLocation3() *locationType {
	//xx := rand.Intn(limitX)
	//yy := rand.Intn(limitY)
	//return newLocation(yy, xx)
	return nil
}

func boardGenerator(gt *gameType) {
	boardType := gt.gameBoard.boardType

	switch boardType {
	case emptyBoard01:
		log.Println("generating empty board")
		break
	case standardBoard01:
		log.Println("generating standard board")
		addPlanets(gt)
		addStarGates(gt)
	case randomBoard01:
		log.Println("generating random board")
		addRandomPlanets(gt)
	default:
		log.Println("unsupported boardType in boardGenerator")
	}
}

func addPlanets(gt *gameType) {
}

func addStarGates(gt *gameType) {
	/*
		for ndx := 0; ndx < 9; ndx++ {
			sg := newStarGate(ndx)
			log.Println(sg)

			gb.StarGates[sg.uuid] = *sg
		}
	*/
}

func addRandomPlanets(gt *gameType) {

}
