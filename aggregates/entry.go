package aggregates

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Entry struct {
	id         uuid.UUID
	distanceID uuid.UUID
	entrantID  uuid.UUID
	eventID    uuid.UUID
	createdAt  time.Time
	updatedAt  time.Time
	archivedAt *time.Time
}

func NewEntry(distanceID, entrantID, eventID uuid.UUID) (*Entry, error) {
	if distanceID == uuid.Nil {
		return nil, errors.New("distance ID is required")
	}

	if entrantID == uuid.Nil {
		return nil, errors.New("entrant ID is required")
	}

	if eventID == uuid.Nil {
		return nil, errors.New("event ID is required")
	}

	return &Entry{
		id:         uuid.New(),
		distanceID: distanceID,
		entrantID:  entrantID,
		eventID:    eventID,
		createdAt:  time.Now(),
		updatedAt:  time.Now(),
	}, nil
}

func (e *Entry) ID() uuid.UUID {
	return e.id
}

func (e *Entry) DistanceID() uuid.UUID {
	return e.distanceID
}

func (e *Entry) EntrantID() uuid.UUID {
	return e.entrantID
}

func (e *Entry) EventID() uuid.UUID {
	return e.eventID
}

func (e *Entry) CreatedAt() time.Time {
	return e.createdAt
}

func (e *Entry) UpdatedAt() time.Time {
	return e.updatedAt
}

func (e *Entry) ArchivedAt() *time.Time {
	return e.archivedAt
}

func (e *Entry) UpdateDistanceID(distanceID uuid.UUID) error {
	if distanceID == uuid.Nil {
		return errors.New("distance ID is required")
	}
	e.distanceID = distanceID
	e.updatedAt = time.Now()
	return nil
}

func (e *Entry) UpdateEntrantID(entrantID uuid.UUID) error {
	if entrantID == uuid.Nil {
		return errors.New("entrant ID is required")
	}
	e.entrantID = entrantID
	e.updatedAt = time.Now()
	return nil
}

func (e *Entry) UpdateEventID(eventID uuid.UUID) error {
	if eventID == uuid.Nil {
		return errors.New("event ID is required")
	}
	e.eventID = eventID
	e.updatedAt = time.Now()
	return nil
}

func (e *Entry) Archive() {
	now := time.Now()
	e.archivedAt = &now
	e.updatedAt = now
}
