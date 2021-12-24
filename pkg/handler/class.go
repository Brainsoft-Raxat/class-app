package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// getAllClasses godoc
// @Summary Get All Classes.
// @Security ApiKeyAuth
// @Tags classes
// @Description get all classes.
// @ID get-all-classes
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Class
// @Failure 500 {object} errorResponse
// @Router /api/classes/get-all-classes [get]
func (h *Handler) getAllClasses(c echo.Context) error {
	classes, err := h.services.Class.GetAllClasses()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse{Status: err.Error()})
	}

	return c.JSON(http.StatusOK, classes)
}

// listClassStudents godoc
// @Summary List Class Students By Class ID.
// @Security ApiKeyAuth
// @Tags classes
// @Description list class students by class id.
// @ID list-class-students
// @Accept */*
// @Produce json
// @Param id path integer true "class id"
// @Success 200 {object} []models.Student
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/classes/{id}/students [get]
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
