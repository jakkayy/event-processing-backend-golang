package pipeline

import (
	"log"

	"event-processing-backend-golang/internal/domain"
)

// create worker
type Worker struct {
	ID    int
	Queue *EventQueue
}

// start worker
func (w *Worker) Start() {
	go func() {
		log.Printf("worker %d start\n", w.ID)

		for event := range w.Queue.Ch {
			w.process(event)
		}
	}()
}

// process
func (w *Worker) process(e domain.Event) {
	log.Printf("worker %d processing event: type=%s message=%s\n", w.ID, e.Type, e.Message)
}
