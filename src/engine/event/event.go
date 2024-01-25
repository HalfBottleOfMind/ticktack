package event

type EventName byte

const (
	GameStarted EventName = iota
	GameFinished
)

type Event struct {
	eventName EventName
}

func (e *Event) EventName() EventName {
	return e.eventName
}

func NewEvent(eventName EventName) Event {
	return Event{
		eventName: eventName,
	}
}
