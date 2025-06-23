package valueobject

import (
	"time"

	"github.com/google/uuid"
)

// Transaction is a value object with no identity.
type Transaction struct {
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
