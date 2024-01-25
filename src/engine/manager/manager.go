package manager

import (
	"fmt"
	"reflect"
	"ticktack/src/engine/command"
	"ticktack/src/engine/event"
	"ticktack/src/engine/log"
	"ticktack/src/engine/state"
	"ticktack/src/engine/trigger"
)

type Manager struct {
	State        *state.State
	CommandQueue []command.Command
	EventQueue   []event.Event
	Triggers     []trigger.Trigger
}

func (m *Manager) AddToCommandQueue(c command.Command) {
	m.CommandQueue = append(m.CommandQueue, c)
	log.GetLogger().Message(fmt.Sprintf("Message: command added to queue\nCommand name: %s\nCurrent queue len: %d\n", reflect.TypeOf(c), len(m.CommandQueue)))
}

func (m *Manager) ProcessCommandQueue() error {
	events, err := m.CommandQueue[0].Execute(m.State)
	if err != nil {
		return err
	}

	m.CommandQueue = m.CommandQueue[1:]
	m.EventQueue = append(m.EventQueue, events...)

	return nil
}

func (m *Manager) ProcessEventQueue() error {
	e := m.EventQueue[0]
	var triggers []trigger.Trigger
	for _, t := range m.Triggers {
		if e.EventName() == t.EventName() {
			triggers = append(triggers, t)
		}
	}

	for _, t := range triggers {
		err := t.Execute(m.State)
		if err != nil {
			return err
		}
	}

	m.EventQueue = m.EventQueue[1:]

	return nil
}

func (m *Manager) AddTrigger(t trigger.Trigger) {
	m.Triggers = append(m.Triggers, t)
}
