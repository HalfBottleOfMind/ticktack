package engine

import (
	"github.com/google/uuid"
	"ticktack/internal/log"
	"ticktack/internal/status"
)

type Game struct {
	Id              uuid.UUID
	Tick            uint
	Status          status.GameStatus
	Logger          log.Logger
	onTickCallbacks []func()
}

func (g *Game) GetStatus() status.GameStatus {
	return g.Status
}

func (g *Game) SetStatus(status status.GameStatus) {
	g.Status = status
}

func (g *Game) Start() {
	if g.Status != status.NotStarted {
		panic("Game cannot be started")
	}
	g.Status = status.InProgress
	g.Tick = 0
	g.Logger.GameStarted()
}

func (g *Game) Finish() {
	if g.Status != status.InProgress {
		panic("Game is not in progress")
	}
	g.Status = status.Finished
	g.Logger.GameFinished()
}

func (g *Game) StopWithError(message string) {
	if g.Status != status.InProgress {
		panic("Game is not in progress")
	}
	g.Status = status.Error
	g.Logger.Error(message)
}

func (g *Game) AddOnTickCallback(callback func()) {
	g.onTickCallbacks = append(g.onTickCallbacks, callback)
}

func (g *Game) OnTick() {
	g.Logger.NewTick(g.Tick)
	for _, callback := range g.onTickCallbacks {
		callback()
	}
}

func (g *Game) NextTick() {
	if g.Status != status.InProgress {
		panic("Game is not in progress")
	}
	g.Tick += 1
	g.OnTick()
}

func NewGame(logger log.Logger) *Game {
	id := uuid.New()
	logger.SetGameId(id)
	return &Game{
		Id:     id,
		Status: status.NotStarted,
		Logger: logger,
	}
}
