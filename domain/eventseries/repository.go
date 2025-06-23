package eventseries

import (
	"errors"
	"github.com/google/uuid"
	"github.com/scottys88/ddd-go/entity"
)

var (
	ErrEventSeriesNotFound       = errors.New("event series not found")
	ErrFailedToAddEventSeries    = errors.New("could not add event series")
	ErrFailedToUpdateEventSeries = errors.New("could not update event series")
)

type EventSeriesRepository interface {
	Get(uuid uuid.UUID) (*entity.EventSeries, error) // Retrieve an event series by its UUID
	Add(entity entity.EventSeries) error
	Update(entity entity.EventSeries) error
}

type eventSeriesRepository struct {
}
