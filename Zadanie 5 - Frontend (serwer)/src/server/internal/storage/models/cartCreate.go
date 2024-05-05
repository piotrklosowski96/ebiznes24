package models

import (
	"Backend/internal/models/carts"
)

// CartCreate ...
type CartCreate struct {
	Id          string   `json:"_id" bson:"_id"`
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	Products    []string `json:"products" bson:"products"`

	CommonFields `json:",inline" bson:",inline"`
}

// FromCartCreateDomainModel ...
func FromCartCreateDomainModel(cartCreate *carts.CartCreate) *CartCreate {
	return &CartCreate{
		Id:          cartCreate.Id,
		Name:        cartCreate.Name,
		Description: cartCreate.Description,
		Products:    cartCreate.Products,
	}
}
