package memory

import (
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/scottys88/ddd-go/aggregates"
	"github.com/scottys88/ddd-go/domain/person"
)

type PersonMemoryRepository struct {
	persons map[uuid.UUID]*aggregates.Person
	mutex   sync.RWMutex
}

func New() *PersonMemoryRepository {
	return &PersonMemoryRepository{
		persons: make(map[uuid.UUID]*aggregates.Person),
	}
}

func (r *PersonMemoryRepository) Get(id uuid.UUID) (*aggregates.Person, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if p, exists := r.persons[id]; exists {
		return p, nil
	}
	return nil, person.ErrPersonNotFound
}

func (r *PersonMemoryRepository) Add(p *aggregates.Person) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.persons[p.ID()]; exists {
		return person.ErrFailedToAddPersonAlreadyExists
	}

	r.persons[p.ID()] = p
	return nil
}

func (r *PersonMemoryRepository) Update(p *aggregates.Person) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.persons[p.ID()]; !exists {
		return person.ErrPersonNotFound
	}

	r.persons[p.ID()] = p
	return nil
}

func (r *PersonMemoryRepository) FindByEmail(email string) (*aggregates.Person, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, p := range r.persons {
		if strings.EqualFold(p.Email(), email) {
			return p, nil
		}
	}
	return nil, person.ErrPersonNotFound
}

func (r *PersonMemoryRepository) FindByNameAndEmail(name, email string) (*aggregates.Person, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, p := range r.persons {
		if strings.EqualFold(p.Name(), name) && strings.EqualFold(p.Email(), email) {
			return p, nil
		}
	}
	return nil, person.ErrPersonNotFound
}

func (r *PersonMemoryRepository) Delete(id uuid.UUID) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.persons[id]; !exists {
		return person.ErrPersonNotFound
	}

	delete(r.persons, id)
	return nil
}
