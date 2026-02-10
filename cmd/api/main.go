package main

import (
	"log"
	"net/http"

	"event-processing-backend-golang/internal/api/handler"
	"event-processing-backend-golang/internal/pipeline"
)

func main() {
	queue := pipeline.NewEventQueue(100)

	processor := &pipeline.LogginProcessor{}

	workerCount := 3

	for i := 1; i <= workerCount; i++ {
		worker := &pipeline.Worker{
			ID:        i,
			Queue:     queue,
			Processor: processor,
		}
		worker.Start()
	}

	h := &handler.EventHandler{
		Queue: queue,
	}

	http.HandleFunc("/events", h.Handle)

	log.Println("event api running on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
