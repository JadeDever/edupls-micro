//go:build wireinject
// +build wireinject

/*
 * @Author: Jadedever
 * @Date: 2022-06-26 20:21:41
 * @LastEditors: Jadedever
 * @LastEditTime: 2022-06-26 20:21:43
 * @FilePath: /edupls-micro/app/admin/service/cmd/server/wire.go
 * @Description:
 *
 * Copyright (c) 2022 by Jadedever, All Rights Reserved.
 */

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"admin/internal/biz"
	"admin/internal/conf"
	"admin/internal/data"
	"admin/internal/server"
	"admin/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init admin application.
func initApp(*conf.Server, *conf.Data, *conf.Auth, *conf.Service, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
