/*
 * @Author: Jadedever
 * @Date: 2022-06-17 23:52:17
 * @LastEditors: Jadedever
 * @LastEditTime: 2022-06-26 19:26:23
 * @FilePath: /edupls-micro/pkg/util/datacopy/datacopy.go
 * @Description: 数据复制，专门用于请求体数据映射到DTO模型。
 * 原因：ShouldBind时需要限制字段，如果直接使用DTO模型接受请求数据则不好限制字段；
 * 又不希望每次都每个字段都显示赋值，所以封装一个基于json包的数据复制功能
 *
 * Copyright (c) 2022 by Jadedever, All Rights Reserved.
 */
package datacopy

import (
	"encoding/json"
	"reflect"
)

// DataCopy 本来想使用reflect包来完成这个功能，json.Marshal()其实也是借助reflect包来实现的.
func DataCopy(data, res interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, res)
}

// DataReflact 传入*reflect.Value 且返回 *reflect.Value ，将类型交给调用者来决定。
func DataReflact(src, dst *reflect.Value) *reflect.Value {
	for i := 0; i < src.NumField(); i++ {
		filed := src.Type().Field(i).Name
		cfiled := dst.FieldByName(filed)
		if cfiled.IsValid() && cfiled.CanSet() { // 可用的属性且可 赋值
			dst.Field(i).Set(src.Field(i))
		}
	}
	return dst
}
