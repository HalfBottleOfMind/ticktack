package commands

import (
	"ticktack/response/engine/state"
)

type SwitchInitiative struct{}

func (si *SwitchInitiative) Execute(s *state.State) {
	if s.GameStatus != state.InProgress {
		panic("Game not in progress")
	}
	s.PlayerOne.SetInitiative(!s.PlayerOne.GetInitiative())
	s.PlayerTwo.SetInitiative(!s.PlayerTwo.GetInitiative())
}
