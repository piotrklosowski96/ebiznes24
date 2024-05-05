package storage

import (
	"context"

	"Backend/internal/models/carts"
)

// ISessionCartsStorage ...
type ISessionCartsStorage interface {
	CreateCart(ctx context.Context, cart *carts.CartCreate) (*carts.Cart, error)
	GetCartById(ctx context.Context, cartId string) (*carts.Cart, error)
	GetCarts(ctx context.Context, cartsFilter *carts.CartsFilter) ([]*carts.Cart, error)
	UpdateCart(ctx context.Context, cartId string, cartUpdateDocument *carts.CartUpdate) (*carts.Cart, error)
	DeleteCart(ctx context.Context, cartId string) error
}
