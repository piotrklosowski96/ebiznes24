package models

// CategoryUpdateRequest ...
type CategoryUpdateRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}
