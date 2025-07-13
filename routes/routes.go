package routes

import (
	"github.com/Artur2912/pet1/handlers"
	"github.com/Artur2912/pet1/middleware"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, h *handlers.Handler) {
	api := e.Group("/api")

	api.POST("/user", h.CreateUser)
	api.POST("/login", h.Login)

	protected := api.Group("", middleware.JWTAuth(h.Config))
	protected.GET("/users", h.GetAllUsers)
	protected.GET("/user/:id", h.GetUser)
	protected.PATCH("/user/:id", h.UpdateUser)
	protected.DELETE("/user/:id", h.DeleteUser)

}
