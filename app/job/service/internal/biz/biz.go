/*
 * @Author: Jadedever
 * @Date: 2022-06-26 21:09:11
 * @LastEditors: Jadedever
 * @LastEditTime: 2022-06-26 21:46:54
 * @FilePath: /edupls-micro/app/job/service/internal/biz/biz.go
 * @Description:
 *
 * Copyright (c) 2022 by Jadedever, All Rights Reserved.
 */
package biz

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewJobUseCase)
