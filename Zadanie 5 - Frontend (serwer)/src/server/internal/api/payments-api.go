package apis

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"

	paymentsModels "Backend/internal/models/payments"
	"Backend/internal/storage"
	"Backend/openapi/gen/backend/models"
	"Backend/openapi/gen/backend/server/operations/payments"
	"Backend/openapi/gen/backend/server/operations/products"
)

const (
	CreatePaymentEP          = "CreatePayment"
	GetPaymentsEP            = "GetPayments"
	GetPaymentByIDEP         = "GetPaymentByID"
	PatchPaymentsPaymentIDEP = "PatchPaymentsPaymentID"
	DeletePaymentByIDEP      = "DeletePaymentByID"
)

// PaymentsAPI ...
type PaymentsAPI struct {
	database storage.IStorage
}

// NewPaymentsAPI ...
func NewPaymentsAPI(database storage.IStorage) *PaymentsAPI {
	return &PaymentsAPI{
		database: database,
	}
}

// CreatePayment ...
func (api *PaymentsAPI) CreatePayment(params payments.CreatePaymentParams) middleware.Responder {
	functionName := CreatePaymentEP
	ctx := params.HTTPRequest.Context()

	session := api.database.Open(functionName)
	startSessionErr := session.StartSession()
	if startSessionErr != nil {
		return payments.NewCreatePaymentInternalServerError().WithPayload(
			&models.Error{Message: startSessionErr.Error()},
		)
	}

	defer session.Close(ctx)

	paymentCreate := paymentsModels.FromAPIModel(params.Body)
	paymentCreate.Id = uuid.NewString()

	createdPayment, createPaymentErr := session.CreatePayment(ctx, paymentCreate)
	if createPaymentErr != nil {
		return payments.NewCreatePaymentInternalServerError().WithPayload(
			&models.Error{Message: createPaymentErr.Error()},
		)
	}

	return payments.NewCreatePaymentCreated().WithPayload(createdPayment.ToAPIModel())
}

// GetPayments ...
func (api *PaymentsAPI) GetPayments(params payments.GetPaymentsParams) middleware.Responder {
	functionName := GetPaymentsEP
	ctx := params.HTTPRequest.Context()

	session := api.database.Open(functionName)
	startSessionErr := session.StartSession()
	if startSessionErr != nil {
		return payments.NewGetPaymentsInternalServerError().WithPayload(
			&models.Error{Message: startSessionErr.Error()},
		)
	}

	defer session.Close(ctx)

	paymentsFilter := &paymentsModels.PaymentsFilter{
		Offset: *params.Offset,
		Limit:  *params.Limit,
	}
	paymentsArray, getPaymentsErr := session.GetPayments(ctx, paymentsFilter)
	if getPaymentsErr != nil {
		return products.NewGetProductsInternalServerError().WithPayload(
			&models.Error{Message: getPaymentsErr.Error()},
		)
	}

	paymentsResponseArray := &models.PaymentResponseArray{
		Pagination: models.Pagination{},
		Payments:   make([]*models.PaymentResponse, len(paymentsArray)),
	}
	for idx := range paymentsArray {
		paymentsResponseArray.Payments[idx] = paymentsArray[idx].ToAPIModel()
	}

	return payments.NewGetPaymentsOK().WithPayload(paymentsResponseArray)
}

// GetPaymentByID ...
func (api *PaymentsAPI) GetPaymentByID(params payments.GetPaymentByIDParams) middleware.Responder {
	functionName := GetPaymentByIDEP
	ctx := params.HTTPRequest.Context()

	session := api.database.Open(functionName)
	startSessionErr := session.StartSession()
	if startSessionErr != nil {
		return payments.NewGetPaymentsInternalServerError().WithPayload(
			&models.Error{Message: startSessionErr.Error()},
		)
	}

	paymentId := params.PaymentID.String()
	payment, getPaymentByIdErr := session.GetPaymentById(ctx, paymentId)
	if getPaymentByIdErr != nil {
		return payments.NewGetPaymentsInternalServerError().WithPayload(
			&models.Error{Message: startSessionErr.Error()},
		)
	}

	return payments.NewGetPaymentByIDOK().WithPayload(payment.ToAPIModel())
}

// UpdatePaymentByID ...
func (api *PaymentsAPI) UpdatePaymentByID(params payments.UpdatePaymentByIDParams) middleware.Responder {
	functionName := PatchPaymentsPaymentIDEP
	ctx := params.HTTPRequest.Context()

	session := api.database.Open(functionName)
	startSessionErr := session.StartSession()
	if startSessionErr != nil {
		return payments.NewUpdatePaymentByIDInternalServerError().WithPayload(
			&models.Error{Message: startSessionErr.Error()},
		)
	}

	paymentId := params.PaymentID.String()
	paymentUpdate := &paymentsModels.PaymentUpdate{
		Status: params.Body.Status,
	}

	updatedPayment, updatePaymentErr := session.UpdatePayment(ctx, paymentId, paymentUpdate)
	if updatePaymentErr != nil {
		return payments.NewUpdatePaymentByIDInternalServerError().WithPayload(
			&models.Error{Message: startSessionErr.Error()},
		)
	}

	return payments.NewUpdatePaymentByIDOK().WithPayload(updatedPayment.ToAPIModel())
}

// DeletePaymentByID ...
func (api *PaymentsAPI) DeletePaymentByID(params payments.DeletePaymentByIDParams) middleware.Responder {
	functionName := DeletePaymentByIDEP
	ctx := params.HTTPRequest.Context()

	session := api.database.Open(functionName)
	startSessionErr := session.StartSession()
	if startSessionErr != nil {
		return payments.NewDeletePaymentByIDInternalServerError().WithPayload(
			&models.Error{Message: startSessionErr.Error()},
		)
	}

	defer session.Close(ctx)

	paymentId := params.PaymentID.String()
	deletePaymentErr := session.DeletePayment(ctx, paymentId)
	if deletePaymentErr != nil {
		return payments.NewDeletePaymentByIDInternalServerError().WithPayload(
			&models.Error{Message: deletePaymentErr.Error()},
		)
	}

	return payments.NewDeletePaymentByIDNoContent()
}
