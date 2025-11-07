package service

import (
	"context"

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

type purchaseImpl struct {
	repo      store.Tabler
	inventory coService.InventoryProducer
}

func (p *purchaseImpl) Search(ctx context.Context, req *coModel.SearchRequest[model.QueryPurchase]) (resp *coModel.SearchResponse[model.EntirePurchase], err error) {
	qf := store.Unmarshal(req.Query)
	resp = &coModel.SearchResponse[model.EntirePurchase]{}
	pf := req.GetPage()
	if resp.Total, err = p.repo.Search(ctx, qf, pf, &(resp.Datas)); err != nil {
		logx.Error("search purchase failed", "error", err)
		return nil, err
	}
	return resp, nil
}

func (p *purchaseImpl) Get(ctx context.Context, id string) (*model.EntirePurchase, error) {
	info := &model.EntirePurchase{}
	if err := p.repo.Query(ctx, store.NewFilter().Eq(field.ID, id), info); err != nil {
		logx.Error("get purchase failed", "error", err)
		return nil, err
	}
	if info.PurchaseID == "" {
		logx.Error("get purchase failed", "error", "purchase not found", "id", id)
		return nil, errx.NewNexist("采购订单不存在")
	}
	return info, nil
}

func (p *purchaseImpl) Create(ctx context.Context, info *model.EntirePurchase) error {
	logx.Info("create purchase", "info", info)
	if err := utils.ProcessAll(ctx, info, cfpx.ProcessCreate); err != nil {
		logx.Error("create purchase failed", "error", err)
		return err
	}
	if err := p.repo.Insert(ctx, info); err != nil {
		logx.Error("create purchase failed", "error", err)
		return err
	}
	return nil
}

func (p *purchaseImpl) Update(ctx context.Context, info *model.EntirePurchase) error {
	if err := utils.ProcessAll(ctx, info, cfpx.ProcessUpdate); err != nil {
		logx.Error("update purchase failed", "error", err)
		return err
	}
	where := store.NewFilter().Eq(field.ID, info.PurchaseID)
	if err := p.repo.Update(ctx, where, info); err != nil {
		logx.Error("update purchase failed", "error", err)
		return err
	}
	return nil
}

func (p *purchaseImpl) Complete(ctx context.Context, id string) error {
	info, err := p.Get(ctx, id)
	if err != nil {
		logx.Error("complete purchase failed", "error", err)
		return err
	}
	if info.StatusCode == enum.StatusCodeCompleted {
		logx.Error("complete purchase failed", "error", "purchase already completed", "id", id)
		return errx.NewErrStatus("采购订单已完成")
	}
	return p.saveInventory(ctx, info)
}

func (p *purchaseImpl) Delete(ctx context.Context, id string) error {
	if err := p.repo.Delete(ctx, store.NewFilter().Eq(field.ID, id)); err != nil {
		logx.Error("delete purchase failed", "error", err)
		return err
	}
	return nil
}
