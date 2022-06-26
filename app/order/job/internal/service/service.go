/*
 * @Author: Jadedever
 * @Date: 2022-06-26 21:09:11
 * @LastEditors: Jadedever
 * @LastEditTime: 2022-06-26 21:48:12
 * @FilePath: /edupls-micro/app/job/service/internal/service/service.go
 * @Description:
 *
 * Copyright (c) 2022 by Jadedever, All Rights Reserved.
 */
package service

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewJobService)
