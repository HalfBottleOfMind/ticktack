package commands

import (
	"ticktack/src/engine/state"
)

type SwitchInitiative struct {
	defaultCommand
}

func (c *SwitchInitiative) Execute(s *state.State) {
	if err := c.validate(s); err != nil {
		panic("Game not in progress")
	}
	s.PlayerOne.Initiative = !s.PlayerOne.Initiative
	s.PlayerTwo.Initiative = !s.PlayerTwo.Initiative
}
