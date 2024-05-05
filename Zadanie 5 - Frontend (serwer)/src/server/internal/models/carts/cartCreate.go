package carts

import "Backend/openapi/gen/backend/models"

// CartCreate ...
type CartCreate struct {
	Id          string
	Name        string
	Description string
	Products    []string
}

// FromCartCreateAPIModel ...
func FromCartCreateAPIModel(cartCreateRequest *models.CartCreate) *CartCreate {
	cartCreate := &CartCreate{
		Name:        cartCreateRequest.Name,
		Description: cartCreateRequest.Description,
		Products:    make([]string, len(cartCreateRequest.ProductIds)),
	}
	for idx := range cartCreateRequest.ProductIds {
		cartCreate.Products[idx] = cartCreateRequest.ProductIds[idx].String()
	}

	return cartCreate
}
