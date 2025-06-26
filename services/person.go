package services

import (
	"github.com/scottys88/ddd-go/aggregates"
	"github.com/scottys88/ddd-go/domain/person"
	"github.com/scottys88/ddd-go/domain/person/memory"
)

type PersonConfiguration func(ps *PersonService) error

type PersonService struct {
	personRepo person.PersonRepository
}

func NewPersonService(options ...PersonConfiguration) (*PersonService, error) {
	ps := &PersonService{}

	for _, option := range options {
		err := option(ps)
		if err != nil {
			return nil, err
		}
	}
	return ps, nil
}

func WithPersonRepository(repo person.PersonRepository) PersonConfiguration {
	return func(ps *PersonService) error {
		if repo == nil {
			return person.ErrFailedToAddPerson
		}
		ps.personRepo = repo
		return nil
	}
}

func WithMemoryPersonRepository() PersonConfiguration {
	repo := memory.New()
	return WithPersonRepository(repo)
}

func (ps *PersonService) Create(name, email, phone string) (*aggregates.Person, error) {
	// Check if person with this email already exists
	existing, err := ps.personRepo.FindByEmail(email)
	if err == nil && existing != nil {
		return nil, person.ErrFailedToAddPersonAlreadyExists
	}

	// Create new person aggregate
	person, err := aggregates.NewPerson(name, email, phone)
	if err != nil {
		return nil, err
	}

	// Save to repository
	if err := ps.personRepo.Add(person); err != nil {
		return nil, err
	}

	return person, nil
}

func (ps *PersonService) GetByEmail(email string) (*aggregates.Person, error) {
	return ps.personRepo.FindByEmail(email)
}

func (ps *PersonService) UpdateContactInfo(personID string, name, email, phone string) error {
	// Implementation would parse personID to uuid.UUID and update person
	// Left as placeholder for now
	return nil
}