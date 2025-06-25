// An in memory implementation of the EventSeriesRepository interface.
package memory

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/scottys88/ddd-go/aggregates"
	"github.com/scottys88/ddd-go/domain/eventseries"
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
	return aggregates.EventSeries{}, errors.New(fmt.Sprint(eventseries.ErrEventSeriesNotFound))
}

func (r *MemoryRepository) Add(eventSeries aggregates.EventSeries) error {
	if r.eventSeries == nil {
		r.Lock()
		r.eventSeries = make(map[uuid.UUID]aggregates.EventSeries)
		r.Unlock()
	}
	// Check if the event series already exists
	if _, exists := r.eventSeries[eventSeries.ID()]; exists {
		return errors.New(fmt.Sprint(eventseries.ErrFailedToAddEventSeriesAlreadyExists))
	}
	r.Lock()
	r.eventSeries[eventSeries.ID()] = eventSeries
	r.Unlock()
	return nil
}
func (r *MemoryRepository) Update(eventSeries aggregates.EventSeries) error {
	// Check if exists
	if _, exists := r.eventSeries[eventSeries.ID()]; !exists {
		return errors.New(fmt.Sprint(eventseries.ErrEventSeriesNotFound))
	}

	r.Lock()
	r.eventSeries[eventSeries.ID()] = eventSeries
	r.Unlock()
	return nil
}
