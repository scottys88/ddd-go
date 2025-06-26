package memory

import (
	"github.com/google/uuid"
	"github.com/scottys88/ddd-go/aggregates"
	"github.com/scottys88/ddd-go/domain/entry"
	"sync"
)

type EntryMemoryRepository struct {
	entries map[uuid.UUID]*aggregates.Entry
	mutex   sync.RWMutex
}

func New() *EntryMemoryRepository {
	return &EntryMemoryRepository{
		entries: make(map[uuid.UUID]*aggregates.Entry),
	}
}

func (r *EntryMemoryRepository) Get(id uuid.UUID) (*aggregates.Entry, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if ent, exists := r.entries[id]; exists {
		return ent, nil
	}
	return nil, entry.ErrEntryNotFound
}

func (r *EntryMemoryRepository) Add(ent *aggregates.Entry) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.entries[ent.ID()]; exists {
		return entry.ErrFailedToAddEntryAlreadyExists
	}

	r.entries[ent.ID()] = ent
	return nil
}

func (r *EntryMemoryRepository) Update(ent *aggregates.Entry) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.entries[ent.ID()]; !exists {
		return entry.ErrEntryNotFound
	}

	r.entries[ent.ID()] = ent
	return nil
}

func (r *EntryMemoryRepository) FindByEntrant(entrantID uuid.UUID) (*aggregates.Entry, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, ent := range r.entries {
		if ent.EntrantID() == entrantID {
			return ent, nil
		}
	}
	return nil, entry.ErrEntryNotFound
}

func (r *EntryMemoryRepository) FindByEvent(eventID uuid.UUID) ([]*aggregates.Entry, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var results []*aggregates.Entry
	for _, ent := range r.entries {
		if ent.EventID() == eventID {
			results = append(results, ent)
		}
	}
	return results, nil
}

func (r *EntryMemoryRepository) FindByDistance(distanceID uuid.UUID) ([]*aggregates.Entry, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var results []*aggregates.Entry
	for _, ent := range r.entries {
		if ent.DistanceID() == distanceID {
			results = append(results, ent)
		}
	}
	return results, nil
}

func (r *EntryMemoryRepository) FindByEventAndDistance(eventID, distanceID uuid.UUID) ([]*aggregates.Entry, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var results []*aggregates.Entry
	for _, ent := range r.entries {
		if ent.EventID() == eventID && ent.DistanceID() == distanceID {
			results = append(results, ent)
		}
	}
	return results, nil
}

func (r *EntryMemoryRepository) Delete(id uuid.UUID) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.entries[id]; !exists {
		return entry.ErrEntryNotFound
	}

	delete(r.entries, id)
	return nil
}