package handler

import (
	"context"

	api "github.com/fun-dotto/faculty-api/generated"
	"github.com/fun-dotto/faculty-api/internal/domain"
)

func (h *FacultyHandler) FacultiesV1Create(ctx context.Context, request api.FacultiesV1CreateRequestObject) (api.FacultiesV1CreateResponseObject, error) {
	faculty := domain.Faculty{
		Name:  request.Body.Name,
		Email: request.Body.Email,
	}

	created, err := h.facultyService.Create(ctx, faculty)
	if err != nil {
		return nil, err
	}

	return api.FacultiesV1Create201JSONResponse{Faculty: toApiFaculty(created)}, nil
}
