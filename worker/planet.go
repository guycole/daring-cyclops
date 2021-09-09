// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found
// in the LICENSE file.
package main

import (
	"github.com/google/uuid"
)

type planetType struct {
	position *locationType
	uuid     string
}

const maxPlanets = 255

type planetArrayType [maxStars]*planetType

func newPlanet(position *locationType) *planetType {
	result := planetType{position: position}
	result.uuid = uuid.NewString()
	return &result
}
