package models

import repositoryModels "zadanie4/internal/repositories/models"

// CartResponse ...
type CartResponse struct {
	CartID      string             `json:"cart_id"`
	Name        string             `json:"name"`
	Description *string            `json:"description,omitempty"`
	Products    []*ProductResponse `json:"products,omitempty"`
}

// FromDatabaseCart ...
func FromDatabaseCart(cartDB *repositoryModels.Cart) *CartResponse {
	products := make([]*ProductResponse, len(cartDB.Products))
	for idx, product := range cartDB.Products {
		products[idx] = FromDatabaseProduct(product)
	}

	return &CartResponse{
		CartID:      cartDB.ID.String(),
		Name:        cartDB.Name,
		Description: cartDB.Description,
		Products:    products,
	}
}
