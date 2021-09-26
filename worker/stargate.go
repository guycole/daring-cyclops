// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import (
	"strings"

	"github.com/google/uuid"
)

type starGateType struct {
	active          bool
	damage          int
	energy          int
	position        *locationType
	gateDestination [9]int
	gateNdx         int
	uuid            string
}

const maxStarGates = 9

type starGateArrayType [maxStarGates]*starGateType

//row column origin = 0,0 lower left
var starGateLocations = [9][2]int{
	{8, 9},
	{8, 35},
	{8, 64},
	{35, 9},
	{35, 35},
	{35, 64},
	{64, 9},
	{64, 35},
	{64, 64},
}

/*
   0 1 2  (gate indices and relative locations)
   3 4 5  (4 = stargate)
   6 7 8
*/

var starGateDestinations = [9][9]int{
	{-1, 6, 8, 1, -1, 2, 4, 3, -1}, //0
	{-1, 7, -1, 2, -1, 0, 5, 4, 3}, //1
	{6, 8, -1, 0, -1, 1, -1, 5, 4}, //2
	{6, 0, -1, 4, -1, 5, 7, 6, -1}, //3
	{2, 6, 0, 5, -1, 3, 8, 7, 6},   //4
	{-1, 2, 6, 3, -1, 4, -1, 8, 7}, //5
	{4, 3, -1, 7, -1, 8, -1, 0, 2}, //6
	{5, 4, 3, 8, -1, 6, -1, 1, -1}, //7
	{-1, 5, 4, 6, -1, 7, 0, 2, -1}, //8
}

//type gateIndicesArray[maxGateIndices] int

// newStarGate convenience function to populate struct
func newStarGate(ndx int) *starGateType {
	result := starGateType{active: true, gateNdx: ndx}
	result.energy = 100 //tweak me
	result.position = newLocation(starGateLocations[ndx][0], starGateLocations[ndx][1])
	result.uuid = uuid.NewString()
	return &result
}

/*
   map origin lower left 1, 1

   0 1 2  (gate indices and relative locations)
   3 4 5
   6 7 8
*/
func starGateAdjacent(candidate *locationType) (gateNdx, locNdx int) {
	for sg := 0; sg < maxStarGates; sg++ {
		current := newLocation(starGateLocations[sg][0], starGateLocations[sg][1])
		ndx := testForAdjacency(current, candidate)
		if ndx >= 0 {
			return sg, ndx
		}
	}

	return -1, -1
}

// planetsAdd(pat *planetArrayType, bat *boardArrayType) {

// generate all stargates
func starGatesAdd(sat *starGateArrayType, bat *boardArrayType) {
	for ndx := 0; ndx < maxStarGates; ndx++ {
		sg := newStarGate(ndx)
		sat[ndx] = sg
		setStarGate(bat[sg.position.yy][sg.position.xx], sg.uuid)
	}
}

// starGateFind returns array index for planet by uuid
func starGateFind(target string, sat starGateArrayType) int {
	for ndx := 0; ndx < maxStarGates; ndx++ {
		if sat[ndx] != nil {
			if strings.Compare(sat[ndx].uuid, target) == 0 {
				return ndx
			}
		}
	}

	return -1
}
