package game

import (
	"testing"
)

func TestFreshLocation(t *testing.T) {
	cases := []struct {
		y, x int
	}{
		{3, 3},
		{3, -3},
		{-3, 3},
		{-3, -3},
	}
	for _, ndx := range cases {
		result := newLocation(ndx.y, ndx.x)
		if result.x != ndx.x {
			t.Errorf("getFreshLocation(%d, %d) failure", ndx.y, ndx.x)
		}
		if result.y != ndx.y {
			t.Errorf("getFreshLocation(%d, %d) failure", ndx.y, ndx.x)
		}
	}
}

func TestDistance(t *testing.T) {
	cases := []struct {
		y1, x1, y2, x2, target int
	}{
		{0, 0, 3, 3, 4},
		{0, 0, 3, -3, 4},
		{0, 0, -3, -3, 4},
		{0, 0, -3, 3, 4},
	}
	for _, ndx := range cases {
		loc1 := newLocation(ndx.y1, ndx.x1)
		loc2 := newLocation(ndx.y2, ndx.x2)
		result := getDistance(loc1, loc2)
		if result != ndx.target {
			t.Errorf("getDistance(%d, %d) failure expect %d got %d", ndx.y1, ndx.x1, ndx.target, result)
		}
	}
}
