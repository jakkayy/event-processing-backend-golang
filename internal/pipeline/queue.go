package pipeline

import "event-processing-backend-golang/internal/domain"

// กล่องเก็บ event
type EventQueue struct {
	Ch chan domain.Event
}

func NewEventQueue(buffer int) *EventQueue {
	return &EventQueue{
		Ch: make(chan domain.Event, buffer),
	}
}
