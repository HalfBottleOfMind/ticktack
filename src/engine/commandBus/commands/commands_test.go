package commands

import (
	"testing"
	"ticktack/src/engine/hero"
	"ticktack/src/engine/player"
	"ticktack/src/engine/state"

	"github.com/stretchr/testify/assert"
)

func TestStartGame_StartingGame(t *testing.T) {
	s := &state.State{GameStatus: state.NotStarted}
	c := StartGame{}

	err := c.Execute(s)

	assert.Nil(t, err)
	assert.Equal(t, state.InProgress, s.GameStatus)
}

func TestStartGame_StartingGameInProgress(t *testing.T) {
	s := &state.State{GameStatus: state.InProgress}
	c := StartGame{}

	err := c.Execute(s)

	assert.Error(t, err, ErrGameAlreadyInProgress)
}

func TestFinishGame_FinishingGame(t *testing.T) {
	s := &state.State{GameStatus: state.InProgress}
	c := FinishGame{}

	err := c.Execute(s)

	assert.Nil(t, err)
	assert.Equal(t, state.Finished, s.GameStatus)
}

func TestFinishGame_FinishingGameNotInProgress(t *testing.T) {
	s := &state.State{GameStatus: state.NotStarted}
	c := FinishGame{}

	err := c.Execute(s)

	assert.Error(t, err, ErrGameNotInProgress)
}

func TestSwitchInitiative(t *testing.T) {
	s := &state.State{
		PlayerOne:  &player.Player{Initiative: true},
		PlayerTwo:  &player.Player{Initiative: false},
		GameStatus: state.InProgress,
	}
	c := SwitchInitiative{}

	err := c.Execute(s)

	assert.Nil(t, err)
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

	err := c.Execute(s)

	assert.Error(t, err, ErrGameNotInProgress)
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

	err := c.Execute(s)

	assert.Nil(t, err)
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

	err := c.Execute(s)

	assert.Nil(t, err)
	assert.Equal(t, uint(3), s.PlayerTwo.Hero.CurrentInfluence)
}

func TestPayInfluence_Execute_GameNotInProgress(t *testing.T) {
	s := &state.State{
		GameStatus: state.NotStarted,
	}
	c := PayInfluence{}

	err := c.Execute(s)

	assert.Error(t, err, ErrGameNotInProgress)
}

func TestPayInfluence_Execute_InvalidTarget(t *testing.T) {
	t.Skip()
}
