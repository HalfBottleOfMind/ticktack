package logger

import (
	"github.com/google/uuid"
	"ticktack/internal/player"
)

type Logger interface {
	SetGameId(gameId uuid.UUID)
	GameStarted()
	NewTick(tick uint)
	GameFinished()
	Win(player player.Player)
	Error(message string)
}
