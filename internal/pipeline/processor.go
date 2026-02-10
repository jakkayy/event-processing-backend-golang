package pipeline

import (
	"event-processing-backend-golang/internal/domain"
	"log"
)

// create processor
type Processor interface {
	Process(event domain.Event)
}

type LogginProcessor struct{}

func (p *LogginProcessor) Process(e domain.Event) {
	log.Printf("[processor] type=%s message=%s time=%s\n", e.Type, e.Message, e.Timestamp.Format("15:04:05"))
}
