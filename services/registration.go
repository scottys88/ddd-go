package services

import (
	"github.com/google/uuid"
	"github.com/scottys88/ddd-go/aggregates"
	"github.com/scottys88/ddd-go/domain/entrant"
	"github.com/scottys88/ddd-go/domain/entry"
)

type RegistrationConfiguration func(rs *RegistrationService) error

type RegistrationService struct {
	entrantRepo entrant.EntrantRepository
	entryRepo   entry.EntryRepository
}

func NewRegistrationService(options ...RegistrationConfiguration) (*RegistrationService, error) {
	rs := &RegistrationService{}

	for _, option := range options {
		err := option(rs)
		if err != nil {
			return nil, err
		}
	}
	return rs, nil
}

func WithEntrantRepository(repo entrant.EntrantRepository) RegistrationConfiguration {
	return func(rs *RegistrationService) error {
		if repo == nil {
			return entrant.ErrFailedToAddEntrant
		}
		rs.entrantRepo = repo
		return nil
	}
}

func WithEntryRepository(repo entry.EntryRepository) RegistrationConfiguration {
	return func(rs *RegistrationService) error {
		if repo == nil {
			return entry.ErrFailedToAddEntry
		}
		rs.entryRepo = repo
		return nil
	}
}

func (rs *RegistrationService) RegisterPersonForDistance(
	personID, eventID, distanceID uuid.UUID,
) (*aggregates.Entry, error) {
	// 1. Check if entrant already exists for this specific combination
	entrant, err := rs.entrantRepo.FindByPersonEventAndDistance(personID, eventID, distanceID)
	if err != nil {
		// Create new entrant for this specific distance
		entrant, err = aggregates.NewEntrant(personID, eventID, distanceID)
		if err != nil {
			return nil, err
		}

		// Save new entrant
		if err := rs.entrantRepo.Add(entrant); err != nil {
			return nil, err
		}
	}

	// 2. Create entry if it doesn't exist
	if !entrant.HasEntry() {
		entry, err := entrant.CreateEntry()
		if err != nil {
			return nil, err
		}

		// Save entry
		if err := rs.entryRepo.Add(entry); err != nil {
			return nil, err
		}

		// Update entrant with entry reference
		if err := rs.entrantRepo.Update(entrant); err != nil {
			return nil, err
		}

		return entry, nil
	}

	// Return existing entry
	return rs.entryRepo.FindByEntrant(entrant.ID())
}

func (rs *RegistrationService) GetPersonRegistrations(personID, eventID uuid.UUID) ([]*aggregates.Entrant, error) {
	return rs.entrantRepo.FindByPersonAndEvent(personID, eventID)
}

func (rs *RegistrationService) GetEventRegistrations(eventID uuid.UUID) ([]*aggregates.Entrant, error) {
	return rs.entrantRepo.FindByEvent(eventID)
}

func (rs *RegistrationService) CancelRegistration(entrantID uuid.UUID) error {
	// Get entrant
	entrant, err := rs.entrantRepo.Get(entrantID)
	if err != nil {
		return err
	}

	// Remove entry if exists
	if entrant.HasEntry() {
		if err := entrant.RemoveEntry(); err != nil {
			return err
		}

		// Update entrant
		if err := rs.entrantRepo.Update(entrant); err != nil {
			return err
		}
	}

	// Archive/delete entrant
	entrant.Archive()
	return rs.entrantRepo.Update(entrant)
}