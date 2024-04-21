package models

// ProductCreateRequest ...
type ProductCreateRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}
