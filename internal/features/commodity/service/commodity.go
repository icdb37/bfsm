package service

import (
	"context"

	"github.com/icdb37/bfsm/internal/constx/field"
	"github.com/icdb37/bfsm/internal/features/commodity/model"
	"github.com/icdb37/bfsm/internal/infra/cfpx"
	"github.com/icdb37/bfsm/internal/infra/errx"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
	coModel "github.com/icdb37/bfsm/internal/model"
	"github.com/icdb37/bfsm/internal/utils"
)

type commodityImpl struct {
	repo store.Tabler
}

// Search 商品列表
func (c *commodityImpl) Search(ctx context.Context, req *coModel.SearchRequest[model.QueryCommodity]) (resp *coModel.SearchResponse[model.EntireCommodity], err error) {
	qf := store.Unmarshal(req.Query)
	resp = &coModel.SearchResponse[model.EntireCommodity]{}
	pf := req.GetPage()
	if resp.Total, err = c.repo.Search(ctx, qf, pf, &(resp.Datas)); err != nil {
		logx.Error("search commodity failed", "error", err)
		return nil, err
	}
	return resp, nil
}

// Get 商品详情
func (c *commodityImpl) Get(ctx context.Context, id string) (*model.EntireCommodity, error) {
	info := &model.EntireCommodity{}
	if err := c.repo.Query(ctx, store.NewFilter().Eq(field.ID, id), info); err != nil {
		logx.Error("get commodity failed", "error", err)
		return nil, err
	}
	if info.ID == "" {
		logx.Error("get commodity failed", "error", "commodity not found", "id", id)
		return nil, errx.NewNexist("商品不存在")
	}
	return info, nil
}

// Create 商品创建
func (c *commodityImpl) Create(ctx context.Context, info *model.EntireCommodity) error {
	logx.Info("create commodity", "info", info)
	if err := utils.ProcessAll(ctx, info, cfpx.ProcessCreate); err != nil {
		logx.Error("create commodity failed", "error", err)
		return err
	}
	if err := c.repo.Insert(ctx, info); err != nil {
		logx.Error("create commodity failed", "error", err)
		return err
	}
	return nil
}

// Update 商品更新
func (c *commodityImpl) Update(ctx context.Context, info *model.EntireCommodity) error {
	if err := utils.ProcessAll(ctx, info, cfpx.ProcessUpdate); err != nil {
		logx.Error("update commodity failed", "error", err)
		return err
	}
	where := store.NewFilter().Eq(field.Hash, info.Hash)
	if err := c.repo.Update(ctx, where, info); err != nil {
		logx.Error("update commodity failed", "error", err)
		return err
	}
	return nil
}

// Delete 商品删除
func (c *commodityImpl) Delete(ctx context.Context, id string) error {
	if err := c.repo.Delete(ctx, store.NewFilter().Eq(field.ID, id)); err != nil {
		logx.Error("delete company failed", "error", err)
		return err
	}
	return nil
}
