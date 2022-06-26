//go:build wireinject
// +build wireinject

/*
 * @Author: Jadedever
 * @Date: 2022-06-26 21:45:25
 * @LastEditors: Jadedever
 * @LastEditTime: 2022-06-26 21:45:27
 * @FilePath: /edupls-micro/app/job/service/cmd/server/wire.go
 * @Description:
 *
 * Copyright (c) 2022 by Jadedever, All Rights Reserved.
 */

package main

import (
	"edupls/app/job/service/internal/biz"
	"edupls/app/job/service/internal/conf"
	"edupls/app/job/service/internal/data"
	"edupls/app/job/service/internal/server"
	"edupls/app/job/service/internal/service"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

func initApp(*conf.Data, *conf.Server, log.Logger) (*kratos.App, func(), error) {
	//	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet, server.ProviderSet, newApp))
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))

}
