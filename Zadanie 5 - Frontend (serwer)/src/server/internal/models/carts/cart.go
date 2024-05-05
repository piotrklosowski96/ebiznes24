package carts

import (
	"github.com/go-openapi/strfmt"

	"Backend/internal/models/products"
	"Backend/openapi/gen/backend/models"
)

// Cart ...
type Cart struct {
	Id          string
	Name        string
	Description string
	Products    []*products.Product
}

// ToAPIModel ...
func (c *Cart) ToAPIModel() *models.CartResponse {
	cartResponse := &models.CartResponse{
		ID:          strfmt.UUID4(c.Id),
		Name:        c.Name,
		Description: c.Description,
		Products:    make([]*models.ProductResponse, len(c.Products)),
	}
	for idx := range c.Products {
		cartResponse.Products[idx] = c.Products[idx].ToAPIModel()
	}

	return cartResponse
}
