package repositories

import (
	"github.com/google/uuid"
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
	categories := make([]*repositoryModels.Category, len(productCreateRequest.CategoryIds))
	for idx, categoryId := range productCreateRequest.CategoryIds {
		categories[idx] = &repositoryModels.Category{
			CommonFields: repositoryModels.CommonFields{
				ID: categoryId,
			},
		}
	}
	product := &repositoryModels.Product{
		Name:        productCreateRequest.Name,
		Description: productCreateRequest.Description,
		Categories:  categories,
	}

	createErr := r.databaseHandle.Scopes(scopes.SkipCategoriesAssociationUpsert).Create(product).Error
	if createErr != nil {
		return nil, errors.HandleDatabaseError(createErr)
	}

	var createdProduct repositoryModels.Product
	findErr := r.databaseHandle.Scopes(
		scopes.WhereId(product.ID.String()),
		scopes.PreloadCategoriesAssociation,
	).First(&createdProduct).Error
	if findErr != nil {
		return nil, errors.HandleDatabaseError(findErr)
	}

	return &createdProduct, nil
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
	categories := make([]*repositoryModels.Category, len(productUpdateRequest.CategoryIds))
	for idx, categoryId := range productUpdateRequest.CategoryIds {
		categories[idx] = &repositoryModels.Category{
			CommonFields: repositoryModels.CommonFields{
				ID: categoryId,
			},
		}
	}
	updateProduct := repositoryModels.Product{
		CommonFields: repositoryModels.CommonFields{
			ID: uuid.MustParse(productId),
		},
	}

	if productUpdateRequest.Name != nil {
		updateProduct.Name = *productUpdateRequest.Name
	}

	if productUpdateRequest.Description != nil {
		updateProduct.Description = productUpdateRequest.Description
	}

	transactionErr := r.databaseHandle.Transaction(func(tx *gorm.DB) error {
		replaceErr := r.databaseHandle.Model(&updateProduct).Scopes(scopes.SkipCategoriesAssociationUpsert).Association("Categories").Replace(&categories)
		if replaceErr != nil {
			return errors.HandleDatabaseError(replaceErr)
		}

		updatesErr := r.databaseHandle.Scopes(scopes.SkipProductsAssociationUpsert).Updates(updateProduct).Error
		if updatesErr != nil {
			return errors.HandleDatabaseError(updatesErr)
		}

		return nil
	})
	if transactionErr != nil {
		return nil, transactionErr
	}

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

// DeleteProduct ...
func (r *ProductsRepository) DeleteProduct(productId string) error {
	deleteErr := r.databaseHandle.Scopes(scopes.WhereId(productId)).Delete(&repositoryModels.Product{}).Error
	if deleteErr != nil {
		return errors.HandleDatabaseError(deleteErr)
	}

	return nil
}
