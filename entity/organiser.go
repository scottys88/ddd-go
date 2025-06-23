package entity

import (
	"github.com/google/uuid"
	"time"
)

type Organiser struct {
	PersonID uuid.UUID
	EventIDs []uuid.UUID // List of event IDs associated with the organiser

	CreatedAt  time.Time  // Timestamp when the entry was created
	UpdatedAt  time.Time  // Timestamp when the entry was last updated
	ArchivedAt *time.Time // Optional timestamp when the entry was archived
}
