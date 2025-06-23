package entity

import (
	"github.com/google/uuid"
	"time"
)

type Person struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Phone    string
	EntryIDs []uuid.UUID // List of event entries associated with the person

	CreatedAt  time.Time  // Timestamp when the entry was created
	UpdatedAt  time.Time  // Timestamp when the entry was last updated
	ArchivedAt *time.Time // Optional timestamp when the entry was archived
}
