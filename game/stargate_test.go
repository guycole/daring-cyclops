package game

import (
	"testing"
)

func TestFreshStargate(t *testing.T) {
	got := newStarGate(4)

	if got.active != true {
		t.Errorf("newStarGate(4) active failure")
	}

	if got.damage != 0 {
		t.Errorf("newStarGate(4) damage failure")
	}

	reference := newLocation(35, 35)
	if got.position.x != reference.x || got.position.y != reference.y {
		t.Errorf("newStarGate(4) position failure")
	}
}
