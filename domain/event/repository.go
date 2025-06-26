package event

import (
	"errors"
	"github.com/google/uuid"
	"github.com/scottys88/ddd-go/aggregates"
)

var (
	ErrEventNotFound                 = errors.New("event not found")
	ErrFailedToAddEvent              = errors.New("could not add event")
	ErrFailedToUpdateEvent           = errors.New("could not update event")
	ErrFailedToAddEventAlreadyExists = errors.New("could not add event, it already exists")
)

type EventRepository interface {
	Get(id uuid.UUID) (*aggregates.Event, error)
	Add(event *aggregates.Event) error
	Update(event *aggregates.Event) error
	FindByOrganiser(organiserID uuid.UUID) ([]*aggregates.Event, error)
	FindByDateRange(startDate, endDate string) ([]*aggregates.Event, error)
	Delete(id uuid.UUID) error
}