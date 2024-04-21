package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"zadanie4/internal/controllers"
	"zadanie4/internal/repositories"
)

var (
	connectionString, _ = os.LookupEnv("DATABASE_CONNECTION_STRING")
)

func main() {
	e := echo.New()
	apiGroup := e.Group("/api/v1")

	databaseHandle, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	productsRepository := repositories.NewProductsRepository(databaseHandle)
	productsController := controllers.NewProductsController(productsRepository)
	productsController.RegisterRoutes(apiGroup)

	e.Logger.Fatal(e.Start(":8080"))
}
