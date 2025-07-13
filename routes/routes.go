package routes


import (
	"github.com/labstack/echo/v4"
	"github.com/yourusername/nextjs-echo-crud-app/handlers"
	"github.com/yourusername/nextjs-echo-crud-app/middleware"
)

func SetupRoutes(e *echo.Echo, h *handlers.Handler){
	api := e.Group("/api")
	
	api.POST("/user", h.CreateUser)
	api.POST("/login", h.Login)



	protected := api.Group("", middleware.JWTAuth(h.Config))
	protected.GET("/users", h.GetAllUsers)
	protected.GET("/user/:id", h.GetUser)
	protected.PATCH("/user/:id", h.UpdateUser)
	protected.DELETE("/user/:id", h.DeleteUser)
	



}