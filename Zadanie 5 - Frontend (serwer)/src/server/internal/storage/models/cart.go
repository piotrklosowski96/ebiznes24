package models

import (
	"Backend/internal/models/carts"
	"Backend/internal/models/products"
)

// Cart ...
type Cart struct {
	Id          string     `json:"_id" bson:"_id"`
	Name        string     `json:"name" bson:"name"`
	Description string     `json:"description" bson:"description"`
	Products    []*Product `json:"products" bson:"products"`

	CommonFields `json:",inline" bson:",inline"`
}

// ToCartDomainModel ...
func (c *Cart) ToCartDomainModel() *carts.Cart {
	cart := &carts.Cart{
		Id:          c.Id,
		Name:        c.Name,
		Description: c.Description,
		Products:    make([]*products.Product, len(c.Products)),
	}
	for idx := range c.Products {
		cart.Products[idx] = c.Products[idx].ToProductDomainModel()
	}

	return cart
}
