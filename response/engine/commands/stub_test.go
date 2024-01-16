package commands

import (
	"ticktack/internal/status"
	"ticktack/response/engine/interfaces"
)

type gameStub struct {
	Status    status.GameStatus
	PlayerOne interfaces.Player
	PlayerTwo interfaces.Player
}

func (gs *gameStub) SetStatus(status status.GameStatus) { gs.Status = status }
func (gs *gameStub) GetStatus() status.GameStatus       { return gs.Status }
func (gs *gameStub) GetPlayerOne() interfaces.Player    { return gs.PlayerOne }
func (gs *gameStub) GetPlayerTwo() interfaces.Player    { return gs.PlayerTwo }

type playerStub struct {
	Initiative bool
}

func (p *playerStub) GetInitiative() bool  { return p.Initiative }
func (p *playerStub) SetInitiative(i bool) { p.Initiative = i }
