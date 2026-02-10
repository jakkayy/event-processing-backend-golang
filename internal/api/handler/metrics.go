package handler

import (
	"encoding/json"
	"net/http"

	"event-processing-backend-golang/internal/pipeline"
)

type MetricsHandler struct {
	WindowAgg *pipeline.WindowAggregationProcessor
}

func (h *MetricsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	snapshot := h.WindowAgg.Snapshot()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(snapshot)
}
