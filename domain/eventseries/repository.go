package eventseries

import (
	"errors"
	"github.com/google/uuid"
	"github.com/scottys88/ddd-go/aggregates"
)

var (
	ErrEventSeriesNotFound                 = errors.New("event series not found")
	ErrFailedToAddEventSeries              = errors.New("could not add event series")
	ErrFailedToUpdateEventSeries           = errors.New("could not update event series")
	ErrFailedToAddEventSeriesAlreadyExists = errors.New("could not add event series, it already exists")
)

type EventSeriesRepository interface {
	Get(uuid uuid.UUID) (aggregates.EventSeries, error) // Retrieve an event series by its UUID
	Add(entity aggregates.EventSeries) error
	Update(entity aggregates.EventSeries) error
}

type eventSeriesRepository struct {
}
