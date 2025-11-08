package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/icdb37/bfsm/internal/constx/field"
	"github.com/icdb37/bfsm/internal/features/bill/model"
	"github.com/icdb37/bfsm/internal/infra/errx"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
	coModel "github.com/icdb37/bfsm/internal/model"
)

// Save - 保存账单批次
func (b *batchImpl) Save(ctx context.Context, param *coModel.BatchDeal) error {
	info := model.BillBatch{
		Datas: param.Datas,
	}
	info.ID = uuid.NewString()
	info.Business = param.Business
	info.Category = param.Category
	info.RefBatch = param.RefBatch
	info.Normalize()
	if err := b.repoBatch.Insert(ctx, info); err != nil {
		logx.Error("save bill batch failed", "error", err)
		return err
	}
	return nil
}

// getSimpleBatch - 同步金额，交易账单批次
func (b *batchImpl) getSimpleBatch(ctx context.Context, id string) (*model.SimpleBillBatch, error) {
	info := &model.SimpleBillBatch{}
	if err := b.repoBatch.Query(ctx, store.NewFilter().Eq(field.ID, id), info); err != nil {
		logx.Error("query bill batch failed", "id", id, "error", err)
		return nil, err
	}
	if info.ID == "" {
		logx.Error("bill batch not found", "id", id)
		return nil, errx.NewStatus("批次账单不存在")
	}
	return info, nil
}
