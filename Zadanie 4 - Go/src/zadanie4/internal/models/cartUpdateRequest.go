package models

import "github.com/google/uuid"

// CartUpdateRequest ...
type CartUpdateRequest struct {
	Name        *string     `json:"name,omitempty"`
	Description *string     `json:"description,omitempty"`
	ProductIds  []uuid.UUID `json:"product_ids,omitempty"`
}
