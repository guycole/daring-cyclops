// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import "strings"

type requestEnum int

// must match order for legalRequestType
const (
	moveAbsoluteRequest requestEnum = iota
	moveComputedRequest
	moveRelativeRequest
	pingRequest
	playerCreateRequest
	playerDeleteRequest
	shipCreateRequest
	shipDeleteRequest
	shutDownRequest
	tellAllRequest
	tellBlueRequest
	tellRedRequest
	unknownRequest
)

// must matchorder for requestEnum
var legalRequests = [...]string{
	"moveAbsoluteRequest",
	"moveComputedRequest",
	"moveRelativeRequest",
	"pingRequest",
	"playerCreateRequest",
	"playerDeleteRequest",
	"shipCreateRequest",
	"shipDeleteRequest",
	"shutDownRequest",
	"tellAllRequest",
	"tellBlueRequest",
	"tellRedRequest",
	"unknownRequest",
}

func findRequest(arg string) requestEnum {
	for ndx := 0; ndx < len(legalRequests); ndx++ {
		if strings.Compare(legalRequests[ndx], arg) == 0 {
			return requestEnum(ndx)
		}
	}

	return requestEnum(unknownRequest)
}

type responseEnum int

// must match order for legalRequestDuration
const (
	moveResponse responseEnum = iota
	pingResponse
	playerCreateResponse
	playerDeleteResponse
	shipCreateResponse
	shipDeleteResponse
	shutDownResponse
	tellResponse
	unknownResponse
)
