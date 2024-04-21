package models

import "github.com/google/uuid"

// ProductCreateRequest ...
type ProductCreateRequest struct {
	Name        string      `json:"name"`
	Description *string     `json:"description,omitempty"`
	CategoryIds []uuid.UUID `json:"category_ids,omitempty"`
}
