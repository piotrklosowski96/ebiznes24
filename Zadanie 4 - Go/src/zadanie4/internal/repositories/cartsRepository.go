package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"zadanie4/internal/models"
	"zadanie4/internal/repositories/errors"
	repositoryModels "zadanie4/internal/repositories/models"
	"zadanie4/internal/repositories/scopes"
)

// CartsRepository ...
type CartsRepository struct {
	databaseHandle *gorm.DB
}

// NewCartsRepository ...
func NewCartsRepository(databaseHandle *gorm.DB) *CartsRepository {
	autoMigrateErr := databaseHandle.AutoMigrate(
		&repositoryModels.Cart{},
		&repositoryModels.Product{},
	)
	if autoMigrateErr != nil {
		panic(autoMigrateErr.Error())
	}

	return &CartsRepository{
		databaseHandle: databaseHandle,
	}
}

// CreateCart ...
func (r *CartsRepository) CreateCart(cartCreateRequest *models.CartCreateRequest) (*repositoryModels.Cart, error) {
	products := make([]*repositoryModels.Product, len(cartCreateRequest.ProductIds))
	for idx, productId := range cartCreateRequest.ProductIds {
		products[idx] = &repositoryModels.Product{
			CommonFields: repositoryModels.CommonFields{
				ID: productId,
			},
		}
	}
	cart := &repositoryModels.Cart{
		Name:        cartCreateRequest.Name,
		Description: cartCreateRequest.Description,
		Products:    products,
	}

	createErr := r.databaseHandle.Scopes(scopes.SkipProductsAssociationUpsert).Create(cart).Error
	if createErr != nil {
		return nil, errors.HandleDatabaseError(createErr)
	}

	// NOTE(Piotr Kłosowski): Check if there is possibility to create model in DB and return it using different model
	// in one statement
	var createdCart repositoryModels.Cart
	findErr := r.databaseHandle.Scopes(
		scopes.WhereId(cart.ID.String()),
		scopes.PreloadProductsAssociation,
	).First(&createdCart).Error
	if findErr != nil {
		return nil, errors.HandleDatabaseError(findErr)
	}

	return &createdCart, nil
}

// GetAllCarts ...
func (r *CartsRepository) GetAllCarts() ([]*repositoryModels.Cart, error) {
	var carts []*repositoryModels.Cart

	findErr := r.databaseHandle.Scopes(scopes.PreloadProductsAssociation).Find(&carts).Error
	if findErr != nil {
		return nil, errors.HandleDatabaseError(findErr)
	}

	return carts, nil
}

// GetCartById ...
func (r *CartsRepository) GetCartById(cartId string) (*repositoryModels.Cart, error) {
	var cart repositoryModels.Cart

	firstErr := r.databaseHandle.Scopes(
		scopes.WhereId(cartId),
		scopes.PreloadProductsAssociation,
	).First(&cart).Error
	if firstErr != nil {
		return nil, errors.HandleDatabaseError(firstErr)
	}

	return &cart, nil
}

// UpdateCart ...
func (r *CartsRepository) UpdateCart(cartId string, cartUpdateRequest *models.CartUpdateRequest) (*repositoryModels.Cart, error) {
	products := make([]*repositoryModels.Product, len(cartUpdateRequest.ProductIds))
	for idx, productId := range cartUpdateRequest.ProductIds {
		products[idx] = &repositoryModels.Product{
			CommonFields: repositoryModels.CommonFields{
				ID: productId,
			},
		}
	}
	updateCart := repositoryModels.Cart{
		CommonFields: repositoryModels.CommonFields{
			ID: uuid.MustParse(cartId),
		},
	}

	if cartUpdateRequest.Name != nil {
		updateCart.Name = *cartUpdateRequest.Name
	}

	if cartUpdateRequest.Description != nil {
		updateCart.Description = cartUpdateRequest.Description
	}

	transactionErr := r.databaseHandle.Transaction(func(tx *gorm.DB) error {
		replaceErr := r.databaseHandle.Model(&updateCart).Scopes(scopes.SkipProductsAssociationUpsert).Association("Products").Replace(&products)
		if replaceErr != nil {
			return errors.HandleDatabaseError(replaceErr)
		}

		updatesErr := r.databaseHandle.Scopes(scopes.SkipProductsAssociationUpsert).Updates(updateCart).Error
		if updatesErr != nil {
			return errors.HandleDatabaseError(updatesErr)
		}

		return nil
	})
	if transactionErr != nil {
		return nil, errors.HandleDatabaseError(transactionErr)
	}

	var cart repositoryModels.Cart
	firstErr := r.databaseHandle.Scopes(
		scopes.WhereId(cartId),
		scopes.PreloadProductsAssociation,
	).First(&cart).Error
	if firstErr != nil {
		return nil, errors.HandleDatabaseError(firstErr)
	}

	return &cart, nil
}

// DeleteCart ...
func (r *CartsRepository) DeleteCart(cartId string) error {
	deleteErr := r.databaseHandle.Scopes(scopes.WhereId(cartId)).Delete(&repositoryModels.Cart{}).Error
	if deleteErr != nil {
		return errors.HandleDatabaseError(deleteErr)
	}

	return nil
}
