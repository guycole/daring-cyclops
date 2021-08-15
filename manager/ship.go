package manager

import (
	"log"

	"github.com/google/uuid"
)

type ShipType int

const (
	Scout ShipType = iota
	Fighter
	Miner
	FlagShip
)

func (st ShipType) String() string {
	return [...]string{"Scout", "Fighter", "Miner", "FlagShip"}[st]
}

type Ship struct {
	Active   bool
	Position Location
	Owner    Player
	Type     ShipType
	Uuid     string

	// ship systems as percentage operational (values 100 to 0)
	Computer    int
	LifeSupport int
	Radio       int
	Shields     int
	TractorBeam int

	// engines
	ImpulseEngines int
	WarpEngines    int

	// weapons
	Phasers      int
	TorpedoTubes int

	// weapons inventory
	Mines    int
	Torpedos int

	// generic ship energy
	Energy int
}

func getFreshShip(st ShipType, owner Player) Ship {
	var result Ship

	result.Active = true
	result.Owner = owner
	result.Type = st
	result.Uuid = uuid.NewString()

	result.Computer = 100
	result.LifeSupport = 100
	result.Radio = 100
	result.Shields = 100
	result.TractorBeam = 100

	result.ImpulseEngines = 100
	result.WarpEngines = 100

	result.Phasers = 100
	result.TorpedoTubes = 100

	switch st {
	case Scout:
		log.Println("scout")
		result.Energy = 1
		result.Mines = 2
		result.Torpedos = 0
	case Fighter:
		log.Println("fighter")
		result.Energy = 1
		result.Mines = 2
		result.Torpedos = 0
	case Miner:
		log.Println("miner")
		result.Energy = 1
		result.Mines = 2
		result.Torpedos = 0
	case FlagShip:
		log.Println("flagship")
		result.Energy = 1
		result.Mines = 2
		result.Torpedos = 0
	default:
		log.Println("must throw error")
	}

	return result
}

func moveShip(ship Ship, location Location) {
	if ship.Active == false {
		log.Println("must throw error")
	}

	//var distance = 123
}
