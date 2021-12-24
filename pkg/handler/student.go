package handler

import (
	"class-app/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// getAllStudents godoc
// @Summary Get All Students.
// @Security ApiKeyAuth
// @Tags students
// @Description get all students.
// @ID get-all-students
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Student
// @Failure 500 {object} errorResponse
// @Router /api/students/get-all-students [get]
func (h *Handler) getAllStudents(c echo.Context) error {

	students, err := h.services.Student.GetAllStudents()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse{Status: err.Error()})
	}

	return c.JSON(http.StatusOK, students)
}

// getStudentById godoc
// @Summary Get Student By ID.
// @Security ApiKeyAuth
// @Tags students
// @Description get student by id.
// @ID get-student-by-id
// @Accept */*
// @Produce json
// @Param id path integer true "student id"
// @Success 200 {object} models.Student
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/students/{id} [get]
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

// createStudent godoc
// @Summary Create Student.
// @Security ApiKeyAuth
// @Tags students
// @Description create student.
// @ID create-student
// @Accept json
// @Produce json
// @Param input body models.Student true "student id"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/students/create-student [post]
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

// deleteAllStudents godoc
// @Summary Delete All Students.
// @Security ApiKeyAuth
// @Tags students
// @Description delete all students.
// @ID delete-all-students
// @Accept */*
// @Produce json
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/students/delete-all-students [delete]
func (h *Handler) deleteAllStudents(c echo.Context) error {
	if err := h.services.Student.DeleteAllStudents(); err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse{Status: err.Error()})
	}

	return c.JSON(http.StatusOK, statusResponse{
		"ok",
	})
}

// deleteStudentById godoc
// @Summary Delete Student By ID.
// @Security ApiKeyAuth
// @Tags students
// @Description delete student by id.
// @ID delete-student-by-id
// @Accept */*
// @Produce json
// @Param id path integer true "student id"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/students/{id} [delete]
func (h *Handler) deleteStudentById(c echo.Context) error {
	id := c.Param("id")
	studentId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse{Status: err.Error()})
	}

	if err := h.services.DeleteStudentById(studentId); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse{Status: err.Error()})
	}

	return c.JSON(http.StatusOK, statusResponse{
		"ok",
	})
}
