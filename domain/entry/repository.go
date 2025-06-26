package entry

import (
	"errors"
	"github.com/google/uuid"
	"github.com/scottys88/ddd-go/aggregates"
)

var (
	ErrEntryNotFound                 = errors.New("entry not found")
	ErrFailedToAddEntry              = errors.New("could not add entry")
	ErrFailedToUpdateEntry           = errors.New("could not update entry")
	ErrFailedToAddEntryAlreadyExists = errors.New("could not add entry, it already exists")
)

type EntryRepository interface {
	Get(id uuid.UUID) (*aggregates.Entry, error)
	Add(entry *aggregates.Entry) error
	Update(entry *aggregates.Entry) error
	FindByEntrant(entrantID uuid.UUID) (*aggregates.Entry, error)
	FindByEvent(eventID uuid.UUID) ([]*aggregates.Entry, error)
	FindByDistance(distanceID uuid.UUID) ([]*aggregates.Entry, error)
	FindByEventAndDistance(eventID, distanceID uuid.UUID) ([]*aggregates.Entry, error)
	Delete(id uuid.UUID) error
}