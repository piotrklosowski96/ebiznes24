package repositories

import (
	"slices"

	"github.com/google/uuid"

	"zadanie4/internal/models"
	"zadanie4/internal/repositories/errors"
	repositoryModels "zadanie4/internal/repositories/models"
)

// ProductsRepository ...
type ProductsRepository struct {
	products []*repositoryModels.Product
}

// NewProductsRepository ...
func NewProductsRepository() *ProductsRepository {
	return &ProductsRepository{
		products: []*repositoryModels.Product{},
	}
}

// CreateProduct ...
func (r *ProductsRepository) CreateProduct(productCreateRequest *models.ProductCreateRequest) (*repositoryModels.Product, error) {
	product := &repositoryModels.Product{
		ProductID:   uuid.NewString(),
		Name:        productCreateRequest.Name,
		Description: productCreateRequest.Description,
	}

	r.products = append(r.products, product)

	return product, nil
}

// GetAllProducts ...
func (r *ProductsRepository) GetAllProducts() ([]*repositoryModels.Product, error) {
	return r.products, nil
}

// GetProductById ...
func (r *ProductsRepository) GetProductById(productId string) (*repositoryModels.Product, error) {
	idx := slices.IndexFunc(r.products, isProductWithIdComparator(productId))
	if idx < 0 {
		return nil, &errors.ResourceNotFoundError{ResourceID: productId}
	}

	return r.products[idx], nil
}

// UpdateProduct ...
func (r *ProductsRepository) UpdateProduct(productId string, productUpdateRequest *models.ProductUpdateRequest) (*repositoryModels.Product, error) {
	idx := slices.IndexFunc(r.products, isProductWithIdComparator(productId))
	if idx < 0 {
		return nil, &errors.ResourceNotFoundError{ResourceID: productId}
	}

	if productUpdateRequest.Name != nil {
		r.products[idx].Name = *productUpdateRequest.Name
	}

	if productUpdateRequest.Description != nil {
		r.products[idx].Description = productUpdateRequest.Description
	}

	return r.products[idx], nil
}

// DeleteProduct ...
func (r *ProductsRepository) DeleteProduct(productId string) error {
	idx := slices.IndexFunc(r.products, isProductWithIdComparator(productId))
	if idx < 0 {
		return nil
	}

	r.products[idx] = r.products[len(r.products)-1]
	r.products = r.products[:len(r.products)-1]

	return nil
}

func isProductWithIdComparator(productId string) func(product *repositoryModels.Product) bool {
	return func(product *repositoryModels.Product) bool {
		return product.ProductID == productId
	}
}
