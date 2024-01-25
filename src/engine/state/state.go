package state

import (
	"errors"
	"ticktack/src/engine/interfaces"
)

type GameStatus byte

const (
	NotStarted GameStatus = iota
	InProgress
	Finished
)

type State struct {
	winner interfaces.Target
	status GameStatus
}

func (s *State) Winner() (interfaces.Target, error) {
	if s.Status() != Finished {
		return 0, errors.New("game is not finished")
	}
	return s.winner, nil
}

func (s *State) Status() GameStatus {
	return s.status
}

func (s *State) SetStatus(status GameStatus) {
	s.status = status
}

func (s *State) SetWinner(w interfaces.Target) error {
	if w != interfaces.PlayerOne && w != interfaces.PlayerTwo {
		return errors.New("invalid target")
	}

	s.winner = w

	return nil
}

func NewState() *State {
	return &State{
		status: NotStarted,
	}
}
