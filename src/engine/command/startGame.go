package command

import (
	"ticktack/src/engine/event"
	"ticktack/src/engine/state"
)

type StartGame struct {
	Command
}

func (c *StartGame) Execute(s *state.State) ([]event.Event, error) {
	var events []event.Event

	s.SetStatus(state.InProgress)
	events = append(events, event.NewEvent(event.GameStarted))

	return events, nil
}
