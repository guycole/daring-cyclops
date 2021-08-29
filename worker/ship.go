package main

import (
	"log"

	"github.com/google/uuid"
)

type shipEnum int

const (
	// ScoutShip ryr
	ScoutShip shipEnum = iota
	// FighterShip ry
	FighterShip
	// MinerShip ry
	MinerShip
	// FlagShip ry
	FlagShip
)

func (se shipEnum) String() string {
	return [...]string{"Scout", "Fighter", "Miner", "FlagShip"}[se]
}

/*
const std::string Ship::kBlueScouts[] =    {"lazor",  "nike",    "rapier",   "saber",    "vanir"};
const std::string Ship::kBlueFighters[] =  {"levant", "nimrod",  "ronin",   "scorpion", "viper"};
const std::string Ship::kBlueMiners[] =    {"lynx",   "napier",  "rigel",   "spartan",  "voyager"};
const std::string Ship::kBlueFlagships[] = {"lotus",  "nemesis", "reliant", "shogun",   "vega"};

const std::string Ship::kRedScouts[] =    {"dirk",    "griffin",  "hornet", "talon",   "wasp"};
const std::string Ship::kRedFighters[] =  {"demon",   "gargoyle", "hunter", "triton",  "wolf"};
const std::string Ship::kRedMiners[] =    {"delphos", "gibbet",   "hansen", "tirade",  "wight"};
const std::string Ship::kRedFlagships[] = {"dagon",   "gordon",   "hydra",  "tendril", "welink"};
*/

// ShipType ry
type ShipType struct {
	active   bool
	position locationType
	owner    string // player UUID
	shipEnum shipEnum
	shipName string
	team     playerTeamEnum
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

// NewShip ry
func NewShip(name, owner string, se shipEnum, team playerTeamEnum) *ShipType {
	result := ShipType{active: true, shipName: name, owner: owner, shipEnum: se, team: team}
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
	case ScoutShip:
		log.Println("scout")
		result.energy = 1
		result.mines = 2
		result.torpedos = 0
	case FighterShip:
		log.Println("fighter")
		result.energy = 1
		result.mines = 2
		result.torpedos = 0
	case MinerShip:
		log.Println("miner")
		result.energy = 1
		result.mines = 2
		result.torpedos = 0
	case FlagShip:
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
