package commandBus

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"ticktack/response/engine/player"
	"ticktack/response/engine/state"
)

type CommandMock struct {
	mock.Mock
}

func (m *CommandMock) Execute(state *state.State) {
	m.Called(state)
}

var p1 = player.NewPlayer(1, "John")
var p2 = player.NewPlayer(2, "Jane")
var State = state.NewState(&p1, &p2, state.NotStarted)

func TestBus_Add(t *testing.T) {
	b := Bus{State: &State}
	assert.Len(t, b.Queue, 0)
	b.Add(&CommandMock{})

	assert.Len(t, b.Queue, 1)
}

func TestBus_Add_GameFinished(t *testing.T) {
	State.GameStatus = state.Finished
	t.Cleanup(func() {
		State.GameStatus = state.NotStarted
	})
	b := Bus{State: &State}

	shouldPanic := func() {
		b.Add(&CommandMock{})
	}

	assert.Panics(t, shouldPanic)
}

func TestBus_ExecuteNext_ExecuteCalling(t *testing.T) {
	b := Bus{State: &State}
	c := CommandMock{}
	c.On("Execute")

	b.Add(&c)
	b.ExecuteNext()

	c.AssertNumberOfCalls(t, "Execute", 1)
}

func TestBus_ExecuteNext_CommandRemovedFromQueue(t *testing.T) {
	b := Bus{State: &State}
	c := CommandMock{}
	c.On("Execute")

	b.Add(&c)
	assert.Contains(t, b.Queue, &c)
	b.ExecuteNext()

	assert.NotContains(t, b.Queue, &c)
}
