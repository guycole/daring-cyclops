package game

import (
	"log"

	"github.com/google/uuid"
)

type shipEnum int

const (
	scoutShip shipEnum = iota
	fighterShip
	minerShip
	flagShip
)

func (se shipEnum) String() string {
	return [...]string{"Scout", "Fighter", "Miner", "FlagShip"}[se]
}

type shipType struct {
	active   bool
	position locationType
	owner    Player
	shipEnum shipEnum
	uuid     string

	// ship systems as percentage operational (values 100 to 0)
	computer    int
	lifeSupport int
	radio       int
	shields     int
	tractorBeam int

	// engines
	impulseEngines int
	warpEngines    int

	// weapons
	phasers      int
	torpedoTubes int

	// weapons inventory
	mines    int
	torpedos int

	// generic ship energy
	energy int
}

func newShip(se shipEnum, owner Player) *shipType {
	result := shipType{active: true, owner: owner, shipEnum: se}
	result.uuid = uuid.NewString()

	result.computer = 100
	result.lifeSupport = 100
	result.radio = 100
	result.shields = 100
	result.tractorBeam = 100

	result.impulseEngines = 100
	result.warpEngines = 100

	result.phasers = 100
	result.torpedoTubes = 100

	switch se {
	case scoutShip:
		log.Println("scout")
		result.energy = 1
		result.mines = 2
		result.torpedos = 0
	case fighterShip:
		log.Println("fighter")
		result.energy = 1
		result.mines = 2
		result.torpedos = 0
	case minerShip:
		log.Println("miner")
		result.energy = 1
		result.mines = 2
		result.torpedos = 0
	case flagShip:
		log.Println("flagship")
		result.energy = 1
		result.mines = 2
		result.torpedos = 0
	default:
		log.Println("must throw error")
	}

	return &result
}

/*
func moveShip(ship Ship, location Location) {
	if ship.Active == false {
		log.Println("must throw error")
	}

	//var distance = 123
}
*/
