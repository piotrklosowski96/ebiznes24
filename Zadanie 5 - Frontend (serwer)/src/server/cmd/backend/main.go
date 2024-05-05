package main

import (
	"context"
	"log"

	"github.com/go-openapi/loads"

	"Backend/internal/configuration"
	"Backend/internal/router"
	"Backend/internal/storage"
	"Backend/openapi/gen/backend/server"
	"Backend/openapi/gen/backend/server/operations"
)

const servicePort = 8080

func main() {
	cfg, err := configuration.NewConfiguration()
	if err != nil {
		log.Panicf("Could not start service due to load configuration error: %s", err.Error())
	}

	storage := storage.NewStorage(cfg)
	defer storage.Disconnect(context.Background())

	swaggerSpec, err := loads.Analyzed(server.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	openapi := operations.NewBackendAPI(swaggerSpec)
	server := server.NewServer(openapi)
	defer server.Shutdown()

	server.Port = servicePort

	router := router.NewRouter(openapi, storage)
	router.RegisterRoutes()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
