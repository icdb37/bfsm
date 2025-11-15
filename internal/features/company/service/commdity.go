package service

import (
	"context"
	"fmt"

	"github.com/icdb37/bfsm/internal/constx/field"
	"github.com/icdb37/bfsm/internal/features/company/model"
	"github.com/icdb37/bfsm/internal/infra/cfpx"
	"github.com/icdb37/bfsm/internal/infra/errx"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
	coModel "github.com/icdb37/bfsm/internal/model"
	"github.com/icdb37/bfsm/internal/utils"
)

// 企业下的商品服务接口

type commodityImpl struct {
	repo store.Tabler
}

// Search 搜索商品
func (c *commodityImpl) Search(ctx context.Context, req *coModel.SearchRequest[model.QueryCommodity]) (resp *coModel.SearchResponse[model.EntireCommodity], err error) {
	qf := store.Unmarshal(req.Query)
	resp = &coModel.SearchResponse[model.EntireCommodity]{}
	pf := req.GetPage()
	if resp.Total, err = c.repo.Search(ctx, qf, pf, &(resp.Data)); err != nil {
		logx.Error("search commodity failed", "error", err)
		return nil, err
	}
	return resp, nil
}

// Get 获取商品
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

// Create 创建商品
func (c *commodityImpl) Create(ctx context.Context, infos []*model.EntireCommodity) error {
	logx.Info("create commodity", "infos", infos)
	for _, info := range infos {
		if err := utils.ProcessAll(ctx, info, cfpx.ProcessCreate); err != nil {
			logx.Error("create commodity failed", "error", err)
			return err
		}
	}
	hashs := []string{}
	for _, info := range infos {
		hashs = append(hashs, info.Hash)
	}
	where := store.NewFilter().
		Eq(field.CompanyID, infos[0].CompanyID).
		In(field.Hash, hashs)
	dupInfos := []*model.EntireCommodity{}
	if err := c.repo.Query(ctx, where, &dupInfos); err != nil {
		logx.Error("create commodity failed", "error", err)
		return err
	}
	if len(dupInfos) > 0 {
		logx.Error("create commodity failed", "error", "commodity already exist", "hashs", hashs)
		return errx.NewMessage("商品已存在：%v",
			utils.Converts(dupInfos,
				func(info *model.EntireCommodity) string {
					return fmt.Sprintf("%s-%s-%s", info.Name, info.Spec, info.Size)
				},
			))
	}
	if err := c.repo.Insert(ctx, utils.Converts(infos, func(v *model.EntireCommodity) any { return v })...); err != nil {
		logx.Error("create commodity failed", "error", err)
		return err
	}
	return nil
}

// Update 更新商品
func (c *commodityImpl) Update(ctx context.Context, info *model.EntireCommodity) error {
	logx.Info("update commodity", "info", info)
	if err := utils.ProcessAll(ctx, info, cfpx.ProcessUpdate); err != nil {
		logx.Error("update commodity failed", "error", err)
		return err
	}
	oldInfo := &model.EntireCommodity{}
	where := store.NewFilter().Eq(field.ID, info.ID)
	if err := c.repo.Query(ctx, where, oldInfo); err != nil {
		logx.Error("update commodity failed", "error", err)
		return err
	}
	if oldInfo.ID == "" {
		logx.Error("update commodity failed", "error", "commodity not found", "id", info.ID)
		return errx.NewNexist("商品不存在")
	}
	info.CompanyID, info.CreatedAt = oldInfo.CompanyID, oldInfo.CreatedAt
	if err := c.repo.Update(ctx, where, info); err != nil {
		logx.Error("update commodity failed", "error", err)
		return err
	}
	return nil
}

// Delete 删除商品
func (c *commodityImpl) Delete(ctx context.Context, id string) error {
	logx.Info("delete commodity", "id", id)
	where := store.NewFilter().
		Eq(field.ID, id)
	if err := c.repo.Delete(ctx, where); err != nil {
		logx.Error("delete commodity failed", "error", err)
		return err
	}
	return nil
}
