package service

import (
	"context"

	"github.com/icdb37/bfsm/internal/constx/featc"
	"github.com/icdb37/bfsm/internal/features/user/model"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
	"github.com/icdb37/bfsm/internal/wire"
)

// Server - 用户服务接口
type Server interface {
	SearchUser(ctx context.Context, req *model.SearchRequest) (resp *model.SearchResponse, err error)
	CreateUser(ctx context.Context, info *model.EntireUser) error
	UpdateUser(ctx context.Context, user *model.EntireUser) error
	DeleteUser(ctx context.Context, id string) error
	GetUser(ctx context.Context, id string) (*model.EntireUser, error)
}

// NewUser 创建用户服务
func New() (Server, error) {
	repo, err := store.NewTable(&model.EntireUser{})
	if err != nil {
		logx.Error("create user repo failed", "error", err)
		return nil, err
	}
	return &UserServer{repo: repo}, nil
}

func Provide() {
	repo, err := store.NewTable(&model.EntireUser{})
	if err != nil {
		logx.Fatal("create user repo failed", "error", err)
	}
	wire.ProvideName(featc.User, func() Server {
		return &UserServer{repo: repo}
	})
}
