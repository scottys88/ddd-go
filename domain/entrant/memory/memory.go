package memory

import (
	"github.com/google/uuid"
	"github.com/scottys88/ddd-go/aggregates"
	"github.com/scottys88/ddd-go/domain/entrant"
	"sync"
)

type EntrantMemoryRepository struct {
	entrants map[uuid.UUID]*aggregates.Entrant
	mutex    sync.RWMutex
}

func New() *EntrantMemoryRepository {
	return &EntrantMemoryRepository{
		entrants: make(map[uuid.UUID]*aggregates.Entrant),
	}
}

func (r *EntrantMemoryRepository) Get(id uuid.UUID) (*aggregates.Entrant, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if ent, exists := r.entrants[id]; exists {
		return ent, nil
	}
	return nil, entrant.ErrEntrantNotFound
}

func (r *EntrantMemoryRepository) Add(ent *aggregates.Entrant) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.entrants[ent.ID()]; exists {
		return entrant.ErrFailedToAddEntrantAlreadyExists
	}

	r.entrants[ent.ID()] = ent
	return nil
}

func (r *EntrantMemoryRepository) Update(ent *aggregates.Entrant) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.entrants[ent.ID()]; !exists {
		return entrant.ErrEntrantNotFound
	}

	r.entrants[ent.ID()] = ent
	return nil
}

func (r *EntrantMemoryRepository) FindByPersonEventAndDistance(personID, eventID, distanceID uuid.UUID) (*aggregates.Entrant, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, ent := range r.entrants {
		if ent.PersonID() == personID && ent.EventID() == eventID && ent.DistanceID() == distanceID {
			return ent, nil
		}
	}
	return nil, entrant.ErrEntrantNotFound
}

func (r *EntrantMemoryRepository) FindByPersonAndEvent(personID, eventID uuid.UUID) ([]*aggregates.Entrant, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var results []*aggregates.Entrant
	for _, ent := range r.entrants {
		if ent.PersonID() == personID && ent.EventID() == eventID {
			results = append(results, ent)
		}
	}
	return results, nil
}

func (r *EntrantMemoryRepository) FindByEvent(eventID uuid.UUID) ([]*aggregates.Entrant, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var results []*aggregates.Entrant
	for _, ent := range r.entrants {
		if ent.EventID() == eventID {
			results = append(results, ent)
		}
	}
	return results, nil
}

func (r *EntrantMemoryRepository) Delete(id uuid.UUID) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.entrants[id]; !exists {
		return entrant.ErrEntrantNotFound
	}

	delete(r.entrants, id)
	return nil
}