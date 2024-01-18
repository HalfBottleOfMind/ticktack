package commands

import (
	"ticktack/src/engine/player"
	"ticktack/src/engine/state"
)

type PayInfluence struct {
	defaultCommand
	Target Target
	Cost   uint
}

func (c *PayInfluence) Execute(s *state.State) error {
	if err := c.validate(s); err != nil {
		return err
	}
	c.getTargetPlayer(s).Hero.CurrentInfluence -= c.Cost
	return nil
}

func (c *PayInfluence) getTargetPlayer(s *state.State) *player.Player {
	if c.Target == PlayerOne {
		return s.PlayerOne
	} else {
		return s.PlayerTwo
	}
}

func (c *PayInfluence) validate(s *state.State) error {
	if err := c.defaultCommand.validate(s); err != nil {
		return err
	} else if c.Target != PlayerOne && c.Target != PlayerTwo {
		return ErrInvalidTarget
	}
	return nil
}
