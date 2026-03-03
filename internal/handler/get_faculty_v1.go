package handler

import (
	"context"
	"errors"

	api "github.com/fun-dotto/faculty-api/generated"
	"github.com/fun-dotto/faculty-api/internal/domain"
)

func (h *FacultyHandler) FacultiesV1Detail(ctx context.Context, request api.FacultiesV1DetailRequestObject) (api.FacultiesV1DetailResponseObject, error) {
	faculty, err := h.facultyService.GetFacultyByID(ctx, request.Id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, err
		}
		return nil, err
	}
	return api.FacultiesV1Detail200JSONResponse{Faculty: toApiFaculty(faculty)}, nil
}
