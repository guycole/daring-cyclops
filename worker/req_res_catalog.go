// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

type requestEnum int

// must match order for legalRequestDuration
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

// must match requestEnum
var legalRequestDuration = [...]int{
	3, // moveAbsoluteRequest
	3, // moveComputeRequest
	3, // moveRelativeRequest
	1, // pingRequest
	0, // playerCreateRequest
	0, // playerDeleteRequest
	0, // shipCreateRequest
	0, // shipDeleteRequest
	0, // shutDownRequest
	1, // tellAllRequest
	1, // tellBlueRequest
	1, // tellRedRequest
	1, // unknownRequest
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
