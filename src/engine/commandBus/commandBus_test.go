package commandBus

import (
	"testing"
	"ticktack/src/engine/player"
	"ticktack/src/engine/state"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CommandMock struct {
	mock.Mock
}

func (m *CommandMock) Execute(state *state.State) error {
	m.Called(state)
	return nil
}

var p1 = player.NewPlayer(1, "John")
var p2 = player.NewPlayer(2, "Jane")
var State = state.NewState(&p1, &p2, state.NotStarted)

func TestBus_Add(t *testing.T) {
	b := Bus{State: &State}
	assert.Len(t, b.Queue, 0)
	err := b.Add(&CommandMock{})

	assert.Nil(t, err)
	assert.Len(t, b.Queue, 1)
}

func TestBus_Add_GameFinished(t *testing.T) {
	State.GameStatus = state.Finished
	t.Cleanup(func() {
		State.GameStatus = state.NotStarted
	})
	b := Bus{State: &State}

	err := b.Add(&CommandMock{})

	assert.Error(t, err, ErrGameIsFinished)
}

func TestBus_ExecuteNext_ExecuteCalling(t *testing.T) {
	b := Bus{State: &State}
	c := CommandMock{}
	c.On("Execute", &State).Return(nil)

	err := b.Add(&c)
	assert.Nil(t, err)
	b.ExecuteNext()

	c.AssertExpectations(t)
}

func TestBus_ExecuteNext_CommandRemovedFromQueue(t *testing.T) {
	b := Bus{State: &State}
	c := CommandMock{}
	c.On("Execute", &State).Return(nil)

	err := b.Add(&c)
	assert.Nil(t, err)
	assert.Contains(t, b.Queue, &c)
	b.ExecuteNext()

	assert.NotContains(t, b.Queue, &c)
}
