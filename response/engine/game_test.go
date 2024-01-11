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

func TestGameStarting(t *testing.T) {
	g := NewGame(
		player.NewPlayer(1, "John"),
		player.NewPlayer(2, "Jane"),
		log.NewNullLogger(),
	)
	assert.False(t, g.GetStatus() == status.InProgress)

	c := &commands.StartGame{}
	c.SetTargets(g)
	g.AddCommandToQueue(c)
	g.ExecuteNextCommand()
	assert.True(t, g.GetStatus() == status.InProgress)
}

func TestStartingGameInProgress(t *testing.T) {
	g := NewGame(
		player.NewPlayer(1, "John"),
		player.NewPlayer(2, "Jane"),
		log.NewNullLogger(),
	)

	c := &commands.StartGame{}
	c.SetTargets(g)
	g.AddCommandToQueue(c)
	g.ExecuteNextCommand()
	g.AddCommandToQueue(c)
	assert.Panics(t, g.ExecuteNextCommand)
}
