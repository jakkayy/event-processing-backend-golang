package main

import (
	"log"
	"net/http"
	"time"

	"event-processing-backend-golang/internal/api/handler"
	"event-processing-backend-golang/internal/pipeline"
)

func main() {
	queue := pipeline.NewEventQueue(100)

	windowAgg := pipeline.NewWindowAggregationProcessor(time.Minute, 5)

	workerCount := 3

	for i := 1; i <= workerCount; i++ {
		worker := &pipeline.Worker{
			ID:        i,
			Queue:     queue,
			Processor: windowAgg,
		}
		worker.Start()
	}

	eventHandler := &handler.EventHandler{
		Queue: queue,
	}

	metricsHandler := &handler.MetricsHandler{
		WindowAgg: windowAgg,
	}

	http.HandleFunc("/events", eventHandler.Handle)
	http.HandleFunc("/metrics", metricsHandler.Handle)

	log.Println("event api running on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
