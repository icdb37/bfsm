package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/icdb37/bfsm/internal/constx/enum"
	"github.com/icdb37/bfsm/internal/constx/field"
	"github.com/icdb37/bfsm/internal/features/inventory/model"
	"github.com/icdb37/bfsm/internal/infra/cfpx"
	"github.com/icdb37/bfsm/internal/infra/errx"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
	coModel "github.com/icdb37/bfsm/internal/model"
	"github.com/icdb37/bfsm/internal/utils"
)

type inventoryImpl struct {
	repoLast store.Tabler
	repoFull store.Tabler
}

func (i *inventoryImpl) SearchLast(ctx context.Context, req *coModel.SearchRequest[model.QueryLastCommodity]) (resp *coModel.SearchResponse[model.LastCommodity], err error) {
	qf := store.Unmarshal(req.Query)
	resp = &coModel.SearchResponse[model.LastCommodity]{}
	pf := req.GetPage()
	if resp.Total, err = i.repoLast.Search(ctx, qf, pf, &(resp.Datas)); err != nil {
		logx.Error("search last commodity failed", "error", err)
		return nil, err
	}
	return resp, nil
}

func (i *inventoryImpl) SearchFull(ctx context.Context, req *coModel.SearchRequest[coModel.QueryProduceCommodity]) (resp *coModel.SearchResponse[coModel.ProduceCommodity], err error) {
	req.Query.Normalize()
	if req.Query.CommodityHash == "" {
		return nil, errx.NewErrParam("query.commodity.hash", "商品哈希必填查询参数")
	}
	qf := store.Unmarshal(req.Query)
	resp = &coModel.SearchResponse[coModel.ProduceCommodity]{}
	pf := req.GetPage()
	if resp.Total, err = i.repoFull.Search(ctx, qf, pf, &(resp.Datas)); err != nil {
		logx.Error("search full commodity failed", "error", err)
		return nil, err
	}
	return resp, nil
}

// Produce 增加库存
func (i *inventoryImpl) Produce(ctx context.Context, info *coModel.ProduceBatch) error {
	info.SourceCode = enum.SourceCodePurchaseProduce
	stmts, err := i.saveProduceStatement(ctx, info)
	if err != nil {
		return err
	}
	if err := store.Transaction(ctx, stmts...); err != nil {
		return err
	}
	return nil
}

// Consume 减少库存
func (i *inventoryImpl) Consume(ctx context.Context, info *coModel.ConsumeBatch) error {
	info.SourceCode = enum.SourceCodeConsume
	stmts, err := i.saveConsumeStatement(ctx, info)
	if err != nil {
		return err
	}
	if err := store.Transaction(ctx, stmts...); err != nil {
		return err
	}
	return nil
}

func (i *inventoryImpl) UpdateFull(ctx context.Context, newFull *coModel.ProduceCommodity) error {
	newFull.Normalize()
	oldFull := &coModel.ProduceCommodity{}
	where := store.NewFilter().Eq(field.ID, newFull.ID)
	if err := i.repoFull.Query(ctx, where, oldFull); err != nil {
		logx.Error("get full commodity failed", "id", newFull.ID, "error", err)
		return err
	}
	if oldFull.ID == "" {
		return errx.NewErrParam("", "商品不存在")
	}
	if oldFull.CommodityHash != newFull.CommodityHash {
		return errx.NewErrParam("", "不支持修改商品名称、商品规格、商品尺寸")
	}
	oldLast := &model.LastCommodity{}
	if err := i.repoLast.Query(ctx, store.NewFilter().Eq(field.CommodityHash, oldFull.CommodityHash), oldLast); err != nil {
		logx.Error("get last commodity failed", "hash", oldFull.CommodityHash, "error", err)
		return err
	}
	if oldLast.ID == "" {
		return errx.NewErrParam("", "商品不存在")
	}
	sig := int32(1)
	if oldFull.SourceCode < 0 {
		sig = -1
	}
	oldLast.CommodityCount += sig * (newFull.CommodityCount - oldFull.CommodityCount)
	newFull.CreatedAt, newFull.SourceCode, newFull.BatchID = oldFull.CreatedAt, oldFull.SourceCode, oldFull.BatchID

	if err := i.repoFull.Update(ctx, where, newFull); err != nil {
		logx.Error("update full commodity failed", "id", newFull.ID, "error", err)
		return err
	}
	if err := i.repoLast.Update(ctx, store.NewFilter().Eq(field.ID, oldLast.ID), oldLast); err != nil {
		logx.Error("update last commodity failed", "id", oldLast.ID, "error", err)
		return err
	}
	return nil
}

func (i *inventoryImpl) UpdateLast(ctx context.Context, newLast *model.LastCommodity) error {
	if err := utils.ProcessAll(ctx, newLast, cfpx.ProcessUpdate, model.ProcessLastCommodity); err != nil {
		logx.Error("create last commodity failed", "error", err)
		return err
	}
	oldLast := &model.LastCommodity{}
	where := store.NewFilter().Eq(field.ID, newLast.ID)
	if err := i.repoLast.Query(ctx, where, oldLast); err != nil {
		logx.Error("get last commodity failed", "id", newLast.ID, "error", err)
		return err
	}
	if oldLast.ID == "" {
		return errx.NewErrParam("", "商品不存在")
	}
	if oldLast.CommodityHash != newLast.CommodityHash {
		where = store.NewFilter().Eq(field.Hash, oldLast.CommodityHash)
		i.repoFull.Update(ctx, where, &coModel.RefSimpleCommodity{
			Hash: oldLast.CommodityHash,
			Name: oldLast.CommodityName,
			Desc: oldLast.CommodityDesc,
			Spec: oldLast.CommoditySpec,
			Size: oldLast.CommoditySize,
		})
	}
	if err := i.repoLast.Update(ctx, where, newLast); err != nil {
		logx.Error("update last commodity failed", "id", newLast.ID, "error", err)
		return err
	}
	if newLast.CommodityCount != oldLast.CommodityCount {
		newFull := &coModel.ProduceCommodity{
			ID:           uuid.NewString(),
			CreatedAt:    newLast.UpdatedAt,
			UpdatedAt:    newLast.UpdatedAt,
			BatchID:      uuid.NewString(),
			BatchDesc:    "手动编辑库存商品信息，自动新增批次",
			RefCommodity: newLast.RefCommodity,
		}
		if newLast.CommodityCount > oldLast.CommodityCount {
			newFull.SourceCode = enum.SourceCodeInventoryUpdateProduce
			newFull.CommodityCount = newLast.CommodityCount - oldLast.CommodityCount
		} else {
			newFull.SourceCode = enum.SourceCodeInventoryUpdateConsume
			newFull.CommodityCount = oldLast.CommodityCount - newLast.CommodityCount
		}
		if err := i.repoFull.Insert(ctx, newFull); err != nil {
			logx.Error("insert full commodity failed", "id", newFull.ID, "error", err)
			return err
		}
	}
	return nil
}
