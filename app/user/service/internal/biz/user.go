package biz

import (
	"context"
	user_proto "edupls/api/user/service/v1"
	"edupls/app/user/service/internal/model"
	"edupls/app/user/service/internal/pkg/util/pwd_md5"
	"edupls/pkg/errors"
	"edupls/pkg/util/token"
)

// ********* 以下实现业务组装，实现service需求 ***********
func (uc *UserUseCase) CreateUser(ctx context.Context, u *model.User) (*user_proto.CreateUserReply, error) {
	res, err := uc.repo.Create(ctx, u)
	if err != nil {
		return &user_proto.CreateUserReply{}, err
	}
	return &user_proto.CreateUserReply{
		Id:       int64(res.ID),
		Mobile:   res.Mobile,
		NickName: res.Name,
		Age:      res.Age,
	}, nil
}

func (uc *UserUseCase) GetUser(ctx context.Context, id int64) (*user_proto.GetUserReply, error) {
	res, err := uc.repo.Get(ctx, id)
	if err != nil {
		return &user_proto.GetUserReply{}, err
	}
	return &user_proto.GetUserReply{
		Id:       int64(res.ID),
		Mobile:   res.Mobile,
		NickName: res.Mobile,
		Age:      res.Age,
	}, nil
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, id int64) (*user_proto.DeleteUserReply, error) {
	_, err := uc.repo.Delete(ctx, id)
	if err != nil {
		return &user_proto.DeleteUserReply{
			Ok: false,
		}, nil
	}
	return &user_proto.DeleteUserReply{
		Ok: true,
	}, nil
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, req *user_proto.UpdateUserRequest) (*user_proto.UpdateUserReply, error) {
	var user model.User
	user.ID = uint(req.Id)
	user.Age = req.Age
	user.Name = req.NickName
	res, err := uc.repo.Update(ctx, &user)
	if err != nil {
		return &user_proto.UpdateUserReply{}, err
	}
	return &user_proto.UpdateUserReply{
		Mobile:   res.Mobile,
		NickName: res.Name,
		Age:      res.Age,
	}, nil
}

func (uc *UserUseCase) UserList(ctx context.Context, pageNum, pageSize int64) (*user_proto.ListUserReply, error) {
	var res []*user_proto.ListUserReply_User
	list, err := uc.repo.List(ctx, pageNum, pageSize)
	if err != nil {
		return &user_proto.ListUserReply{}, err
	}

	// 拼接数据
	for _, v := range list {
		res = append(res, &user_proto.ListUserReply_User{
			Id:       int64(v.ID),
			Mobile:   v.Mobile,
			NickName: v.Name,
		})
	}
	return &user_proto.ListUserReply{Users: res}, nil
}

func (uc *UserUseCase) Login(ctx context.Context, u *user_proto.GetTokenRequest) (res *user_proto.GetTokenReply, err error) {
	user, err := uc.repo.GetUserByMobile(ctx, u.Mobile)
	if err != nil {
		return res, err
	}

	if user.Pass != pwd_md5.Base64Md5(u.Pass) {
		return res, errors.InvalidParams
	}

	t, err := token.NewJWT().CreateToken(token.CustomClaims{
		ID: int(user.ID),
	})

	if err != nil {
		return res, errors.MakeTokenFailed
	}
	return &user_proto.GetTokenReply{Token: t}, nil
}
