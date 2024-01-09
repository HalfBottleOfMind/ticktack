package logger

import (
	"ticktack/internal/logger"
	"ticktack/response/player"
)

type Logger interface {
	logger.Logger
	Win(player player.Player)
}
