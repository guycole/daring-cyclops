package game

import (
	"github.com/google/uuid"
)

const maxGateIndices = 9

var gate0yx = [...]int{8, 9}
var gate1yx = [...]int{8, 35}
var gate2yx = [...]int{8, 64}
var gate3yx = [...]int{35, 9}
var gate4yx = [...]int{35, 35}
var gate5yx = [...]int{35, 64}
var gate6yx = [...]int{64, 9}
var gate7yx = [...]int{64, 35}
var gate8yx = [...]int{64, 64}

//type gateIndicesArray[maxGateIndices] int

type starGateType struct {
	active   bool
	damage   int
	energy   int
	position *locationType
	gateID   int // 0 to 8
	uuid     string
}

func newStarGate(id int) *starGateType {
	result := starGateType{active: true, gateID: id}
	result.energy = 100 //tweak me
	result.uuid = uuid.NewString()

	switch id {
	case 0:
		result.position = newLocation(8, 9)
	case 1:
		result.position = newLocation(8, 35)
	case 2:
		result.position = newLocation(8, 64)
	case 3:
		result.position = newLocation(35, 9)
	case 4:
		result.position = newLocation(35, 35)
	case 5:
		result.position = newLocation(35, 64)
	case 6:
		result.position = newLocation(64, 9)
	case 7:
		result.position = newLocation(64, 35)
	case 8:
		result.position = newLocation(64, 64)
	}

	return &result
}

/*
   map origin lower left 1, 1

   0 1 2  (gate indices and relative locations)
   3 4 5
   6 7 8
*/

func starGateAdjacent(position *locationType) bool {
	var x, y int

	for ndx := 0; ndx < maxGateIndices; ndx++ {
		switch ndx {
		case 0:
			x = position.x - 1
			y = position.y + 1
		case 1:
			x = position.x
			y = position.y + 1
		case 2:
			x = position.x + 1
			y = position.y + 1
		case 3:
			x = position.x - 1
			y = position.y
		case 4:
			x = position.x
			y = position.y
		case 5:
			x = position.x + 1
			y = position.y
		case 6:
			x = position.x - 1
			y = position.y - 1
		case 7:
			x = position.x
			y = position.y - 1
		case 8:
			x = position.x + 1
			y = position.y - 1
		}

	}

	return false
}
