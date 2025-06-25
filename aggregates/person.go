package aggregates

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Person struct {
	id         uuid.UUID
	name       string
	email      string
	phone      string
	entryIDs   []uuid.UUID
	createdAt  time.Time
	updatedAt  time.Time
	archivedAt *time.Time
}

func NewPerson(name, email, phone string) (Person, error) {
	if name == "" {
		return Person{}, errors.New("name is required")
	}

	if email == "" {
		return Person{}, errors.New("email is required")
	}

	return Person{
		id:        uuid.New(),
		name:      name,
		email:     email,
		phone:     phone,
		entryIDs:  []uuid.UUID{},
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func (p *Person) ID() uuid.UUID {
	return p.id
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) Email() string {
	return p.email
}

func (p *Person) Phone() string {
	return p.phone
}

func (p *Person) EntryIDs() []uuid.UUID {
	return p.entryIDs
}

func (p *Person) UpdateName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}
	p.name = name
	p.updatedAt = time.Now()
	return nil
}

func (p *Person) UpdateEmail(email string) error {
	if email == "" {
		return errors.New("email is required")
	}
	p.email = email
	p.updatedAt = time.Now()
	return nil
}

func (p *Person) UpdatePhone(phone string) {
	p.phone = phone
	p.updatedAt = time.Now()
}

func (p *Person) AddEntryID(entryID uuid.UUID) error {
	if entryID == uuid.Nil {
		return errors.New("entry ID cannot be nil")
	}
	p.entryIDs = append(p.entryIDs, entryID)
	p.updatedAt = time.Now()
	return nil
}

func (p *Person) RemoveEntryID(entryID uuid.UUID) error {
	for i, id := range p.entryIDs {
		if id == entryID {
			p.entryIDs = append(p.entryIDs[:i], p.entryIDs[i+1:]...)
			p.updatedAt = time.Now()
			return nil
		}
	}
	return errors.New("entry ID not found")
}

func (p *Person) SetEntryIDs(entryIDs []uuid.UUID) {
	p.entryIDs = entryIDs
	p.updatedAt = time.Now()
}
