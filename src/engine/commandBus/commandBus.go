package commandBus

import (
	"ticktack/src/engine/commandBus/commands"
	"ticktack/src/engine/state"
)

type Bus struct {
	State *state.State
	Queue []commands.Command
}

func (b *Bus) Add(c commands.Command) {
	if b.State.IsFinished() {
		panic("Cannot add command to finished game.")
	}
	b.Queue = append(b.Queue, c)
}

func (b *Bus) ExecuteNext() {
	b.Queue[0].Execute(b.State)
	b.Queue = b.Queue[1:]
}
