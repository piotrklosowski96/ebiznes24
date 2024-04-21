package repositories

import (
	"gorm.io/gorm"

	"zadanie4/internal/models"
	"zadanie4/internal/repositories/errors"
	repositoryModels "zadanie4/internal/repositories/models"
	"zadanie4/internal/repositories/scopes"
)

// ProductsRepository ...
type ProductsRepository struct {
	databaseHandle *gorm.DB
}

// NewProductsRepository ...
func NewProductsRepository(databaseHandle *gorm.DB) *ProductsRepository {
	autoMigrateErr := databaseHandle.AutoMigrate(
		&repositoryModels.Product{},
		&repositoryModels.Category{},
	)
	if autoMigrateErr != nil {
		panic(autoMigrateErr.Error())
	}

	return &ProductsRepository{
		databaseHandle: databaseHandle,
	}
}

// CreateProduct ...
func (r *ProductsRepository) CreateProduct(productCreateRequest *models.ProductCreateRequest) (*repositoryModels.Product, error) {
	product := &repositoryModels.Product{
		Name:        productCreateRequest.Name,
		Description: productCreateRequest.Description,
	}

	createErr := r.databaseHandle.Scopes(scopes.SkipCategoriesAssociationUpsert).Create(product).Error
	if createErr != nil {
		return nil, errors.HandleDatabaseError(createErr)
	}

	return product, nil
}

// GetAllProducts ...
func (r *ProductsRepository) GetAllProducts() ([]*repositoryModels.Product, error) {
	var products []*repositoryModels.Product

	findErr := r.databaseHandle.Scopes(scopes.PreloadCategoriesAssociation).Find(&products).Error
	if findErr != nil {
		return nil, errors.HandleDatabaseError(findErr)
	}

	return products, nil
}

// GetProductById ...
func (r *ProductsRepository) GetProductById(productId string) (*repositoryModels.Product, error) {
	var product repositoryModels.Product

	firstErr := r.databaseHandle.Scopes(
		scopes.WhereId(productId),
		scopes.PreloadCategoriesAssociation,
	).First(&product).Error
	if firstErr != nil {
		return nil, errors.HandleDatabaseError(firstErr)
	}

	return &product, nil
}

// UpdateProduct ...
func (r *ProductsRepository) UpdateProduct(productId string, productUpdateRequest *models.ProductUpdateRequest) (*repositoryModels.Product, error) {
	var product repositoryModels.Product

	var updateProduct repositoryModels.Product
	if productUpdateRequest.Name != nil {
		updateProduct.Name = *productUpdateRequest.Name
	}

	if productUpdateRequest.Description != nil {
		updateProduct.Description = productUpdateRequest.Description
	}

	updatesErr := r.databaseHandle.Scopes(scopes.WhereId(productId)).Updates(updateProduct).Error
	if updatesErr != nil {
		return nil, errors.HandleDatabaseError(updatesErr)
	}

	firstErr := r.databaseHandle.Scopes(
		scopes.WhereId(productId),
		scopes.PreloadCategoriesAssociation,
	).First(&product).Error
	if firstErr != nil {
		return nil, errors.HandleDatabaseError(firstErr)
	}

	return &product, nil
}

// DeleteProduct ...
func (r *ProductsRepository) DeleteProduct(productId string) error {
	deleteErr := r.databaseHandle.Scopes(scopes.WhereId(productId)).Delete(&repositoryModels.Product{}).Error
	if deleteErr != nil {
		return errors.HandleDatabaseError(deleteErr)
	}

	return nil
}
