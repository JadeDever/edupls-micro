/*
 * @Author: Jadedever
 * @Date: 2022-06-17 23:52:17
 * @LastEditors: Jadedever
 * @LastEditTime: 2022-06-26 19:26:07
 * @FilePath: /edupls-micro/pkg/util/snowflake/snowflake_test.go
 * @Description:
 *
 * Copyright (c) 2022 by Jadedever, All Rights Reserved.
 */
package snowflake

import (
	"fmt"
	"testing"
)

func TestNewSnow(t *testing.T) {
	fmt.Println(RandomUID()) // go test -v snowflake_test.go snowflake.go
}
