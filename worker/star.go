// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found
// in the LICENSE file.
package main

import (
	"github.com/google/uuid"
)

type starType struct {
	position *locationType
	uuid     string
}

const maxStars = 9

type starArrayType [maxStars]*starType

//row column origin = 0,0 lower left
var starLocations = [9][2]int{
	{0, 32},
	{1, 21},
	{2, 27},
	{2, 45},
	{4, 34},
	{6, 39},
	{8, 25},
	{8, 43},
	{8, 50},
}

func newStar(ndx int) *starType {
	result := starType{}
	result.position = newLocation(starLocations[ndx][0], starLocations[ndx][1])
	result.uuid = uuid.NewString()

	return &result
}
