package main

import (
	"fmt"
	"ticktack/src/engine/command"
	"ticktack/src/engine/manager"
	"ticktack/src/engine/state"
	"ticktack/src/engine/trigger"
)

func main() {
	m := manager.Manager{
		State: state.NewState(),
	}
	m.AddTrigger(trigger.NewTestTrigger())
	m.State.SetWinner(1)

	m.AddToCommandQueue(&command.StartGame{})
	m.ProcessCommandQueue()
	for len(m.EventQueue) != 0 {
		m.ProcessEventQueue()
	}

	fmt.Print(m.State.Winner())
}
