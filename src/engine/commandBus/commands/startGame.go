package commands

import (
	"errors"
	"ticktack/src/engine/state"
)

type StartGame struct {
	defaultCommand
}

func (c *StartGame) Execute(s *state.State) {
	if err := c.validate(s); err != nil {
		panic("Game cannot be started")
	}
	s.GameStatus = state.InProgress
}

func (c *StartGame) validate(s *state.State) error {
	if s.GameStatus != state.InProgress {
		return errors.New("game cannot be started")
	}
	return nil
}
