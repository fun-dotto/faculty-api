package handler

import (
	"context"
	"errors"

	api "github.com/fun-dotto/faculty-api/generated"
	"github.com/fun-dotto/faculty-api/internal/domain"
)

func (h *FacultyHandler) FacultiesV1Update(ctx context.Context, request api.FacultiesV1UpdateRequestObject) (api.FacultiesV1UpdateResponseObject, error) {
	faculty := domain.Faculty{
		ID:    request.Id,
		Name:  request.Body.Name,
		Email: request.Body.Email,
	}

	updated, err := h.facultyService.Update(ctx, faculty)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, err
		}
		return nil, err
	}

	return api.FacultiesV1Update200JSONResponse{Faculty: toApiFaculty(updated)}, nil
}
