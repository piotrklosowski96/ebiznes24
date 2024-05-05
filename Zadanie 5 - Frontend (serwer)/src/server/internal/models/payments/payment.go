package payments

import (
	"github.com/AlekSi/pointer"
	"github.com/go-openapi/strfmt"

	"Backend/openapi/gen/backend/models"
)

// Payment ...
type Payment struct {
	Id     string
	Status string
}

// ToAPIModel ...
func (p *Payment) ToAPIModel() *models.PaymentResponse {
	return &models.PaymentResponse{
		ID:     pointer.To(strfmt.UUID4(p.Id)),
		Status: &p.Status,
	}
}

// FromAPIModel ...
func FromAPIModel(paymentCreateModel *models.PaymentCreate) *Payment {
	return &Payment{
		Status: paymentCreateModel.Status,
	}
}
