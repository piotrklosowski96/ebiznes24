package apis

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"

	cartsModels "Backend/internal/models/carts"
	"Backend/internal/storage"
	"Backend/openapi/gen/backend/models"
	"Backend/openapi/gen/backend/server/operations/carts"
)

const (
	CreateCartEP       = "CreateCart"
	GetCartsEP         = "GetCarts"
	GetCartByIDEP      = "GetCartByID"
	PatchCartsCartIDEP = "PatchCartsCartID"
	DeleteCartByIDEP   = "DeleteCartByID"
)

// CartsAPI ...
type CartsAPI struct {
	database storage.IStorage
}

// NewCartsAPI ...
func NewCartsAPI(database storage.IStorage) *CartsAPI {
	return &CartsAPI{
		database: database,
	}
}

// CreateCart ...
func (api *CartsAPI) CreateCart(params carts.CreateCartParams) middleware.Responder {
	functionName := CreateCartEP
	ctx := params.HTTPRequest.Context()

	session := api.database.Open(functionName)
	startSessionErr := session.StartSession()
	if startSessionErr != nil {
		return carts.NewCreateCartInternalServerError().WithPayload(
			&models.Error{Message: startSessionErr.Error()},
		)
	}

	defer session.Close(ctx)

	cartCreate := cartsModels.FromCartCreateAPIModel(params.Body)
	cartCreate.Id = uuid.NewString()

	createdCart, createCartErr := session.CreateCart(ctx, cartCreate)
	if createCartErr != nil {
		return carts.NewCreateCartInternalServerError().WithPayload(
			&models.Error{Message: createCartErr.Error()},
		)
	}

	return carts.NewCreateCartCreated().WithPayload(createdCart.ToAPIModel())
}

// GetCarts ...
func (api *CartsAPI) GetCarts(params carts.GetCartsParams) middleware.Responder {
	functionName := GetCartsEP
	ctx := params.HTTPRequest.Context()

	session := api.database.Open(functionName)
	startSessionErr := session.StartSession()
	if startSessionErr != nil {
		return carts.NewGetCartsInternalServerError().WithPayload(
			&models.Error{Message: startSessionErr.Error()},
		)
	}

	defer session.Close(ctx)

	cartsFilter := &cartsModels.CartsFilter{
		Offset: *params.Offset,
		Limit:  *params.Limit,
	}
	cartsArray, getCartsErr := session.GetCarts(ctx, cartsFilter)
	if getCartsErr != nil {
		return carts.NewGetCartsInternalServerError().WithPayload(
			&models.Error{Message: getCartsErr.Error()},
		)
	}

	cartsResponseArray := &models.CartResponseArray{
		Pagination: models.Pagination{},
		Carts:      make([]*models.CartResponse, len(cartsArray)),
	}
	for idx := range cartsArray {
		cartsResponseArray.Carts[idx] = cartsArray[idx].ToAPIModel()
	}

	return carts.NewGetCartsOK().WithPayload(cartsResponseArray)
}

// GetCartByID ...
func (api *CartsAPI) GetCartByID(params carts.GetCartByIDParams) middleware.Responder {
	functionName := GetCartByIDEP
	ctx := params.HTTPRequest.Context()

	session := api.database.Open(functionName)
	startSessionErr := session.StartSession()
	if startSessionErr != nil {
		return carts.NewGetCartByIDInternalServerError().WithPayload(
			&models.Error{Message: startSessionErr.Error()},
		)
	}

	cartId := params.CartID.String()
	cart, getCartByIdErr := session.GetCartById(ctx, cartId)
	if getCartByIdErr != nil {
		return carts.NewGetCartByIDInternalServerError().WithPayload(
			&models.Error{Message: startSessionErr.Error()},
		)
	}

	return carts.NewGetCartByIDOK().WithPayload(cart.ToAPIModel())
}

// PatchCartsCartID ...
func (api *CartsAPI) PatchCartsCartID(params carts.PatchCartsCartIDParams) middleware.Responder {
	functionName := PatchCartsCartIDEP
	ctx := params.HTTPRequest.Context()

	session := api.database.Open(functionName)
	startSessionErr := session.StartSession()
	if startSessionErr != nil {
		return carts.NewPatchCartsCartIDInternalServerError().WithPayload(
			&models.Error{Message: startSessionErr.Error()},
		)
	}

	cartId := params.CartID.String()
	cartUpdate := &cartsModels.CartUpdate{}

	updatedCart, updateCartErr := session.UpdateCart(ctx, cartId, cartUpdate)
	if updateCartErr != nil {
		return carts.NewPatchCartsCartIDInternalServerError().WithPayload(
			&models.Error{Message: startSessionErr.Error()},
		)
	}

	return carts.NewPatchCartsCartIDOK().WithPayload(updatedCart.ToAPIModel())
}

// DeleteCartByID ...
func (api *CartsAPI) DeleteCartByID(params carts.DeleteCartByIDParams) middleware.Responder {
	functionName := DeleteCartByIDEP
	ctx := params.HTTPRequest.Context()

	session := api.database.Open(functionName)
	startSessionErr := session.StartSession()
	if startSessionErr != nil {
		return carts.NewDeleteCartByIDInternalServerError().WithPayload(
			&models.Error{Message: startSessionErr.Error()},
		)
	}

	defer session.Close(ctx)

	cartId := params.CartID.String()
	deletePaymentErr := session.DeletePayment(ctx, cartId)
	if deletePaymentErr != nil {
		return carts.NewDeleteCartByIDInternalServerError().WithPayload(
			&models.Error{Message: deletePaymentErr.Error()},
		)
	}

	return carts.NewDeleteCartByIDNoContent()
}
