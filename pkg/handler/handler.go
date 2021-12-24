package handler

import (
	_ "class-app/docs"
	"class-app/pkg/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
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
			students.POST("/create-student", h.createStudent)
			students.GET("/get-all-students", h.getAllStudents)
			students.DELETE("/delete-all-students", h.deleteAllStudents)
			students.GET("/:id", h.getStudentById)
			students.DELETE("/:id", h.deleteStudentById)
		}

		classes := api.Group("/classes")
		{
			classes.GET("/get-all-classes", h.getAllClasses)
			class := classes.Group("/")
			{
				class.GET(":id/students", h.listClassStudents)
			}
		}
	}
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
