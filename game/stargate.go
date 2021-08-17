package game

import (
	"log"

	"github.com/google/uuid"
)

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

type starGateType struct {
	active          bool
	damage          int
	energy          int
	position        *locationType
	gateDestination [9]int
	gateNdx         int
	uuid            string
}

func newStarGate(ndx int) *starGateType {
	result := starGateType{active: true, gateNdx: ndx}
	result.energy = 100 //tweak me
	result.position = newLocation(starGateLocations[ndx][0], starGateLocations[ndx][1])
	result.uuid = uuid.NewString()

	log.Println(starGateLocations[ndx][0])

	return &result
}

/*
   map origin lower left 1, 1

   0 1 2  (gate indices and relative locations)
   3 4 5
   6 7 8
*/

// starGateAdjacent discover if next to a SG.  Returns index into starGateDestinations
func starGateAdjacent(shipPosition, starGatePosition *locationType) int {
	var x, y int

	for ndx := 0; ndx < 9; ndx++ {
		switch ndx {
		case 0:
			x = starGatePosition.x - 1
			y = starGatePosition.y + 1
		case 1:
			x = starGatePosition.x
			y = starGatePosition.y + 1
		case 2:
			x = starGatePosition.x + 1
			y = starGatePosition.y + 1
		case 3:
			x = starGatePosition.x - 1
			y = starGatePosition.y
		case 4: // should never match
			x = starGatePosition.x
			y = starGatePosition.y
		case 5:
			x = starGatePosition.x + 1
			y = starGatePosition.y
		case 6:
			x = starGatePosition.x - 1
			y = starGatePosition.y - 1
		case 7:
			x = starGatePosition.x
			y = starGatePosition.y - 1
		case 8:
			x = starGatePosition.x + 1
			y = starGatePosition.y - 1
		}

		temp := newLocation(y, x)
		if temp.x == shipPosition.x && temp.y == shipPosition.y {
			return ndx
		}

	}

	return -1
}
