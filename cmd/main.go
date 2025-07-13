package main

import (
	"fmt"
	"log"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"github.com/yourusername/nextjs-echo-crud-app/config"
	"github.com/yourusername/nextjs-echo-crud-app/handlers"
	"github.com/yourusername/nextjs-echo-crud-app/routes"
)

func main(){
	cfg, err := config.LoadConfig()
	if err != nil{
		log.Fatalf("Error; %v", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())  
    e.Use(middleware.Recover())
	h := handlers.NewHandler(cfg)
	routes.SetupRoutes(e, h)
	port := cfg.AppConfig.Port
	if port == ""{
		port = "8080"
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s",port)))
}