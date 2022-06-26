/*
 * @Author: Jadedever
 * @Date: 2022-06-17 23:52:17
 * @LastEditors: Jadedever
 * @LastEditTime: 2022-06-26 19:23:57
 * @FilePath: /edupls-micro/app/user/service/internal/data/user.go
 * @Description:
 *
 * Copyright (c) 2022 by Jadedever, All Rights Reserved.
 */
package data

import (
	"context"
	"edupls/app/user/service/internal/biz"
	"edupls/app/user/service/internal/model"
	"edupls/app/user/service/internal/pkg/util/pwd_md5"
	"edupls/pkg/errors"
	"edupls/pkg/util/pagination"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*UserRepo)(nil)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}

func (r *UserRepo) Create(ctx context.Context, b *model.User) (*model.User, error) {
	user := &model.User{Name: b.Name, Age: b.Age, Mobile: b.Mobile, Pass: pwd_md5.Base64Md5(b.Pass)}
	err := r.data.db.WithContext(ctx).Create(user).First(user).Error
	if err != nil {
		r.log.Errorf("[data.Create] err : %#v", err)
		return &model.User{}, errors.UnknownError
	}
	return user, nil
}

func (r *UserRepo) Get(ctx context.Context, id int64) (*model.User, error) {
	user := model.User{}
	err := r.data.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return &model.User{}, errors.RecordNotFound
	}
	return &user, nil
}

func (r *UserRepo) Update(ctx context.Context, b *model.User) (*model.User, error) {
	user := model.User{}
	if err := r.data.db.Updates(b).First(&user).Error; err != nil {
		return &model.User{}, errors.UnknownError
	}
	return &user, nil
}

func (r *UserRepo) Delete(ctx context.Context, id int64) (*model.User, error) {
	user := model.User{}
	user.ID = uint(id)
	result := r.data.db.WithContext(ctx).First(&user).Delete(&user, id).Error
	if result != nil {
		return &model.User{}, errors.UnknownError
	}
	return &user, nil
}

func (r *UserRepo) List(ctx context.Context, pageNum, pageSize int64) ([]*model.User, error) {
	var userList []*model.User
	result := r.data.db.WithContext(ctx).
		Limit(int(pageSize)).
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Find(&userList)
	if result.Error != nil {
		return nil, errors.UnknownError
	}

	return userList, nil
}

func (r *UserRepo) GetUserByMobile(ctx context.Context, mobile string) (user *model.User, err error) {
	if err = r.data.db.WithContext(ctx).Where("mobile = ?", mobile).First(&user).Error; err != nil {
		return &model.User{}, errors.RecordNotFound
	}

	return
}
