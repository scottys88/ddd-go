package entity

import (
	"github.com/google/uuid"
	"time"
)

type Entrant struct {
	ID         uuid.UUID
	PersonID   uuid.UUID
	EventID    uuid.UUID
	DistanceID uuid.UUID // ID of the distance associated with the entrant

	CreatedAt  time.Time  // Timestamp when the entry was created
	UpdatedAt  time.Time  // Timestamp when the entry was last updated
	ArchivedAt *time.Time // Optional timestamp when the entry was archived
}
