// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found
// in the LICENSE file.
package main

type gameMasterType struct {
	seqId     int
	shipPop   int
	blueScore int
	redScore  int
}

const maxGames = 5

type masterArrayType [maxGames]*gameMasterType
