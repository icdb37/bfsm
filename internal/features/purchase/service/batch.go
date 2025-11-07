package service

import (
	"context"
	"time"

	"github.com/icdb37/bfsm/internal/constx/enum"
	"github.com/icdb37/bfsm/internal/constx/field"
	"github.com/icdb37/bfsm/internal/features/purchase/model"
	"github.com/icdb37/bfsm/internal/infra/cfpx"
	"github.com/icdb37/bfsm/internal/infra/errx"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
	coModel "github.com/icdb37/bfsm/internal/model"
	coService "github.com/icdb37/bfsm/internal/service"
	"github.com/icdb37/bfsm/internal/utils"
)

type batchImpl struct {
	repoBatch store.Tabler
	repoGoods store.Tabler
	inventory coService.InventorySaver
}

func (p *batchImpl) Search(ctx context.Context, req *coModel.SearchRequest[model.QueryPurchase]) (resp *coModel.SearchResponse[model.SimplePurchase], err error) {
	qf := store.Unmarshal(req.Query)
	resp = &coModel.SearchResponse[model.SimplePurchase]{}
	pf := req.GetPage()
	if resp.Total, err = p.repoBatch.Search(ctx, qf, pf, &(resp.Datas)); err != nil {
		logx.Error("search purchase batch failed", "error", err)
		return nil, err
	}
	return resp, nil
}

func (p *batchImpl) Get(ctx context.Context, id string) (*model.PurchaseBatch, error) {
	info := &model.PurchaseBatch{}
	if err := p.repoBatch.Query(ctx, store.NewFilter().Eq(field.ID, id), info); err != nil {
		logx.Error("get purchase batch failed", "error", err)
		return nil, err
	}
	if info.ID == "" {
		logx.Error("get purchase batch not found", "id", id)
		return nil, errx.NewNexist("采购订单不存在")
	}
	return info, nil
}

func (p *batchImpl) Create(ctx context.Context, info *model.PurchaseBatch) error {
	logx.Info("create purchase batch", "info", info)
	if err := utils.ProcessAll(ctx, info, cfpx.ProcessCreate); err != nil {
		logx.Error("create purchase batch failed", "error", err)
		return err
	}
	if err := p.repoBatch.Insert(ctx, info); err != nil {
		logx.Error("create purchase batch failed", "error", err)
		return err
	}
	return nil
}

func (p *batchImpl) Update(ctx context.Context, info *model.PurchaseBatch) error {
	if err := utils.ProcessAll(ctx, info, cfpx.ProcessUpdate); err != nil {
		logx.Error("update purchase batch failed", "error", err)
		return err
	}
	where := store.NewFilter().Eq(field.ID, info.ID)
	if err := p.repoBatch.Update(ctx, where, info); err != nil {
		logx.Error("update purchase batch failed", "error", err)
		return err
	}
	return nil
}

func (p *batchImpl) UpdateStatus(ctx context.Context, req *model.UpdateBatchStatus) error {
	info, err := p.Get(ctx, req.ID)
	if err != nil {
		logx.Error("complete purchase failed", "error", err)
		return err
	}
	if info.StatusCode == enum.StatusCodeCompleted {
		logx.Error("complete purchase batch already completed", "id", req.ID)
		return errx.NewErrStatus("采购订单已完成")
	}
	if info.StatusCode == req.StatusCode {
		logx.Error("complete purchase batch status not changed", "id", req.ID)
		return errx.NewErrStatus("采购订单状态未改变")
	}
	info.StatusCode = req.StatusCode
	info.UpdatedAt = time.Now()
	if err := p.repoBatch.Update(ctx, store.NewFilter().Eq(field.ID, info.ID), info); err != nil {
		logx.Error("update purchase batch failed", "error", err)
		return err
	}
	return p.saveInventory(ctx, info)
}

func (p *batchImpl) Delete(ctx context.Context, id string) error {
	if err := p.repoBatch.Delete(ctx, store.NewFilter().Eq(field.ID, id)); err != nil {
		logx.Error("delete purchase batch failed", "error", err)
		return err
	}
	return nil
}
