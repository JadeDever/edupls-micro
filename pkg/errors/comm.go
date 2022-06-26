/*
 * @Author: Jadedever
 * @Date: 2022-06-17 23:52:17
 * @LastEditors: Jadedever
 * @LastEditTime: 2022-06-26 19:10:25
 * @FilePath: /edupls-micro/pkg/errors/comm.go
 * @Description:
 *
 * Copyright (c) 2022 by Jadedever, All Rights Reserved.
 */
package errors

import "github.com/go-kratos/kratos/v2/errors"

var (
	InvalidParams   = errors.New(400, "InvalidParams", "Missing Params")
	RecordNotFound  = errors.New(404, "RecordNotFound", "Record Not Found")
	UnknownError    = errors.New(500, "UnknownError", "Unknown Errors")
	MakeTokenFailed = errors.New(500, "MakeTokenFailed", "Make Token Faild")
)
