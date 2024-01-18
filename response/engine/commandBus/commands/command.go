package commands

import "ticktack/response/engine/state"

type Target uint

const (
	PlayerOne Target = iota
	PlayerTwo
)

type Command interface {
	Execute(s *state.State)
}
