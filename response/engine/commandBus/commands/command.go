package commands

import "ticktack/response/engine/state"

type Command interface {
	Execute(s *state.State)
}
