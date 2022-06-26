/*
 * @Author: Jadedever
 * @Date: 2022-06-17 23:52:17
 * @LastEditors: Jadedever
 * @LastEditTime: 2022-06-26 19:22:57
 * @FilePath: /edupls-micro/app/user/service/internal/data/data.go
 * @Description:
 *
 * Copyright (c) 2022 by Jadedever, All Rights Reserved.
 */
package data

import (
	"edupls/app/user/service/internal/conf"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/gorm"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRd, NewUserRepo)

// Data .
type Data struct {
	rd  *redis.Client
	db  *gorm.DB
	log *log.Helper
}

// NewData .
func NewData(db *gorm.DB, rd *redis.Client, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "user-service/data"))

	d := &Data{
		rd:  rd,
		db:  db,
		log: log,
	}

	// 监听配置文件并处理
	go func() {
		for v := range conf.ConfCh {
			fmt.Println("初始化自定义配置文件：", v.EduplsConf)
		}
	}()

	// 启动定时任务
	go InitTimer(*d)

	return d, func() {

	}, nil
}
