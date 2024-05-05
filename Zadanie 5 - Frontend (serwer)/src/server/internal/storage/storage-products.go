package storage

import (
	"Backend/internal/models/products"
	"Backend/internal/storage/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// CreateProduct ...
func (s *Session) CreateProduct(ctx context.Context, product *products.Product) (*products.Product, error) {
	productDB := models.FromProductDomainModel(product)
	productDB.CreatedAt = time.Now()
	productDB.UpdatedAt = time.Now()
	productDB.SoftDelete = false

	_, insertOneErr := s.storage.productsCollection.InsertOne(ctx, productDB)
	if insertOneErr != nil {
		return nil, insertOneErr
	}

	filter := bson.M{"_id": productDB.Id, "soft_delete": false}
	findOneResult := s.storage.productsCollection.FindOne(ctx, filter)
	if findOneResult.Err() != nil {
		return nil, findOneResult.Err()
	}

	var insertedProduct models.Product
	decodeErr := findOneResult.Decode(&insertedProduct)
	if decodeErr != nil {
		return nil, decodeErr
	}

	return insertedProduct.ToProductDomainModel(), nil
}

// GetProductById ...
func (s *Session) GetProductById(ctx context.Context, productId string) (*products.Product, error) {
	filter := bson.M{"_id": productId, "soft_delete": false}
	findOneResult := s.storage.productsCollection.FindOne(ctx, filter)
	if findOneResult.Err() != nil {
		return nil, findOneResult.Err()
	}

	var insertedProduct models.Product
	decodeErr := findOneResult.Decode(&insertedProduct)
	if decodeErr != nil {
		return nil, decodeErr
	}

	return insertedProduct.ToProductDomainModel(), nil
}

// GetProducts ...
func (s *Session) GetProducts(ctx context.Context, productsFilter *products.ProductsFilter) ([]*products.Product, error) {
	filter := bson.M{"soft_delete": false}
	findOptions := options.Find().
		SetSkip(productsFilter.Offset).
		SetLimit(productsFilter.Limit)
	findResult, findErr := s.storage.productsCollection.Find(ctx, filter, findOptions)
	if findErr != nil {
		return nil, findErr
	}

	productsDBResult := make([]*models.Product, productsFilter.Limit)
	decodeErr := findResult.All(ctx, &productsDBResult)
	if decodeErr != nil {
		return nil, decodeErr
	}

	products := make([]*products.Product, len(productsDBResult))
	for idx := range productsDBResult {
		products[idx] = productsDBResult[idx].ToProductDomainModel()
	}

	return products, nil
}

// UpdateProduct ...
func (s *Session) UpdateProduct(ctx context.Context, productId string, productUpdateDocument *products.ProductUpdate) (*products.Product, error) {
	filter := bson.M{"_id": productId}
	update := s.prepareProductUpdateBSON(productUpdateDocument)
	findOneAndUpdateOptions := options.FindOneAndUpdate().SetReturnDocument(options.After)
	findOneAndUpdateResult := s.storage.productsCollection.FindOneAndUpdate(ctx, filter, update, findOneAndUpdateOptions)
	if findOneAndUpdateResult.Err() != nil {
		return nil, findOneAndUpdateResult.Err()
	}

	var productDB models.Product
	decodeErr := findOneAndUpdateResult.Decode(&productDB)
	if decodeErr != nil {
		return nil, decodeErr
	}

	return productDB.ToProductDomainModel(), nil
}

// DeleteProduct ...
func (s *Session) DeleteProduct(ctx context.Context, productId string) error {
	update := bson.M{
		"$set": bson.M{"soft_delete": true},
	}
	_, updateByIDErr := s.storage.productsCollection.UpdateByID(ctx, productId, update)
	if updateByIDErr != nil {
		return updateByIDErr
	}

	return nil
}

func (s *Session) prepareProductUpdateBSON(productUpdateDocument *products.ProductUpdate) bson.M {
	fields := bson.M{}

	setIfNotNil(&fields, "name", productUpdateDocument.Name)
	setIfNotNil(&fields, "description", productUpdateDocument.Description)

	return bson.M{"$set": fields}
}
