package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) getAllClasses(c echo.Context) error {
	classes, err := h.services.Class.GetAllClasses()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse{Status: err.Error()})
	}

	return c.JSON(http.StatusOK, classes)
}

func (h *Handler) listClassStudents(c echo.Context) error {
	idParam := c.Param("id")
	classId, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse{Status: "missing or invalid id parameter in url"})
	}

	students, err := h.services.Class.ListClassStudents(classId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse{Status: err.Error()})
	}

	return c.JSON(http.StatusOK, students)
}
