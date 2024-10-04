// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

type boardCellType struct {
	acheronVoidFlag bool          // acheron void is impassable
	blackHoleFlag   bool          // black hole is impassable
	tokenKey        *tokenKeyType // occupied by this token
	tokenType       tokenEnum     // occupied by this token
}

func newBoardCell() *boardCellType {
	result := boardCellType{acheronVoidFlag: false, blackHoleFlag: false, tokenKey: nil, tokenType: emptyToken}
	return &result
}

func (bc *boardCellType) setOccupied(tokenKey *tokenKeyType, tokenType tokenEnum) {
	bc.tokenKey = tokenKey
	bc.tokenType = tokenType
}

func (bc *boardCellType) clearOccupied() {
	bc.tokenKey = nil
	bc.tokenType = emptyToken
}
