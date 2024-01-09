package engine

type Status uint8

const (
	NotStarted Status = iota
	InProgress
	Finished
	Error
)
