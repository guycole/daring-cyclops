// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import (
	"log"
	"math/rand"
	"strings"

	"github.com/google/uuid"
)

type planetType struct {
	build    int
	name     string
	team     teamEnum
	position *locationType
	uuid     string
}

const maxPlanets = 255

// planetArrayType contains all active planets
type planetArrayType [maxPlanets]*planetType

// newPlanet convenience function to populate struct
func newPlanet(position *locationType) *planetType {
	result := planetType{position: position, team: neutralTeam}
	result.name = nameGenerator(false)
	result.uuid = uuid.NewString()
	return &result
}

// generate all planets
func planetsAdd(pat *planetArrayType, bat *boardArrayType) {
	quarter := int(maxPlanets / 4)
	planetPopulation := 3*quarter + rand.Intn(quarter)
	log.Printf("planetPopulation:%d", planetPopulation)
	/*
		for ndx := 0; ndx < planetPopulation; ndx++ {
			position := randomCelestialLocation(*bat)
			if position == nil {
				log.Println("skipping nil position for planet")
			} else {
				planet := newPlanet(position)
				pat[ndx] = planet
				setPlanet(bat[planet.position.yy][planet.position.xx], planet.uuid)
			}
		}
	*/
}

// return current planet census
func planetCensus(pat planetArrayType) (int, int, int) {
	neutralPopulation := 0
	bluePopulation := 0
	redPopulation := 0

	for ndx := 0; ndx < maxPlanets; ndx++ {
		if pat[ndx] != nil {
			switch pat[ndx].team {
			case neutralTeam:
				neutralPopulation++
			case blueTeam:
				bluePopulation++
			case redTeam:
				redPopulation++
			}
		}
	}

	return neutralPopulation, bluePopulation, redPopulation
}

// planetDelete removes planet from map
func planetDelete(target string, pat *planetArrayType, bat *boardArrayType) int {
	log.Printf("planetDelete:%s", target)

	/*
		for ndx := 0; ndx < maxPlanets; ndx++ {
			if pat[ndx] != nil {
				if strings.Compare(pat[ndx].uuid, target) == 0 {
					// remove planet from map
					bc := bat[pat[ndx].position.yy][pat[ndx].position.xx]
					clearPlanet(bc)
					// remove from planet array
					pat[ndx] = nil
					return ndx
				}
			}
		}
	*/

	return -1
}

// planetDump diagnostic
func planetDump(pat planetArrayType) {
	log.Println("=-=-=-= planetDump =-=-=-=")

	for ndx := 0; ndx < maxPlanets; ndx++ {
		if pat[ndx] == nil {
			log.Printf("%d nil", ndx)
		} else {
			team := pat[ndx].team.string()

			log.Printf("%d %s %s %d %d %s", ndx, pat[ndx].name, team, pat[ndx].position.yy, pat[ndx].position.xx, pat[ndx].uuid)
		}
	}

	log.Println("=-=-=-= planetDump =-=-=-=")
}

// planetFind returns array index for planet by uuid
func planetFind(target string, pat planetArrayType) int {
	for ndx := 0; ndx < maxPlanets; ndx++ {
		if pat[ndx] != nil {
			if strings.Compare(pat[ndx].uuid, target) == 0 {
				return ndx
			}
		}
	}

	return -1
}
