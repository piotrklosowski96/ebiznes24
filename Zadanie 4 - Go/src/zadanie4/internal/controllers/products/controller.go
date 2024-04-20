package products

import (
	"net/http"

	"github.com/labstack/echo/v4"

	productsRepository "zadanie4/internal/repositories/products"
)

// Controller ...
type Controller struct {
	repository *productsRepository.Repository
}

// New ...
func New(repository *productsRepository.Repository) *Controller {
	return &Controller{repository: repository}
}

// Get ...
func (c *Controller) Get(ctx echo.Context) error {
	products := c.repository.GetAll()

	return ctx.JSON(http.StatusOK, products)
}

// GetById ...
func (c *Controller) GetById(ctx echo.Context) error {
	productId := ctx.Param("id")
	product := c.repository.GetById(productId)

	return ctx.JSON(http.StatusOK, product)
}

// Create ...
func (c *Controller) Create(ctx echo.Context) error {
	product := &productsRepository.Product{}
	_ = ctx.Bind(product)

	c.repository.Add(product)

	return ctx.JSON(http.StatusOK, product)
}

// Delete ...
func (c *Controller) Delete(ctx echo.Context) error {
	productId := ctx.Param("id")
	c.repository.Remove(productId)

	return ctx.NoContent(http.StatusNoContent)
}

// Update ...
func (c *Controller) Update(ctx echo.Context) error {
	productId := ctx.Param("id")
	product := &productsRepository.Product{}
	_ = ctx.Bind(product)

	updatedProduct := c.repository.Update(productId, product)
	return ctx.JSON(http.StatusOK, updatedProduct)
}
