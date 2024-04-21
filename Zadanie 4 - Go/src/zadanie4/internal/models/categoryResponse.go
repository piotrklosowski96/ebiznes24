package models

import repositoryModels "zadanie4/internal/repositories/models"

// CategoryResponse ...
type CategoryResponse struct {
	CategoryID  string  `json:"category_id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// FromDatabaseCategory ...
func FromDatabaseCategory(productDB *repositoryModels.Category) *CategoryResponse {
	return &CategoryResponse{
		CategoryID:  productDB.ID.String(),
		Name:        productDB.Name,
		Description: productDB.Description,
	}
}
