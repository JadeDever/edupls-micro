/*
 * @Author: Jadedever
 * @Date: 2022-06-26 21:09:38
 * @LastEditors: Jadedever
 * @LastEditTime: 2022-06-26 21:26:20
 * @FilePath: /edupls-micro/app/task/service/internal/biz/biz.go
 * @Description:
 *
 * Copyright (c) 2022 by Jadedever, All Rights Reserved.
 */
package biz

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewTaskUseCase)
