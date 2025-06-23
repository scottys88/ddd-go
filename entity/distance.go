package entity

import (
	"github.com/google/uuid"
	"time"
)

type Distance struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Length  float64   `json:"length"`
	EventID uuid.UUID `json:"event_id"`
	Medals  []Medal   `json:"medal_id"`

	CreatedAt  time.Time  // Timestamp when the entry was created
	UpdatedAt  time.Time  // Timestamp when the entry was last updated
	ArchivedAt *time.Time // Optional timestamp when the entry was archived
}
