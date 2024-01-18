package commands

import "ticktack/src/engine/state"

type Target uint

const (
	PlayerOne Target = iota
	PlayerTwo
)

type Command interface {
	Execute(s *state.State)
}
