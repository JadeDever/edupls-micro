/*
 * @Author: Jadedever
 * @Date: 2022-06-26 21:09:11
 * @LastEditors: Jadedever
 * @LastEditTime: 2022-06-26 21:48:27
 * @FilePath: /edupls-micro/app/job/service/internal/service/job.go
 * @Description:
 *
 * Copyright (c) 2022 by Jadedever, All Rights Reserved.
 */
package service

import (
	"context"
	v1 "edupls/api/job/service/v1"
	"edupls/app/job/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type JobService struct {
	v1.UnimplementedJobServer

	ac  *biz.JobUseCase
	log *log.Helper
}

func NewJobService(ac *biz.JobUseCase, logger log.Logger) *JobService {
	return &JobService{
		ac:  ac,
		log: log.NewHelper(log.With(logger, "module", "service/job")),
	}

}

func (s *JobService) IntegratingJob(ctx context.Context, req *v1.IntegratingJobReq) (*v1.IntegratingJobReply, error) {
	return &v1.IntegratingJobReply{}, s.ac.IntegratingJob(ctx)
}
