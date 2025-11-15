package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/icdb37/bfsm/internal/constx/enum"
	"github.com/icdb37/bfsm/internal/constx/field"
	"github.com/icdb37/bfsm/internal/features/bill/model"
	"github.com/icdb37/bfsm/internal/infra/cfpx"
	"github.com/icdb37/bfsm/internal/infra/errx"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
	coModel "github.com/icdb37/bfsm/internal/model"
	"github.com/icdb37/bfsm/internal/utils"
)

type batchImpl struct {
	repoBatch store.Tabler
	repoDeal  store.Tabler
}

func (b *batchImpl) Search(ctx context.Context, req *coModel.SearchRequest[model.QueryBillBatch]) (resp *coModel.SearchResponse[model.SimpleBillBatch], err error) {
	qf := store.Unmarshal(req.Query)
	resp = &coModel.SearchResponse[model.SimpleBillBatch]{}
	pf := req.GetPage()
	if resp.Total, err = b.repoBatch.Search(ctx, qf, pf, &(resp.Data)); err != nil {
		logx.Error("search bill batch failed", "error", err)
		return nil, err
	}
	return resp, nil
}

// Create - 创建，交易账单批次
func (b *batchImpl) Get(ctx context.Context, id string) (info *model.BillBatch, err error) {
	logx.Info("get bill batch", "id", id)
	info = &model.BillBatch{}
	if err := b.repoBatch.Query(ctx, store.NewFilter().Eq(field.ID, id), info); err != nil {
		logx.Error("get bill batch failed", "id", id, "error", err)
		return nil, err
	}
	if info.ID == "" {
		logx.Error("bill batch not found", "id", id)
		return nil, errx.NewStatus("批次账单不存在")
	}
	return info, nil
}

// Create - 创建，交易账单批次
func (b *batchImpl) Create(ctx context.Context, info *model.BillBatch) error {
	logx.Info("create bill batch", "info", info)
	if err := utils.ProcessAll(ctx, info, cfpx.ProcessCreate); err != nil {
		logx.Error("create bill batch failed", "error", err)
		return err
	}
	if err := b.repoBatch.Insert(ctx, info); err != nil {
		logx.Error("create bill batch failed", "error", err)
		return err
	}
	return nil
}

// Update - 修改，交易账单批次
func (b *batchImpl) Update(ctx context.Context, newInfo *model.BillBatch) error {
	logx.Info("update bill batch", "info", newInfo)
	if err := utils.ProcessAll(ctx, newInfo, cfpx.ProcessUpdate); err != nil {
		logx.Error("update bill batch failed", "error", err)
		return err
	}
	oldInfo := coModel.UpdateStatus{}
	if err := b.repoBatch.Query(ctx, store.NewFilter().Eq(field.ID, newInfo.ID), &oldInfo); err != nil {
		logx.Error("query bill batch failed", "id", newInfo.ID, "error", err)
		return err
	}
	if oldInfo.Status > enum.StatusCodeSubmitted {
		logx.Error("bill batch status not allow update", "id", newInfo.ID, "status", oldInfo.Status)
		return errx.NewStatus("批次账单禁止修改")
	}
	if err := b.repoBatch.Update(ctx, store.NewFilter().Eq(field.ID, newInfo.ID), newInfo); err != nil {
		logx.Error("update bill batch failed", "error", err)
		return err
	}
	return nil
}

// UpdateAmount - 更新金额，交易账单批次
func (b *batchImpl) UpdateAmount(ctx context.Context, info *coModel.UpdateAmount) error {
	logx.Info("update bill batch amount", "info", info)
	oldInfo, err := b.getSimpleBatch(ctx, info.ID)
	if err != nil {
		return err
	}
	if oldInfo.Status >= enum.StatusCodeCompleted {
		logx.Error("bill batch status not allow update", "id", info.ID, "status", oldInfo.Status)
		return errx.NewStatus("批次账单已归档，禁止修改")
	}
	info.ClearedAt, info.AmountStatus = oldInfo.ClearedAt, oldInfo.AmountStatus
	info.AmountLeft = info.AmountTotal - info.AmountClear
	if info.AmountLeft <= 0 {
		info.ClearedAt = time.Now()
		info.AmountStatus = enum.AmountStatusPaid
	}
	if err := b.repoDeal.Update(ctx, store.NewFilter().Eq(field.ID, info.ID), info); err != nil {
		logx.Error("update bill deal amount failed", "error", err)
		return err
	}
	return nil
}

// SyncAmount - 同步金额，交易账单批次
func (b *batchImpl) SyncAmount(ctx context.Context, id string) error {
	logx.Info("sync bill batch amount", "id", id)
	infoBatch, err := b.getSimpleBatch(ctx, id)
	if err != nil {
		return err
	}
	infoDeals := []*coModel.UpdateAmount{}
	if err := b.repoDeal.Query(ctx, store.NewFilter().Eq(field.BatchID, infoBatch.BatchID), &infoDeals); err != nil {
		logx.Error("query bill deal failed", field.BatchID, infoBatch.BatchID, "error", err)
		return err
	}
	amountTotal, amountClear := int32(0), int32(0)
	for _, d := range infoDeals {
		amountTotal += d.AmountTotal
		amountClear += d.AmountClear
	}
	amountLeft := amountTotal - amountClear
	if infoBatch.AmountLeft > 0 && amountLeft <= 0 {
		amountLeft = 0
		infoBatch.ClearedAt = time.Now()
		infoBatch.AmountStatus = enum.AmountStatusPaid
	}
	infoBatch.AmountTotal = amountTotal
	infoBatch.AmountClear = amountClear
	infoBatch.AmountLeft = amountLeft
	if err := b.repoBatch.Update(ctx,
		store.NewFilter().Eq(field.ID, infoBatch.ID),
		infoBatch,
	); err != nil {
		logx.Error("update bill batch failed", "id", infoBatch.ID, "error", err)
		return err
	}
	return nil
}

// Approve - 审核通过，交易账单细化
func (b *batchImpl) Approve(ctx context.Context, param *coModel.UpdateStatus) error {
	logx.Info("approve bill batch", "id", param.ID)
	info := &model.BillBatch{}
	if err := b.repoBatch.Query(ctx, store.NewFilter().Eq(field.ID, param.ID), info); err != nil {
		logx.Error("query bill batch failed", "id", param.ID, "error", err)
		return err
	}
	if info.Status != enum.StatusCodeSubmitted {
		logx.Error("bill batch status not submitted", "id", param.ID, "status", info.Status)
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
		store.NewFilter().Eq(field.ID, param.ID),
		param,
	); err != nil {
		logx.Error("update bill batch status failed", "id", param.ID, "error", err)
		return err
	}
	return nil
}

// Complete - 完成，交易账单细化
func (b *batchImpl) Complete(ctx context.Context, param *coModel.UpdateStatus) error {
	logx.Info("complete bill batch", "id", param.ID)
	if err := b.repoBatch.Query(ctx, store.NewFilter().Eq(field.ID, param.ID), param); err != nil {
		logx.Error("query bill batch failed", "id", param.ID, "error", err)
		return err
	}
	if param.Status >= enum.StatusCodeCompleted {
		logx.Error("bill batch status invalid", "id", param.ID, "status", param.Status)
		return errx.NewStatus("批次账单状态错误")
	}
	param.UpdatedAt = time.Now()
	if err := b.repoBatch.Update(ctx,
		store.NewFilter().Eq(field.ID, param.ID),
		param,
	); err != nil {
		logx.Error("update bill batch status failed", "id", param.ID, "error", err)
		return err
	}
	return nil
}

// Cancel - 取消，交易账单细化
func (b *batchImpl) Cancel(ctx context.Context, param *coModel.UpdateStatus) error {
	logx.Info("cancel bill batch", "id", param.ID)
	if err := b.repoBatch.Query(ctx, store.NewFilter().Eq(field.ID, param.ID), param); err != nil {
		logx.Error("query bill batch failed", "id", param.ID, "error", err)
		return err
	}
	if param.Status >= enum.StatusCodeCanceled {
		logx.Error("bill batch status invalid", "id", param.ID, "status", param.Status)
		return errx.NewStatus("批次账单状态错误")
	}
	param.UpdatedAt = time.Now()
	if err := b.repoBatch.Update(ctx,
		store.NewFilter().Eq(field.ID, param.ID),
		param,
	); err != nil {
		logx.Error("update bill batch status failed", "id", param.ID, "error", err)
		return err
	}
	return nil
}

// Close - 关闭，交易账单细化
func (b *batchImpl) Close(ctx context.Context, param *coModel.UpdateStatus) error {
	logx.Info("close bill batch", "id", param.ID)
	if err := b.repoBatch.Query(ctx, store.NewFilter().Eq(field.ID, param.ID), param); err != nil {
		logx.Error("query bill batch failed", "id", param.ID, "error", err)
		return err
	}
	if param.Status == enum.StatusCodeClosed {
		logx.Error("bill batch status invalid", "id", param.ID, "status", param.Status)
		return errx.NewStatus("批次账单状态错误")
	}
	param.UpdatedAt = time.Now()
	if err := b.repoBatch.Update(ctx,
		store.NewFilter().Eq(field.ID, param.ID),
		param,
	); err != nil {
		logx.Error("update bill batch status failed", "id", param.ID, "error", err)
		return err
	}
	return nil
}
