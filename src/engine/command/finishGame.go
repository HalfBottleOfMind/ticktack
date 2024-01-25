package command

import (
	"ticktack/src/engine/event"
	"ticktack/src/engine/state"
)

type FinishGame struct {
	Command
}

func (c *FinishGame) Execute(s *state.State) ([]event.Event, error) {
	var events []event.Event

	s.SetStatus(state.Finished)
	events = append(events, event.NewEvent(event.GameFinished))

	return events, nil
}
