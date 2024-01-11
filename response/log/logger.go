package log

import (
	"ticktack/internal/log"
	"ticktack/response/player"
)

type Logger interface {
	log.Logger
	NewTurn(turn uint)
	Win(player player.Player)
	DamageDoneToPlayer(player player.Player)
}
