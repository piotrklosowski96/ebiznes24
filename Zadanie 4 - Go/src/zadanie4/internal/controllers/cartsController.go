package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"zadanie4/internal/models"
	repositoryErrors "zadanie4/internal/repositories/errors"

	"zadanie4/internal/repositories"
)

// Controller ...
type Controller struct {
	repository *repositories.CartsRepository
}

// NewCartsController ...
func NewCartsController(repository *repositories.CartsRepository) *Controller {
	return &Controller{repository: repository}
}

// RegisterRoutes ...
func (c *Controller) RegisterRoutes(endpointsGroup *echo.Group) {
	endpointsGroup.GET("/carts", c.GetAllCarts)
	endpointsGroup.GET("/carts/:cartId", c.GetCartById)
	endpointsGroup.POST("/carts", c.CreateCart)
	endpointsGroup.PUT("/carts/:cartId", c.UpdateCart)
	endpointsGroup.DELETE("/carts/:cartId", c.DeleteCart)
}

// CreateCart ...
func (c *Controller) CreateCart(ctx echo.Context) error {
	cartCreateRequest := &models.CartCreateRequest{}
	bindErr := ctx.Bind(cartCreateRequest)
	if bindErr != nil {
		return c.handleCreateCartError(ctx, bindErr)
	}

	createdCart, createCartErr := c.repository.CreateCart(cartCreateRequest)
	if createCartErr != nil {
		return c.handleCreateCartError(ctx, createCartErr)
	}

	return ctx.JSON(http.StatusOK, models.FromDatabaseCart(createdCart))
}

func (c *Controller) handleCreateCartError(ctx echo.Context, createCartErr error) error {
	var resourceNotFoundError *repositoryErrors.ResourceNotFoundError
	if errors.As(createCartErr, &resourceNotFoundError) {
		// INFO(Piotr Kłosowski): Some other, more robust error handling should be done
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error_message": fmt.Sprintf("cart (id: '%s') does not exist", ctx.Param("productId")),
		})
	}

	var foreignKeyConstraintViolated *repositoryErrors.ForeignKeyConstraintViolated
	if errors.As(createCartErr, &foreignKeyConstraintViolated) {
		// INFO(Piotr Kłosowski): Some other, more robust error handling should be done
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error_message": fmt.Sprintf("foreign key constraint failed"),
		})
	}

	return ctx.JSON(http.StatusInternalServerError, map[string]string{
		"error_message": "unknown error has occurred",
	})
}

// GetAllCarts ...
func (c *Controller) GetAllCarts(ctx echo.Context) error {
	cartsDB, getAllCartsErr := c.repository.GetAllCarts()
	if getAllCartsErr != nil {
		return c.handleGetAllCartsError(ctx, getAllCartsErr)
	}

	carts := make([]*models.CartResponse, len(cartsDB))
	for index, cart := range cartsDB {
		carts[index] = models.FromDatabaseCart(cart)
	}

	return ctx.JSON(http.StatusOK, carts)
}

func (c *Controller) handleGetAllCartsError(ctx echo.Context, _ error) error {
	return ctx.JSON(http.StatusInternalServerError, map[string]string{
		"error_message": "unknown error has occurred",
	})
}

// GetCartById ...
func (c *Controller) GetCartById(ctx echo.Context) error {
	cartId := ctx.Param("cartId")
	cartDB, getCartByIdErr := c.repository.GetCartById(cartId)
	if getCartByIdErr != nil {
		return c.handleGetCartByIdError(ctx, getCartByIdErr)
	}

	return ctx.JSON(http.StatusOK, models.FromDatabaseCart(cartDB))
}

func (c *Controller) handleGetCartByIdError(ctx echo.Context, getCartByIdErr error) error {
	var resourceNotFoundError *repositoryErrors.ResourceNotFoundError
	if errors.As(getCartByIdErr, &resourceNotFoundError) {
		// INFO(Piotr Kłosowski): Some other, more robust error handling should be done
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error_message": fmt.Sprintf("cart (id: '%s') does not exist", ctx.Param("cartId")),
		})
	}

	return ctx.JSON(http.StatusInternalServerError, map[string]string{
		"error_message": "unknown error has occurred",
	})
}

// UpdateCart ...
func (c *Controller) UpdateCart(ctx echo.Context) error {
	cartId := ctx.Param("cartId")
	cartUpdateRequest := &models.CartUpdateRequest{}
	bindErr := ctx.Bind(cartUpdateRequest)
	if bindErr != nil {
		return c.handleUpdateCartError(ctx, bindErr)
	}

	updatedCartDB, updateCartErr := c.repository.UpdateCart(cartId, cartUpdateRequest)
	if updateCartErr != nil {
		return c.handleUpdateCartError(ctx, updateCartErr)
	}

	return ctx.JSON(http.StatusOK, models.FromDatabaseCart(updatedCartDB))
}

func (c *Controller) handleUpdateCartError(ctx echo.Context, updateCartErr error) error {
	var resourceNotFoundError *repositoryErrors.ResourceNotFoundError
	if errors.As(updateCartErr, &resourceNotFoundError) {
		// INFO(Piotr Kłosowski): Some other, more robust error handling should be done
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error_message": fmt.Sprintf("cart (id: '%s') does not exist", ctx.Param("cartId")),
		})
	}

	var foreignKeyConstraintViolated *repositoryErrors.ForeignKeyConstraintViolated
	if errors.As(updateCartErr, &foreignKeyConstraintViolated) {
		// INFO(Piotr Kłosowski): Some other, more robust error handling should be done
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error_message": fmt.Sprintf("foreign key constraint failed"),
		})
	}

	return ctx.JSON(http.StatusInternalServerError, map[string]string{
		"error_message": "unknown error has occurred",
	})
}

// DeleteCart ...
func (c *Controller) DeleteCart(ctx echo.Context) error {
	cartId := ctx.Param("cartId")
	deleteProductErr := c.repository.DeleteCart(cartId)
	if deleteProductErr != nil {
		return c.handleDeleteCartError(ctx, deleteProductErr)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// NOTE(Piotr Kłosowski): When more errors are handles change '_' into some meaningful variable name
func (c *Controller) handleDeleteCartError(ctx echo.Context, _ error) error {
	return ctx.JSON(http.StatusInternalServerError, map[string]string{
		"error_message": "unknown error has occurred",
	})
}
