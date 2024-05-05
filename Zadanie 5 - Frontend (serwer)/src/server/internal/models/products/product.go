package products

import "Backend/openapi/gen/backend/models"

// Product ...
type Product struct {
	Id          string
	Name        string
	Description string
}

// ToAPIModel ...
func (p *Product) ToAPIModel() *models.ProductResponse {
	return &models.ProductResponse{
		ID:          p.Id,
		Name:        p.Name,
		Description: p.Description,
	}
}

// FromAPIModel ...
func FromAPIModel(productCreateModel *models.ProductCreate) *Product {
	return &Product{
		Name:        *productCreateModel.Name,
		Description: productCreateModel.Description,
	}
}
