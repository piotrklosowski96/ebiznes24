package main

import (
	"github.com/labstack/echo/v4"

	"zadanie4/internal/controllers"
	"zadanie4/internal/repositories"
)

func main() {
	e := echo.New()
	apiGroup := e.Group("/api/v1")

	productsRepository := repositories.NewProductsRepository()
	productsController := controllers.NewProductsController(productsRepository)
	productsController.RegisterRoutes(apiGroup)

	e.Logger.Fatal(e.Start(":8080"))
}
