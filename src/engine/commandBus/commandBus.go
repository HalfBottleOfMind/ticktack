package commandBus

import (
	"errors"
	"ticktack/src/engine/commandBus/commands"
	"ticktack/src/engine/state"
)

var ErrGameIsFinished = errors.New("game is finished")

type Bus struct {
	State *state.State
	Queue []commands.Command
}

func (b *Bus) Add(c commands.Command) error {
	if b.State.IsFinished() {
		return ErrGameIsFinished
	}
	b.Queue = append(b.Queue, c)
	return nil
}

func (b *Bus) ExecuteNext() {
	if b.Queue[0].Execute(b.State) == nil {
		b.Queue = b.Queue[1:]
	} else {
		b.Error()
	}
}

func (b *Bus) Error() {
	b.State.GameStatus = state.Error
}
