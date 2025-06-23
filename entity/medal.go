package entity

import (
	"github.com/google/uuid"
	"time"
)

type Medal struct {
	ID          uuid.UUID `json:"id"`          // Unique identifier for the medal
	Name        string    `json:"name"`        // Name of the medal
	Description string    `json:"description"` // Description of the medal
	DistanceID  uuid.UUID `json:"distance_id"` // ID of the distance associated with the medal
	EntryID     uuid.UUID `json:"entry_id"`    // Which entry won the medal

	CreatedAt  time.Time  // Timestamp when the entry was created
	UpdatedAt  time.Time  // Timestamp when the entry was last updated
	ArchivedAt *time.Time // Optional timestamp when the entry was archived
}
