package models

// CategoryCreateRequest ...
type CategoryCreateRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}
