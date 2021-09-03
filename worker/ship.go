package main

import (
	"log"
	"strings"

	"github.com/google/uuid"
)

type conditionEnum int

const (
	unknownCondition conditionEnum = iota
	greenCondition
	yellowCondition
	redCondition
)

// must match order for conditionEnum
func (ce conditionEnum) string() string {
	return [...]string{"Unknown", "Green", "Yellow", "Red"}[ce]
}

type shipEnum int

// must match order for legalShips
const (
	unknownShip shipEnum = iota
	scoutShip
	fighterShip
	minerShip
	flagShip
)

// must match order for shipEnum
var legalShips = [...]string{
	"unknownShip",
	"scout",
	"fighter",
	"miner",
	"flag",
}

// must match order for shipEnum
func (se shipEnum) string() string {
	return [...]string{"Unknown", "Scout", "Fighter", "Miner", "FlagShip"}[se]
}

func findShipClass(arg string) shipEnum {
	for ndx := 0; ndx < len(legalShips); ndx++ {
		if legalShips[ndx] == arg {
			return shipEnum(ndx)
		}
	}

	return shipEnum(unknownShip)
}

type nameEnum int

// must match order for legalNames
const (
	unknownName nameEnum = iota
	lazorName
	nikeName
	rapierName
	saberName
	vanirName
	levantName
	nimrodName
	roninName
	scorpionName
	viperName
	lynxName
	napierName
	rigelName
	spartanName
	voyagerName
	lotusName
	nemesisName
	reliantName
	shogunName
	vegaName
	dirkName
	griffinName
	hornetName
	talonName
	waspName
	demonName
	gargoyleName
	hunterName
	tritonName
	wolfName
	delphosName
	gibbetName
	hansenName
	tiradeName
	wightName
	dagonName
	gordonName
	hydraName
	tendrilName
	welinkName
)

// must match order for nameEnum
var legalNames = [...]string{
	"unknown",
	"lazor",
	"nike",
	"rapier",
	"saber",
	"vanir",
	"levant",
	"nimrod",
	"ronin",
	"scorpion",
	"viper",
	"lynx",
	"napier",
	"rigel",
	"spartan",
	"voyager",
	"lotus",
	"nemesis",
	"reliant",
	"shogun",
	"vega",
	"dirk",
	"griffin",
	"hornet",
	"talon",
	"wasp",
	"demon",
	"gargoyle",
	"hunter",
	"triton",
	"wolf",
	"delphos",
	"gibbet",
	"hansen",
	"tirade",
	"wight",
	"dagon",
	"gordon",
	"hydra",
	"tendril",
	"welink",
}

// must match order for playerRankEnum
func (ne nameEnum) string() string {
	return [...]string{"Unknown",
		"Lazor", "Nike", "Rapier", "Saber", "Vanir",
		"Levant", "Nimrod", "Ronin", "Scorpion", "Viper",
		"Lynx", "Napier", "Rigel", "Spartan", "Voyager",
		"Lotus", "Nemesis", "Reliant", "Shogun", "Vega",
		"Dirk", "Griffin", "Hornet", "Talon", "Wasp",
		"Demon", "Gargoyle", "Hunter", "Triton", "Wolf",
		"Delphos", "Gibbet", "Hansen", "Tirade", "Wight",
		"Dagon", "Gordon", "Hydra", "Tendril", "Welink"}[ne]
}

func findShipName(arg string) nameEnum {
	for ndx := 0; ndx < len(legalNames); ndx++ {
		if legalNames[ndx] == arg {
			return nameEnum(ndx)
		}
	}

	return nameEnum(unknownName)
}

func findShipTeam(arg nameEnum) (playerTeamEnum, shipEnum) {
	switch arg {
	case lazorName:
		fallthrough
	case nikeName:
		fallthrough
	case rapierName:
		fallthrough
	case saberName:
		fallthrough
	case vanirName:
		return blueTeam, scoutShip
	case levantName:
		fallthrough
	case nimrodName:
		fallthrough
	case roninName:
		fallthrough
	case scorpionName:
		fallthrough
	case viperName:
		return blueTeam, fighterShip
	case lynxName:
		fallthrough
	case napierName:
		fallthrough
	case rigelName:
		fallthrough
	case spartanName:
		fallthrough
	case voyagerName:
		return blueTeam, minerShip
	case lotusName:
		fallthrough
	case nemesisName:
		fallthrough
	case reliantName:
		fallthrough
	case shogunName:
		fallthrough
	case vegaName:
		return blueTeam, flagShip
	case dirkName:
		fallthrough
	case griffinName:
		fallthrough
	case hornetName:
		fallthrough
	case talonName:
		fallthrough
	case waspName:
		return redTeam, scoutShip
	case demonName:
		fallthrough
	case gargoyleName:
		fallthrough
	case hunterName:
		fallthrough
	case tritonName:
		fallthrough
	case wolfName:
		return redTeam, fighterShip
	case delphosName:
		fallthrough
	case gibbetName:
		fallthrough
	case hansenName:
		fallthrough
	case tiradeName:
		fallthrough
	case wightName:
		return redTeam, minerShip
	case dagonName:
		fallthrough
	case gordonName:
		fallthrough
	case hydraName:
		fallthrough
	case tendrilName:
		fallthrough
	case welinkName:
		return redTeam, flagShip
	default:
		return unknownTeam, unknownShip
	}
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
type shipType struct {
	active    bool
	condition conditionEnum
	docked    bool
	name      nameEnum
	position  locationType
	owner     string // player UUID
	shipClass shipEnum
	team      playerTeamEnum
	uuid      string // ship UUID

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

func shipAdd(st shipType, gt *gameType) int {
	for ndx := 0; ndx < maxShips; ndx++ {
		if gt.ships[ndx].active == false {
			gt.ships[ndx] = st
			return ndx
		}
	}

	return -1
}

func shipCensus(gt gameType) (int, int) {
	blue := 0
	red := 0

	for ndx := 0; ndx < maxShips; ndx++ {
		if gt.players[ndx].active == true {
			switch gt.ships[ndx].team {
			case blueTeam:
				blue++
			case redTeam:
				red++
			}
		}
	}

	return blue, red
}

func shipDelete(target string, gt *gameType) int {
	for ndx := 0; ndx < maxShips; ndx++ {
		if strings.Compare(gt.ships[ndx].uuid, target) == 0 {
			gt.ships[ndx].active = false
			return ndx
		}
	}

	return -1
}

func shipDump(gt gameType) {
	log.Println("=-=-=-= shipDump =-=-=-=")

	for ndx := 0; ndx < maxShips; ndx++ {
		shipClass := gt.ships[ndx].shipClass.string()
		shipName := gt.ships[ndx].name.string()
		shipTeam := gt.ships[ndx].team.string()
		log.Printf("%d %t %s %s %s %s", ndx, gt.ships[ndx].active, shipName, shipClass, shipTeam, gt.ships[ndx].uuid)
	}

	log.Println("=-=-=-= shipDump =-=-=-=")
}

func shipFind(target string, gt *gameType) int {
	for ndx := 0; ndx < maxShips; ndx++ {
		if gt.ships[ndx].active == true {
			if strings.Compare(gt.ships[ndx].uuid, target) == 0 {
				return ndx
			}
		}
	}

	return -1
}

func shipFindByOwner(target string, gt *gameType) int {
	for ndx := 0; ndx < maxShips; ndx++ {
		if gt.ships[ndx].active == true {
			if strings.Compare(gt.ships[ndx].owner, target) == 0 {
				return ndx
			}
		}
	}

	return -1
}

func commandCreateShip(command commandType, gt *gameType) {
	log.Println("create ship")

	// convert structure
	st := shipType{active: true, condition: greenCondition, docked: false, owner: command.player}
	st.name = findShipName(command.args[1])
	st.team, st.shipClass = findShipTeam(st.name)
	st.uuid = uuid.NewString()

	st.computer = 100
	st.lifeSupport = 100
	st.radio = 100
	st.shields = 100
	st.tractorBeam = 100

	st.impulseEngines = 100
	st.warpEngines = 100

	st.phasers = 100
	st.torpedoTubes = 100

	switch st.shipClass {
	case scoutShip:
		log.Println("scout")
		st.energy = 1
		st.mines = 2
		st.torpedos = 0
	case fighterShip:
		log.Println("fighter")
		st.energy = 1
		st.mines = 2
		st.torpedos = 0
	case minerShip:
		log.Println("miner")
		st.energy = 1
		st.mines = 2
		st.torpedos = 0
	case flagShip:
		log.Println("flagship")
		st.energy = 1
		st.mines = 2
		st.torpedos = 0
	default:
		log.Println("must throw error")
	}

	// add to ships array
	ndx := shipAdd(st, gt)
	log.Println(ndx)

	shipDump(*gt)
}

func commandMoveShip(command commandType, gt *gameType) {
	log.Println("move ship")

	playerNdx := playerFind(command.player, gt)
	if playerNdx < 0 {
		log.Println("unknown player")
		// return
	}

	log.Println("player noted")
	player := gt.players[playerNdx]
	log.Println(player.name)

	shipNdx := shipFindByOwner(command.player, gt)
	if shipNdx < 0 {
		log.Println("unknown ship")
		// return
	}
	ship := gt.ships[shipNdx]
	log.Println(ship)

	//	{"player":"player1uuid", "requestId":"request1uuid", "command":["move", "3", "3"]}
}

// NewShip ry
/*
func NewShip(name, owner string, se shipEnum, team playerTeamEnum) *shipType {
	result := shipType{active: true, name: name, owner: owner, shipClass: se, team: team}
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
*/

/*
func moveShip(ship Ship, location Location) {
	if ship.Active == false {
		log.Println("must throw error")
	}

	//var distance = 123
}
*/
