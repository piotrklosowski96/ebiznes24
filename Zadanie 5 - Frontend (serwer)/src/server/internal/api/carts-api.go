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
	CreateCartEP            = "CreateCart"
	GetCartsEP              = "GetCarts"
	GetCartByIDEP           = "GetCartByID"
	PatchCartsCartIDEP      = "PatchCartsCartID"
	DeleteCartByIDEP        = "DeleteCartByID"
	AddProductToCartEP      = "AddProductToCart"
	DeleteProductFromCartEP = "DeleteProductFromCart"
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
	deleteCartErr := session.DeleteCart(ctx, cartId)
	if deleteCartErr != nil {
		return carts.NewCreateCartInternalServerError().WithPayload(
			&models.Error{Message: deleteCartErr.Error()},
		)
	}

	return carts.NewDeleteCartByIDNoContent()
}

// AddProductToCart ...
func (api *CartsAPI) AddProductToCart(params carts.AddProductToCartParams) middleware.Responder {
	functionName := AddProductToCartEP
	ctx := params.HTTPRequest.Context()

	session := api.database.Open(functionName)
	startSessionErr := session.StartSession()
	if startSessionErr != nil {
		return carts.NewAddProductToCartInternalServerError().WithPayload(
			&models.Error{Message: startSessionErr.Error()},
		)
	}

	defer session.Close(ctx)

	cartId := params.CartID.String()
	productId := params.ProductID.String()
	updatedCart, addProductToCartErr := session.AddProductToCart(ctx, cartId, productId)
	if addProductToCartErr != nil {
		return carts.NewAddProductToCartInternalServerError().WithPayload(
			&models.Error{Message: addProductToCartErr.Error()},
		)
	}

	return carts.NewAddProductToCartOK().WithPayload(updatedCart.ToAPIModel())
}

// DeleteProductFromCart ...
func (api *CartsAPI) DeleteProductFromCart(params carts.DeleteProductFromCartParams) middleware.Responder {
	functionName := DeleteProductFromCartEP
	ctx := params.HTTPRequest.Context()

	session := api.database.Open(functionName)
	startSessionErr := session.StartSession()
	if startSessionErr != nil {
		return carts.NewDeleteProductFromCartInternalServerError().WithPayload(
			&models.Error{Message: startSessionErr.Error()},
		)
	}

	defer session.Close(ctx)

	cartId := params.CartID.String()
	productId := params.ProductID.String()
	updatedCart, removeProductFromCartErr := session.RemoveProductFromCart(ctx, cartId, productId)
	if removeProductFromCartErr != nil {
		return carts.NewDeleteProductFromCartInternalServerError().WithPayload(
			&models.Error{Message: removeProductFromCartErr.Error()},
		)
	}

	return carts.NewDeleteProductFromCartOK().WithPayload(updatedCart.ToAPIModel())
}
