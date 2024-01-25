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

	m.AddToCommandQueue(&command.StartGame{})
	m.AddToCommandQueue(&command.FinishGame{})

	for len(m.CommandQueue) > 0 {
		if err := m.ProcessCommandQueue(); err != nil {
			panic(err)
		}

		for len(m.EventQueue) > 0 {
			if err := m.ProcessEventQueue(); err != nil {
				panic(err)
			}
		}
	}

	fmt.Println(m.State.Winner())
}
