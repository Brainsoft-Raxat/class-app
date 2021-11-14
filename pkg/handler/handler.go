package handler

import (
	"class-app/pkg/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(e *echo.Echo) {
	students := e.Group("/students")
	{
		students.POST("/", h.create)
		students.GET("/", h.getAll)
		students.DELETE("/", h.deleteAll)
		students.GET("/:id", h.getById)
		students.DELETE("/:id", h.deleteById)
	}
}
