package service

import (
	"context"

	"github.com/icdb37/bfsm/internal/constx/field"
	"github.com/icdb37/bfsm/internal/features/company/model"
	"github.com/icdb37/bfsm/internal/infra/errx"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
	coModel "github.com/icdb37/bfsm/internal/model"
	"github.com/icdb37/bfsm/internal/utils"
)

type companyImpl struct {
	repo store.Tabler
}

func (c *companyImpl) Search(ctx context.Context, req *coModel.SearchRequest[model.QueryCompany]) (resp *coModel.SearchResponse[model.SimpleCompany], err error) {
	qf := store.Unmarshal(req.Query)
	resp = &coModel.SearchResponse[model.SimpleCompany]{}
	pf := req.GetPage()
	if resp.Total, err = c.repo.Search(ctx, qf, pf, &(resp.Datas)); err != nil {
		logx.Error("search companies failed", "error", err)
		return nil, err
	}
	return resp, nil
}
func (c *companyImpl) Get(ctx context.Context, id string) (*model.EntireCompany, error) {
	info := &model.EntireCompany{}
	if err := c.repo.Query(ctx, store.NewFilter().Eq(field.ID, id), info); err != nil {
		logx.Error("get company failed", "error", err)
		return nil, err
	}
	if info.ID == "" {
		logx.Error("get company failed", "error", "company not found", "id", id)
		return nil, errx.NewNexist("企业不存在")
	}
	return info, nil
}

func (c *companyImpl) Create(ctx context.Context, info *model.EntireCompany) error {
	if err := utils.ProcessAll(ctx, info, processCompanyCreate); err != nil {
		logx.Error("create company failed", "error", err)
		return err
	}
	if err := c.repo.Insert(ctx, info); err != nil {
		logx.Error("create company failed", "error", err)
		return err
	}
	return nil
}

func (c *companyImpl) Update(ctx context.Context, info *model.EntireCompany) error {
	if err := utils.ProcessAll(ctx, info, processCompanyUpdate); err != nil {
		logx.Error("update company failed", "error", err)
		return err
	}
	where := store.NewFilter().Eq(field.ID, info.ID)
	if err := c.repo.Update(ctx, where, info); err != nil {
		logx.Error("update company failed", "error", err)
		return err
	}
	return nil
}

func (c *companyImpl) Delete(ctx context.Context, id string) error {
	if err := c.repo.Delete(ctx, store.NewFilter().Eq(field.ID, id)); err != nil {
		logx.Error("delete company failed", "error", err)
		return err
	}
	return nil
}
