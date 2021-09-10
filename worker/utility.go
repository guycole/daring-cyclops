// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import "math/rand"

var consonants = [...]string{
	"b",
	"c",
	"d",
	"f",
	"g",
	"h",
	"j",
	"k",
	"l",
	"m",
	"n",
	"p",
	"q",
	"r",
	"s",
	"t",
	"v",
	"w",
	"x",
	"y",
	"z",
}

var vowels = [...]string{
	"a",
	"e",
	"i",
	"o",
	"u",
}

func nameGenerator(vowelFirst bool) string {
	var flag bool
	var result string

	cLimit := len(consonants)
	vLimit := len(vowels)

	if vowelFirst {
		flag = false
	} else {
		flag = true
	}

	for ndx := 0; ndx < 7; ndx++ {
		if flag {
			result += consonants[rand.Intn(cLimit)]
		} else {
			result += vowels[rand.Intn(vLimit)]
		}

		flag = !flag
	}

	return result
}
