package storage

import (
	"context"
)

// IStorage ...
type IStorage interface {
	Open(functionName string) ISession
	Disconnect(ctx context.Context) error
	Validate() error
}

// ISession ...
type ISession interface {
	Close(ctx context.Context)
	StartSession() error

	ISessionProductsStorage
	ISessionCartsStorage
	ISessionPaymentsStorage
}
