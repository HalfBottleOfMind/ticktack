package engine

import (
	"github.com/google/uuid"
	"ticktack/internal/logger"
)

type Game struct {
	Id   uuid.UUID
	Tick uint
	Status
	Logger logger.Logger
}

func (g *Game) Start() {
	if g.Status != NotStarted {
		panic("Game cannot be started")
	}
	g.Status = InProgress
	g.Tick = 0
	g.Logger.GameStarted()
}

func (g *Game) Finish() {
	if g.Status != InProgress {
		panic("Game is not in progress")
	}
	g.Status = Finished
	g.Logger.GameFinished()
}

func (g *Game) StopWithError(message string) {
	if g.Status != InProgress {
		panic("Game is not in progress")
	}
	g.Status = Error
	g.Logger.Error(message)
}

func (g *Game) onTick() {
	g.Logger.NewTick(g.Tick)
}

func (g *Game) NextTick() {
	if g.Status != InProgress {
		panic("Game not in progress")
	}
	g.Tick += 1
	g.onTick()
}

func NewGame(logger logger.Logger) *Game {
	id := uuid.New()
	logger.SetGameId(id)
	return &Game{
		Id:     id,
		Status: NotStarted,
		Logger: logger,
	}
}
