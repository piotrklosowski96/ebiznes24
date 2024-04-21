package models

// ProductResponse ...
type ProductResponse struct {
	ProductID   string  `json:"product_id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}
