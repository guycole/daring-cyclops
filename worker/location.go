// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found
// in the LICENSE file.
package main

import (
	"math"
)

//row column origin = 0,0 lower left corner of map
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
