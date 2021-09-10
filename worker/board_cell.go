package main

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
	// celestial objects without uuid
	acheronVoid bool
	blackHole   bool

	// celestial objects with uuid
	planet      bool
	star        bool
	starGate    bool
	celestialID string // uuid

	// player object without uuid
	mine bool

	// player object with uuid
	ship       bool
	shipSymbol string

	tokenID string
}

func newBoardCell() *boardCellType {
	result := boardCellType{}
	return &result
}

/*
func setAcheronVoid(bc boardCellType) {
	bc.acheronVoid = true
}

func setBlackHole(bc boardCellType) {
	bc.blackHole = true
}
*/

func setPlanet(bc *boardCellType, uuid string) {
	if testForCelestial(*bc) {
		log.Println("unable to set planet because cell is occupied")
		return
	}

	bc.planet = true
	bc.celestialID = uuid
}

func setShip(bc *boardCellType, symbol, uuid string) {
	if !testForEmpty(*bc) {
		log.Println("unable to set ship because cell is occupied")
		return
	}

	bc.ship = true
	bc.shipSymbol = symbol
	bc.tokenID = uuid
}

func setStar(bc *boardCellType, uuid string) {
	if testForCelestial(*bc) {
		log.Println("unable to set star because cell is occupied")
		return
	}

	bc.star = true
	bc.celestialID = uuid
}

func setStarGate(bc *boardCellType, uuid string) {
	if testForCelestial(*bc) {
		log.Println("unable to set starGate because cell is occupied")
		return
	}

	bc.starGate = true
	bc.celestialID = uuid
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
