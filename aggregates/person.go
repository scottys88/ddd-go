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
	createdAt  time.Time
	updatedAt  time.Time
	archivedAt *time.Time
}

func NewPerson(name, email, phone string) (*Person, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	if email == "" {
		return nil, errors.New("email is required")
	}

	return &Person{
		id:        uuid.New(),
		name:      name,
		email:     email,
		phone:     phone,
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

func (p *Person) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Person) UpdatedAt() time.Time {
	return p.updatedAt
}

func (p *Person) ArchivedAt() *time.Time {
	return p.archivedAt
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

func (p *Person) Archive() {
	now := time.Now()
	p.archivedAt = &now
	p.updatedAt = now
}
