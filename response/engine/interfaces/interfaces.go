package interfaces

import (
	"ticktack/response/engine/state"
)

type HasStatus interface {
	GetStatus() state.GameStatus
	SetStatus(state.GameStatus)
}

type Player interface {
	HasInitiative
}

type HasInitiative interface {
	GetInitiative() bool
	SetInitiative(bool)
}

type HasPlayers interface {
	GetPlayerOne() Player
	GetPlayerTwo() Player
}
