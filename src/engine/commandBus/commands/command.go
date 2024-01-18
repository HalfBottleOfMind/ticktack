package commands

import (
	"errors"
	"ticktack/src/engine/state"
)

type Target uint

const (
	PlayerOne Target = iota
	PlayerTwo
)

type Command interface {
	Execute(s *state.State)
}

type defaultCommand struct{}

func (c *defaultCommand) validate(s *state.State) error {
	if s.GameStatus != state.InProgress {
		return errors.New("game not in progress")
	}
	return nil
}
