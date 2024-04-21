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

// CategoriesController ...
type CategoriesController struct {
	repository *repositories.CategoriesRepository
}

// NewCategoriesController ...
func NewCategoriesController(repository *repositories.CategoriesRepository) *CategoriesController {
	return &CategoriesController{
		repository: repository,
	}
}

// RegisterRoutes ...
func (c *CategoriesController) RegisterRoutes(endpointsGroup *echo.Group) {
	endpointsGroup.POST("/categories", c.CreateCategory)
	endpointsGroup.GET("/categories", c.GetAllCategories)
	endpointsGroup.GET("/categories/:categoryId", c.GetCategoryById)
	endpointsGroup.PUT("/categories/:categoryId", c.UpdateCategory)
	endpointsGroup.DELETE("/categories/:categoryId", c.DeleteCategory)
}

// CreateCategory ...
func (c *CategoriesController) CreateCategory(ctx echo.Context) error {
	categoryCreateRequest := &models.CategoryCreateRequest{}
	bindErr := ctx.Bind(categoryCreateRequest)
	if bindErr != nil {
		return c.handleCreateCategoryError(ctx, bindErr)
	}

	createdCategory, createCategoryErr := c.repository.CreateCategory(categoryCreateRequest)
	if createCategoryErr != nil {
		return c.handleCreateCategoryError(ctx, createCategoryErr)
	}

	return ctx.JSON(http.StatusOK, models.FromDatabaseCategory(createdCategory))
}

func (c *CategoriesController) handleCreateCategoryError(ctx echo.Context, createCategoryErr error) error {
	var resourceNotFoundError *repositoryErrors.ResourceNotFoundError
	if errors.As(createCategoryErr, &resourceNotFoundError) {
		// INFO(Piotr Kłosowski): Some other, more robust error handling should be done
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error_message": fmt.Sprintf("product (id: '%s') does not exist", ctx.Param("productId")),
		})
	}

	return ctx.JSON(http.StatusInternalServerError, map[string]string{
		"error_message": "unknown error has occurred",
	})
}

// GetAllCategories ...
func (c *CategoriesController) GetAllCategories(ctx echo.Context) error {
	categoriesDB, getAllCategoriesErr := c.repository.GetAllCategories()
	if getAllCategoriesErr != nil {
		return c.handleGetAllCategoriesError(ctx, getAllCategoriesErr)
	}

	categories := make([]*models.CategoryResponse, len(categoriesDB))
	for index, category := range categoriesDB {
		categories[index] = models.FromDatabaseCategory(category)
	}

	return ctx.JSON(http.StatusOK, categories)
}

// NOTE(Piotr Kłosowski): When more errors are handled change '_' into some meaningful variable name
func (c *CategoriesController) handleGetAllCategoriesError(ctx echo.Context, _ error) error {
	return ctx.JSON(http.StatusInternalServerError, map[string]string{
		"error_message": "unknown error has occurred",
	})
}

// GetCategoryById ...
func (c *CategoriesController) GetCategoryById(ctx echo.Context) error {
	categoryId := ctx.Param("categoryId")
	categoryDB, getCategoryByIdErr := c.repository.GetCategoryById(categoryId)
	if getCategoryByIdErr != nil {
		return c.handleGetCategoryByIdError(ctx, getCategoryByIdErr)
	}

	return ctx.JSON(http.StatusOK, models.FromDatabaseCategory(categoryDB))
}

func (c *CategoriesController) handleGetCategoryByIdError(ctx echo.Context, getCategoryByIdErr error) error {
	var resourceNotFoundError *repositoryErrors.ResourceNotFoundError
	if errors.As(getCategoryByIdErr, &resourceNotFoundError) {
		// INFO(Piotr Kłosowski): Some other, more robust error handling should be done
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error_message": fmt.Sprintf("category (id: '%s') does not exist", ctx.Param("categoryId")),
		})
	}

	return ctx.JSON(http.StatusInternalServerError, map[string]string{
		"error_message": "unknown error has occurred",
	})
}

// UpdateCategory ...
func (c *CategoriesController) UpdateCategory(ctx echo.Context) error {
	categoryId := ctx.Param("categoryId")
	categoryUpdateRequest := &models.CategoryUpdateRequest{}
	bindErr := ctx.Bind(categoryUpdateRequest)
	if bindErr != nil {
		return c.handleUpdateCategoryError(ctx, bindErr)
	}

	updatedCategoryDB, updateCategoryErr := c.repository.UpdateCategory(categoryId, categoryUpdateRequest)
	if updateCategoryErr != nil {
		return c.handleUpdateCategoryError(ctx, updateCategoryErr)
	}

	return ctx.JSON(http.StatusOK, models.FromDatabaseCategory(updatedCategoryDB))
}

func (c *CategoriesController) handleUpdateCategoryError(ctx echo.Context, updateCategoryErr error) error {
	var resourceNotFoundError *repositoryErrors.ResourceNotFoundError
	if errors.As(updateCategoryErr, &resourceNotFoundError) {
		// INFO(Piotr Kłosowski): Some other, more robust error handling should be done
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error_message": fmt.Sprintf("category (id: '%s') does not exist", ctx.Param("categoryId")),
		})
	}

	return ctx.JSON(http.StatusInternalServerError, map[string]string{
		"error_message": "unknown error has occurred",
	})
}

// DeleteCategory ...
func (c *CategoriesController) DeleteCategory(ctx echo.Context) error {
	categoryId := ctx.Param("categoryId")
	deleteCategoryErr := c.repository.DeleteCategory(categoryId)
	if deleteCategoryErr != nil {
		return c.handleDeleteCategoryError(ctx, deleteCategoryErr)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// NOTE(Piotr Kłosowski): When more errors are handled change '_' into some meaningful variable name
func (c *CategoriesController) handleDeleteCategoryError(ctx echo.Context, _ error) error {
	return ctx.JSON(http.StatusInternalServerError, map[string]string{
		"error_message": "unknown error has occurred",
	})
}
