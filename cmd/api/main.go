package main

import (
	"log"
	"net/http"

	"event-processing-backend-golang/internal/api/handler"
	"event-processing-backend-golang/internal/pipeline"
)

func main() {
	queue := pipeline.NewEventQueue(100)

	worker := &pipeline.Worker{
		ID:    1,
		Queue: queue,
	}
	worker.Start()

	h := &handler.EventHandler{
		Queue: queue,
	}

	http.HandleFunc("/events", h.Handle)

	log.Println("event api running")
	http.ListenAndServe("port 8080", nil)
}
