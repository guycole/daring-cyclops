// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

type boardCellType struct {
	acheronVoidFlag bool // Acheron void is a special type of cell that is impassable
	blackHoleFlag   bool // Black hole is a special type of cell that is impassable
	occupiedFlag    bool // True, there is a game token in the cell
}

func newBoardCell() *boardCellType {
	result := boardCellType{acheronVoidFlag: false, blackHoleFlag: false, occupiedFlag: false}
	return &result
}

func (bc *boardCellType) isOccupied() bool {
	return bc.occupiedFlag
}

func (bc *boardCellType) setOccupiedFlag() {
	bc.occupiedFlag = true
}

func (bc *boardCellType) clearOccupiedFlag() {
	bc.occupiedFlag = false
}
