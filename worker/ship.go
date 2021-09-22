// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found
// in the LICENSE file.
package main

import (
	"errors"
	"log"
	"strings"

	"github.com/google/uuid"
)

// shipConditionEnum
type shipConditionEnum int

const (
	unknownCondition shipConditionEnum = iota
	greenCondition
	yellowCondition
	redCondition
)

// must match order for shipConditionEnum
func (sce shipConditionEnum) string() string {
	return [...]string{"unknown", "green", "yellow", "red"}[sce]
}

type shipClassEnum int

// must match order for legalShipClasses
const (
	unknownShip shipClassEnum = iota
	scoutShip
	fighterShip
	minerShip
	flagShip
)

// must match order for shipEnum
var legalShipClasses = [...]string{
	"unknownShip",
	"scout",
	"fighter",
	"miner",
	"flag",
}

// must match order for shipClassEnum
func (sce shipClassEnum) string() string {
	return [...]string{"unknown", "scout", "fighter", "miner", "flagShip"}[sce]
}

func findShipClass(arg string) shipClassEnum {
	for ndx := 0; ndx < len(legalShipClasses); ndx++ {
		if legalShipClasses[ndx] == arg {
			return shipClassEnum(ndx)
		}
	}

	return shipClassEnum(unknownShip)
}

type shipNameEnum int

// must match order for legalShipNames
const (
	unknownShipName shipNameEnum = iota
	lazorShipName
	nikeShipName
	rapierShipName
	saberShipName
	vanirShipName
	levantShipName
	nimrodShipName
	roninShipName
	scorpionShipName
	viperShipName
	lynxShipName
	napierShipName
	rigelShipName
	spartanShipName
	voyagerShipName
	lotusShipName
	nemesisShipName
	reliantShipName
	shogunShipName
	vegaShipName
	dirkShipName
	griffinShipName
	hornetShipName
	talonShipName
	waspShipName
	demonShipName
	gargoyleShipName
	hunterShipName
	tritonShipName
	wolfShipName
	delphosShipName
	gibbetShipName
	hansenShipName
	tiradeShipName
	wightShipName
	dagonShipName
	gordonShipName
	hydraShipName
	tendrilShipName
	welinkShipName
)

// must match order for shipNameEnum
var legalShipNames = [...]string{
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

// must match order for shipNameEnum
func (sne shipNameEnum) string() string {
	return [...]string{"unknown",
		"Lazor", "Nike", "Rapier", "Saber", "Vanir",
		"Levant", "Nimrod", "Ronin", "Scorpion", "Viper",
		"Lynx", "Napier", "Rigel", "Spartan", "Voyager",
		"Lotus", "Nemesis", "Reliant", "Shogun", "Vega",
		"Dirk", "Griffin", "Hornet", "Talon", "Wasp",
		"Demon", "Gargoyle", "Hunter", "Triton", "Wolf",
		"Delphos", "Gibbet", "Hansen", "Tirade", "Wight",
		"Dagon", "Gordon", "Hydra", "Tendril", "Welink"}[sne]
}

func findShipName(arg string) shipNameEnum {
	for ndx := 0; ndx < len(legalShipNames); ndx++ {
		if legalShipNames[ndx] == arg {
			return shipNameEnum(ndx)
		}
	}

	return shipNameEnum(unknownShipName)
}

func findShipClassTeam(arg shipNameEnum) (shipClassEnum, teamEnum) {
	switch arg {
	case lazorShipName:
		fallthrough
	case nikeShipName:
		fallthrough
	case rapierShipName:
		fallthrough
	case saberShipName:
		fallthrough
	case vanirShipName:
		return scoutShip, blueTeam
	case levantShipName:
		fallthrough
	case nimrodShipName:
		fallthrough
	case roninShipName:
		fallthrough
	case scorpionShipName:
		fallthrough
	case viperShipName:
		return fighterShip, blueTeam
	case lynxShipName:
		fallthrough
	case napierShipName:
		fallthrough
	case rigelShipName:
		fallthrough
	case spartanShipName:
		fallthrough
	case voyagerShipName:
		return minerShip, blueTeam
	case lotusShipName:
		fallthrough
	case nemesisShipName:
		fallthrough
	case reliantShipName:
		fallthrough
	case shogunShipName:
		fallthrough
	case vegaShipName:
		return flagShip, blueTeam
	case dirkShipName:
		fallthrough
	case griffinShipName:
		fallthrough
	case hornetShipName:
		fallthrough
	case talonShipName:
		fallthrough
	case waspShipName:
		return scoutShip, redTeam
	case demonShipName:
		fallthrough
	case gargoyleShipName:
		fallthrough
	case hunterShipName:
		fallthrough
	case tritonShipName:
		fallthrough
	case wolfShipName:
		return fighterShip, redTeam
	case delphosShipName:
		fallthrough
	case gibbetShipName:
		fallthrough
	case hansenShipName:
		fallthrough
	case tiradeShipName:
		fallthrough
	case wightShipName:
		return minerShip, redTeam
	case dagonShipName:
		fallthrough
	case gordonShipName:
		fallthrough
	case hydraShipName:
		fallthrough
	case tendrilShipName:
		fallthrough
	case welinkShipName:
		return flagShip, redTeam
	default:
		return unknownShip, unknownTeam
	}
}

type shipType struct {
	condition shipConditionEnum
	docked    bool
	shipName  shipNameEnum
	position  *locationType
	owner     string // player UUID
	shipClass shipClassEnum
	team      teamEnum
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

const maxShipTeam = 5
const maxShips = maxShipTeam * 2

// shipArrayType contains all active ships
type shipArrayType [maxShips]*shipType

// newShip convenience function to populate struct
func newShip(shipName, shipOwner string, position *locationType) (*shipType, error) {
	if position == nil {
		return nil, errors.New("nil position")
	}

	if len(shipName) < 1 {
		return nil, errors.New("empty ship name")
	}

	if len(shipOwner) < 1 {
		return nil, errors.New("empty ship owner")
	}

	shipName2 := findShipName(shipName)
	if shipName2 == unknownShipName {
		return nil, errors.New("unknown ship name")
	}

	shipClass, playerTeam := findShipClassTeam(shipName2)

	st := shipType{condition: greenCondition, shipName: shipName2, owner: shipOwner, shipClass: shipClass, position: position, team: playerTeam}
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

	// TODO resolve inventory values
	switch shipClass {
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
		return nil, errors.New("unknown ship class")
	}

	return &st, nil
}

const testShipName1 = "nike"
const testShipUuid1 = "ship1uuid"

// testShip1 returns test ship1
func testShip1(gt *gameType) *shipType {
	position := randomShipLocation(gt.board)
	ns1, _ := newShip(testShipName1, testPlayerID1, position)
	ns1.uuid = testShipUuid1
	return ns1
}

const testShipName2 = "welink"
const testShipUuid2 = "ship2uuid"

// testShip2 returns test ship2
func testShip2(gt *gameType) *shipType {
	position := randomShipLocation(gt.board)
	ns2, _ := newShip(testShipName2, testPlayerID2, position)
	ns2.uuid = testShipUuid2
	return ns2
}

// shipAdd adds ship to array
func shipAdd(st *shipType, sat *shipArrayType, bat *boardArrayType) int {
	log.Printf("shipAdd:%s %s", st.shipName.string(), st.uuid)

	symbol := st.shipName.string()[0:1]

	for ndx := 0; ndx < maxShips; ndx++ {
		if sat[ndx] == nil {
			setShip(bat[st.position.yy][st.position.xx], symbol, st.uuid)
			sat[ndx] = st
			return ndx
		}
	}

	return -1
}

// shipCensus returns population of red/blue ships
func shipCensus(sat shipArrayType) (int, int) {
	bluePopulation := 0
	redPopulation := 0

	for ndx := 0; ndx < maxShips; ndx++ {
		if sat[ndx] != nil {
			switch sat[ndx].team {
			case blueTeam:
				bluePopulation++
			case redTeam:
				redPopulation++
			}
		}
	}

	return bluePopulation, redPopulation
}

// shipDelete removes ship from array
func shipDelete(target string, sat *shipArrayType, bat *boardArrayType) int {
	log.Printf("shipDelete:%s", target)

	for ndx := 0; ndx < maxShips; ndx++ {
		if sat[ndx] != nil {
			if strings.Compare(sat[ndx].uuid, target) == 0 {
				clearShip(bat[sat[ndx].position.yy][sat[ndx].position.xx])
				sat[ndx] = nil
				return ndx
			}
		}
	}

	return -1
}

// shipDump diagnostic
func shipDump(sat shipArrayType) {
	log.Println("=-=-=-= shipDump =-=-=-=")

	for ndx := 0; ndx < maxShips; ndx++ {
		if sat[ndx] == nil {
			log.Printf("%d nil", ndx)
		} else {
			shipClass := sat[ndx].shipClass.string()
			shipShipName := sat[ndx].shipName.string()
			shipTeam := sat[ndx].team.string()
			log.Printf("%d %s %s %s %s", ndx, shipShipName, shipClass, shipTeam, sat[ndx].uuid)
		}
	}

	log.Println("=-=-=-= shipDump =-=-=-=")
}

// shipFind returns array index for ship by uuid
func shipFind(target string, sat shipArrayType) int {
	for ndx := 0; ndx < maxShips; ndx++ {
		if sat[ndx] != nil {
			if strings.Compare(sat[ndx].uuid, target) == 0 {
				return ndx
			}
		}
	}

	return -1
}

// shipFindByName returns array index for ship by name
func shipFindByName(target shipNameEnum, sat shipArrayType) int {
	for ndx := 0; ndx < maxShips; ndx++ {
		if sat[ndx] != nil {
			if sat[ndx].shipName == target {
				return ndx
			}
		}
	}

	return -1
}

// shipFindByOwner returns array index for ship by owner uuid
func shipFindByOwner(target string, sat shipArrayType) int {
	for ndx := 0; ndx < maxShips; ndx++ {
		if sat[ndx] != nil {
			if strings.Compare(sat[ndx].owner, target) == 0 {
				return ndx
			}
		}
	}

	return -1
}

// shipMove
func shipMove(shipID string, newLoc locationType, sat *shipArrayType, bat *boardArrayType) error {
	log.Printf("shipMove:%s", shipID)

	ndx := shipFind(shipID, *sat)
	if ndx < 0 {
		return errors.New("moveShip ship not found")
	}

	log.Println(sat[ndx])

	clearShip(bat[sat[ndx].position.yy][sat[ndx].position.xx])

	// need collision logic

	sat[ndx].position = &newLoc
	symbol := sat[ndx].shipName.string()[0:1]

	setShip(bat[sat[ndx].position.yy][sat[ndx].position.xx], symbol, sat[ndx].uuid)

	return nil
}

/*
// commandShipCreate services command
func commandShipCreate(ct commandType, gt *gameType) error {
	position := randomShipLocation(gt.board)
	st, err := newShip(ct.args[1], ct.player, position)
	if err != nil {
		return errors.New("commandShip creation failure")
	}

	singleOwner := shipFindByOwner(ct.player, gt.ships)
	if singleOwner >= 0 {
		return errors.New("commandShip duplicate player id")
	}

	duplicateShip := shipFindByName(st.shipName, gt.ships)
	if duplicateShip >= 0 {
		return errors.New("commandShip duplicate ship name")
	}

	bluePopulation, redPopulation := shipCensus(gt.ships)
	if st.team == blueTeam {
		if bluePopulation >= maxShipTeam {
			return errors.New("commandShip blue team population limit")
		}
	} else {
		if redPopulation >= maxShipTeam {
			return errors.New("commandShip red team population limit")
		}
	}

	shipAdd(st, &gt.ships, &gt.board)

	return nil
}
*/

/*
// commandShipDelete services command
func commandShipDelete(ct commandType, gt *gameType) error {
	owner := shipFindByOwner(ct.player, gt.ships)
	if owner < 0 {
		return errors.New("deleteShip player id not found")
	}

	shipDelete(gt.ships[owner].uuid, &gt.ships, &gt.board)

	return nil
}
*/

/*
// commandMoveShip services command
func commandMoveShip(ct commandType, gt *gameType) error {
	log.Println("move ship")

	ndx := shipFindByOwner(ct.player, gt.ships)
	if ndx < 0 {
		return errors.New("moveShip ship not found")
	}

	ship := gt.ships[ndx]
	log.Println(ship)

	//	{"player":"player1uuid", "requestId":"request1uuid", "command":["move", "3", "3"]}

	return nil
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
