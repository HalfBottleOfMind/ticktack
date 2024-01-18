package commands

import (
	"ticktack/response/engine/state"
)

type StartGame struct{}

func (c *StartGame) Execute(s *state.State) {
	if s.GameStatus != state.NotStarted {
		panic("Game cannot be started")
	}
	s.GameStatus = state.InProgress
}
