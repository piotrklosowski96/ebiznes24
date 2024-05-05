package models

import "Backend/internal/models/products"

// Product ...
type Product struct {
	Id          string `json:"_id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`

	CommonFields `json:",inline" bson:",inline"`
}

// ToProductDomainModel ...
func (p *Product) ToProductDomainModel() *products.Product {
	return &products.Product{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
	}
}

// FromProductDomainModel ...
func FromProductDomainModel(product *products.Product) *Product {
	return &Product{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
	}
}
