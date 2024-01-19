package commands

import (
	"ticktack/src/engine/card"
	"ticktack/src/engine/state"
)

type PlayCard struct {
	defaultCommand
	Card   card.Card
	Player Target // Player that plays card
	Target Target // Target of played card
}

func (c *PlayCard) Execute(s *state.State) error {
	if err := c.validate(s); err != nil {
		return err
	}

	p := PayInfluence{
		Target: c.Player,
		Cost:   c.Card.Cost,
	}
	if err := p.Execute(s); err != nil {
		return err
	}

	return nil
}

func (c *PlayCard) validate(s *state.State) error {
	if err := c.defaultCommand.validate(s); err != nil {
		return err
	} else if c.Player != PlayerOne && c.Player != PlayerTwo {
		return ErrInvalidTarget
	}
	return nil
}
