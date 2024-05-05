package storage

import (
	"Backend/internal/models/payments"
	"context"
)

// ISessionPaymentsStorage ...
type ISessionPaymentsStorage interface {
	CreatePayment(ctx context.Context, payment *payments.Payment) (*payments.Payment, error)
	GetPaymentById(ctx context.Context, paymentId string) (*payments.Payment, error)
	GetPayments(ctx context.Context, paymentsFilter *payments.PaymentsFilter) ([]*payments.Payment, error)
	UpdatePayment(ctx context.Context, paymentId string, paymentUpdateDocument *payments.PaymentUpdate) (*payments.Payment, error)
	DeletePayment(ctx context.Context, paymentId string) error
}
