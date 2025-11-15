package service

import (
	"context"

	"github.com/icdb37/bfsm/internal/constx/field"
	"github.com/icdb37/bfsm/internal/features/bill/model"
	"github.com/icdb37/bfsm/internal/infra/cfpx"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
	coModel "github.com/icdb37/bfsm/internal/model"
	"github.com/icdb37/bfsm/internal/utils"
)

type dealImpl struct {
	repoDeal store.Tabler
}

// Search 查询账单处理记录
func (d *dealImpl) Search(ctx context.Context, req *coModel.SearchRequest[model.QueryBillDeal]) (resp *coModel.SearchResponse[model.BillDeal], err error) {
	qf := store.Unmarshal(req.Query)
	resp = &coModel.SearchResponse[model.BillDeal]{}
	pf := req.GetPage()
	if resp.Total, err = d.repoDeal.Search(ctx, qf, pf, &(resp.Data)); err != nil {
		logx.Error("search bill deal failed", "error", err)
		return nil, err
	}
	return resp, nil
}

// Update 更新账单处理记录
func (d *dealImpl) Update(ctx context.Context, newInfo *model.BillDeal) error {
	logx.Info("update bill deal", "info", newInfo)
	if err := utils.ProcessAll(ctx, newInfo, cfpx.ProcessUpdate, d.processUpdate); err != nil {
		logx.Error("update bill deal failed", "error", err)
		return err
	}
	if err := d.repoDeal.Update(ctx, store.NewFilter().Eq(field.ID, newInfo.ID), newInfo); err != nil {
		logx.Error("update bill deal failed", "error", err)
		return err
	}
	return nil
}
