package engine

import (
	"github.com/google/uuid"
	"ticktack/internal/log"
	"ticktack/internal/status"
)

type Game struct {
	Id              uuid.UUID
	Tick            uint
	status          status.Status
	Logger          log.Logger
	onTickCallbacks []func()
}

func (g *Game) GetStatus() status.Status {
	return g.status
}

func (g *Game) SetStatus(status status.Status) {
	g.status = status
}

func (g *Game) Start() {
	if g.status != status.NotStarted {
		panic("Game cannot be started")
	}
	g.status = status.InProgress
	g.Tick = 0
	g.Logger.GameStarted()
}

func (g *Game) Finish() {
	if g.status != status.InProgress {
		panic("Game is not in progress")
	}
	g.status = status.Finished
	g.Logger.GameFinished()
}

func (g *Game) StopWithError(message string) {
	if g.status != status.InProgress {
		panic("Game is not in progress")
	}
	g.status = status.Error
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
	if g.status != status.InProgress {
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
		status: status.NotStarted,
		Logger: logger,
	}
}
