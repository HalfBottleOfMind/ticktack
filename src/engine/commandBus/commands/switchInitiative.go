package commands

import (
	"ticktack/src/engine/state"
)

type SwitchInitiative struct {
	defaultCommand
}

func (c *SwitchInitiative) Execute(s *state.State) error {
	if err := c.validate(s); err != nil {
		return err
	}
	s.PlayerOne.Initiative = !s.PlayerOne.Initiative
	s.PlayerTwo.Initiative = !s.PlayerTwo.Initiative
	return nil
}
