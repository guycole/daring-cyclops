package game

// PlayerRank comment
type PlayerRank int

const (
	// Cadet is beginner
	Cadet PlayerRank = iota
	// Lieutenant arg
	Lieutenant
	// Captain arg
	Captain
	// Admiral arg
	Admiral
)

func (pr PlayerRank) String() string {
	return [...]string{"Cadet", "Lieutenant", "Captain", "Admiral"}[pr]
}

// PlayerTeam comment
type PlayerTeam int

const (
	// Neutral no team
	NeutralTeam PlayerTeam = iota
	// Blue team
	BlueTeam
	// Red team
	RedTeam
)

func (pt PlayerTeam) String() string {
	return [...]string{"Neutral", "Blue", "Red"}[pt]
}

// Player comment
type Player struct {
	Active bool
	Name   string
	Rank   PlayerRank
	Team   PlayerTeam
	UUID   string
}

func getFreshPlayer(name, id string, rank PlayerRank, team PlayerTeam) Player {
	var result Player
	result.Active = true
	result.Name = name
	result.Rank = rank
	result.Team = team
	result.UUID = id
	return result
}
