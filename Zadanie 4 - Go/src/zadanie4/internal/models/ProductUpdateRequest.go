package models

// ProductUpdateRequest ...
type ProductUpdateRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}
