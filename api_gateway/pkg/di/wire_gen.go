// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"X-TENTIONCREW/api_gateway/pkg/api"
	"X-TENTIONCREW/api_gateway/pkg/api/handlers"
	"X-TENTIONCREW/api_gateway/pkg/client"
	"X-TENTIONCREW/api_gateway/pkg/config"
	"X-TENTIONCREW/api_gateway/pkg/service"
)

// Injectors from wire.go:

func InitializeAPI(c *config.Config) (*api.Server, error) {
	clients, err := service.InitClient(c)
	if err != nil {
		return nil, err
	}
	authClient := client.NewauthClient(clients)
	authHandler := handlers.NewUserHandler(authClient)
	server, err := api.NewServerHTTP(c, authHandler)
	if err != nil {
		return nil, err
	}
	return server, nil
}