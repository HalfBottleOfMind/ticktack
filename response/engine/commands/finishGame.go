package commands

import "ticktack/internal/status"

type FinishGame struct {
	Target interface {
		SetStatus(status.Status)
		GetStatus() status.Status
	}
}

func (c *FinishGame) Execute() {
	if c.Target.GetStatus() != status.InProgress {
		panic("Game is not in progress")
	}
	c.Target.SetStatus(status.Finished)
}

func (c *FinishGame) SetTarget(target interface {
	SetStatus(status.Status)
	GetStatus() status.Status
}) {
	c.Target = target
}
