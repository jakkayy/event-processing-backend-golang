package pipeline

import (
	"log"
)

// create worker
type Worker struct {
	ID        int
	Queue     *EventQueue
	Processor Processor
}

// start worker
func (w *Worker) Start() {
	go func() {
		log.Printf("worker %d start\n", w.ID)

		for event := range w.Queue.Ch {
			w.Processor.Process(event)
		}
	}()
}
