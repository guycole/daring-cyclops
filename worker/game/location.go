package game

import "math"

type locationType struct {
	x int
	y int
}

func newLocation(y, x int) *locationType {
	result := locationType{y: y, x: x}
	return &result
}

func getDistance(origin, destination *locationType) int {
	var x = float64(origin.x - destination.x)
	var y = float64(origin.y - destination.y)
	return int(math.Round(math.Sqrt(x*x + y*y)))
}
