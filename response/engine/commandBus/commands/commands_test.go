package commands

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ticktack/response/cards/hero"
	"ticktack/response/engine/player"
	"ticktack/response/engine/state"
)

func TestStartGame_StartingGame(t *testing.T) {
	s := &state.State{GameStatus: state.NotStarted}
	c := StartGame{}

	c.Execute(s)

	assert.Equal(t, state.InProgress, s.GameStatus)
}

func TestStartGame_StartingGameInProgress(t *testing.T) {
	s := &state.State{GameStatus: state.InProgress}
	c := StartGame{}

	shouldPanic := func() {
		c.Execute(s)
	}

	assert.Panics(t, shouldPanic)
}

func TestFinishGame_FinishingGame(t *testing.T) {
	s := &state.State{GameStatus: state.InProgress}
	c := FinishGame{}

	c.Execute(s)

	assert.Equal(t, state.Finished, s.GameStatus)
}

func TestFinishGame_FinishingGameNotInProgress(t *testing.T) {
	s := &state.State{GameStatus: state.NotStarted}
	c := FinishGame{}

	shouldPanic := func() {
		c.Execute(s)
	}

	assert.Panics(t, shouldPanic)
}

func TestSwitchInitiative(t *testing.T) {
	s := &state.State{
		PlayerOne:  &player.Player{Initiative: true},
		PlayerTwo:  &player.Player{Initiative: false},
		GameStatus: state.InProgress,
	}
	c := SwitchInitiative{}

	c.Execute(s)

	assert.False(t, s.PlayerOne.Initiative)
	assert.True(t, s.PlayerTwo.Initiative)
}

func TestSwitchInitiative_GameNotInProgress(t *testing.T) {
	s := &state.State{
		PlayerOne:  &player.Player{Initiative: true},
		PlayerTwo:  &player.Player{Initiative: false},
		GameStatus: state.Finished,
	}
	c := SwitchInitiative{}

	shouldPanic := func() {
		c.Execute(s)
	}

	assert.Panics(t, shouldPanic)
}

func TestPayInfluence_Execute_PlayerOne(t *testing.T) {
	h := hero.NewHero(5)
	s := &state.State{
		PlayerOne:  &player.Player{Hero: &h},
		GameStatus: state.InProgress,
	}
	c := PayInfluence{
		Target: PlayerOne,
		Cost:   2,
	}

	c.Execute(s)

	assert.Equal(t, uint(3), s.PlayerOne.Hero.CurrentInfluence)
}

func TestPayInfluence_Execute_PlayerTow(t *testing.T) {
	h := hero.NewHero(5)
	s := &state.State{
		PlayerTwo:  &player.Player{Hero: &h},
		GameStatus: state.InProgress,
	}
	c := PayInfluence{
		Target: PlayerTwo,
		Cost:   2,
	}

	c.Execute(s)

	assert.Equal(t, uint(3), s.PlayerTwo.Hero.CurrentInfluence)
}

func TestPayInfluence_Execute_GameNotInProgress(t *testing.T) {
	s := &state.State{
		GameStatus: state.NotStarted,
	}
	c := PayInfluence{}

	shouldPanic := func() {
		c.Execute(s)
	}

	assert.Panics(t, shouldPanic)
}

func TestPayInfluence_Execute_InvalidTarget(t *testing.T) {
	t.Skip()
}
