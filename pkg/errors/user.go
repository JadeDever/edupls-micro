/*
 * @Author: Jadedever
 * @Date: 2022-06-17 23:52:17
 * @LastEditors: Jadedever
 * @LastEditTime: 2022-06-26 19:10:19
 * @FilePath: /edupls-micro/pkg/errors/user.go
 * @Description:
 *
 * Copyright (c) 2022 by Jadedever, All Rights Reserved.
 */
package errors

import "github.com/go-kratos/kratos/v2/errors"

var (
	UserNotExit = errors.New(400, "UserNotExit", "User Not Exit")
	ErrAuthFail = errors.New(401, "Authentication failed", "Missing token or token incorrect")
)
