package commands

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ticktack/internal/status"
)

func TestStartGame_StartingGame(t *testing.T) {
	g := &gameStub{Status: status.NotStarted}
	c := StartGame{g}

	c.Execute()

	assert.Equal(t, status.InProgress, g.Status)
}

func TestStartGame_StartingGameInProgress(t *testing.T) {
	g := &gameStub{Status: status.InProgress}
	c := StartGame{g}

	assert.Panics(t, c.Execute)
}

func TestStartGame_SettingTarget(t *testing.T) {
	g := &gameStub{Status: status.NotStarted}
	c := StartGame{}
	c.SetTarget(g)

	assert.Same(t, g, c.Target)
}
