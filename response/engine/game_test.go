package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ticktack/internal/status"
	"ticktack/response/engine/commands"
	"ticktack/response/log"
	"ticktack/response/player"
)

func TestNewGame(t *testing.T) {
	p1 := player.NewPlayer(1, "John")
	p2 := player.NewPlayer(2, "Jane")
	logger := log.NewNullLogger()

	g := NewGame(p1, p2, logger)

	assert.Same(t, p1, g.PlayerOne)
	assert.Same(t, p2, g.PlayerTwo)
	assert.Same(t, logger, g.Logger)
}

func TestCommandsQueue(t *testing.T) {
	g := NewGame(
		player.NewPlayer(1, "John"),
		player.NewPlayer(2, "Jane"),
		log.NewNullLogger(),
	)
	g.AddCommandToQueue(&commands.StartGame{g})
	g.AddCommandToQueue(&commands.FinishGame{g})

	assert.Equal(t, status.NotStarted, g.GetStatus())
	g.ExecuteNextCommand()
	assert.Equal(t, status.InProgress, g.GetStatus())
	g.ExecuteNextCommand()
	assert.Equal(t, status.Finished, g.GetStatus())
}

func TestGameFinishing(t *testing.T) {
	g := NewGame(
		player.NewPlayer(1, "John"),
		player.NewPlayer(2, "Jane"),
		log.NewNullLogger(),
	)

	c1 := &commands.StartGame{g}
	g.AddCommandToQueue(c1)
	g.ExecuteNextCommand()

	c := &commands.FinishGame{g}
	g.AddCommandToQueue(c)
	g.ExecuteNextCommand()
	assert.True(t, g.GetStatus() == status.Finished)
}

func TestFinishingFinishedGame(t *testing.T) {
	g := NewGame(
		player.NewPlayer(1, "John"),
		player.NewPlayer(2, "Jane"),
		log.NewNullLogger(),
	)

	c1 := &commands.StartGame{g}
	g.AddCommandToQueue(c1)
	g.ExecuteNextCommand()

	c2 := &commands.FinishGame{g}
	g.AddCommandToQueue(c2)
	g.AddCommandToQueue(c2)
	g.ExecuteNextCommand()
	assert.Panics(t, g.ExecuteNextCommand)
}
