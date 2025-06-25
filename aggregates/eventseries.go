package aggregates

import (
	"errors"
	"github.com/google/uuid"
	"github.com/scottys88/ddd-go/entity"
	"time"
)

type EventSeries struct {
	id          uuid.UUID
	name        string
	description string
	events      []entity.Event
	organiserID uuid.UUID
	createdAt   time.Time
	updatedAt   time.Time
	archivedAt  *time.Time
}

func NewEventSeries(name, description string, organiserID uuid.UUID) (*EventSeries, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	if organiserID == uuid.Nil {
		return nil, errors.New("organiser ID is required")
	}

	if description == "" {
		return nil, errors.New("description is required")
	}

	return &EventSeries{
		id:          uuid.New(),
		name:        name,
		description: description,
		organiserID: organiserID,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
		events:      []entity.Event{},
	}, nil
}

func (es *EventSeries) ID() uuid.UUID {
	return es.id
}

func (es *EventSeries) Name() string {
	return es.name
}

func (es *EventSeries) Description() string {
	return es.description
}

func (es *EventSeries) Events() []entity.Event {
	return es.events
}

func (es *EventSeries) OrganiserID() uuid.UUID {
	return es.organiserID
}

func (es *EventSeries) CreatedAt() time.Time {
	return es.createdAt
}

func (es *EventSeries) UpdatedAt() time.Time {
	return es.updatedAt
}

func (es *EventSeries) ArchivedAt() *time.Time {
	return es.archivedAt
}

func (es *EventSeries) UpdateName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}
	es.name = name
	es.updatedAt = time.Now()
	return nil
}

func (es *EventSeries) UpdateDescription(description string) error {
	if description == "" {
		return errors.New("description is required")
	}
	es.description = description
	es.updatedAt = time.Now()
	return nil
}

func (es *EventSeries) AddEvent(event entity.Event) {
	es.events = append(es.events, event)
	es.updatedAt = time.Now()
}

func (es *EventSeries) UpdateOrganiserID(organiserID uuid.UUID) error {
	if organiserID == uuid.Nil {
		return errors.New("organiser ID is required")
	}
	es.organiserID = organiserID
	es.updatedAt = time.Now()
	return nil
}

func (es *EventSeries) RemoveEvent(eventID uuid.UUID) error {
	for i, event := range es.events {
		id := event.ID
		if id == eventID {
			es.events = append(es.events[:i], es.events[i+1:]...)
			es.updatedAt = time.Now()
			return nil
		}
	}
	return errors.New("event not found")
}

func (es *EventSeries) SetEvents(events []entity.Event) {
	es.events = events
	es.updatedAt = time.Now()
}

func (es *EventSeries) Archive() {
	now := time.Now()
	es.archivedAt = &now
	es.updatedAt = now
}
