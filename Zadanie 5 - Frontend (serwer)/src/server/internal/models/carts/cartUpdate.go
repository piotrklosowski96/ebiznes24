package carts

import "Backend/openapi/gen/backend/models"

// CartUpdate ...
type CartUpdate struct {
	Name        *string
	Description *string
	ProductIds  []string
}

// FromCartUpdateAPIModel ...
func FromCartUpdateAPIModel(cartUpdateRequest *models.CartUpdate) *CartUpdate {
	cartCreate := &CartUpdate{
		Name:        cartUpdateRequest.Name,
		Description: cartUpdateRequest.Description,
		ProductIds:  make([]string, len(cartUpdateRequest.ProductIds)),
	}
	for idx := range cartUpdateRequest.ProductIds {
		cartCreate.ProductIds[idx] = cartUpdateRequest.ProductIds[idx].String()
	}

	return cartCreate
}
