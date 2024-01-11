package status

type Status byte

const (
	NotStarted Status = iota
	InProgress
	Finished
	Error
)
