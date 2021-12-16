package handler

import (
	"class-app/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) signUp(c echo.Context) error {
	var input models.Teacher
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse{Status: err.Error()})
	}

	id, err := h.services.Authorization.CreateTeacher(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse{Status: err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c echo.Context) error {
	var input signInInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse{Status: err.Error()})
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse{Status: err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
