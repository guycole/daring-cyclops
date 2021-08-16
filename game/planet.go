package game

import (
	"github.com/google/uuid"
)

type planetType struct {
	active   bool
	position locationType
	team     PlayerTeam
	uuid     string
}

func newPlanet(location locationType) *planetType {
	result := planetType{active: true, position: location, team: NeutralTeam}
	result.uuid = uuid.NewString()
	return &result
}
