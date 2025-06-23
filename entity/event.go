package entity

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	ID          uuid.UUID
	Name        string
	Date        time.Time
	Distances   []Distance
	Entries     []Entry // List of entry IDs associated with the event
	OrganiserID uuid.UUID

	CreatedAt  time.Time  // Timestamp when the entry was created
	UpdatedAt  time.Time  // Timestamp when the entry was last updated
	ArchivedAt *time.Time // Optional timestamp when the entry was archived
}
