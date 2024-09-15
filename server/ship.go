// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package server

import (
	"errors"
	"strings"
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

// must match order for shipClassEnum
var legalShipClasses = [...]string{
	"unknownShip",
	"scout",
	"fighter",
	"miner",
	"flag",
}

type legalShipInventoryType struct {
	energy    int
	mines     int
	torpedoes int
}

// must match order for shipClassEnum
var legalShipInventory = [...]legalShipInventoryType{
	{0, 0, 0},
	{1, 2, 3},
	{10, 20, 30},
	{100, 200, 300},
	{1000, 2000, 3000},
}

// must match order for shipClassEnum
func (sce shipClassEnum) string() string {
	return [...]string{"unknown", "scout", "fighter", "miner", "flagShip"}[sce]
}

func findShipClass(arg string) shipClassEnum {
	for ndx := 0; ndx < len(legalShipClasses); ndx++ {
		if strings.Compare(legalShipClasses[ndx], arg) == 0 {
			return shipClassEnum(ndx)
		}
	}

	return shipClassEnum(unknownShip)
}

type shipNameEnum int

// must match order for legalShips
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

type legalShipType struct {
	name      string
	shipClass shipClassEnum
	symbol    string
	team      teamEnum
}

// must match order for shipNameEnum
var legalShips = [...]legalShipType{
	{"unknown", unknownShip, "?", unknownTeam},
	{"lazor", scoutShip, "L", blueTeam},
	{"nike", scoutShip, "N", blueTeam},
	{"rapier", scoutShip, "R", blueTeam},
	{"saber", scoutShip, "S", blueTeam},
	{"vanir", scoutShip, "V", blueTeam},
	{"levant", fighterShip, "L", blueTeam},
	{"nimrod", fighterShip, "N", blueTeam},
	{"ronin", fighterShip, "R", blueTeam},
	{"scorpion", fighterShip, "S", blueTeam},
	{"viper", fighterShip, "V", blueTeam},
	{"lynx", minerShip, "L", blueTeam},
	{"napier", minerShip, "N", blueTeam},
	{"rigel", minerShip, "R", blueTeam},
	{"spartan", minerShip, "S", blueTeam},
	{"voyager", minerShip, "V", blueTeam},
	{"lotus", flagShip, "L", blueTeam},
	{"nemesis", flagShip, "N", blueTeam},
	{"reliant", flagShip, "R", blueTeam},
	{"shogun", flagShip, "S", blueTeam},
	{"vega", flagShip, "V", blueTeam},
	{"dirk", scoutShip, "D", redTeam},
	{"griffin", scoutShip, "G", redTeam},
	{"hornet", scoutShip, "H", redTeam},
	{"talon", scoutShip, "T", redTeam},
	{"wasp", scoutShip, "W", redTeam},
	{"demon", fighterShip, "D", redTeam},
	{"gargoyle", fighterShip, "G", redTeam},
	{"hunter", fighterShip, "H", redTeam},
	{"triton", fighterShip, "T", redTeam},
	{"wolf", fighterShip, "W", redTeam},
	{"delphos", minerShip, "D", redTeam},
	{"gibbet", minerShip, "G", redTeam},
	{"hansen", minerShip, "H", redTeam},
	{"tirade", minerShip, "T", redTeam},
	{"wight", minerShip, "W", redTeam},
	{"dagon", flagShip, "D", redTeam},
	{"gordon", flagShip, "G", redTeam},
	{"hydra", flagShip, "H", redTeam},
	{"tendril", flagShip, "T", redTeam},
	{"welink", flagShip, "W", redTeam},
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
	for ndx := 0; ndx < len(legalShips); ndx++ {
		if strings.Compare(legalShips[ndx].name, arg) == 0 {
			return shipNameEnum(ndx)
		}
	}

	return shipNameEnum(unknownShipName)
}

type shipType struct {
	classEnum shipClassEnum
	condition shipConditionEnum
	docked    bool
	nameEnum  shipNameEnum
	//	position  *locationType
	owner  string
	symbol string
	team   teamEnum
	uuid   string // ship UUID

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
	mines     int
	torpedoes int

	// generic ship energy
	energy int
}

/*
const maxTeamShips = maxTeamPlayers
const maxShips = maxTeamShips * 2

// shipArrayType contains all active ships
type shipArrayType [maxShips]*shipType
*/

// newShip convenience function to populate struct
func newShip(name shipNameEnum) (*shipType, error) {
	st := shipType{}
	return &st, nil
}

func newShip2(shipName, shipOwner string, position *locationType) (*shipType, error) {
	if position == nil {
		return nil, errors.New("nil position")
	}

	if len(shipName) < 1 {
		return nil, errors.New("empty ship name")
	}

	if len(shipOwner) < 1 {
		return nil, errors.New("empty ship owner")
	}

	shipEnum := findShipName(shipName)
	if shipEnum == unknownShipName {
		return nil, errors.New("unknown ship name")
	}

	st := shipType{}

	/*
		st := shipType{condition: greenCondition, owner: shipOwner, position: position}
		st.nameEnum = shipEnum
		st.uuid = uuid.NewString()

		legalShip := legalShips[shipEnum]
		st.classEnum = legalShip.shipClass
		st.symbol = legalShip.symbol
		st.team = legalShip.team

		// all systems 100 percent effective
		st.computer = 100
		st.lifeSupport = 100
		st.radio = 100
		st.shields = 100
		st.tractorBeam = 100

		st.impulseEngines = 100
		st.warpEngines = 100

		st.phasers = 100
		st.torpedoTubes = 100

		// inventory
		inventory := legalShipInventory[legalShip.shipClass]
		st.energy = inventory.energy
		st.mines = inventory.mines
		st.torpedoes = inventory.torpedoes
	*/

	return &st, nil
}

/*
// shipAdd adds ship to array
func (sat *shipArrayType) add(st *shipType, bat *boardArrayType) int {
	log.Printf("shipAdd:%s %s", st.nameEnum.string(), st.uuid)

	for ndx := 0; ndx < maxShips; ndx++ {
		if sat[ndx] == nil {
			bc := bat[st.position.yy][st.position.xx]
			bc.setShip(st.symbol, st.uuid)
			sat[ndx] = st
			return ndx
		}
	}

	return -1
}
*/

/*
// shipCensus returns population of red/blue ships
func (sat shipArrayType) census() (int, int) {
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
*/

/*
// shipDelete removes ship from array
func (sat *shipArrayType) delete(target string, bat *boardArrayType) int {
	log.Printf("shipDelete:%s", target)

	for ndx := 0; ndx < maxShips; ndx++ {
		if sat[ndx] != nil {
			if strings.Compare(sat[ndx].uuid, target) == 0 {
				bc := bat[sat[ndx].position.yy][sat[ndx].position.xx]
				bc.clearShip()
				sat[ndx] = nil
				return ndx
			}
		}
	}

	return -1
}
*/

/*
// shipDump diagnostic
func (sat shipArrayType) dump() {
	log.Println("=-=-=-= shipDump =-=-=-=")

	for ndx := 0; ndx < maxShips; ndx++ {
		if sat[ndx] == nil {
			log.Printf("%d nil", ndx)
		} else {
			shipClass := sat[ndx].classEnum.string()
			shipName := sat[ndx].nameEnum.string()
			shipTeam := sat[ndx].team.string()
			log.Printf("%d %s %s %s %s", ndx, shipName, shipClass, shipTeam, sat[ndx].uuid)
		}
	}

	log.Println("=-=-=-= shipDump =-=-=-=")
}
*/

/*
// shipFind returns array index for ship by uuid
func (sat shipArrayType) find(target string) int {
	for ndx := 0; ndx < maxShips; ndx++ {
		if sat[ndx] != nil {
			if strings.Compare(sat[ndx].uuid, target) == 0 {
				return ndx
			}
		}
	}

	return -1
}
*/

/*
// shipFindByName returns array index for ship by name
func (sat shipArrayType) findByName(target shipNameEnum) int {
	for ndx := 0; ndx < maxShips; ndx++ {
		if sat[ndx] != nil {
			if sat[ndx].nameEnum == target {
				return ndx
			}
		}
	}

	return -1
}
*/

/*
// shipFindByOwner returns array index for ship by owner uuid
func (sat shipArrayType) findByOwner(target string) int {
	for ndx := 0; ndx < maxShips; ndx++ {
		if sat[ndx] != nil {
			if strings.Compare(sat[ndx].owner, target) == 0 {
				return ndx
			}
		}
	}

	return -1
}
*/

/*
// shipMove
func (sat *shipArrayType) move(shipID string, newLoc *locationType, bat *boardArrayType) error {
	log.Printf("shipMove:%s:%d:%d", shipID, newLoc.yy, newLoc.xx)

	ndx := sat.find(shipID)
	if ndx < 0 {
		return errors.New("moveShip ship not found")
	}

	log.Println(sat[ndx])

	bc := bat[sat[ndx].position.yy][sat[ndx].position.xx]

	bc.clearShip()

	// TODO need collision logic

	sat[ndx].position = newLoc

	bc = bat[sat[ndx].position.yy][sat[ndx].position.xx]

	bc.setShip(sat[ndx].symbol, sat[ndx].uuid)

	return nil
}
*/

/*
func (sat *shipArrayType) condition(shipID string) error {
	log.Printf("condition:%s", shipID)

	ndx := sat.find(shipID)
	if ndx < 0 {
		return errors.New("condition ship not found")
	}

	sat[ndx].condition = greenCondition

	return nil
}
*/

/*
func commandShipCreate(tnt *turnNodeType, bat *boardArrayType, sat *shipArrayType) error {
	singleOwner := sat.findByOwner(tnt.name)
	if singleOwner >= 0 {
		return errors.New("commandShip duplicate player id")
	}

	position := bat.randomShipLocation()

	st, err := newShip(tnt.arguments[1], tnt.name, position)
	if err != nil {
		return errors.New("commandShip creation failure")
	}

	duplicateShip := sat.findByName(st.nameEnum)
	if duplicateShip >= 0 {
		return errors.New("commandShip duplicate ship name")
	}

	bluePopulation, redPopulation := sat.census()
	if st.team == blueTeam {
		if bluePopulation >= maxTeamShips {
			return errors.New("commandShip blue team population limit")
		}
	} else {
		if redPopulation >= maxTeamShips {
			return errors.New("commandShip red team population limit")
		}
	}

	sat.add(st, bat)

	return nil
}
*/

/*
func commandShipDelete(tnt *turnNodeType, bat *boardArrayType, sat *shipArrayType) error {
	owner := sat.findByOwner(tnt.name)
	if owner < 0 {
		return errors.New("deleteShip player id not found")
	}

	sat.delete(sat[owner].uuid, bat)

	return nil
}
*/

/*
func commandShipMove(tnt *turnNodeType, bat *boardArrayType, sat *shipArrayType) (*RequestType, error) {
	ndx := sat.findByOwner(tnt.name)
	if ndx < 0 {
		return nil, errors.New("moveShip player id not found")
	}

	newLocation := stringLocation(tnt.arguments[1], tnt.arguments[2])
	if newLocation == nil {
		return nil, errors.New("moveShip bad location")
	}

	err := sat.move(sat[ndx].uuid, newLocation, bat)
	if err != nil {
		return nil, err
	}

	err = sat.condition(sat[ndx].uuid)
	return nil, err

	/////
	var commands argumentArrayType
	commands[0] = "pong"

	ct := newRequest(tnt.name, tnt.request, 1, commands)

	return ct, nil
}
*/
