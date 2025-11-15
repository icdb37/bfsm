package service

import (
	"context"

	"github.com/icdb37/bfsm/internal/constx/featc"
	"github.com/icdb37/bfsm/internal/features/company/model"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
	coModel "github.com/icdb37/bfsm/internal/model"
	"github.com/icdb37/bfsm/internal/wire"
)

// CompanyServer - 企业服务接口
type CompanyServer interface {
	SelectAll(ctx context.Context) (resp []*model.SimpleCompany, err error)
	Search(ctx context.Context, req *coModel.SearchRequest[model.QueryCompany]) (resp *coModel.SearchResponse[model.EntireCompany], err error)
	Create(ctx context.Context, info *model.EntireCompany) error
	Update(ctx context.Context, info *model.EntireCompany) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*model.EntireCompany, error)
}

// CommodityServer - 商品服务接口
type CommodityServer interface {
	Search(ctx context.Context, req *coModel.SearchRequest[model.QueryCommodity]) (resp *coModel.SearchResponse[model.EntireCommodity], err error)
	Create(ctx context.Context, infos []*model.EntireCommodity) error
	Update(ctx context.Context, info *model.EntireCommodity) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*model.EntireCommodity, error)
}

func Provide() {
	wire.ProvideName(featc.CompanyCompany, func() CompanyServer {
		repo, err := store.NewTable(&model.EntireCompany{})
		if err != nil {
			logx.Fatal("create company repo failed", "error", err)
		}
		return &companyImpl{repo: repo}
	})
	wire.ProvideName(featc.CompanyCommodity, func() CommodityServer {
		repo, err := store.NewTable(&model.EntireCommodity{})
		if err != nil {
			logx.Fatal("create commodity repo failed", "error", err)
		}
		return &commodityImpl{repo: repo}
	})
}
