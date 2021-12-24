package handler

import (
	"class-app/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// signUp godoc
// @Summary SignUp
// @Tags auth
// @Description create account for teacher.
// @ID sign-up
// @Accept json
// @Produce json
// @Param input body models.Teacher true "account info"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /auth/sign-up [post]
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

// signIn godoc
// @Summary SignIn
// @Tags auth
// @Description Sign in as a teacher.
// @ID sign-in
// @Accept json
// @Produce json
// @Param input body signInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /auth/sign-in [post]
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
