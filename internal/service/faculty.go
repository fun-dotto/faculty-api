package service

import (
	"context"

	"github.com/fun-dotto/faculty-api/internal/domain"
)

type FacultyRepository interface {
	List(ctx context.Context, ids []string) ([]domain.Faculty, error)
	GetByID(ctx context.Context, id string) (domain.Faculty, error)
	Create(ctx context.Context, faculty domain.Faculty) (domain.Faculty, error)
	Update(ctx context.Context, faculty domain.Faculty) (domain.Faculty, error)
	Delete(ctx context.Context, id string) error
}

type FacultyService struct {
	facultyRepository FacultyRepository
}

func NewFacultyService(facultyRepository FacultyRepository) *FacultyService {
	return &FacultyService{facultyRepository: facultyRepository}
}

func (s *FacultyService) ListFaculties(ctx context.Context, ids []string) ([]domain.Faculty, error) {
	return s.facultyRepository.List(ctx, ids)
}

func (s *FacultyService) GetFacultyByID(ctx context.Context, id string) (domain.Faculty, error) {
	return s.facultyRepository.GetByID(ctx, id)
}

func (s *FacultyService) CreateFaculty(ctx context.Context, faculty domain.Faculty) (domain.Faculty, error) {
	return s.facultyRepository.Create(ctx, faculty)
}

func (s *FacultyService) UpdateFaculty(ctx context.Context, faculty domain.Faculty) (domain.Faculty, error) {
	return s.facultyRepository.Update(ctx, faculty)
}

func (s *FacultyService) DeleteFaculty(ctx context.Context, id string) error {
	return s.facultyRepository.Delete(ctx, id)
}
