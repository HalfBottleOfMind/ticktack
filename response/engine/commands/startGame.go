package commands

import (
	"ticktack/internal/status"
)

type target interface {
	SetStatus(status.Status)
	GetStatus() status.Status
}

type StartGame struct {
	t target
}

func (c *StartGame) Execute() {
	if c.t.GetStatus() != status.NotStarted {
		panic("Game cannot be started")
	}
	c.t.SetStatus(status.InProgress)
}

func (c *StartGame) SetTargets(targets ...any) {
	if target, ok := targets[0].(target); ok {
		c.t = target
	} else {
		panic("Invalid target")
	}
}

func (c *StartGame) GetTargets() []any {
	return []any{c.t}
}
