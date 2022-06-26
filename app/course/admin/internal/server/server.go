/*
 * @Author: Jadedever
 * @Date: 2022-06-26 20:11:52
 * @LastEditors: Jadedever
 * @LastEditTime: 2022-06-26 20:26:15
 * @FilePath: /edupls-micro/app/admin/service/internal/server/server.go
 * @Description:
 *
 * Copyright (c) 2022 by Jadedever, All Rights Reserved.
 */
package server

import (
	"github.com/google/wire"
)

// ProviderSet is s providers.
var ProviderSet = wire.NewSet(NewHTTPServer)
