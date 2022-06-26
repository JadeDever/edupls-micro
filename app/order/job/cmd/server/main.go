/*
 * @Author: Jadedever
 * @Date: 2022-06-26 21:09:10
 * @LastEditors: Jadedever
 * @LastEditTime: 2022-06-26 22:21:12
 * @FilePath: /edupls-micro/app/job/service/cmd/server/main.go
 * @Description:
 *
 * Copyright (c) 2022 by Jadedever, All Rights Reserved.
 */
package main

import (
	"flag"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/goxiaoy/go-saas-kit/sys/private/conf"
)

var (
	flagconf string
	id, _    = os.Hostname()
	Name     = "edupls.service.job"

	// go build -ldflags "-X main.Version=x.y.z"
	Version string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path,eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
		),
	)
}

func main() {
	flag.Parse()

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()
	if err := c.Load(); err != nil {
		panic(err)
	}
	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := initApp(bc.Data, bc.Server, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	if err := app.Run(); err != nil {
		panic(err)
	}

}
