package main

import "math"

type locationType struct {
	xx int // column
	yy int // row
}

func newLocation(y, x int) *locationType {
	result := locationType{yy: y, xx: x}
	return &result
}

func getDistance(origin, destination *locationType) int {
	var x = float64(origin.xx - destination.xx)
	var y = float64(origin.yy - destination.yy)
	return int(math.Round(math.Sqrt(x*x + y*y)))
}
