package models

import "github.com/google/uuid"

// CartCreateRequest ...
type CartCreateRequest struct {
	Name        string      `json:"name"`
	Description *string     `json:"description,omitempty"`
	ProductIds  []uuid.UUID `json:"product_ids,omitempty"`
}
