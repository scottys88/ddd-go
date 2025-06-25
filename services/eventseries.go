package services

import (
	"github.com/google/uuid"
	"github.com/scottys88/ddd-go/aggregates"
	"github.com/scottys88/ddd-go/domain/eventseries"
	"github.com/scottys88/ddd-go/domain/eventseries/memory"
)

type EventSeriesConfiguration func(es *EventService) error

type EventService struct {
	eventSeries eventseries.EventSeriesRepository
}

func NewEventService(options ...EventSeriesConfiguration) (*EventService, error) {
	es := &EventService{}

	for _, option := range options {
		err := option(es)

		if err != nil {
			return nil, err
		}
	}
	return es, nil
}

func WithEventSeriesRepository(repo eventseries.EventSeriesRepository) EventSeriesConfiguration {
	return func(es *EventService) error {
		if repo == nil {
			return eventseries.ErrFailedToAddEventSeries
		}
		es.eventSeries = repo
		return nil
	}
}

func WithMemoryEventSeriesRepository() EventSeriesConfiguration {
	repo := memory.New()
	return WithEventSeriesRepository(repo)
}

func (es *EventService) Create(organiseID uuid.UUID) (aggregates.EventSeries, error) {
	// fetch the organiser from the customer repo to be built
	// ensure they exist
	return aggregates.EventSeries{}, nil
}
