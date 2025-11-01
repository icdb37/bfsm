package service

import (
	"context"

	"github.com/icdb37/bfsm/internal/constx/featc"
	"github.com/icdb37/bfsm/internal/features/commodity/model"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
	coModel "github.com/icdb37/bfsm/internal/model"
	"github.com/icdb37/bfsm/internal/wire"
)

// TemplateServer - 商品模板服务接口
type TemplateServer interface {
	Search(ctx context.Context, req *coModel.SearchRequest[model.QueryTemplate]) (resp *coModel.SearchResponse[model.EntireTemplate], err error)
	Create(ctx context.Context, info *model.EntireTemplate) error
	Update(ctx context.Context, info *model.EntireTemplate) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*model.EntireTemplate, error)
}

// CommodityServer - 商品服务接口
type CommodityServer interface {
	Search(ctx context.Context, req *coModel.SearchRequest[model.QueryCommodity]) (resp *coModel.SearchResponse[model.EntireCommodity], err error)
	Create(ctx context.Context, info *model.EntireCommodity) error
	Update(ctx context.Context, info *model.EntireCommodity) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*model.EntireCommodity, error)
}

func Provide() {
	wire.ProvideName(featc.CommodityCommodity, func() CommodityServer {
		repo, err := store.NewTable(&model.EntireCommodity{})
		if err != nil {
			logx.Fatal("create commodity repo failed", "error", err)
		}
		return &commodityImpl{repo: repo}
	})
	wire.ProvideName(featc.CommodityTemplate, func() TemplateServer {
		repo, err := store.NewTable(&model.EntireTemplate{})
		if err != nil {
			logx.Fatal("create company repo failed", "error", err)
		}
		return &templateImpl{repo: repo}
	})
}
