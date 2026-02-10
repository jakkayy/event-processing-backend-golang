package main

import (
	"log"
	"net/http"

	"event-processing-backend-golang/internal/api/handler"
	"event-processing-backend-golang/internal/pipeline"
)

func main() {
	queue := pipeline.NewEventQueue(100)

	h := &handler.EventHandler{
		Queue: queue,
	}

	http.HandleFunc("/events", h.Handle)

	log.Println("event api running")
	http.ListenAndServe("port 8080", nil)
}
