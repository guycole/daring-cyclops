// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import (
	"log"
	"math/rand"
	"strings"

	"github.com/google/uuid"
)

type starType struct {
	name     string
	position *locationType
	uuid     string
}

const maxStars = 255

// contains all active stars
type starArrayType [maxStars]*starType

// newStar convenience function to populate struct
func newStar(position *locationType) *starType {
	result := starType{position: position}
	result.name = nameGenerator(true)
	result.uuid = uuid.NewString()
	return &result
}

// generate all stars
func starsAdd(sat *starArrayType, bat *boardArrayType) {
	quarter := int(maxStars / 4)
	starPopulation := 3*quarter + rand.Intn(quarter)
	log.Printf("starPopulation:%d", starPopulation)
	for ndx := 0; ndx < starPopulation; ndx++ {
		position := randomCelestialLocation(*bat)
		if position == nil {
			log.Println("skipping nil position for star")
		} else {
			star := newStar(position)
			sat[ndx] = star
			setStar(bat[star.position.yy][star.position.xx], star.uuid)
		}
	}
}

// starDelete changes star to black hole
func starDelete(target string, sat *starArrayType, bat *boardArrayType) int {
	log.Printf("starDelete:%s", target)

	for ndx := 0; ndx < maxStars; ndx++ {
		if sat[ndx] != nil {
			if strings.Compare(sat[ndx].uuid, target) == 0 {
				// convert to black hole
				bc := bat[sat[ndx].position.yy][sat[ndx].position.xx]
				starToBlackHole(bc)
				// remove from star array
				sat[ndx] = nil
				return ndx
			}
		}
	}

	return -1
}

// starDump diagnostic
func starDump(sat starArrayType) {
	log.Println("=-=-=-= starDump =-=-=-=")

	for ndx := 0; ndx < maxStars; ndx++ {
		if sat[ndx] == nil {
			log.Printf("%d nil", ndx)
		} else {
			log.Printf("%d %s %d %d %s", ndx, sat[ndx].name, sat[ndx].position.yy, sat[ndx].position.xx, sat[ndx].uuid)
		}
	}

	log.Println("=-=-=-= starDump =-=-=-=")
}

// starFind returns array index for star by uuid
func starFind(target string, sat starArrayType) int {
	for ndx := 0; ndx < maxStars; ndx++ {
		if sat[ndx] != nil {
			if strings.Compare(sat[ndx].uuid, target) == 0 {
				return ndx
			}
		}
	}

	return -1
}
