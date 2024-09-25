// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package server

import "log"

type boardTokenEnum int

const (
	vacantToken boardTokenEnum = iota
	mineToken
	planetToken
	shipToken
	starGateToken
	voidToken
)

func (bte boardTokenEnum) String() string {
	return [...]string{"vacant", "mine", "planet", "ship", "starGate", "void"}[bte]
}

type boardCellType struct {
	// celestial objects without token uuid
	acheronVoid bool
	blackHole   bool

	// celestial objects with token uuid
	planet   bool
	star     bool
	starGate bool

	// player object without token uuid
	mine bool

	// player object with token uuid
	ship       bool
	shipSymbol string

	tokenID string
}

func newBoardCell() *boardCellType {
	result := boardCellType{}
	return &result
}

func (bc *boardCellType) setAcheronVoid() {
	bc.acheronVoid = true
}

func (bc *boardCellType) setBlackHole() {
	bc.blackHole = true
}

func (bc *boardCellType) clearPlanet() {
	if !bc.planet {
		log.Println("attempting to clearPlanet when none declared")
	}

	bc.planet = false
}

func (bc *boardCellType) setPlanet(uuid string) {
	if testForCelestial(*bc) {
		log.Println("unable to set planet because cell is occupied")
		return
	}

	bc.planet = true
	bc.tokenID = uuid
}

func (bc *boardCellType) clearShip() {
	log.Println("clear ship clear ship")
	if !bc.ship {
		log.Println("attempting to clearShip when none declared")
	}

	bc.ship = false
	bc.shipSymbol = ""
	bc.tokenID = ""
}

func (bc *boardCellType) setShip(symbol, uuid string) {
	if !testForEmpty(*bc) {
		log.Println("unable to set ship because cell is occupied")
		return
	}

	bc.ship = true
	bc.shipSymbol = symbol
	bc.tokenID = uuid
}

func (bc *boardCellType) setStar(uuid string) {
	if testForCelestial(*bc) {
		log.Println("unable to set star because cell is occupied")
		return
	}

	bc.star = true
	bc.tokenID = uuid
}

// convert a star to black hole
func (bc *boardCellType) starToBlackHole() {
	if !bc.blackHole {
		log.Println("attempt to convert start to black hole when none declared")
	}

	bc.blackHole = true
	bc.star = false
}

func (bc *boardCellType) setStarGate(uuid string) {
	if testForCelestial(*bc) {
		log.Println("unable to set starGate because cell is occupied")
		return
	}

	bc.starGate = true
	bc.tokenID = uuid
}

// return true if cell contains celestial object
func testForCelestial(arg boardCellType) bool {
	if arg.acheronVoid || arg.blackHole || arg.planet || arg.star || arg.starGate {
		return true
	}

	return false
}

// return true if empty cell
func testForEmpty(arg boardCellType) bool {
	if testForCelestial(arg) {
		return false
	}

	if arg.mine || arg.ship {
		return false
	}

	return true
}

func boardCellToken(arg boardCellType) string {
	if arg.acheronVoid {
		return " "
	}

	if arg.blackHole {
		return " "
	}

	// mine
	if arg.mine {
		return "#"
	}

	if arg.planet {
		return "@"
		// "++" or "--"
	}

	if arg.ship {
		return arg.shipSymbol
	}

	if arg.star {
		return "*"
	}

	if arg.starGate {
		return "X"
	}

	return "."
}
