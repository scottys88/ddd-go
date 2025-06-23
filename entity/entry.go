package entity

import (
	"time"

	"github.com/google/uuid"
)

type Entry struct {
	ID         uuid.UUID
	DistanceID uuid.UUID
	EntrantID  uuid.UUID
	EventID    uuid.UUID // ID of the event associated with the entry

	CreatedAt  time.Time  // Timestamp when the entry was created
	UpdatedAt  time.Time  // Timestamp when the entry was last updated
	ArchivedAt *time.Time // Optional timestamp when the entry was archived
}
