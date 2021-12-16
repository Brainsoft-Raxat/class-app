package handler

import (
	"class-app/configs"
	"class-app/pkg/service"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

const (
	teacherCtx = "teacherId"
)

var (
	config = middleware.JWTConfig{
		Claims:     &service.TokenClaims{},
		SigningKey: []byte(configs.SigningKey),
	}
)

func (h *Handler) userIdentity(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*service.TokenClaims)
	TeacherId := claims.TeacherId
	c.Set(teacherCtx, TeacherId)
	return c.JSON(http.StatusOK, "Allowed")
}
