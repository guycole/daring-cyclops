package main

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
	acheronVoid bool
	blackHole   bool
	mine        bool
	planet      bool
	planetID    string
	ship        bool
	shipID      string
	star        bool
	starGate    bool
	starGateID  string

	//position *locationType
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

func setPlanet(bc boardCellType, uuid string) {
	bc.planet = true
	bc.planetID = uuid
}

func setStarGate(bc *boardCellType, uuid string) {
	bc.starGate = true
	bc.starGateID = uuid
}
*/

func isEmptyCell(arg boardCellType) bool {
	return true
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
		return " @"
		// "++" or "--"
	}

	if arg.ship {
		// find ship, return first character of name
		return " N"
	}

	if arg.star {
		return " *"
	}

	if arg.starGate {
		return " X"
	}

	return " ."
}
