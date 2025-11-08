package service

import (
	"context"

	"github.com/icdb37/bfsm/internal/features/bill/model"
	"github.com/icdb37/bfsm/internal/infra/logx"
	coModel "github.com/icdb37/bfsm/internal/model"
)

// Save - 保存账单批次
func (b *batchImpl) Save(ctx context.Context, param *coModel.BatchDeal) error {
	info := model.BillBatch{
		Category: param.Category,
		RefBatch: param.RefBatch,
		Datas:    param.Datas,
	}
	info.Normalize()
	if err := b.repoBatch.Insert(ctx, info); err != nil {
		logx.Error("save bill batch failed", "error", err)
		return err
	}
	return nil
}
