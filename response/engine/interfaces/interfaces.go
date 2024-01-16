package interfaces

import (
	"ticktack/internal/status"
)

type HasStatus interface {
	GetStatus() status.GameStatus
	SetStatus(status.GameStatus)
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
