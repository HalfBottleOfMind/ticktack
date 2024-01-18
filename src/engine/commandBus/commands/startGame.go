package commands

import (
	"ticktack/src/engine/state"
)

type StartGame struct {
	defaultCommand
}

func (c *StartGame) Execute(s *state.State) error {
	if err := c.validate(s); err != nil {
		return err
	}
	s.GameStatus = state.InProgress
	return nil
}

func (c *StartGame) validate(s *state.State) error {
	if s.GameStatus != state.NotStarted {
		return ErrGameAlreadyInProgress
	}
	return nil
}
