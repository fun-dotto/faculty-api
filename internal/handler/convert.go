package handler

import (
	api "github.com/fun-dotto/faculty-api/generated"
	"github.com/fun-dotto/faculty-api/internal/domain"
)

func toApiFaculty(faculty domain.Faculty) api.Faculty {
	return api.Faculty{
		Id:    faculty.ID,
		Name:  faculty.Name,
		Email: faculty.Email,
	}
}

func toDomainFacultyFromRequest(id string, req api.FacultyRequest) domain.Faculty {
	return domain.Faculty{
		ID:    id,
		Name:  req.Name,
		Email: req.Email,
	}
}
