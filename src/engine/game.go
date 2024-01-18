package engine

import (
	"ticktack/src/engine/commandBus"
	"ticktack/src/engine/player"
	"ticktack/src/engine/state"

	"github.com/google/uuid"
)

type Game struct {
	Id    uuid.UUID
	State *state.State
	Bus   commandBus.Bus
}

func NewGame(playerOne, playerTwo *player.Player) Game {
	newState := state.NewState(playerOne, playerTwo, state.NotStarted)
	return Game{
		Id:    uuid.New(),
		State: &newState,
		Bus:   commandBus.Bus{State: &newState},
	}
}
