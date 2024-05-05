package storage

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"Backend/internal/configuration"
)

const (
	productsCollectionName = "products"
	cartsCollectionName    = "carts"
	paymentsCollectionName = "payments"
)

// Storage ...
type Storage struct {
	client       *mongo.Client
	databaseName string

	productsCollection *mongo.Collection
	cartsCollection    *mongo.Collection
	paymentsCollection *mongo.Collection
}

// Session ...
type Session struct {
	session      mongo.Session
	storage      *Storage
	functionName string
}

// NewStorage ...
func NewStorage(cfg *configuration.Configuration) IStorage {
	mongoDBName := cfg.GetMongoDBName()
	client := getClient(cfg)

	return &Storage{
		client:       client,
		databaseName: mongoDBName,

		productsCollection: client.Database(mongoDBName).Collection(productsCollectionName),
		cartsCollection:    client.Database(mongoDBName).Collection(cartsCollectionName),
		paymentsCollection: client.Database(mongoDBName).Collection(paymentsCollectionName),
	}
}

// Open ...
func (s *Storage) Open(functionName string) ISession {
	return &Session{
		functionName: functionName,
		storage:      s,
	}
}

// Disconnect ...
func (s *Storage) Disconnect(ctx context.Context) error {
	return s.client.Disconnect(ctx)
}

// Close ...
func (s *Session) Close(ctx context.Context) {
	if s.session != nil {
		s.session.EndSession(ctx)
	}
}

// StartSession ...
func (s *Session) StartSession() error {
	var err error
	s.session, err = s.storage.client.StartSession()

	return err
}

// Validate ...
func (s *Storage) Validate() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return s.client.Ping(ctx, readpref.Primary())
}

func getClient(cfg *configuration.Configuration) *mongo.Client {
	uri := fmt.Sprintf("mongodb://%s:%s@%s",
		cfg.GetMongoDBUser(),
		cfg.GetMongoDBPass(),
		cfg.GetMongoDBHost(),
	)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	var result bson.M
	err = client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result)
	if err != nil {
		panic(err)
	}

	return client
}
