package commands

import (
	"ticktack/src/engine/state"
)

type FinishGame struct {
	defaultCommand
}

func (c *FinishGame) Execute(s *state.State) {
	if err := c.validate(s); err != nil {
		panic("Game cannot be finished")
	}
	s.GameStatus = state.Finished
}
