package engine

import (
	"github.com/google/uuid"
	"ticktack/internal/engine"
	"ticktack/response/engine/commands"
	"ticktack/response/log"
	"ticktack/response/player"
)

type Game struct {
	*engine.Game
	Turn         uint
	Logger       log.Logger
	PlayerOne    *player.Player
	PlayerTwo    *player.Player
	CommandQueue []commands.Command
}

func (g *Game) AddCommandToQueue(c commands.Command) {
	g.CommandQueue = append(g.CommandQueue, c)
}

func (g *Game) ExecuteNextCommand() {
	g.CommandQueue[0].Execute()
	g.CommandQueue = g.CommandQueue[1:]
}

//func (g *Game) Start() {
//	if g.GetStatus() != status.NotStarted {
//		panic("Game cannot be started")
//	}
//	g.SetStatus(status.InProgress)
//	g.Tick = 0
//	g.Turn = 0
//	g.Logger.GameStarted()
//}
//
//func (g *Game) NextTurn() {
//	if g.GetStatus() != status.InProgress {
//		panic("Game is not in progress")
//	}
//	g.Turn += 1
//	g.OnTurn()
//	g.NextTick()
//}
//
//func (g *Game) DealDamageToPlayer(player *player.Player) {
//	player.Health -= 1
//	g.Logger.DamageDoneToPlayer(*player)
//	g.NextTick()
//	g.CheckWinCondition()
//}
//
//func (g *Game) CheckWinCondition() {
//	if g.PlayerOne.Health < 1 {
//		g.Win(g.PlayerTwo)
//	} else if g.PlayerTwo.Health < 1 {
//		g.Win(g.PlayerOne)
//	}
//}
//
//func (g *Game) Win(player *player.Player) {
//	g.Finish()
//	g.Logger.Win(*player)
//}
//
//func (g *Game) OnTurn() {
//	g.Logger.NewTurn(g.Turn)
//}

func NewGame(playerOne, playerTwo *player.Player, logger log.Logger) *Game {
	id := uuid.New()
	logger.SetGameId(id)
	return &Game{
		Game:      engine.NewGame(logger),
		Logger:    logger,
		PlayerOne: playerOne,
		PlayerTwo: playerTwo,
	}
}
