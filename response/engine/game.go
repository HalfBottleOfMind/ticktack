package engine

import (
	"github.com/google/uuid"
	"ticktack/response/engine/commandBus"
	"ticktack/response/engine/player"
	"ticktack/response/engine/state"
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
