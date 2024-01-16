package engine

import (
	"github.com/google/uuid"
	"ticktack/internal/engine"
	"ticktack/response/engine/commands"
	"ticktack/response/engine/interfaces"
	"ticktack/response/log"
)

type Game struct {
	*engine.Game
	Turn         uint
	Logger       log.Logger
	PlayerOne    interfaces.Player
	PlayerTwo    interfaces.Player
	CommandQueue []commands.Command
}

func (g *Game) GetPlayerOne() interfaces.Player {
	return g.PlayerOne
}

func (g *Game) GetPlayerTwo() interfaces.Player {
	return g.PlayerTwo
}

func (g *Game) AddCommandToQueue(c commands.Command) {
	g.CommandQueue = append(g.CommandQueue, c)
}

func (g *Game) ExecuteNextCommand() {
	g.CommandQueue[0].Execute()
	g.CommandQueue = g.CommandQueue[1:]
}

func NewGame(playerOne, playerTwo interfaces.Player, logger log.Logger) *Game {
	id := uuid.New()
	logger.SetGameId(id)
	return &Game{
		Game:      engine.NewGame(logger),
		Logger:    logger,
		PlayerOne: playerOne,
		PlayerTwo: playerTwo,
	}
}
