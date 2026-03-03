package handler

import (
	"github.com/fun-dotto/faculty-api/internal/service"
)

type FacultyHandler struct {
	facultyService *service.FacultyService
}

func NewFacultyHandler(facultyService *service.FacultyService) *FacultyHandler {
	return &FacultyHandler{facultyService: facultyService}
}
