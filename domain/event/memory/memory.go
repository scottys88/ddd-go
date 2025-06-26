package memory

import (
	"github.com/google/uuid"
	"github.com/scottys88/ddd-go/aggregates"
	"github.com/scottys88/ddd-go/domain/event"
	"sync"
)

type EventMemoryRepository struct {
	events map[uuid.UUID]*aggregates.Event
	mutex  sync.RWMutex
}

func New() *EventMemoryRepository {
	return &EventMemoryRepository{
		events: make(map[uuid.UUID]*aggregates.Event),
	}
}

func (r *EventMemoryRepository) Get(id uuid.UUID) (*aggregates.Event, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if evt, exists := r.events[id]; exists {
		return evt, nil
	}
	return nil, event.ErrEventNotFound
}

func (r *EventMemoryRepository) Add(evt *aggregates.Event) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.events[evt.ID()]; exists {
		return event.ErrFailedToAddEventAlreadyExists
	}

	r.events[evt.ID()] = evt
	return nil
}

func (r *EventMemoryRepository) Update(evt *aggregates.Event) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.events[evt.ID()]; !exists {
		return event.ErrEventNotFound
	}

	r.events[evt.ID()] = evt
	return nil
}

func (r *EventMemoryRepository) FindByOrganiser(organiserID uuid.UUID) ([]*aggregates.Event, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var results []*aggregates.Event
	for _, evt := range r.events {
		if evt.OrganiserID() == organiserID {
			results = append(results, evt)
		}
	}
	return results, nil
}

func (r *EventMemoryRepository) FindByDateRange(startDate, endDate string) ([]*aggregates.Event, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	// TODO: Implement date range filtering when Event aggregate has date field
	var results []*aggregates.Event
	for _, evt := range r.events {
		results = append(results, evt)
	}
	return results, nil
}

func (r *EventMemoryRepository) Delete(id uuid.UUID) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.events[id]; !exists {
		return event.ErrEventNotFound
	}

	delete(r.events, id)
	return nil
}