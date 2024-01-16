package commands

import (
	"ticktack/internal/status"
	"ticktack/response/engine/interfaces"
)

type FinishGame struct {
	Target interfaces.HasStatus
}

func (c *FinishGame) Execute() {
	if c.Target.GetStatus() != status.InProgress {
		panic("Game is not in progress")
	}
	c.Target.SetStatus(status.Finished)
}

func (c *FinishGame) SetTarget(target interfaces.HasStatus) {
	c.Target = target
}
