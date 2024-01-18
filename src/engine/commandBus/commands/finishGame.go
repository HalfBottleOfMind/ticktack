package commands

import (
	"ticktack/src/engine/state"
)

type FinishGame struct {
	defaultCommand
}

func (c *FinishGame) Execute(s *state.State) error {
	if err := c.validate(s); err != nil {
		return err
	}
	s.GameStatus = state.Finished
	return nil
}
