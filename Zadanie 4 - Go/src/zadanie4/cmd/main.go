package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	productsController "zadanie4/internal/controllers/products"
	productsRepository "zadanie4/internal/repositories/products"
)

func main() {
	e := echo.New()
	apiGroup := e.Group("/api/v1")

	productsRepository := productsRepository.New()
	productsController := productsController.New(productsRepository)

	apiGroup.GET("/products", productsController.Get)
	apiGroup.GET("/products/:id", productsController.GetById)
	apiGroup.POST("/products", productsController.Create)
	apiGroup.PUT("/products/:id", productsController.Update)
	apiGroup.DELETE("/products/:id", productsController.Delete)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
