package commands

import "ticktack/internal/status"

type gameStub struct {
	Status status.Status
}

func (gs *gameStub) SetStatus(status status.Status) { gs.Status = status }
func (gs *gameStub) GetStatus() status.Status       { return gs.Status }
