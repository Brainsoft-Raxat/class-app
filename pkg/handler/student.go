package handler

import (
	"class-app/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) getAllStudents(c echo.Context) error {

	students, err := h.services.Student.GetAllStudents()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse{Status: err.Error()})
	}

	return c.JSON(http.StatusOK, students)
}

func (h *Handler) getStudentById(c echo.Context) error {
	id := c.Param("id")
	studentId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse{Status: err.Error()})
	}
	student, err := h.services.Student.GetStudentById(studentId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse{Status: err.Error()})
	}
	return c.JSON(http.StatusOK, student)
}

func (h *Handler) createStudent(c echo.Context) error {
	var student models.Student
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse{Status: err.Error()})
	}

	id, err := h.services.Student.CreateStudent(student)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse{Status: err.Error()})
	}

	return c.JSON(http.StatusOK, id)
}

func (h *Handler) deleteAllStudents(c echo.Context) error {
	if err := h.services.Student.DeleteAllStudents(); err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse{Status: err.Error()})
	}

	return c.JSON(http.StatusOK, statusResponse{
		"ok",
	})
}

func (h *Handler) deleteStudentById(c echo.Context) error {
	id := c.Param("id")
	studentId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse{Status: err.Error()})
	}

	if err := h.services.DeleteStudentById(studentId); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse{
			Status: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, statusResponse{
		"ok",
	})
}
