package manager

type PlayerRank int

const (
	Cadet PlayerRank = iota
	Lieutenant
	Captain
	Admiral
)

func (pr PlayerRank) String() string {
	return [...]string{"Cadet", "Lieutenant", "Captain", "Admiral"}[pr]
}

type PlayerTeam int

const (
	Neutral PlayerTeam = iota
	Blue
	Red
)

func (pt PlayerTeam) String() string {
	return [...]string{"Neutral", "Blue", "Red"}[pt]
}

type Player struct {
	Active   bool
	RandomId string
	Name     string
	Rank     PlayerRank
	Team     PlayerTeam
}
