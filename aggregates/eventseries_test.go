package aggregates_test

import (
	"errors"
	"github.com/scottys88/ddd-go/aggregates"
	"testing"

	"github.com/google/uuid"
)

func TestNewEventSeries_NewEventSeries(t *testing.T) {
	type args struct {
		name          string
		description   string
		organiserID   string
		expectedError error
	}

	testCases := []args{
		{
			name:          "Test Event Series",
			description:   "This is a test event series",
			organiserID:   "123e4567-e89b-12d3-a456-426614174000", // Example UUID
			expectedError: nil,
		},
		{
			name:          "",
			description:   "This event series has no name",
			organiserID:   "123e4567-e89b-12d3-a456-426614174000", // Example UUID
			expectedError: errors.New("name is required"),
		},
		{
			name:          "No description",
			description:   "",
			organiserID:   "123e4567-e89b-12d3-a456-426614174000", // Example UUID
			expectedError: errors.New("description is required"),
		}, {
			name:          "No orgID",
			description:   "Description",
			organiserID:   "", // Example UUID
			expectedError: errors.New("organiser ID is required"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			orgID, err := uuid.Parse(tc.organiserID) // Convert string to UUID, handle error if needed
			es, err := aggregates.NewEventSeries(tc.name, tc.description, orgID)

			if tc.expectedError == nil {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if es.Name() != tc.name {
					t.Errorf("Expected name %s, got %s", tc.name, es.Name())
				}
				if es.Description() != tc.description {
					t.Errorf("Expected description %s, got %s", tc.description, es.Description())
				}
				if es.OrganiserID() != orgID {
					t.Errorf("Expected organiser ID %s, got %s", orgID, es.OrganiserID())
				}
			} else {
				if err == nil || err.Error() != tc.expectedError.Error() {
					t.Errorf("Expected error '%v', got '%v'", tc.expectedError, err)
				}
			}
		})
	}
}