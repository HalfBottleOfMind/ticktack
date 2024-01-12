package commands

import (
	"ticktack/internal/status"
)

type StartGame struct {
	Target interface {
		SetStatus(status.Status)
		GetStatus() status.Status
	}
}

func (c *StartGame) Execute() {
	if c.Target.GetStatus() != status.NotStarted {
		panic("Game cannot be started")
	}
	c.Target.SetStatus(status.InProgress)
}

func (c *StartGame) SetTarget(target interface {
	SetStatus(status.Status)
	GetStatus() status.Status
}) {
	c.Target = target
}
