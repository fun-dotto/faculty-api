package repository

import (
	"context"
	"errors"

	"github.com/fun-dotto/faculty-api/internal/database"
	"github.com/fun-dotto/faculty-api/internal/domain"
	"gorm.io/gorm"
)

type facultyRepository struct {
	db *gorm.DB
}

func NewFacultyRepository(db *gorm.DB) *facultyRepository {
	return &facultyRepository{db: db}
}

func (r *facultyRepository) List(ctx context.Context, ids []string) ([]domain.Faculty, error) {
	var dbFaculties []database.Faculty
	query := r.db.WithContext(ctx)
	if len(ids) > 0 {
		query = query.Where("id IN ?", ids)
	}
	if err := query.Find(&dbFaculties).Error; err != nil {
		return nil, err
	}

	domainFaculties := make([]domain.Faculty, len(dbFaculties))
	for i, dbFaculty := range dbFaculties {
		domainFaculties[i] = dbFaculty.ToDomain()
	}

	return domainFaculties, nil
}

func (r *facultyRepository) GetByID(ctx context.Context, id string) (domain.Faculty, error) {
	var dbFaculty database.Faculty
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&dbFaculty).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Faculty{}, domain.ErrNotFound
		}
		return domain.Faculty{}, err
	}
	return dbFaculty.ToDomain(), nil
}

func (r *facultyRepository) Create(ctx context.Context, faculty domain.Faculty) (domain.Faculty, error) {
	dbFaculty := database.FromDomain(faculty)
	if err := r.db.WithContext(ctx).Create(&dbFaculty).Error; err != nil {
		return domain.Faculty{}, err
	}
	return dbFaculty.ToDomain(), nil
}

func (r *facultyRepository) Update(ctx context.Context, faculty domain.Faculty) (domain.Faculty, error) {
	dbFaculty := database.FromDomain(faculty)
	result := r.db.WithContext(ctx).Model(&database.Faculty{}).Where("id = ?", faculty.ID).Updates(map[string]interface{}{
		"name":  dbFaculty.Name,
		"email": dbFaculty.Email,
	})
	if result.Error != nil {
		return domain.Faculty{}, result.Error
	}
	if result.RowsAffected == 0 {
		return domain.Faculty{}, domain.ErrNotFound
	}
	return r.GetByID(ctx, faculty.ID)
}

func (r *facultyRepository) Delete(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&database.Faculty{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return nil
}
