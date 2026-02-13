package domain

import "time"

type Event struct {
	ID        string
	Type      string
	Message   string
	Timestamp time.Time
}
