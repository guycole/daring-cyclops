// Copyright 2024 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package server

type stubRequestType struct{}

// convenience factory
func newStub0Request(playerKey *playerKeyType) *commandType {
	ct := commandType{command: stubCommand0, sourcePlayerKey: playerKey}
	ct.stubRequest = &stubRequestType{}
	return &ct
}

// convenience factory
func newStub1Request(playerKey *playerKeyType) *commandType {
	ct := commandType{command: stubCommand1, sourcePlayerKey: playerKey}
	ct.stubRequest = &stubRequestType{}
	return &ct
}

// convenience factory
func newStub2Request(playerKey *playerKeyType) *commandType {
	ct := commandType{command: stubCommand2, sourcePlayerKey: playerKey}
	ct.stubRequest = &stubRequestType{}
	return &ct
}

// convenience factory
func newStub3Request(playerKey *playerKeyType) *commandType {
	ct := commandType{command: stubCommand3, sourcePlayerKey: playerKey}
	ct.stubRequest = &stubRequestType{}
	return &ct
}

type stubResponseType struct{}

func (gt *gameType) stubCommand(ct *commandType) *commandType {
	gt.sugarLog.Debug("stubCommand")

	response := stubResponseType{}

	ct.stubResponse = &response

	ct.destinationPlayerKeys = append(ct.destinationPlayerKeys, ct.sourcePlayerKey)

	return ct
}
