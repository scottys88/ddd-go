package aggregates

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/scottys88/ddd-go/entity"
)

type Event struct {
	id          uuid.UUID
	name        string
	date        time.Time
	distances   []entity.Distance
	organiserID uuid.UUID
	createdAt   time.Time
	updatedAt   time.Time
	archivedAt  *time.Time
}

func NewEvent(name string, date time.Time, organiserID uuid.UUID) (*Event, error) {
	if name == "" {
		return nil, errors.New("event name is required")
	}

	if organiserID == uuid.Nil {
		return nil, errors.New("organiser ID is required")
	}

	if date.IsZero() {
		return nil, errors.New("event date is required")
	}

	return &Event{
		id:          uuid.New(),
		name:        name,
		date:        date,
		organiserID: organiserID,
		distances:   []entity.Distance{},
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}, nil
}

func (e *Event) ID() uuid.UUID {
	return e.id
}

func (e *Event) Name() string {
	return e.name
}

func (e *Event) Date() time.Time {
	return e.date
}

func (e *Event) Distances() []entity.Distance {
	return e.distances
}

func (e *Event) OrganiserID() uuid.UUID {
	return e.organiserID
}

func (e *Event) CreatedAt() time.Time {
	return e.createdAt
}

func (e *Event) UpdatedAt() time.Time {
	return e.updatedAt
}

func (e *Event) ArchivedAt() *time.Time {
	return e.archivedAt
}

func (e *Event) UpdateName(name string) error {
	if name == "" {
		return errors.New("event name is required")
	}
	e.name = name
	e.updatedAt = time.Now()
	return nil
}

func (e *Event) UpdateDate(date time.Time) error {
	if date.IsZero() {
		return errors.New("event date is required")
	}
	e.date = date
	e.updatedAt = time.Now()
	return nil
}

func (e *Event) UpdateOrganiserID(organiserID uuid.UUID) error {
	if organiserID == uuid.Nil {
		return errors.New("organiser ID is required")
	}
	e.organiserID = organiserID
	e.updatedAt = time.Now()
	return nil
}

func (e *Event) AddDistance(distance entity.Distance) error {
	// Check if distance already exists
	for _, d := range e.distances {
		if d.ID == distance.ID {
			return errors.New("distance already exists in event")
		}
	}

	e.distances = append(e.distances, distance)
	e.updatedAt = time.Now()
	return nil
}

func (e *Event) RemoveDistance(distanceID uuid.UUID) error {
	for i, distance := range e.distances {
		if distance.ID == distanceID {
			e.distances = append(e.distances[:i], e.distances[i+1:]...)
			e.updatedAt = time.Now()
			return nil
		}
	}
	return errors.New("distance not found in event")
}

func (e *Event) GetDistance(distanceID uuid.UUID) (*entity.Distance, error) {
	for _, distance := range e.distances {
		if distance.ID == distanceID {
			return &distance, nil
		}
	}
	return nil, errors.New("distance not found in event")
}

func (e *Event) HasDistance(distanceID uuid.UUID) bool {
	for _, distance := range e.distances {
		if distance.ID == distanceID {
			return true
		}
	}
	return false
}

func (e *Event) SetDistances(distances []entity.Distance) {
	e.distances = distances
	e.updatedAt = time.Now()
}

func (e *Event) Archive() {
	now := time.Now()
	e.archivedAt = &now
	e.updatedAt = now
}
