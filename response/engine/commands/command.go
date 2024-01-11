package commands

type Command interface {
	Execute()
	SetTargets(targets ...any)
	GetTargets() []any
}
