package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"zadanie4/internal/repositories"
	repositoryErrors "zadanie4/internal/repositories/errors"
)

// ProductsController ...
type ProductsController struct {
	repository *repositories.ProductsRepository
}

// NewProductsController ...
func NewProductsController(repository *repositories.ProductsRepository) *ProductsController {
	return &ProductsController{repository: repository}
}

// RegisterRoutes ...
func (c *ProductsController) RegisterRoutes(endpointsGroup *echo.Group) {
	endpointsGroup.POST("/products", c.CreateProduct)
	endpointsGroup.GET("/products", c.GetAllProducts)
	endpointsGroup.GET("/products/:productId", c.GetProductById)
	endpointsGroup.PUT("/products/:productId", c.UpdateProduct)
	endpointsGroup.DELETE("/products/:productId", c.DeleteProduct)
}

// CreateProduct ...
func (c *ProductsController) CreateProduct(ctx echo.Context) error {
	product := &repositories.Product{}
	_ = ctx.Bind(product)

	newProduct, createProductErr := c.repository.CreateProduct(product)
	if createProductErr != nil {
		return c.handleCreateProductError(ctx, createProductErr)
	}

	return ctx.JSON(http.StatusOK, newProduct)
}

func (c *ProductsController) handleCreateProductError(ctx echo.Context, createProductErr error) error {
	var resourceNotFoundError *repositoryErrors.ResourceNotFoundError
	if errors.As(createProductErr, &resourceNotFoundError) {
		// INFO(Piotr Kłosowski): Some other, more robust error handling should be done
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error_message": fmt.Sprintf("product (id: '%s') does not exist", resourceNotFoundError.ResourceID),
		})
	}

	return ctx.JSON(http.StatusInternalServerError, map[string]string{
		"error_message": "unknown error has occurred",
	})
}

// GetAllProducts ...
func (c *ProductsController) GetAllProducts(ctx echo.Context) error {
	products, getAllProductsErr := c.repository.GetAllProducts()
	if getAllProductsErr != nil {
		return c.handleGetAllProductsError(ctx, getAllProductsErr)
	}

	return ctx.JSON(http.StatusOK, products)
}

// NOTE(Piotr Kłosowski): When more errors are handles change '_' into some meaningful variable name
func (c *ProductsController) handleGetAllProductsError(ctx echo.Context, _ error) error {
	return ctx.JSON(http.StatusInternalServerError, map[string]string{
		"error_message": "unknown error has occurred",
	})
}

// GetProductById ...
func (c *ProductsController) GetProductById(ctx echo.Context) error {
	productId := ctx.Param("productId")
	product, getProductByIdErr := c.repository.GetProductById(productId)
	if getProductByIdErr != nil {
		return c.handleGetProductByIdError(ctx, getProductByIdErr)
	}

	return ctx.JSON(http.StatusOK, product)
}

func (c *ProductsController) handleGetProductByIdError(ctx echo.Context, getProductByIdErr error) error {
	var resourceNotFoundError *repositoryErrors.ResourceNotFoundError
	if errors.As(getProductByIdErr, &resourceNotFoundError) {
		// INFO(Piotr Kłosowski): Some other, more robust error handling should be done
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error_message": fmt.Sprintf("product (id: '%s') does not exist", resourceNotFoundError.ResourceID),
		})
	}

	return ctx.JSON(http.StatusInternalServerError, map[string]string{
		"error_message": "unknown error has occurred",
	})
}

// UpdateProduct ...
func (c *ProductsController) UpdateProduct(ctx echo.Context) error {
	productId := ctx.Param("productId")
	product := &repositories.Product{}
	_ = ctx.Bind(product)

	updatedProduct, updateProductId := c.repository.UpdateProduct(productId, product)
	if updateProductId != nil {
		return c.handleUpdateProductError(ctx, updateProductId)
	}

	return ctx.JSON(http.StatusOK, updatedProduct)
}

func (c *ProductsController) handleUpdateProductError(ctx echo.Context, getProductByIdErr error) error {
	var resourceNotFoundError *repositoryErrors.ResourceNotFoundError
	if errors.As(getProductByIdErr, &resourceNotFoundError) {
		// INFO(Piotr Kłosowski): Some other, more robust error handling should be done
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error_message": fmt.Sprintf("product (id: '%s') does not exist", resourceNotFoundError.ResourceID),
		})
	}

	return ctx.JSON(http.StatusInternalServerError, map[string]string{
		"error_message": "unknown error has occurred",
	})
}

// DeleteProduct ...
func (c *ProductsController) DeleteProduct(ctx echo.Context) error {
	productId := ctx.Param("productId")
	deleteProductErr := c.repository.DeleteProduct(productId)
	if deleteProductErr != nil {
		return c.handleDeleteProductError(ctx, deleteProductErr)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// NOTE(Piotr Kłosowski): When more errors are handles change '_' into some meaningful variable name
func (c *ProductsController) handleDeleteProductError(ctx echo.Context, _ error) error {
	return ctx.JSON(http.StatusInternalServerError, map[string]string{
		"error_message": "unknown error has occurred",
	})
}
