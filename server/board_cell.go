// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

/*
type boardTokenEnum int

const (
	vacantToken boardTokenEnum = iota
	mineToken
	planetToken
	shipToken
	starGateToken
	voidToken
)

func (bte boardTokenEnum) String() string {
	return [...]string{"vacant", "mine", "planet", "ship", "starGate", "void"}[bte]
}
*/

type boardCellType struct {
	// celestial objects without token uuid
	acheronVoid bool
	blackHole   bool

	// if not nil, look in catalog
	key *catalogKeyType
}

func newBoardCell() *boardCellType {
	result := boardCellType{}
	return &result
}

func (bc *boardCellType) setAcheronVoid() {
	bc.acheronVoid = true
}

func (bc *boardCellType) setBlackHole() {
	bc.blackHole = true
}
