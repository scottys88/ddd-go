package entrant

import (
	"errors"
	"github.com/google/uuid"
	"github.com/scottys88/ddd-go/aggregates"
)

var (
	ErrEntrantNotFound                 = errors.New("entrant not found")
	ErrFailedToAddEntrant              = errors.New("could not add entrant")
	ErrFailedToUpdateEntrant           = errors.New("could not update entrant")
	ErrFailedToAddEntrantAlreadyExists = errors.New("could not add entrant, it already exists")
)

type EntrantRepository interface {
	Get(id uuid.UUID) (*aggregates.Entrant, error)
	Add(entrant *aggregates.Entrant) error
	Update(entrant *aggregates.Entrant) error
	FindByPersonEventAndDistance(personID, eventID, distanceID uuid.UUID) (*aggregates.Entrant, error)
	FindByPersonAndEvent(personID, eventID uuid.UUID) ([]*aggregates.Entrant, error)
	FindByEvent(eventID uuid.UUID) ([]*aggregates.Entrant, error)
	Delete(id uuid.UUID) error
}