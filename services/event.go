package services

import (
	"time"
	
	"github.com/google/uuid"
	"github.com/scottys88/ddd-go/aggregates"
	"github.com/scottys88/ddd-go/domain/event"
	"github.com/scottys88/ddd-go/domain/event/memory"
)

type EventConfiguration func(es *EventService) error

type EventService struct {
	eventRepo event.EventRepository
}

func NewEventService(options ...EventConfiguration) (*EventService, error) {
	es := &EventService{}

	for _, option := range options {
		err := option(es)

		if err != nil {
			return nil, err
		}
	}
	return es, nil
}

func WithEventRepository(repo event.EventRepository) EventConfiguration {
	return func(es *EventService) error {
		if repo == nil {
			return event.ErrFailedToAddEvent
		}
		es.eventRepo = repo
		return nil
	}
}

func WithMemoryEventRepository() EventConfiguration {
	repo := memory.New()
	return WithEventRepository(repo)
}

func (es *EventService) Create(name string, eventDate time.Time, organiserID uuid.UUID) (*aggregates.Event, error) {
	// TODO: Validate organiser exists

	// Create new event aggregate
	evt, err := aggregates.NewEvent(name, eventDate, organiserID)
	if err != nil {
		return nil, err
	}

	// Save to repository
	if err := es.eventRepo.Add(evt); err != nil {
		return nil, err
	}

	return evt, nil
}
