// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/api"
	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/api/service"
	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/client"
	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/config"
	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/db"
	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/repository"
	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	cartRepository := repository.NewCartRepository(gormDB)
	cartUseCase := usecase.NewCartUseCase(cartRepository)
	productServiceClient := client.InitProductServiceClient(cfg)
	cartService := service.NewCartService(cartUseCase, productServiceClient)
	serverHTTP := http.NewServerHTTP(cartService)
	return serverHTTP, nil
}
