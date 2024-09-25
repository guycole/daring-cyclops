// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"strings"

	"github.com/google/uuid"
)

// catalog object
type catalogTokenEnum int

const (
	mineToken catalogTokenEnum = iota
	planetToken
	shipToken
	starGateToken
)

func (cte catalogTokenEnum) String() string {
	return [...]string{"mine", "planet", "ship", "starGate"}[cte]
}

type catalogKeyType struct {
	key string
}

// convenience factory
func newCatalogKey(key string) *catalogKeyType {
	var result catalogKeyType

	temp := strings.TrimSpace(key)
	if len(temp) < 36 {
		result = catalogKeyType{key: uuid.NewString()}
	} else {
		result = catalogKeyType{key: temp}
	}

	return &result
}

type catalogType struct {
	cte      catalogTokenEnum
	key      *catalogKeyType
	location *locationType
}
