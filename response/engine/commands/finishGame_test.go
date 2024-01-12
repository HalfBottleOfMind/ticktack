package commands

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ticktack/internal/status"
)

func TestFinishGame_ExecuteGame_FinishingGame(t *testing.T) {
	g := &gameStub{status.InProgress}
	c := FinishGame{g}

	c.Execute()

	assert.Equal(t, status.Finished, g.Status)
}

func TestFinishGame_ExecuteGame_FinishingGameNotInProgress(t *testing.T) {
	g := &gameStub{status.NotStarted}
	c := FinishGame{g}

	assert.Panics(t, c.Execute)
}

func TestFinishGame_ExecuteGame_SettingTarget(t *testing.T) {
	g := &gameStub{status.InProgress}
	c := FinishGame{}
	c.SetTarget(g)

	assert.Same(t, g, c.Target)
}
