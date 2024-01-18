package commands

import (
	"ticktack/response/engine/state"
)

type FinishGame struct{}

func (c *FinishGame) Execute(s *state.State) {
	if s.GameStatus != state.InProgress {
		panic("Game cannot be started")
	}
	s.GameStatus = state.Finished
}
