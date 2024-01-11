package log

import (
	"github.com/google/uuid"
	"ticktack/response/player"
)

type NullLogger struct {
	GameId uuid.UUID
}

func NewNullLogger() *NullLogger {
	return &NullLogger{}
}

func (l *NullLogger) SetGameId(gameId uuid.UUID) {
	l.GameId = gameId
}

func (l *NullLogger) GameStarted() {}

func (l *NullLogger) NewTick(tick uint) {}

func (l *NullLogger) NewTurn(turn uint) {}

func (l *NullLogger) GameFinished() {}

func (l *NullLogger) Win(player player.Player) {}

func (l *NullLogger) Error(message string) {}

func (l *NullLogger) DamageDoneToPlayer(player player.Player) {}
