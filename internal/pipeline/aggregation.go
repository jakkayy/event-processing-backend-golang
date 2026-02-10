package pipeline

import (
	"log"
	"sync"

	"event-processing-backend-golang/internal/domain"
)

type AggregationProcessor struct {
	mu       sync.Mutex
	counters map[string]int
}

func NewAggregationProcessor() *AggregationProcessor {
	return &AggregationProcessor{
		counters: make(map[string]int),
	}
}

func (p *AggregationProcessor) Process(e domain.Event) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.counters[e.Type]++

	log.Printf("[aggregate] type=%s count=%d\n", e.Type, p.counters[e.Type])
}

func (p *AggregationProcessor) Snapshot() map[string]int {
	p.mu.Lock()
	defer p.mu.Unlock()

	result := make(map[string]int)
	for k, v := range p.counters {
		result[k] = v
	}
	return result
}
