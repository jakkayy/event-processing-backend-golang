package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"event-processing-backend-golang/internal/domain"
	"event-processing-backend-golang/internal/pipeline"

	"github.com/google/uuid"
)

type EventHandler struct {
	Queue *pipeline.EventQueue
}

func (h *EventHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Type    string
		Message string
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	event := domain.Event{
		ID:        uuid.NewString(),
		Type:      input.Type,
		Message:   input.Message,
		Timestamp: time.Now(),
	}

	h.Queue.Ch <- event

	w.WriteHeader(http.StatusAccepted)
}
