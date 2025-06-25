package memory

import (
	"errors"
	"github.com/google/uuid"
	"github.com/scottys88/ddd-go/aggregates"
	"testing"
)

func TestMemoryRepository_Get(t *testing.T) {
	type TestCase struct {
		name          string
		id            uuid.UUID
		expectedError error
	}

	es, err := aggregates.NewEventSeries("Test Series", "A test event series", uuid.New())

	if es == nil || err != nil {
		t.Fatalf("Failed to create event series: %v", err)
	}

	repo := MemoryRepository{
		eventSeries: map[uuid.UUID]aggregates.EventSeries{
			es.ID(): *es,
		},
	}

	testCases := []TestCase{
		{
			name:          "Valid ID",
			id:            es.ID(),
			expectedError: nil,
		},
		{
			name:          "Invalid ID",
			id:            uuid.MustParse("00000000-0000-0000-0000-000000000000"),
			expectedError: errors.New("Event series not found"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			eventSeries, err := repo.Get(tc.id)
			if err != nil && err.Error() != tc.expectedError.Error() {
				t.Errorf("Expected error %v, got %v", tc.expectedError, err)
			}
			if err == nil && eventSeries.ID() != tc.id {
				t.Errorf("Expected event series ID %v, got %v", tc.id, eventSeries.ID())
			}
		})
	}
}
