package models

import (
	repositoryModels "zadanie4/internal/repositories/models"
)

// ProductResponse ...
type ProductResponse struct {
	ProductID   string              `json:"product_id"`
	Name        string              `json:"name"`
	Description *string             `json:"description"`
	Categories  []*CategoryResponse `json:"categories"`
}

// FromDatabaseProduct ...
func FromDatabaseProduct(productDB *repositoryModels.Product) *ProductResponse {
	productCategories := make([]*CategoryResponse, len(productDB.Categories))
	for idx, category := range productDB.Categories {
		productCategories[idx] = FromDatabaseCategory(category)
	}

	return &ProductResponse{
		ProductID:   productDB.ID.String(),
		Name:        productDB.Name,
		Description: productDB.Description,
		Categories:  productCategories,
	}
}
