package trigger

import (
	"ticktack/src/engine/event"
	"ticktack/src/engine/interfaces"
	"ticktack/src/engine/state"
)

type Trigger interface {
	EventName() event.EventName
	Execute(*state.State) error
}

type defaultTrigger struct {
	eventName event.EventName
	Target    interfaces.Target
	Action    func(*state.State) error
}

func (t *defaultTrigger) EventName() event.EventName {
	return t.eventName
}

func (t *defaultTrigger) Execute(s *state.State) error {
	return t.Action(s)
}
