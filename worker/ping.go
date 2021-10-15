// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

func pingReqRes(tnt *turnNodeType) (*ResponseType, error) {
	argSize, argArray := okArgument()

	nr := newResponse(pingResponse, tnt.request, argSize, argArray)

	return nr, nil
}
