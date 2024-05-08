package storage

import (
	"context"
	"time"

	"Backend/internal/models/carts"
	"Backend/internal/storage/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateCart ...
func (s *Session) CreateCart(ctx context.Context, cartCreate *carts.CartCreate) (*carts.Cart, error) {
	cartDB := models.FromCartCreateDomainModel(cartCreate)
	cartDB.CreatedAt = time.Now()
	cartDB.UpdatedAt = time.Now()
	cartDB.SoftDelete = false

	_, insertOneErr := s.storage.cartsCollection.InsertOne(ctx, cartDB)
	if insertOneErr != nil {
		return nil, insertOneErr
	}

	return s.GetCartById(ctx, cartCreate.Id)
}

// GetCartById ...
func (s *Session) GetCartById(ctx context.Context, cartId string) (*carts.Cart, error) {
	aggregatePipeline := mongo.Pipeline{
		{{"$match", bson.M{
			"_id":         cartId,
			"soft_delete": false,
		}}},
		{{"$lookup", bson.M{
			"from":         productsCollectionName,
			"localField":   "products",
			"foreignField": "_id",
			"as":           "products",
		}}},
	}
	aggregateResult, aggregateErr := s.storage.cartsCollection.Aggregate(ctx, aggregatePipeline)
	if aggregateErr != nil {
		return nil, aggregateErr
	}

	var cartsResult []*models.Cart
	allErr := aggregateResult.All(ctx, &cartsResult)
	if allErr != nil {
		return nil, allErr
	}

	return cartsResult[0].ToCartDomainModel(), nil
}

// GetCarts ...
func (s *Session) GetCarts(ctx context.Context, cartsFilter *carts.CartsFilter) ([]*carts.Cart, error) {
	return nil, nil
}

// UpdateCart ...
func (s *Session) UpdateCart(ctx context.Context, cartId string, cartUpdateDocument *carts.CartUpdate) (*carts.Cart, error) {
	return nil, nil
}

// DeleteCart ...
func (s *Session) DeleteCart(ctx context.Context, cartId string) error {
	return nil
}

// AddProductToCart ...
func (s *Session) AddProductToCart(ctx context.Context, cartId string, productId string) (*carts.Cart, error) {
	update := bson.M{"$addToSet": bson.M{"products": productId}}
	_, updateByIDErr := s.storage.cartsCollection.UpdateByID(ctx, cartId, update)
	if updateByIDErr != nil {
		return nil, updateByIDErr
	}

	return s.GetCartById(ctx, cartId)
}

// RemoveProductFromCart ...
func (s *Session) RemoveProductFromCart(ctx context.Context, cartId string, productId string) (*carts.Cart, error) {
	update := bson.M{"$pull": bson.M{"products": productId}}
	_, updateByIDErr := s.storage.cartsCollection.UpdateByID(ctx, cartId, update)
	if updateByIDErr != nil {
		return nil, updateByIDErr
	}

	return s.GetCartById(ctx, cartId)
}
