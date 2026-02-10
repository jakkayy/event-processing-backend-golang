package pipeline

import (
	"log"
	"sync"
	"time"

	"event-processing-backend-golang/internal/domain"
)

type WindowAggregationProcessor struct {
	mu      sync.Mutex
	window  time.Duration
	maxKeep int
	windows map[time.Time]map[string]int
}

func NewWindowAggregationProcessor(window time.Duration, maxKeep int) *WindowAggregationProcessor {
	return &WindowAggregationProcessor{
		window:  window,
		maxKeep: maxKeep,
		windows: make(map[time.Time]map[string]int),
	}
}

func (p *WindowAggregationProcessor) Process(e domain.Event) {
	p.mu.Lock()
	defer p.mu.Unlock()

	bucket := e.Timestamp.Truncate(p.window)

	if _, ok := p.windows[bucket]; !ok {
		p.windows[bucket] = make(map[string]int)
	}
	p.windows[bucket][e.Type]++

	p.cleanupLocked()

	log.Printf("[window] %s type=%s count=%d",
		bucket.Format("15:04"),
		e.Type,
		p.windows[bucket][e.Type],
	)
}

// delete very od bucket
func (p *WindowAggregationProcessor) cleanupLocked() {
	if len(p.windows) <= p.maxKeep {
		return
	}

	// search oldest bucket
	var oldest time.Time
	for t := range p.windows {
		if oldest.IsZero() || t.Before(oldest) {
			oldest = t
		}
	}
	delete(p.windows, oldest)
}

func (p *WindowAggregationProcessor) Snapshot() map[string]map[string]int {
	p.mu.Lock()
	defer p.mu.Unlock()

	result := make(map[string]map[string]int)
	for t, counters := range p.windows {
		key := t.Format("15:04")
		result[key] = make(map[string]int)
		for k, v := range counters {
			result[key][k] = v
		}
	}
	return result
}
