/*
 * @Author: Jadedever
 * @Date: 2022-06-17 23:52:17
 * @LastEditors: Jadedever
 * @LastEditTime: 2022-06-26 19:25:18
 * @FilePath: /edupls-micro/pkg/util/snowflake/snowflake.go
 * @Description:
 *
 * Copyright (c) 2022 by Jadedever, All Rights Reserved.
 */
package snowflake

import (
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node = nil

func init() {
	if nodes, err := snowflake.NewNode(1); err != nil {
		panic("snowflake init faild")
	} else {
		node = nodes
	}
}

// RandomUID 简易发号器
func RandomUID() int64 {
	return node.Generate().Int64()
}
