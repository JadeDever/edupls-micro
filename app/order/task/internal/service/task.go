/*
 * @Author: Jadedever
 * @Date: 2022-06-26 21:09:38
 * @LastEditors: Jadedever
 * @LastEditTime: 2022-06-26 21:28:43
 * @FilePath: /edupls-micro/app/task/service/internal/service/task.go
 * @Description:
 *
 * Copyright (c) 2022 by Jadedever, All Rights Reserved.
 */
package service

import (
	"context"
	v1 "edupls/api/task/service/v1"
	"edupls/app/task/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type TaskService struct {
	v1.UnimplementedTaskServer

	ac  *biz.TaskUseCase
	log *log.Helper
}

func NewTaskService(ac *biz.TaskUseCase, logger log.Logger) *TaskService {
	return &TaskService{
		ac:  ac,
		log: log.NewHelper(log.With(logger, "module", "service/task")),
	}

}

func (s *TaskService) IntegratingCount(ctx context.Context, req *v1.IntegratingCountReq) (*v1.IntegratingCountReply, error) {
	return &v1.IntegratingCountReply{}, s.ac.IntegratingCount(ctx)
}

func (s *TaskService) GetIntegrating(ctx context.Context, req *v1.GetIntegratingReq) (*v1.GetIntegratingReply, error) {
	rst, err := s.ac.GetIntegrating(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &v1.GetIntegratingReply{
		Grade: rst,
	}, nil
}
