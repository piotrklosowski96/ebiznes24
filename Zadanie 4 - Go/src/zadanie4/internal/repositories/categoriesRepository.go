package repositories

import (
	"gorm.io/gorm"
	"zadanie4/internal/models"
	"zadanie4/internal/repositories/errors"
	repositoryModels "zadanie4/internal/repositories/models"
)

// CategoriesRepository ...
type CategoriesRepository struct {
	databaseHandle *gorm.DB
}

// NewCategoriesRepository ...
func NewCategoriesRepository(databaseHandle *gorm.DB) *CategoriesRepository {
	autoMigrateErr := databaseHandle.AutoMigrate(
		&repositoryModels.Category{},
		&repositoryModels.Product{},
	)
	if autoMigrateErr != nil {
		panic(autoMigrateErr.Error())
	}

	return &CategoriesRepository{
		databaseHandle: databaseHandle,
	}
}

// CreateCategory ...
func (r *CategoriesRepository) CreateCategory(categoryCreateRequest *models.CategoryCreateRequest) (*repositoryModels.Category, error) {
	category := &repositoryModels.Category{
		Name:        categoryCreateRequest.Name,
		Description: categoryCreateRequest.Description,
	}

	createErr := r.databaseHandle.Create(category).Error
	if createErr != nil {
		return nil, errors.HandleDatabaseError(createErr)
	}

	return category, nil
}

// GetAllCategories ...
func (r *CategoriesRepository) GetAllCategories() ([]*repositoryModels.Category, error) {
	var categories []*repositoryModels.Category

	findErr := r.databaseHandle.Find(&categories).Error
	if findErr != nil {
		return nil, errors.HandleDatabaseError(findErr)
	}

	return categories, nil
}

// GetCategoryById ...
func (r *CategoriesRepository) GetCategoryById(categoryId string) (*repositoryModels.Category, error) {
	var category repositoryModels.Category

	firstErr := r.databaseHandle.First(&category, "id = ?", categoryId).Error
	if firstErr != nil {
		return nil, errors.HandleDatabaseError(firstErr)
	}

	return &category, nil
}

// UpdateCategory ...
func (r *CategoriesRepository) UpdateCategory(categoryId string, cartUpdateRequest *models.CategoryUpdateRequest) (*repositoryModels.Category, error) {
	var updateCategory repositoryModels.Category
	if cartUpdateRequest.Name != nil {
		updateCategory.Name = *cartUpdateRequest.Name
	}

	if cartUpdateRequest.Description != nil {
		updateCategory.Description = cartUpdateRequest.Description
	}

	updatesErr := r.databaseHandle.Where("id = ?", categoryId).Updates(updateCategory).Error
	if updatesErr != nil {
		return nil, errors.HandleDatabaseError(updatesErr)
	}

	var category repositoryModels.Category
	firstErr := r.databaseHandle.First(&category, "id = ?", categoryId).Error
	if firstErr != nil {
		return nil, errors.HandleDatabaseError(firstErr)
	}

	return &category, nil
}

// DeleteCategory ...
func (r *CategoriesRepository) DeleteCategory(categoryId string) error {
	deleteErr := r.databaseHandle.Delete(&repositoryModels.Category{}, "id = ?", categoryId).Error
	if deleteErr != nil {
		return errors.HandleDatabaseError(deleteErr)
	}

	return nil
}
