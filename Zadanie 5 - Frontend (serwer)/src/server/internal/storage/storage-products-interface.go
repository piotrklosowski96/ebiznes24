package storage

import (
	"context"

	"Backend/internal/models/products"
)

// ISessionProductsStorage ...
type ISessionProductsStorage interface {
	CreateProduct(ctx context.Context, product *products.Product) (*products.Product, error)
	GetProductById(ctx context.Context, productId string) (*products.Product, error)
	GetProducts(ctx context.Context, productsFilter *products.ProductsFilter) ([]*products.Product, error)
	UpdateProduct(ctx context.Context, productId string, productUpdateDocument *products.ProductUpdate) (*products.Product, error)
	DeleteProduct(ctx context.Context, productId string) error
}
