/*
 * @Author: Jadedever
 * @Date: 2022-06-26 21:09:38
 * @LastEditors: Jadedever
 * @LastEditTime: 2022-06-26 21:28:08
 * @FilePath: /edupls-micro/app/task/service/internal/server/server.go
 * @Description:
 *
 * Copyright (c) 2022 by Jadedever, All Rights Reserved.
 */
package server

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewGRPCServer)
