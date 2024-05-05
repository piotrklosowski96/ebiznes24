package models

import (
	"Backend/internal/models/payments"
)

// Payment ...
type Payment struct {
	Id     string `json:"_id" bson:"_id"`
	Status string `json:"status" bson:"status"`

	CommonFields `json:",inline" bson:",inline"`
}

// ToPaymentDomainModel ...
func (p *Payment) ToPaymentDomainModel() *payments.Payment {
	return &payments.Payment{
		Id:     p.Id,
		Status: p.Status,
	}
}

// FromPaymentDomainModel ...
func FromPaymentDomainModel(payment *payments.Payment) *Payment {
	return &Payment{
		Id:     payment.Id,
		Status: payment.Status,
	}
}
