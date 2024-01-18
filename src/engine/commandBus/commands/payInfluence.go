package commands

import (
	"fmt"
	"ticktack/src/engine/player"
	"ticktack/src/engine/state"
)

type PayInfluence struct {
	Target Target
	Cost   uint
}

func (c *PayInfluence) Execute(s *state.State) {
	if s.GameStatus != state.InProgress {
		panic("Game not in progress")
	}
	c.getTargetPlayer(s).Hero.CurrentInfluence -= c.Cost
}

func (c *PayInfluence) getTargetPlayer(s *state.State) *player.Player {
	if c.Target == PlayerOne {
		return s.PlayerOne
	} else if c.Target == PlayerTwo {
		return s.PlayerTwo
	} else {
		panic(fmt.Sprintf("Invalid target: %d", c.Target))
	}
}
