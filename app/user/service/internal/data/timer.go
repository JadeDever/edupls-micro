/*
 * @Author: Jadedever
 * @Date: 2022-06-17 23:52:17
 * @LastEditors: Jadedever
 * @LastEditTime: 2022-06-26 19:23:25
 * @FilePath: /edupls-micro/app/user/service/internal/data/timer.go
 * @Description:
 *
 * Copyright (c) 2022 by Jadedever, All Rights Reserved.
 */
package data

import (
	"fmt"

	"github.com/robfig/cron"
)

// 定时任务

type ChangedShopLog struct {
	Data Data
}

func (c *ChangedShopLog) Run() {
	fmt.Println("定时任务:cron")
}

func InitTimer(data Data) {
	Corns := cron.New() // 定时任务
	Corns.Start()
	// 每五秒： */5 * * * * ?
	// 每隔1分钟执行一次："0 */1 * * * ?"
	// 每天23点执行一次："0 0 23 * * ?"
	// 每天凌晨1点执行一次："0 0 1 * * ?"
	// 每月1号凌晨1点执行一次："0 0 1 1 * ?"
	// 在26分、29分、33分执行一次："0 26,29,33 * * * ?"
	// 每天的0点、13点、18点、21点都执行一次："0 0 0,13,18,21 * * ?"

	Corns.AddJob("*/5 * * * * ?", &ChangedShopLog{
		Data: data,
	})
}
