package aggregates

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Entrant struct {
	id         uuid.UUID
	personID   uuid.UUID
	eventID    uuid.UUID
	distanceID uuid.UUID
	entryID    *uuid.UUID // Optional - set when Entry is created
	createdAt  time.Time
	updatedAt  time.Time
	archivedAt *time.Time
}

func NewEntrant(personID, eventID, distanceID uuid.UUID) (*Entrant, error) {
	if personID == uuid.Nil {
		return nil, errors.New("person ID is required")
	}

	if eventID == uuid.Nil {
		return nil, errors.New("event ID is required")
	}

	if distanceID == uuid.Nil {
		return nil, errors.New("distance ID is required")
	}

	return &Entrant{
		id:         uuid.New(),
		personID:   personID,
		eventID:    eventID,
		distanceID: distanceID,
		createdAt:  time.Now(),
		updatedAt:  time.Now(),
	}, nil
}

func (e *Entrant) ID() uuid.UUID {
	return e.id
}

func (e *Entrant) PersonID() uuid.UUID {
	return e.personID
}

func (e *Entrant) EventID() uuid.UUID {
	return e.eventID
}

func (e *Entrant) DistanceID() uuid.UUID {
	return e.distanceID
}

func (e *Entrant) EntryID() *uuid.UUID {
	return e.entryID
}

func (e *Entrant) CreatedAt() time.Time {
	return e.createdAt
}

func (e *Entrant) UpdatedAt() time.Time {
	return e.updatedAt
}

func (e *Entrant) ArchivedAt() *time.Time {
	return e.archivedAt
}

func (e *Entrant) CreateEntry() (*Entry, error) {
	if e.entryID != nil {
		return nil, errors.New("entrant already has an entry")
	}

	entry, err := NewEntry(e.distanceID, e.id, e.eventID)
	if err != nil {
		return nil, err
	}

	entryID := entry.ID()
	e.entryID = &entryID
	e.updatedAt = time.Now()

	return entry, nil
}

func (e *Entrant) HasEntry() bool {
	return e.entryID != nil
}

func (e *Entrant) UpdateDistanceID(distanceID uuid.UUID) error {
	if distanceID == uuid.Nil {
		return errors.New("distance ID is required")
	}

	if e.entryID != nil {
		return errors.New("cannot change distance after entry is created")
	}

	e.distanceID = distanceID
	e.updatedAt = time.Now()
	return nil
}

func (e *Entrant) RemoveEntry() error {
	if e.entryID == nil {
		return errors.New("no entry to remove")
	}

	e.entryID = nil
	e.updatedAt = time.Now()
	return nil
}

func (e *Entrant) Archive() {
	now := time.Now()
	e.archivedAt = &now
	e.updatedAt = now
}
