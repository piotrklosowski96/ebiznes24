package apis

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"

	productsModels "Backend/internal/models/products"
	"Backend/internal/storage"
	"Backend/openapi/gen/backend/models"
	"Backend/openapi/gen/backend/server/operations/products"
)

const (
	CreateProductEP          = "CreateProduct"
	GetProductsEP            = "GetProducts"
	GetProductByIDEP         = "GetProductByID"
	PatchProductsProductIDEP = "PatchProductsProductID"
	DeleteProductByIDEP      = "DeleteProductByID"
)

// ProductsAPI ...
type ProductsAPI struct {
	database storage.IStorage
}

// NewProductsAPI ...
func NewProductsAPI(database storage.IStorage) *ProductsAPI {
	return &ProductsAPI{
		database: database,
	}
}

// CreateProduct ...
func (api *ProductsAPI) CreateProduct(params products.CreateProductParams) middleware.Responder {
	functionName := CreateProductEP
	ctx := params.HTTPRequest.Context()

	session := api.database.Open(functionName)
	startSessionErr := session.StartSession()
	if startSessionErr != nil {
		return products.NewCreateProductInternalServerError().WithPayload(
			&models.Error{Message: startSessionErr.Error()},
		)
	}

	defer session.Close(ctx)

	productCreate := productsModels.FromAPIModel(params.Body)
	productCreate.Id = uuid.NewString()

	product, createProductErr := session.CreateProduct(ctx, productCreate)
	if createProductErr != nil {
		return products.NewCreateProductInternalServerError().WithPayload(
			&models.Error{Message: createProductErr.Error()},
		)
	}

	return products.NewCreateProductCreated().WithPayload(product.ToAPIModel())
}

// GetProducts ...
func (api *ProductsAPI) GetProducts(params products.GetProductsParams) middleware.Responder {
	functionName := GetProductsEP
	ctx := params.HTTPRequest.Context()

	session := api.database.Open(functionName)
	startSessionErr := session.StartSession()
	if startSessionErr != nil {
		return products.NewGetProductsInternalServerError().WithPayload(
			&models.Error{Message: startSessionErr.Error()},
		)
	}

	defer session.Close(ctx)

	productsFilter := &productsModels.ProductsFilter{
		Offset: *params.Offset,
		Limit:  *params.Limit,
	}
	productsArray, getProductsErr := session.GetProducts(ctx, productsFilter)
	if getProductsErr != nil {
		return products.NewGetProductsInternalServerError().WithPayload(
			&models.Error{Message: getProductsErr.Error()},
		)
	}

	productsResponseArray := &models.ProductResponseArray{
		Pagination: models.Pagination{},
		Products:   make([]*models.ProductResponse, len(productsArray)),
	}
	for idx := range productsArray {
		productsResponseArray.Products[idx] = productsArray[idx].ToAPIModel()
	}

	return products.NewGetProductsOK().WithPayload(productsResponseArray)
}

// GetProductByID ...
func (api *ProductsAPI) GetProductByID(params products.GetProductByIDParams) middleware.Responder {
	functionName := GetProductByIDEP
	ctx := params.HTTPRequest.Context()

	session := api.database.Open(functionName)
	startSessionErr := session.StartSession()
	if startSessionErr != nil {
		return products.NewGetProductByIDInternalServerError().WithPayload(
			&models.Error{Message: startSessionErr.Error()},
		)
	}

	productId := params.ProductID.String()
	product, getProductByIdErr := session.GetProductById(ctx, productId)
	if getProductByIdErr != nil {
		return products.NewGetProductByIDInternalServerError().WithPayload(
			&models.Error{Message: startSessionErr.Error()},
		)
	}

	return products.NewGetProductByIDOK().WithPayload(product.ToAPIModel())
}

// PatchProductsProductID ...
func (api *ProductsAPI) PatchProductsProductID(params products.PatchProductsProductIDParams) middleware.Responder {
	functionName := PatchProductsProductIDEP
	ctx := params.HTTPRequest.Context()

	session := api.database.Open(functionName)
	startSessionErr := session.StartSession()
	if startSessionErr != nil {
		return products.NewPatchProductsProductIDInternalServerError().WithPayload(
			&models.Error{Message: startSessionErr.Error()},
		)
	}

	productId := params.ProductID.String()
	productUpdate := &productsModels.ProductUpdate{
		Name:        params.Body.Name,
		Description: params.Body.Description,
	}

	updatedProduct, updateProductErr := session.UpdateProduct(ctx, productId, productUpdate)
	if updateProductErr != nil {
		return products.NewPatchProductsProductIDInternalServerError().WithPayload(
			&models.Error{Message: startSessionErr.Error()},
		)
	}

	return products.NewPatchProductsProductIDOK().WithPayload(updatedProduct.ToAPIModel())
}

// DeleteProductByID ...
func (api *ProductsAPI) DeleteProductByID(params products.DeleteProductByIDParams) middleware.Responder {
	functionName := DeleteProductByIDEP
	ctx := params.HTTPRequest.Context()

	session := api.database.Open(functionName)
	startSessionErr := session.StartSession()
	if startSessionErr != nil {
		return products.NewDeleteProductByIDInternalServerError().WithPayload(
			&models.Error{Message: startSessionErr.Error()},
		)
	}

	defer session.Close(ctx)

	productId := params.ProductID.String()
	deleteProductErr := session.DeleteProduct(ctx, productId)
	if deleteProductErr != nil {
		return products.NewDeleteProductByIDInternalServerError().WithPayload(
			&models.Error{Message: deleteProductErr.Error()},
		)
	}

	return products.NewDeleteProductByIDNoContent()
}
