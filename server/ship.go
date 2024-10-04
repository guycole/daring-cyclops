// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package server

import (
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
	energy    uint16
	mines     uint16
	torpedoes uint16
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
	computer       uint16 // ship system as percentage operational (values 100 to 0)
	condition      shipConditionEnum
	dockedFlag     bool
	energy         uint16 // generic ship power
	impulseEngines uint16
	key            *tokenKeyType // unique identifier
	lifeSupport    uint16        // ship systems as percentage operational (values 100 to 0)
	location       *locationType // current ship location
	mines          uint16        // inventory
	name           shipNameEnum
	phasers        uint16 // ship systems as percentage operational (values 100 to 0)
	radio          uint16 // ship systems as percentage operational (values 100 to 0)
	shields        uint16 // ship systems as percentage operational (values 100 to 0)
	shipClass      shipClassEnum
	symbol         string
	torpedoes      uint16 // inventory
	torpedoTubes   uint16 // ship systems as percentage operational (values 100 to 0)
	team           teamEnum
	tractorBeam    uint16 // ship systems as percentage operational (values 100 to 0)
	warpEngines    uint16 // ship systems as percentage operational (values 100 to 0)
}

// newShip convenience function to populate struct
func newShip(name shipNameEnum) *shipType {
	st := shipType{shipClass: legalShips[name].shipClass, name: name, symbol: legalShips[name].symbol, team: legalShips[name].team}

	st.computer = 100
	st.condition = greenCondition
	st.dockedFlag = false
	st.energy = legalShipInventory[st.shipClass].energy
	st.impulseEngines = 100
	st.key = newTokenKey("")
	st.lifeSupport = 100
	st.location = randomLocation(maxBoardSideRow, maxBoardSideCol)
	st.mines = legalShipInventory[st.shipClass].mines
	st.phasers = 100
	st.radio = 100
	st.shields = 100
	st.torpedoTubes = 100
	st.torpedoes = legalShipInventory[st.shipClass].torpedoes
	st.tractorBeam = 100
	st.warpEngines = 100

	return &st
}
