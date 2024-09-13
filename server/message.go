// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

import (
	"strings"

	"github.com/google/uuid"
)

type messageKeyType struct {
	key string
}

// convenience factory
func newMessageKey(key string) *messageKeyType {
	var result messageKeyType

	temp := strings.TrimSpace(key)
	if len(temp) < 36 {
		result = messageKeyType{key: uuid.NewString()}
	} else {
		result = messageKeyType{key: temp}
	}

	return &result
}

// player for this game
type messageType struct {
	destination *playerKeyType
	key         *messageKeyType
	next        *messageType
	payload     string
	source      *playerKeyType
}

// convenience factory
func newMessage(destination *playerKeyType, source *playerKeyType, payload string) *messageType {
	key := newMessageKey("")
	result := messageType{destination: destination, key: key, payload: payload, source: source}
	return &result
}
