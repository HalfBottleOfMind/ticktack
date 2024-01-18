package commands

import (
	"ticktack/src/engine/state"
)

type Target uint

const (
	PlayerOne Target = iota
	PlayerTwo
)

type Command interface {
	Execute(s *state.State) error
}

type defaultCommand struct{}

func (c *defaultCommand) validate(s *state.State) error {
	if s.GameStatus != state.InProgress {
		return ErrGameNotInProgress
	}
	return nil
}
