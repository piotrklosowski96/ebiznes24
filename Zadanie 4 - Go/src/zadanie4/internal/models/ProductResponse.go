package models

import (
	repositoryModels "zadanie4/internal/repositories/models"
)

// ProductResponse ...
type ProductResponse struct {
	ProductID   string  `json:"product_id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// FromDatabaseProduct ...
func FromDatabaseProduct(productDB *repositoryModels.Product) *ProductResponse {
	return &ProductResponse{
		ProductID:   productDB.ProductID,
		Name:        productDB.Name,
		Description: productDB.Description,
	}
}
