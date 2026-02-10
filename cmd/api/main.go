package main

import (
	"log"
	"net/http"

	"event-processing-backend-golang/internal/api/handler"
	"event-processing-backend-golang/internal/pipeline"
)

func main() {
	queue := pipeline.NewEventQueue(100)

	aggregator := pipeline.NewAggregationProcessor()

	workerCount := 3

	for i := 1; i <= workerCount; i++ {
		worker := &pipeline.Worker{
			ID:        i,
			Queue:     queue,
			Processor: aggregator,
		}
		worker.Start()
	}

	eventHandler := &handler.EventHandler{
		Queue: queue,
	}

	metricsHandler := &handler.MetricsHandler{
		Aggregator: aggregator,
	}

	http.HandleFunc("/events", eventHandler.Handle)
	http.HandleFunc("/metrics", metricsHandler.Handle)

	log.Println("event api running on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
