// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
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
	/*
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
	*/
}

func TestNewOkShip(t *testing.T) {
	gt := newGame("testGame", emptyBoard)

	position := randomShipLocation(gt.board)

	result, err := newShip("nike", testPlayerName1, position)
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

		if result.nameEnum != nikeShipName {
			t.Error("newShip name failure")
		}

		if result.classEnum != scoutShip {
			t.Error("newShip class failure")
		}

		if result.team != blueTeam {
			t.Error("newShip team failure")
		}

		if result.owner != testPlayerName1 {
			t.Error("newShip owner failure")
		}
	} else {
		t.Error("newPlayer returns nil")
	}
}

func TestNewBadShip01(t *testing.T) {
	gt := newGame("testGame", emptyBoard)

	position := randomShipLocation(gt.board)

	result, err := newShip("", testPlayerName2, position)
	if err == nil {
		t.Error("newShip error:expecting bad shipName")
	}

	if result != nil {
		t.Error("newShip error expecting nil")
	}
}

func TestNewBadShip02(t *testing.T) {
	gt := newGame("testGame", emptyBoard)

	position := randomShipLocation(gt.board)

	result, err := newShip("nike", "", position)
	if err == nil {
		t.Error("newShipq error:expecting bad player name")
	}

	if result != nil {
		t.Error("newShip error expecting nil")
	}
}

func TestNewBadShip03(t *testing.T) {
	gt := newGame("testGame", emptyBoard)

	position := randomShipLocation(gt.board)

	result, err := newShip("bogus", testPlayerName2, position)
	if err == nil {
		t.Error("newShip error:expecting bad shipName")
	}

	if result != nil {
		t.Error("newShip error expecting nil")
	}
}

func TestShipArray(t *testing.T) {
	gt := newGame("testGame", emptyBoard)

	bluePopulation, redPopulation := gt.ships.census()
	if bluePopulation != 0 && redPopulation != 0 {
		t.Errorf("shipCensus error:%d %d", bluePopulation, redPopulation)
	}

	ns1 := gt.testShip1()
	if ns1 == nil {
		t.Error("testShip1 returns nil")
	}

	ns2 := gt.testShip2()
	if ns2 == nil {
		t.Error("testShip2 returns nil")
	}

	// add ship to ship array, should be first array element
	ndx := gt.ships.add(ns1, &gt.board)
	if ndx != 0 {
		t.Errorf("shipAdd returns wrong index %d", ndx)
	}

	// add ship to ship array, should be second array element
	ndx = gt.ships.add(ns2, &gt.board)
	if ndx != 1 {
		t.Errorf("shipAddAdd returns wrong index %d", ndx)
	}

	bluePopulation, redPopulation = gt.ships.census()
	if bluePopulation != 1 && redPopulation != 1 {
		t.Errorf("shipCensus error:%d %d", bluePopulation, redPopulation)
	}

	// shipDump(sat)

	ndx = gt.ships.find(testShipUuid1)
	if ndx != 0 {
		t.Errorf("shipFind returns wrong index %d", ndx)
	}

	ndx = gt.ships.findByName(nikeShipName)
	if ndx != 0 {
		t.Errorf("shipFindByName returns wrong index %d", ndx)
	}

	ndx = gt.ships.findByOwner(testPlayerName1)
	if ndx != 0 {
		t.Errorf("shipFindByOwner returns wrong index %d", ndx)
	}

	ndx = gt.ships.find(testShipUuid2)
	if ndx != 1 {
		t.Errorf("shipFind returns wrong index %d", ndx)
	}

	ndx = gt.ships.findByName(welinkShipName)
	if ndx != 1 {
		t.Errorf("shipFindByName returns wrong index %d", ndx)
	}

	ndx = gt.ships.findByOwner(testPlayerName2)
	if ndx != 1 {
		t.Errorf("shipFindByOwner returns wrong index %d", ndx)
	}

	ndx = gt.ships.find("bogus")
	if ndx >= 0 {
		t.Errorf("shipFind returns wrong index %d", ndx)
	}

	ndx = gt.ships.findByOwner("bogus")
	if ndx >= 0 {
		t.Errorf("shipFindByOwner returns wrong index %d", ndx)
	}

	ndx = gt.ships.delete(testShipUuid1, &gt.board)
	if ndx != 0 {
		t.Errorf("shipDelete returns wrong index %d", ndx)
	}

	bluePopulation, redPopulation = gt.ships.census()
	if bluePopulation != 0 && redPopulation != 1 {
		t.Errorf("shipCensus error:%d %d", bluePopulation, redPopulation)
	}
}

func TestCreateDeleteShip(t *testing.T) {
	gt := newGame("testGame", emptyBoard)

	var commands commandArrayType
	commands[0] = "shipCreate"
	commands[1] = "nike"

	ct := newCommand(testPlayerName1, "reqId", 2, commands)
	tnt := newTurnNode(ct)

	// create fresh ship
	err := commandShipCreate(tnt, &gt.board, &gt.ships)
	if err != nil {
		t.Errorf("commandCreateShip error:%s", err)
	}

	gt.ships.dump()

	// verify ship population
	bluePopulation, redPopulation := gt.ships.census()
	if bluePopulation != 1 && redPopulation != 0 {
		t.Errorf("shipCensus error:%d %d", bluePopulation, redPopulation)
	}

	// should reject duplicate ship
	err = commandShipCreate(tnt, &gt.board, &gt.ships)
	if err == nil {
		t.Errorf("commandShipCreate should have duplicate error")
	}

	gt.ships.dump()

	commands[0] = "shipDelete"

	ct = newCommand(testPlayerName2, "reqId", 1, commands)
	tnt = newTurnNode(ct)

	// delete of nonexisting player should fail
	err = commandShipDelete(tnt, &gt.board, &gt.ships)
	if err == nil {
		t.Error("commandDeleteShip should have complained about nonexist player")
	}

	ct = newCommand(testPlayerName1, "reqId", 1, commands)
	tnt = newTurnNode(ct)

	// delete of ship should succeed
	err = commandShipDelete(tnt, &gt.board, &gt.ships)
	if err != nil {
		t.Error("commandDeleteShip should have succeeded")
	}

	gt.ships.dump()

	// verify ship population
	bluePopulation, redPopulation = gt.ships.census()
	if bluePopulation != 0 && redPopulation != 0 {
		t.Errorf("shipCensus error:%d %d", bluePopulation, redPopulation)
	}
}

func TestCreateMoveShip(t *testing.T) {
	gt := newGame("testGame", emptyBoard)

	position1 := newLocation(36, 36)
	position2 := newLocation(40, 33)

	ns1, err := newShip("nike", testPlayerName1, position1)
	if err != nil {
		t.Errorf("newShip error:%s", err)
	}

	gt.ships.add(ns1, &gt.board)

	bc := gt.board[position1.yy][position1.xx]
	if !bc.ship {
		t.Error("board not contain ship at position1")
	}

	gt.ships.move(ns1.uuid, position2, &gt.board)

	bc = gt.board[position1.yy][position1.xx]
	if bc.ship {
		t.Error("board should not contain ship at position1")
	}

	bc = gt.board[position2.yy][position2.xx]
	if !bc.ship {
		t.Error("board not contain ship at position2")
	}

	if ns1.position.yy != position2.yy || ns1.position.xx != position2.xx {
		t.Error("ship position not match expected position")
	}
}
