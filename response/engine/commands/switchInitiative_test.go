package commands

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ticktack/internal/status"
)

func TestSwitchInitiative_POneToPTwo(t *testing.T) {
	g := &gameStub{
		Status:    status.InProgress,
		PlayerOne: &playerStub{true},
		PlayerTwo: &playerStub{false},
	}
	c := SwitchInitiative{g}

	c.Execute()

	assert.False(t, g.GetPlayerOne().GetInitiative())
	assert.True(t, g.GetPlayerTwo().GetInitiative())
}

//func TestSwitchInitiative_PTwoToPOne(t *testing.T) {
//	g := &gameStub{
//		Status:    status.InProgress,
//		PlayerOne: &playerStub{false},
//		PlayerTwo: &playerStub{true},
//	}
//	c := SwitchInitiative{g}
//
//	c.Execute()
//
//	assert.True(t, g.GetPlayerOne().GetInitiative())
//	assert.False(t, g.GetPlayerTwo().GetInitiative())
//}
//
//func TestSwitchInitiative_GameNotInProgress(t *testing.T) {
//	g := &gameStub{
//		Status:    status.InProgress,
//		PlayerOne: &playerStub{false},
//		PlayerTwo: &playerStub{true},
//	}
//	c := SwitchInitiative{g}
//
//	c.Execute()
//
//	assert.True(t, g.GetPlayerOne().GetInitiative())
//	assert.False(t, g.GetPlayerTwo().GetInitiative())
//}
