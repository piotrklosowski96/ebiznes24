package storage

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"Backend/internal/models/payments"
	"Backend/internal/storage/models"
)

// CreatePayment ...
func (s *Session) CreatePayment(ctx context.Context, payment *payments.Payment) (*payments.Payment, error) {
	paymentDB := models.FromPaymentDomainModel(payment)
	paymentDB.CreatedAt = time.Now()
	paymentDB.UpdatedAt = time.Now()
	paymentDB.SoftDelete = false

	_, insertOneErr := s.storage.paymentsCollection.InsertOne(ctx, paymentDB)
	if insertOneErr != nil {
		return nil, insertOneErr
	}

	filter := bson.M{"_id": paymentDB.Id, "soft_delete": false}
	findOneResult := s.storage.paymentsCollection.FindOne(ctx, filter)
	if findOneResult.Err() != nil {
		return nil, findOneResult.Err()
	}

	var insertedPayment models.Payment
	decodeErr := findOneResult.Decode(&insertedPayment)
	if decodeErr != nil {
		return nil, decodeErr
	}

	return insertedPayment.ToPaymentDomainModel(), nil
}

// GetPaymentById ...
func (s *Session) GetPaymentById(ctx context.Context, paymentId string) (*payments.Payment, error) {
	filter := bson.M{"_id": paymentId, "soft_delete": false}
	findOneResult := s.storage.paymentsCollection.FindOne(ctx, filter)
	if findOneResult.Err() != nil {
		return nil, findOneResult.Err()
	}

	var insertedPayment models.Payment
	decodeErr := findOneResult.Decode(&insertedPayment)
	if decodeErr != nil {
		return nil, decodeErr
	}

	return insertedPayment.ToPaymentDomainModel(), nil
}

// GetPayments ...
func (s *Session) GetPayments(ctx context.Context, paymentsFilter *payments.PaymentsFilter) ([]*payments.Payment, error) {
	filter := bson.M{"soft_delete": false}
	findOptions := options.Find().
		SetSkip(paymentsFilter.Offset).
		SetLimit(paymentsFilter.Limit)
	findResult, findErr := s.storage.paymentsCollection.Find(ctx, filter, findOptions)
	if findErr != nil {
		return nil, findErr
	}

	paymentsDBResult := make([]*models.Payment, paymentsFilter.Limit)
	decodeErr := findResult.All(ctx, &paymentsDBResult)
	if decodeErr != nil {
		return nil, decodeErr
	}

	payments := make([]*payments.Payment, len(paymentsDBResult))
	for idx := range paymentsDBResult {
		payments[idx] = paymentsDBResult[idx].ToPaymentDomainModel()
	}

	return payments, nil
}

// UpdatePayment ...
func (s *Session) UpdatePayment(ctx context.Context, productId string, productUpdateDocument *payments.PaymentUpdate) (*payments.Payment, error) {
	filter := bson.M{"_id": productId}
	update := s.preparePaymentUpdateBSON(productUpdateDocument)
	findOneAndUpdateOptions := options.FindOneAndUpdate().SetReturnDocument(options.After)
	findOneAndUpdateResult := s.storage.paymentsCollection.FindOneAndUpdate(ctx, filter, update, findOneAndUpdateOptions)
	if findOneAndUpdateResult.Err() != nil {
		return nil, findOneAndUpdateResult.Err()
	}

	var paymentDB models.Payment
	decodeErr := findOneAndUpdateResult.Decode(&paymentDB)
	if decodeErr != nil {
		return nil, decodeErr
	}

	return paymentDB.ToPaymentDomainModel(), nil
}

// DeletePayment ...
func (s *Session) DeletePayment(ctx context.Context, paymentId string) error {
	update := bson.M{
		"$set": bson.M{"soft_delete": true},
	}
	_, updateByIDErr := s.storage.paymentsCollection.UpdateByID(ctx, paymentId, update)
	if updateByIDErr != nil {
		return updateByIDErr
	}

	return nil
}

func (s *Session) preparePaymentUpdateBSON(paymentUpdateDocument *payments.PaymentUpdate) bson.M {
	fields := bson.M{}

	setIfNotNil(&fields, "status", paymentUpdateDocument.Status)

	return bson.M{"$set": fields}
}
