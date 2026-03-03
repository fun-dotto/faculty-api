package database

import (
	"time"

	"github.com/fun-dotto/faculty-api/internal/domain"
)

type Faculty struct {
	ID    string `gorm:"primaryKey;type:uuid"`
	Name  string `gorm:"not null"`
	Email string `gorm:"not null;"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (m *Faculty) ToDomain() domain.Faculty {
	return domain.Faculty{
		ID:    m.ID,
		Name:  m.Name,
		Email: m.Email,
	}
}

func FromDomain(faculty domain.Faculty) Faculty {
	return Faculty{
		ID:    faculty.ID,
		Name:  faculty.Name,
		Email: faculty.Email,
	}
}
