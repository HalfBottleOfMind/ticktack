package commands

type Target byte

const (
	Game Target = iota
	PlayerOne
	PlayerTwo
)

type Command interface {
	Execute()
}
