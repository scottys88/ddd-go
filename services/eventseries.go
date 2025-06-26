package services

import (
	"github.com/google/uuid"
	"github.com/scottys88/ddd-go/aggregates"
	"github.com/scottys88/ddd-go/domain/eventseries"
	"github.com/scottys88/ddd-go/domain/eventseries/memory"
)

type EventSeriesConfiguration func(es *EventSeriesService) error

type EventSeriesService struct {
	eventSeries eventseries.EventSeriesRepository
}

func NewEventSeriesService(options ...EventSeriesConfiguration) (*EventSeriesService, error) {
	es := &EventSeriesService{}

	for _, option := range options {
		err := option(es)

		if err != nil {
			return nil, err
		}
	}
	return es, nil
}

func WithEventSeriesRepository(repo eventseries.EventSeriesRepository) EventSeriesConfiguration {
	return func(es *EventSeriesService) error {
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

func (es *EventSeriesService) Create(name, description string, organiserID uuid.UUID) (*aggregates.EventSeries, error) {
	// TODO: fetch the organiser from the customer repo to be built
	// TODO: ensure they exist

	// Create new event series aggregate
	eventSeries, err := aggregates.NewEventSeries(name, description, organiserID)
	if err != nil {
		return nil, err
	}

	// Save to repository
	if err := es.eventSeries.Add(*eventSeries); err != nil {
		return nil, err
	}

	return eventSeries, nil
}
