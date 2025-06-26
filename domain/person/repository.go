package person

import (
	"errors"
	"github.com/google/uuid"
	"github.com/scottys88/ddd-go/aggregates"
)

var (
	ErrPersonNotFound                 = errors.New("person not found")
	ErrFailedToAddPerson              = errors.New("could not add person")
	ErrFailedToUpdatePerson           = errors.New("could not update person")
	ErrFailedToAddPersonAlreadyExists = errors.New("could not add person, it already exists")
)

type PersonRepository interface {
	Get(id uuid.UUID) (*aggregates.Person, error)
	Add(person *aggregates.Person) error
	Update(person *aggregates.Person) error
	FindByEmail(email string) (*aggregates.Person, error)
	FindByNameAndEmail(name, email string) (*aggregates.Person, error)
	Delete(id uuid.UUID) error
}