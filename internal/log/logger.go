package log

import (
	"github.com/google/uuid"
)

type Logger interface {
	SetGameId(gameId uuid.UUID)
	GameStarted()
	NewTick(tick uint)
	GameFinished()
	Error(message string)
}
