package router

import (
	apis "Backend/internal/api"
	"Backend/internal/storage"
	"Backend/openapi/gen/backend/server/operations"
	"Backend/openapi/gen/backend/server/operations/carts"
	"Backend/openapi/gen/backend/server/operations/payments"
	"Backend/openapi/gen/backend/server/operations/products"
)

// Router ...
type Router struct {
	swaggerAPI  *operations.BackendAPI
	productsApi *apis.ProductsAPI
	cartsApi    *apis.CartsAPI
	paymentsApi *apis.PaymentsAPI
}

// NewRouter ...
func NewRouter(
	swaggerAPI *operations.BackendAPI,
	database storage.IStorage,
) *Router {
	router := &Router{
		swaggerAPI:  swaggerAPI,
		productsApi: apis.NewProductsAPI(database),
		cartsApi:    apis.NewCartsAPI(database),
		paymentsApi: apis.NewPaymentsAPI(database),
	}

	return router
}

// RegisterRoutes ...
func (router *Router) RegisterRoutes() {
	router.swaggerAPI.ProductsCreateProductHandler = products.CreateProductHandlerFunc(router.productsApi.CreateProduct)
	router.swaggerAPI.ProductsGetProductsHandler = products.GetProductsHandlerFunc(router.productsApi.GetProducts)
	router.swaggerAPI.ProductsGetProductByIDHandler = products.GetProductByIDHandlerFunc(router.productsApi.GetProductByID)
	router.swaggerAPI.ProductsPatchProductsProductIDHandler = products.PatchProductsProductIDHandlerFunc(router.productsApi.PatchProductsProductID)
	router.swaggerAPI.ProductsDeleteProductByIDHandler = products.DeleteProductByIDHandlerFunc(router.productsApi.DeleteProductByID)

	router.swaggerAPI.CartsCreateCartHandler = carts.CreateCartHandlerFunc(router.cartsApi.CreateCart)
	router.swaggerAPI.CartsGetCartsHandler = carts.GetCartsHandlerFunc(router.cartsApi.GetCarts)
	router.swaggerAPI.CartsGetCartByIDHandler = carts.GetCartByIDHandlerFunc(router.cartsApi.GetCartByID)
	router.swaggerAPI.CartsPatchCartsCartIDHandler = carts.PatchCartsCartIDHandlerFunc(router.cartsApi.PatchCartsCartID)
	router.swaggerAPI.CartsDeleteCartByIDHandler = carts.DeleteCartByIDHandlerFunc(router.cartsApi.DeleteCartByID)
	router.swaggerAPI.CartsAddProductToCartHandler = carts.AddProductToCartHandlerFunc(router.cartsApi.AddProductToCart)
	router.swaggerAPI.CartsDeleteProductFromCartHandler = carts.DeleteProductFromCartHandlerFunc(router.cartsApi.DeleteProductFromCart)

	router.swaggerAPI.PaymentsCreatePaymentHandler = payments.CreatePaymentHandlerFunc(router.paymentsApi.CreatePayment)
	router.swaggerAPI.PaymentsGetPaymentsHandler = payments.GetPaymentsHandlerFunc(router.paymentsApi.GetPayments)
	router.swaggerAPI.PaymentsGetPaymentByIDHandler = payments.GetPaymentByIDHandlerFunc(router.paymentsApi.GetPaymentByID)
	router.swaggerAPI.PaymentsUpdatePaymentByIDHandler = payments.UpdatePaymentByIDHandlerFunc(router.paymentsApi.UpdatePaymentByID)
	router.swaggerAPI.PaymentsDeletePaymentByIDHandler = payments.DeletePaymentByIDHandlerFunc(router.paymentsApi.DeletePaymentByID)
}
