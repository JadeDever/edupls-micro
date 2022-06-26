/*
 * @Author: Jadedever
 * @Date: 2022-06-26 21:09:11
 * @LastEditors: Jadedever
 * @LastEditTime: 2022-06-26 21:47:57
 * @FilePath: /edupls-micro/app/job/service/internal/server/server.go
 * @Description:
 *
 * Copyright (c) 2022 by Jadedever, All Rights Reserved.
 */
package server

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewGRPCServer)
