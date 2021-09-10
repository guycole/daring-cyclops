// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found
// in the LICENSE file.
package main

import (
	"testing"
)

func TestShipClass(t *testing.T) {
	tests := []struct {
		candidate string
		answer    shipClassEnum
	}{
		{"scout", scoutShip},
		{"flag", flagShip},
		{"bogus", unknownShip},
	}
	for _, ndx := range tests {
		result := findShipClass(ndx.candidate)
		if result != ndx.answer {
			t.Errorf("findShipClass(%s) failure", ndx.candidate)
		}
	}
}

func TestShipName(t *testing.T) {
	tests := []struct {
		candidate string
		answer    shipNameEnum
	}{
		{"lazor", lazorShipName},
		{"welink", welinkShipName},
		{"bogus", unknownShipName},
	}
	for _, ndx := range tests {
		result := findShipName(ndx.candidate)
		if result != ndx.answer {
			t.Errorf("findShipName(%s) failure", ndx.candidate)
		}
	}
}

func TestShipClassTeam(t *testing.T) {
	tests := []struct {
		candidate  shipNameEnum
		shipClass  shipClassEnum
		playerTeam teamEnum
	}{
		{lazorShipName, scoutShip, blueTeam},
		{welinkShipName, flagShip, redTeam},
	}
	for _, ndx := range tests {
		shipClass, playerTeam := findShipClassTeam(ndx.candidate)
		if shipClass != ndx.shipClass {
			t.Errorf("findShipClassTeam(%s) class failure", ndx.candidate.string())
		}
		if playerTeam != ndx.playerTeam {
			t.Errorf("findShipClassTeam(%s) team failure", ndx.candidate.string())
		}
	}
}

func TestNewOkShip(t *testing.T) {
	gt := newGame("testGame", emptyBoard)

	result, err := newShip("nike", testPlayerID1, gt)
	if err != nil {
		t.Errorf("newShip error:%s", err)
	}

	if result != nil {
		if result.condition != greenCondition {
			t.Error("newShip condition failure")
		}

		if result.docked != false {
			t.Error("newShip docked failure")
		}

		if result.shipName != nikeShipName {
			t.Error("newShip name failure")
		}

		if result.shipClass != scoutShip {
			t.Error("newShip class failure")
		}

		if result.team != blueTeam {
			t.Error("newShip team failure")
		}

		if result.owner != testPlayerID1 {
			t.Error("newShip owner failure")
		}
	} else {
		t.Error("newPlayer returns nil")
	}
}

func TestNewBadShip01(t *testing.T) {
	gt := newGame("testGame", emptyBoard)

	result, err := newShip("", testPlayerID2, gt)
	if err == nil {
		t.Error("newShip error:expecting bad shipName")
	}

	if result != nil {
		t.Error("newShip error expecting nil")
	}
}

func TestNewBadShip02(t *testing.T) {
	gt := newGame("testGame", emptyBoard)

	result, err := newShip("nike", "", gt)
	if err == nil {
		t.Error("newShipq error:expecting bad player ID")
	}

	if result != nil {
		t.Error("newShip error expecting nil")
	}
}

func TestNewBadShip03(t *testing.T) {
	gt := newGame("testGame", emptyBoard)

	result, err := newShip("bogus", testPlayerID2, gt)
	if err == nil {
		t.Error("newShip error:expecting bad shipName")
	}

	if result != nil {
		t.Error("newShip error expecting nil")
	}
}

func TestShipArray(t *testing.T) {
	var sat shipArrayType

	bluePopulation, redPopulation := shipCensus(sat)
	if bluePopulation != 0 && redPopulation != 0 {
		t.Errorf("shipCensus error:%d %d", bluePopulation, redPopulation)
	}

	gt := newGame("testGame", emptyBoard)

	ns1 := testShip1(gt)
	if ns1 == nil {
		t.Error("testShip1 returns nil")
	}

	ns2 := testShip2(gt)
	if ns2 == nil {
		t.Error("testShip2 returns nil")
	}

	// add ship to ship array, should be first array element
	ndx := shipAdd(ns1, &sat)
	if ndx != 0 {
		t.Errorf("shipAdd returns wrong index %d", ndx)
	}

	// add ship to ship array, should be second array element
	ndx = shipAdd(ns2, &sat)
	if ndx != 1 {
		t.Errorf("shipAddAdd returns wrong index %d", ndx)
	}

	bluePopulation, redPopulation = shipCensus(sat)
	if bluePopulation != 1 && redPopulation != 1 {
		t.Errorf("shipCensus error:%d %d", bluePopulation, redPopulation)
	}

	// shipDump(sat)

	ndx = shipFind(testShipUuid1, sat)
	if ndx != 0 {
		t.Errorf("shipFind returns wrong index %d", ndx)
	}

	ndx = shipFindByName(nikeShipName, sat)
	if ndx != 0 {
		t.Errorf("shipFindByName returns wrong index %d", ndx)
	}

	ndx = shipFindByOwner(testPlayerID1, sat)
	if ndx != 0 {
		t.Errorf("shipFindByOwner returns wrong index %d", ndx)
	}

	ndx = shipFind(testShipUuid2, sat)
	if ndx != 1 {
		t.Errorf("shipFind returns wrong index %d", ndx)
	}

	ndx = shipFindByName(welinkShipName, sat)
	if ndx != 1 {
		t.Errorf("shipFindByName returns wrong index %d", ndx)
	}

	ndx = shipFindByOwner(testPlayerID2, sat)
	if ndx != 1 {
		t.Errorf("shipFindByOwner returns wrong index %d", ndx)
	}

	ndx = shipFind("bogus", sat)
	if ndx >= 0 {
		t.Errorf("shipFind returns wrong index %d", ndx)
	}

	ndx = shipFindByOwner("bogus", sat)
	if ndx >= 0 {
		t.Errorf("shipFindByOwner returns wrong index %d", ndx)
	}

	ndx = shipDelete(testShipUuid1, &sat)
	if ndx != 0 {
		t.Errorf("shipDelete returns wrong index %d", ndx)
	}

	bluePopulation, redPopulation = shipCensus(sat)
	if bluePopulation != 0 && redPopulation != 1 {
		t.Errorf("shipCensus error:%d %d", bluePopulation, redPopulation)
	}
}

func TestCreateDeleteShip(t *testing.T) {
	ct := commandType{player: testPlayerID2, request: "requestId"}
	ct.args = []string{"shipCreate", "nimrod"}
	ct.command = shipCreateCommand

	gt := newGame("testGame", emptyBoard)

	err := commandShipCreate(ct, gt)
	if err != nil {
		t.Errorf("commandCreateShip error:%s", err)
	}

	// shipDump(gt.ships)

	bluePopulation, redPopulation := shipCensus(gt.ships)
	if bluePopulation != 1 && redPopulation != 0 {
		t.Errorf("shipCensus error:%d %d", bluePopulation, redPopulation)
	}

	err = commandShipCreate(ct, gt)
	if err == nil {
		t.Errorf("commandShipCreate should have duplicate error")
	}

	//shipDump(gt.ships)

	ct = commandType{player: testPlayerID2, request: "requestId"}
	ct.args = []string{"shipDelete"}
	ct.command = shipDeleteCommand

	commandShipDelete(ct, gt)

	//shipDump(gt.ships)
}
