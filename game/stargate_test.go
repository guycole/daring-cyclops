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

func TestStargateAdjacent(t *testing.T) {
	sg := newStarGate(4)

	var got = starGateAdjacent(newLocation(36, 34), sg.position)
	if got != 0 {
		t.Errorf("starGateAdjacent failure top left")
	}

	got = starGateAdjacent(newLocation(36, 35), sg.position)
	if got != 1 {
		t.Errorf("starGateAdjacent failure top")
	}

	got = starGateAdjacent(newLocation(36, 36), sg.position)
	if got != 2 {
		t.Errorf("starGateAdjacent failure top right")
	}

	got = starGateAdjacent(newLocation(35, 34), sg.position)
	if got != 3 {
		t.Errorf("starGateAdjacent failure left")
	}

	got = starGateAdjacent(newLocation(35, 35), sg.position)
	if got != 4 {
		t.Errorf("starGateAdjacent failure middle")
	}

	got = starGateAdjacent(newLocation(35, 36), sg.position)
	if got != 5 {
		t.Errorf("starGateAdjacent failure right")
	}

	got = starGateAdjacent(newLocation(34, 34), sg.position)
	if got != 6 {
		t.Errorf("starGateAdjacent failure bottom left")
	}

	got = starGateAdjacent(newLocation(34, 35), sg.position)
	if got != 7 {
		t.Errorf("starGateAdjacent bottom middle")
	}

	got = starGateAdjacent(newLocation(34, 36), sg.position)
	if got != 8 {
		t.Errorf("starGateAdjacent bottom right")
	}

	got = starGateAdjacent(newLocation(3, 3), sg.position)
	if got != -1 {
		t.Errorf("starGateAdjacent stupid")
	}
}
