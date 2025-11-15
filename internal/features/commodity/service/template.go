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

type templateImpl struct {
	repo store.Tabler
}

// Search 商品模板列表
func (t *templateImpl) Search(ctx context.Context, req *coModel.SearchRequest[model.QueryTemplate]) (resp *coModel.SearchResponse[model.EntireTemplate], err error) {
	qf := store.Unmarshal(req.Query)
	resp = &coModel.SearchResponse[model.EntireTemplate]{}
	pf := req.GetPage()
	if resp.Total, err = t.repo.Search(ctx, qf, pf, &(resp.Data)); err != nil {
		logx.Error("search template commodity failed", "error", err)
		return nil, err
	}
	return resp, nil
}

// Get 商品模板详情
func (t *templateImpl) Get(ctx context.Context, id string) (*model.EntireTemplate, error) {
	info := &model.EntireTemplate{}
	if err := t.repo.Query(ctx, store.NewFilter().Eq(field.ID, id), info); err != nil {
		logx.Error("get template failed", "error", err)
		return nil, err
	}
	if info.ID == "" {
		logx.Error("get template failed", "error", "template not found", "id", id)
		return nil, errx.NewNexist("商品不存在")
	}
	return info, nil
}

// Create 商品模板创建
func (t *templateImpl) Create(ctx context.Context, info *model.EntireTemplate) error {
	logx.Info("create commodity", "info", info)
	if err := utils.ProcessAll(ctx, info, cfpx.ProcessCreate); err != nil {
		logx.Error("create template failed", "error", err)
		return err
	}
	if err := t.repo.Insert(ctx, info); err != nil {
		logx.Error("create template failed", "error", err)
		return err
	}
	return nil
}

// Update 商品模板更新
func (t *templateImpl) Update(ctx context.Context, info *model.EntireTemplate) error {
	if err := utils.ProcessAll(ctx, info, cfpx.ProcessUpdate); err != nil {
		logx.Error("update template failed", "error", err)
		return err
	}
	where := store.NewFilter().Eq(field.ID, info.ID)
	if err := t.repo.Update(ctx, where, info); err != nil {
		logx.Error("update template failed", "error", err)
		return err
	}
	return nil
}

// Delete 商品模板删除
func (t *templateImpl) Delete(ctx context.Context, id string) error {
	if err := t.repo.Delete(ctx, store.NewFilter().Eq(field.ID, id)); err != nil {
		logx.Error("delete template failed", "error", err)
		return err
	}
	return nil
}
