package commands

import (
	"ticktack/response/engine/state"
)

type SwitchInitiative struct{}

func (si *SwitchInitiative) Execute(s *state.State) {
	if s.GameStatus != state.InProgress {
		panic("Game not in progress")
	}
	s.PlayerOne.Initiative = !s.PlayerOne.Initiative
	s.PlayerTwo.Initiative = !s.PlayerTwo.Initiative
}
