package engine

import (
	"github.com/google/uuid"
	"ticktack/internal/engine"
	"ticktack/response/logger"
	"ticktack/response/player"
)

type Game struct {
	*engine.Game
	Logger    logger.Logger
	PlayerOne player.Player
	PlayerTwo player.Player
}

func (g *Game) Win(player player.Player) {
	g.Finish()
	g.Logger.Win(player)
}

func NewGame(playerOne, playerTwo player.Player, logger logger.Logger) *Game {
	id := uuid.New()
	logger.SetGameId(id)
	return &Game{
		Game:      engine.NewGame(logger),
		Logger:    logger,
		PlayerOne: playerOne,
		PlayerTwo: playerTwo,
	}
}
