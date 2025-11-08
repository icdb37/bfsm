package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/icdb37/bfsm/internal/constx/enum"
	"github.com/icdb37/bfsm/internal/constx/field"
	"github.com/icdb37/bfsm/internal/features/bill/model"
	"github.com/icdb37/bfsm/internal/infra/errx"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
	coModel "github.com/icdb37/bfsm/internal/model"
)

type batchImpl struct {
	repoBatch store.Tabler
	repoDeal  store.Tabler
}

// Approve - 审核通过，交易账单细化
func (b *batchImpl) Approve(ctx context.Context, id string) error {
	info := &model.BillBatch{}
	if err := b.repoBatch.Query(ctx, store.NewFilter().Eq(field.ID, id), info); err != nil {
		logx.Error("query bill batch failed", "id", id, "error", err)
		return err
	}
	if info.Status != enum.StatusCodeSubmitted {
		logx.Error("bill batch status not submitted", "id", id, "status", info.Status)
		return errx.NewStatus("批次账单状态错误")
	}
	ds := []any{}
	nowTime := time.Now()
	for _, d := range info.Datas {
		bd := &model.BillDeal{
			ID:        uuid.NewString(),
			Desc:      info.Desc,
			CreatedAt: nowTime,
			UpdatedAt: nowTime,
			Category:  info.Category,
			RefBatch:  info.RefBatch,
			RefDeal:   *d,
		}
		if bd.AmountClear >= bd.AmountTotal {
			bd.ClearedAt = nowTime
		}
		ds = append(ds, bd)
	}
	if err := b.repoDeal.Insert(ctx, ds...); err != nil {
		logx.Error("save bill deal failed", "error", err)
		return err
	}
	if err := b.repoBatch.Update(ctx,
		store.NewFilter().Eq(field.ID, id),
		&coModel.UpdateStatus{
			ID:        id,
			Status:    enum.StatusCodeApproved,
			UpdatedAt: nowTime,
		}); err != nil {
		logx.Error("update bill batch status failed", "id", id, "error", err)
		return err
	}
	return nil
}
