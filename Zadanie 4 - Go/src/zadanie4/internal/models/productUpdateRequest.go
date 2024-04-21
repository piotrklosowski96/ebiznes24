package models

import "github.com/google/uuid"

// ProductUpdateRequest ...
type ProductUpdateRequest struct {
	Name        *string     `json:"name,omitempty"`
	Description *string     `json:"description,omitempty"`
	CategoryIds []uuid.UUID `json:"category_ids,omitempty"`
}
