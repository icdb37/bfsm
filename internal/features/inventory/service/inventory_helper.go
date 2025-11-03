package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/icdb37/bfsm/internal/constx/field"
	"github.com/icdb37/bfsm/internal/features/inventory/model"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
	coModel "github.com/icdb37/bfsm/internal/model"
)

func (i *inventoryImpl) saveFull(ctx context.Context, info *coModel.EntireBatch) error {
	var fulls []any
	nowTime := time.Now()
	for _, c := range info.Commodity {
		fulls = append(fulls, &model.FullCommodity{
			ID:         uuid.NewString(),
			BatchDesc:  info.Desc,
			CreatedAt:  nowTime,
			UpdatedAt:  nowTime,
			BatchID:    info.ID,
			Storage:    info.Storage,
			Company:    &coModel.RefCompany{CompanyID: info.Company.ID, CompanyName: info.Company.Name},
			Commodity:  c,
			SourceCode: info.SourceCode,
		})
	}
	if err := i.repoFull.Insert(ctx, fulls...); err != nil {
		logx.Error("inventory insert full fail", "error", err)
		return err
	}
	return nil
}

func (i *inventoryImpl) saveLast(ctx context.Context, info *coModel.EntireBatch) error {
	hashs := []string{}
	for _, c := range info.Commodity {
		c.Hash = c.GetHash()
		hashs = append(hashs, c.Hash)
	}
	lasts := []*model.LastCommodity{}
	if err := i.repoLast.Query(ctx, store.NewFilter().In(field.Hash, hashs), &lasts); err != nil {
		logx.Error("get last commodity failed", "ids", hashs, "error", err)
		return err
	}
	nowTime := time.Now()
	sig := int32(1)
	if info.SourceCode < 0 {
		sig = -1
	}
	oldInfos := map[string]*model.LastCommodity{}
	for _, l := range lasts {
		oldInfos[l.Hash] = l
	}
	newInfos := map[string]*model.LastCommodity{}
	for _, c := range info.Commodity {
		if l, ok := oldInfos[c.Hash]; ok {
			l.UpdatedAt = nowTime
			l.Count += sig * c.Count
			continue
		}
		l, ok := newInfos[c.Hash]
		if !ok {
			l := &model.LastCommodity{
				ID:        uuid.NewString(),
				CreatedAt: nowTime,
				UpdatedAt: nowTime,
				Commodity: *c,
			}
			newInfos[c.Hash] = l
		}
		l.Count = sig * c.Count
	}
	for _, l := range oldInfos {
		if err := i.repoLast.Update(ctx, store.NewFilter().Eq(field.Hash, l.Hash), l); err != nil {
			logx.Error("inventory upate last fail", "id", l.ID, "error", err)
			return err
		}
	}
	for _, l := range newInfos {
		if err := i.repoLast.Insert(ctx, l); err != nil {
			logx.Error("inventory insert last fail", "info", info, "error", err)
			return err
		}
	}
	return nil
}

// Consume 减少库存
func (i *inventoryImpl) save(ctx context.Context, info *coModel.EntireBatch) error {
	if err := i.saveFull(ctx, info); err != nil {
		return err
	}
	if err := i.saveLast(ctx, info); err != nil {
		return err
	}
	return nil
}
