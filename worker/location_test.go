// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found
// in the LICENSE file.
package main

import (
	"testing"
)

func TestFreshLocation(t *testing.T) {
	tests := []struct {
		yy, xx int
	}{
		{3, 3},
		{3, -3},
		{-3, 3},
		{-3, -3},
	}
	for _, ndx := range tests {
		result := newLocation(ndx.yy, ndx.xx)
		if result.xx != ndx.xx {
			t.Errorf("getFreshLocation(%d, %d) failure", ndx.yy, ndx.xx)
		}
		if result.yy != ndx.yy {
			t.Errorf("getFreshLocation(%d, %d) failure", ndx.yy, ndx.xx)
		}
	}
}

func TestDistance(t *testing.T) {
	tests := []struct {
		y1, x1, y2, x2, target int
	}{
		{0, 0, 3, 3, 4},
		{0, 0, 3, -3, 4},
		{0, 0, -3, -3, 4},
		{0, 0, -3, 3, 4},
	}
	for _, ndx := range tests {
		loc1 := newLocation(ndx.y1, ndx.x1)
		loc2 := newLocation(ndx.y2, ndx.x2)
		result := getDistance(loc1, loc2)
		if result != ndx.target {
			t.Errorf("getDistance(%d, %d) failure expect %d got %d", ndx.y1, ndx.x1, ndx.target, result)
		}
	}
}
