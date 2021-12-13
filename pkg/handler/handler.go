package handler

import (
	"class-app/pkg/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(e *echo.Echo) {
	auth := e.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := e.Group("/api", middleware.JWTWithConfig(config))
	api.GET("", h.userIdentity)
	{
		students := api.Group("/students")
		{
			students.POST("/", h.create)
			students.GET("/", h.getAll)
			students.DELETE("/", h.deleteAll)
			students.GET("/:id", h.getById)
			students.DELETE("/:id", h.deleteById)
		}
	}
}
