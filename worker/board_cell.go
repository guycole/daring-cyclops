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
	// celestial objects
	acheronVoid bool
	blackHole   bool
	planet      bool
	star        bool
	starGate    bool
	celestialID string // uuid

	// player objects
	mine   bool
	ship   bool
	shipID string
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

func boardCellToken(arg boardCellType) string {
	if arg.acheronVoid {
		return "  "
	}

	if arg.blackHole {
		return "  "
	}

	// mine?

	if arg.planet {
		return "@"
		// "++" or "--"
	}

	if arg.ship {
		// find ship, return first character of name
		return " N"
	}

	if arg.star {
		return "*"
	}

	if arg.starGate {
		return "X"
	}

	return "."
}
