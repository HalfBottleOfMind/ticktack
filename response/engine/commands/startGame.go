package commands

import (
	"ticktack/internal/status"
	"ticktack/response/engine/interfaces"
)

type StartGame struct {
	Target interfaces.HasStatus
}

func (c *StartGame) Execute() {
	if c.Target.GetStatus() != status.NotStarted {
		panic("Game cannot be started")
	}
	c.Target.SetStatus(status.InProgress)
}

func (c *StartGame) SetTarget(target interfaces.HasStatus) {
	c.Target = target
}
