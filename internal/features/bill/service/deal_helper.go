package service

import (
	"context"

	"github.com/icdb37/bfsm/internal/constx/enum"
	"github.com/icdb37/bfsm/internal/constx/field"
	"github.com/icdb37/bfsm/internal/features/bill/model"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
)

func (d *dealImpl) processUpdate(ctx context.Context, newInfo *model.BillDeal) (err error) {
	oldInfo := &model.BillDeal{}
	if err := d.repoDeal.Query(ctx, store.NewFilter().Eq(field.ID, newInfo.ID), oldInfo); err != nil {
		logx.Error("get bill deal failed", "id", newInfo.ID, "error", err)
		return err
	}
	// 禁止修改字段
	newInfo.RefBatch = oldInfo.RefBatch
	newInfo.CreatedAt = oldInfo.CreatedAt
	newInfo.AmountLeft = newInfo.AmountTotal - newInfo.AmountClear
	if newInfo.AmountLeft <= 0 && oldInfo.AmountLeft > 0 {
		newInfo.AmountLeft = 0
		newInfo.AmountStatus = enum.AmountStatusPaid
		newInfo.ClearedAt = newInfo.UpdatedAt
	}
	return nil
}
