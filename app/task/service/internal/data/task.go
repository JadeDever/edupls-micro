/*
 * @Author: Jadedever
 * @Date: 2022-06-26 21:27:39
 * @LastEditors: Jadedever
 * @LastEditTime: 2022-06-26 21:36:07
 * @FilePath: /edupls-micro/app/task/service/internal/data/task.go
 * @Description:
 *
 * Copyright (c) 2022 by Jadedever, All Rights Reserved.
 */
package data

import (
	"context"
	"edupls/app/task/service/internal/biz"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var _ biz.TaskRepo = (*taskRepo)(nil)

// PO
type IntegratingCount struct {
	gorm.Model
	Id       int64
	UserId   int64
	Grade    int64
	CreateAt time.Time
}

type Integrating struct {
	gorm.Model
	Id       int64
	UserId   int64
	OrderId  int64
	Grade    int64
	CreateAt time.Time
}

// 实现 biz 定义的 AlbumRepo 接口
type taskRepo struct {
	data *Data
	log  *log.Helper
}

func NewTaskRepo(data *Data, logger log.Logger) biz.TaskRepo {
	return &taskRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/task")),
	}
}

func (t *taskRepo) IntegratingCount(ctx context.Context) error {

	var integratings = make([]Integrating, 0)

	var rst []struct {
		UserId int64
		Grade  int64
	}

	dbRst := t.data.db.WithContext(ctx).Raw("select user_id,sum(grade) as grade from integrating group by user_id").Scan(&rst)
	if dbRst.Error != nil {
		return dbRst.Error
	}

	for _, r := range rst {
		integratings = append(integratings, Integrating{
			UserId:   r.UserId,
			Grade:    r.Grade,
			CreateAt: time.Now(),
		})
	}
	t.data.db.WithContext(ctx).Create(&integratings)
	return nil
}

func (t *taskRepo) GetIntegrating(ctx context.Context, userId int64) (int64, error) {
	var cnt IntegratingCount
	cnt.UserId = userId
	result := t.data.db.WithContext(ctx).First(&cnt)
	if result.Error != nil {
		return 0, result.Error
	}
	return cnt.Grade, nil
}
