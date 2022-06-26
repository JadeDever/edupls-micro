// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"edupls/app/user/service/internal/biz"
	"edupls/app/user/service/internal/conf"
	"edupls/app/user/service/internal/data"
	"edupls/app/user/service/internal/server"
	"edupls/app/user/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	db := data.NewDB(confData, logger)
	client := data.NewRd(confData, logger)
	dataData, cleanup, err := data.NewData(db, client, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUseCase := biz.NewUserUseCase(userRepo, logger)
	userService := service.NewUserService(userUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, logger, tracerProvider, userService)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
