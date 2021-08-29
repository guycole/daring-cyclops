package main

// playerRankEnum comment
type playerRankEnum int

const (
	// CadetRank is beginner
	CadetRank playerRankEnum = iota
	// LieutenantRank arg
	LieutenantRank
	// CaptainRank arg
	CaptainRank
	// AdmiralRank arg
	AdmiralRank
)

func (pre playerRankEnum) string() string {
	return [...]string{"Cadet", "Lieutenant", "Captain", "Admiral"}[pre]
}

// playerTeam comment
type playerTeamEnum int

const (
	// NeutralTeam no team
	NeutralTeam playerTeamEnum = iota
	// BlueTeam ry
	BlueTeam
	// RedTeam ry
	RedTeam
)

func (pte playerTeamEnum) string() string {
	return [...]string{"Neutral", "Blue", "Red"}[pte]
}

// PlayerType comment
type PlayerType struct {
	name string
	rank playerRankEnum
	team playerTeamEnum
	uuid string
}

// NewPlayer ryryr
func NewPlayer(name, id string, rank playerRankEnum, team playerTeamEnum) *PlayerType {
	result := PlayerType{name: name, rank: rank, team: team, uuid: id}
	return &result
}
