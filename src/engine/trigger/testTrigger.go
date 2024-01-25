package trigger

import (
	"ticktack/src/engine/event"
	"ticktack/src/engine/interfaces"
	"ticktack/src/engine/state"
)

type TestTrigger struct {
	defaultTrigger
}

func NewTestTrigger() *TestTrigger {
	return &TestTrigger{
		defaultTrigger{
			eventName: event.GameStarted,
			Action: func(s *state.State) error {
				return s.SetWinner(interfaces.PlayerTwo)
			},
		},
	}
}
