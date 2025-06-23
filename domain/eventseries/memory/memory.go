// An in memory implementation of the EventSeriesRepository interface.
package memory

import (
	"github.com/google/uuid"
	"github.com/scottys88/ddd-go/aggregates"
	"sync"
)

type MemoryRepository struct {
	eventSeries map[uuid.UUID]aggregates.EventSeries
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		eventSeries: make(map[uuid.UUID]aggregates.EventSeries),
	}
}

func (r *MemoryRepository) Get(id uuid.UUID) (aggregates.EventSeries, error) {
	if eventSeries, exists := r.eventSeries[id]; exists {
		return eventSeries, nil
	}
	return aggregates.EventSeries{}, nil
}

func (r *MemoryRepository) Add(eventSeries aggregates.EventSeries) error {
	return nil
}
func (r *MemoryRepository) Update(eventSeries aggregates.EventSeries) error {
	return nil
}
