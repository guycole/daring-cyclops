// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

func commandPing(tnt *turnNodeType) *CommandType {
	var commands commandArrayType
	commands[0] = "pong"

	ct := newCommand(tnt.name, tnt.request, 1, commands)

	return ct
}
