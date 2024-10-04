// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"strings"

	"github.com/google/uuid"
)

// catalog object
type tokenEnum int

const (
	emptyToken tokenEnum = iota
	mineToken
	planetToken
	shipToken
	starGateToken
)

func (te tokenEnum) String() string {
	return [...]string{"empty", "mine", "planet", "ship", "starGate"}[te]
}

type tokenKeyType struct {
	key string
}

// convenience factory
func newTokenKey(key string) *tokenKeyType {
	var result tokenKeyType

	temp := strings.TrimSpace(key)
	if len(temp) < 36 {
		result = tokenKeyType{key: uuid.NewString()}
	} else {
		result = tokenKeyType{key: temp}
	}

	return &result
}
