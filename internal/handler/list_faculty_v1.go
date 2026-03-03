package handler

import (
	"context"

	api "github.com/fun-dotto/faculty-api/generated"
)

func (h *FacultyHandler) FacultiesV1List(ctx context.Context, request api.FacultiesV1ListRequestObject) (api.FacultiesV1ListResponseObject, error) {
	var ids []string
	if request.Params.Ids != nil {
		ids = *request.Params.Ids
	}

	faculties, err := h.facultyService.List(ctx, ids)
	if err != nil {
		return nil, err
	}

	apiFaculties := make([]api.Faculty, len(faculties))
	for i, faculty := range faculties {
		apiFaculties[i] = toApiFaculty(faculty)
	}

	return api.FacultiesV1List200JSONResponse{Faculties: apiFaculties}, nil
}
