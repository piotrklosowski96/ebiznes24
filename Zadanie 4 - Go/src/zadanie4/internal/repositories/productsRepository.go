package repositories

import (
	"slices"

	"github.com/google/uuid"

	"zadanie4/internal/repositories/errors"
)

// Product ...
type Product struct {
	ProductID   string  `json:"product_id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// ProductsRepository ...
type ProductsRepository struct {
	products []*Product
}

// NewProductsRepository ...
func NewProductsRepository() *ProductsRepository {
	return &ProductsRepository{
		products: []*Product{},
	}
}

// CreateProduct ...
func (r *ProductsRepository) CreateProduct(product *Product) (*Product, error) {
	product.ProductID = uuid.NewString()
	r.products = append(r.products, product)

	return product, nil
}

// GetAllProducts ...
func (r *ProductsRepository) GetAllProducts() ([]*Product, error) {
	return r.products, nil
}

// GetProductById ...
func (r *ProductsRepository) GetProductById(productId string) (*Product, error) {
	idx := slices.IndexFunc(r.products, isProductWithIdComparator(productId))
	if idx < 0 {
		return nil, &errors.ResourceNotFoundError{ResourceID: productId}
	}

	return r.products[idx], nil
}

// UpdateProduct ...
func (r *ProductsRepository) UpdateProduct(productId string, updatedProduct *Product) (*Product, error) {
	idx := slices.IndexFunc(r.products, isProductWithIdComparator(productId))
	if idx < 0 {
		return nil, &errors.ResourceNotFoundError{ResourceID: productId}
	}

	r.products[idx] = updatedProduct

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

func isProductWithIdComparator(productId string) func(product *Product) bool {
	return func(product *Product) bool {
		return product.ProductID == productId
	}
}
