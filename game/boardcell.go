package game

type boardTokenEnum int

const (
	vacantToken boardTokenEnum = iota
	mineToken
	planetToken
	shipToken
	starGateToken
)

func (bte boardTokenEnum) String() string {
	return [...]string{"Vacant", "Mine", "Planet", "Ship", "StarGate"}[bte]
}

type boardCellType struct {
	acheronVoid bool
	blackHole   bool
	mine        bool
	planet      bool
	planetID    string
	ship        bool
	shipID      string
	starGate    bool
	starGateID  string

	position *locationType
}

func newBoardCell(position *locationType) *boardCellType {
	result := boardCellType{position: position}

	result.acheronVoid = false
	result.blackHole = false
	result.mine = false
	result.planet = false
	result.planetID = "bogus"
	result.ship = false
	result.shipID = "bogus"
	result.starGate = false
	result.starGateID = "bogus"

	return &result
}

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
