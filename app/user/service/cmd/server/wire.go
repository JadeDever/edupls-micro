//go:build wireinject
// +build wireinject

/*
 * @Author: Jadedever
 * @Date: 2022-06-17 23:52:17
 * @LastEditors: Jadedever
 * @LastEditTime: 2022-06-26 19:17:28
 * @FilePath: /edupls-micro/app/user/service/cmd/server/wire.go
 * @Description:
 *
 * Copyright (c) 2022 by Jadedever, All Rights Reserved.
 */

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"edupls/app/user/service/internal/biz"
	"edupls/app/user/service/internal/conf"
	"edupls/app/user/service/internal/data"
	"edupls/app/user/service/internal/server"
	"edupls/app/user/service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
