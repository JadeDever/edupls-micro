package service

import (
	"context"

	pb "edupls/api/user/service/v1"
	"edupls/app/user/service/internal/model"
	"edupls/pkg/errors"
)

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	// 数据校验
	if req.NickName == "" || req.Mobile == "" {
		return &pb.CreateUserReply{}, errors.InvalidParams
	}
	// 调用业务用例
	return s.uc.CreateUser(ctx, &model.User{Mobile: req.Mobile, Pass: req.Pass, Name: req.NickName, Age: req.Age})
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	// 数据校验
	if req.Id == 0 {
		return &pb.GetUserReply{}, errors.InvalidParams
	}
	// 调用业务用例
	return s.uc.GetUser(ctx, req.Id)
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	// 数据校验
	if req.Id == 0 {
		return &pb.UpdateUserReply{}, errors.InvalidParams
	}
	// 调用业务用例
	return s.uc.UpdateUser(ctx, req)
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	// 数据校验
	if req.Id == 0 {
		return &pb.DeleteUserReply{}, errors.InvalidParams
	}
	// 调用业务用例
	return s.uc.DeleteUser(ctx, req.Id)
}

func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	// 数据校验
	if req.Page == 0 || req.Limit == 0 {
		return &pb.ListUserReply{}, errors.InvalidParams
	}
	// 调用业务用例
	return s.uc.UserList(ctx, req.Page, req.Limit)
}

func (s *UserService) GetToken(ctx context.Context, req *pb.GetTokenRequest) (*pb.GetTokenReply, error) {
	// 数据校验
	if req.Mobile == "" || req.Pass == "" {
		return &pb.GetTokenReply{}, errors.InvalidParams
	}
	// 调用业务用例
	return s.uc.Login(ctx, req)
}
