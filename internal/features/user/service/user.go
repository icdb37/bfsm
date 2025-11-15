// Package service 用户服务
package service

import (
	"context"

	"github.com/icdb37/bfsm/internal/constx/field"
	"github.com/icdb37/bfsm/internal/features/user/model"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
	"github.com/icdb37/bfsm/internal/utils"
)

type UserServer struct {
	repo store.Tabler
}

// 搜索用户
func (u *UserServer) SearchUser(ctx context.Context, req *model.SearchRequest) (resp *model.SearchResponse, err error) {
	qf := store.Unmarshal(req.Query)
	resp = &model.SearchResponse{}
	pf := req.GetPage()
	if resp.Total, err = u.repo.Search(ctx, qf, pf, &(resp.Data)); err != nil {
		logx.Error("search users failed", "error", err)
		return nil, err
	}
	return resp, nil
}

// 创建用户
func (u *UserServer) CreateUser(ctx context.Context, info *model.EntireUser) error {
	if err := utils.ProcessAll(ctx, info, processCheck, processCreate); err != nil {
		logx.Error("create user failed", "error", err)
		return err
	}
	if err := u.repo.Insert(ctx, info); err != nil {
		logx.Error("create user failed", "error", err)
		return err
	}
	return nil
}

// 修改用户
func (u *UserServer) UpdateUser(ctx context.Context, user *model.EntireUser) error {
	if err := utils.ProcessAll(ctx, user, processCheck, processUpdate); err != nil {
		logx.Error("update user failed", "error", err)
		return err
	}
	if err := u.repo.Upsert(ctx, store.NewFilter().Eq(field.ID, user.ID), user); err != nil {
		logx.Error("update user failed", "error", err)
		return err
	}
	return nil
}

// 删除用户
func (u *UserServer) DeleteUser(ctx context.Context, id string) error {
	if err := u.repo.Delete(ctx, store.NewFilter().Eq(field.ID, id)); err != nil {
		logx.Error("delete user failed", "error", err)
		return err
	}
	return nil
}

// 获取用户
func (u *UserServer) GetUser(ctx context.Context, id string) (*model.EntireUser, error) {
	info := &model.EntireUser{}
	if err := u.repo.Query(ctx, store.NewFilter().Eq(field.ID, id), info); err != nil {
		logx.Error("get user failed", "error", err)
		return nil, err
	}
	return info, nil
}
