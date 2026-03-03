package handler

import (
	"context"
	"errors"

	api "github.com/fun-dotto/faculty-api/generated"
	"github.com/fun-dotto/faculty-api/internal/domain"
)

func (h *FacultyHandler) FacultiesV1Delete(ctx context.Context, request api.FacultiesV1DeleteRequestObject) (api.FacultiesV1DeleteResponseObject, error) {
	err := h.facultyService.Delete(ctx, request.Id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, err
		}
		return nil, err
	}

	return api.FacultiesV1Delete204Response{}, nil
}
