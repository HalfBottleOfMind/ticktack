package state

type GameStatus byte

const (
	NotStarted GameStatus = iota
	InProgress
	Finished
	Error
)
