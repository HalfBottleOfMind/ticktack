package command

import (
	"ticktack/src/engine/event"
	"ticktack/src/engine/state"
)

type Command interface {
	Execute(state *state.State) ([]event.Event, error)
}
