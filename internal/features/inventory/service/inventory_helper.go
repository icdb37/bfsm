package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/icdb37/bfsm/internal/constx/field"
	"github.com/icdb37/bfsm/internal/features/inventory/model"
	"github.com/icdb37/bfsm/internal/infra/errx"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
	coModel "github.com/icdb37/bfsm/internal/model"
	"github.com/icdb37/bfsm/internal/utils"
)

func (i *inventoryImpl) b2pCommodity(info *coModel.ProduceBatch) func(p *coModel.ProduceCommodity) *model.FullCommodity {
	nowTime := time.Now()
	return func(p *coModel.ProduceCommodity) *model.FullCommodity {
		fc := &model.FullCommodity{
			ID:           uuid.NewString(),
			BatchDesc:    info.Desc,
			CreatedAt:    nowTime,
			UpdatedAt:    nowTime,
			BatchID:      info.ID,
			Storage:      info.Storage,
			SourceCode:   info.SourceCode,
			RefCommodity: p.RefCommodity,
			RefCompany:   p.RefCompany,
			LeftCount:    p.CommodityCount,
		}
		fc.Normalize()
		return fc
	}
}
func (i *inventoryImpl) getProduceCommodityHash(c *model.FullCommodity) string {
	return c.CommodityHash
}

func (i *inventoryImpl) saveProduceStatement(ctx context.Context, info *coModel.ProduceBatch) (stmts []*store.SessionStatement, err error) {
	nowTime := time.Now()
	// newFulls 批次商品扁平化
	newFulls := utils.Converts(info.Commodity, i.b2pCommodity(info))
	hashs := utils.Converts(newFulls, i.getProduceCommodityHash)
	oldLasts := []*model.LastCommodity{}
	if err := i.repoLast.Query(ctx, store.NewFilter().In(field.CommodityHash, hashs), &oldLasts); err != nil {
		logx.Error("get last commodity failed", "hashs", hashs, "error", err)
		return nil, err
	}
	// id2Lasts 商品HASH -> 聚合商品
	id2Lasts := utils.Convertm(oldLasts, func(l *model.LastCommodity) string { return l.CommodityHash })
	newLasts := map[string]*model.LastCommodity{}
	// 将批次商品转换为聚合商品，修改老商品，新增新商品
	for _, c := range newFulls {
		if l, ok := id2Lasts[c.CommodityHash]; ok {
			l.UpdatedAt = nowTime
			l.CommodityCount += c.CommodityCount
			continue
		}
		l, ok := newLasts[c.CommodityHash]
		if !ok {
			l = &model.LastCommodity{
				ID:           uuid.NewString(),
				CreatedAt:    nowTime,
				UpdatedAt:    nowTime,
				RefCommodity: c.RefCommodity,
				LeftCount:    c.CommodityCount,
			}
			newLasts[c.CommodityHash] = l
		}
		l.CommodityCount = c.CommodityCount
	}
	// 老聚合商品更新商品总量
	stmts = append(stmts, utils.Converts(oldLasts, func(l *model.LastCommodity) *store.SessionStatement {
		return &store.SessionStatement{
			Repo: i.repoLast,
			Process: func(ctx context.Context, tab store.Tabler) error {
				return tab.Update(ctx,
					store.NewFilter().Eq(field.CommodityHash, l.CommodityHash),
					&coModel.UpdateCommodityCount{
						UpdatedAt:      nowTime,
						CommodityCount: l.CommodityCount,
						UsedCount:      l.UsedCount,
						LeftCount:      l.CommodityCount - l.UsedCount,
					})
			},
		}
	})...)
	// 新聚合商品插入
	stmts = append(stmts, &store.SessionStatement{
		Repo: i.repoLast,
		Process: func(ctx context.Context, tab store.Tabler) error {
			return tab.Insert(ctx, utils.Converts(utils.MapValues(newLasts), func(l *model.LastCommodity) any { return l })...)
		},
	})
	// 新批次商品插入
	stmts = append(stmts, &store.SessionStatement{
		Repo: i.repoFull,
		Process: func(ctx context.Context, tab store.Tabler) error {
			return tab.Insert(ctx, utils.Converts(newFulls, func(p *model.FullCommodity) any { return p })...)
		},
	})
	return
}

func (i *inventoryImpl) saveConsumeStatement(ctx context.Context, info *coModel.ConsumeBatch) (stmts []*store.SessionStatement, err error) {
	nowTime := time.Now()
	refProduceIDs := utils.Converts(info.Commodity, func(c *coModel.ConsumeCommodity) string { return c.RefFullID })
	// 修改批次库存
	oldProduces := []*model.FullCommodity{}
	if err = i.repoFull.Query(ctx, store.NewFilter().In(field.ID, refProduceIDs), &oldProduces); err != nil {
		logx.Error("get last commodity failed", "ref_produce_ids", refProduceIDs, "error", err)
		return nil, err
	}
	var hashs []string
	for _, c := range info.Commodity {
		var p *model.FullCommodity
		for _, tp := range oldProduces {
			if tp.ID == c.RefFullID {
				p = tp
				break
			}
		}
		if p == nil {
			logx.Error("get produce commodity failed", "ref_produce_id", c.RefFullID, "error", err)
			return nil, errx.NewErrParam("", "引用商品【%s】不存在", c.RefFullID)
		}
		p.UsedCount += c.CommodityCount
		if p.UsedCount > p.CommodityCount {
			logx.Error("produce commodity count not enough", "left_count", p.CommodityCount-p.UsedCount, "consume_count", c.CommodityCount, "ref_produce_id", c.RefFullID, "error", err)
			return nil, errx.NewErrParam("", "引用商品【%s】库存不足，批次【%s】库存剩余 %d", p.CommodityName, p.BatchDesc, p.CommodityCount-p.UsedCount)
		}
		c.RefBatchID = p.BatchID
		c.CommodityHash = p.CommodityHash
		hashs = append(hashs, p.CommodityHash)
	}
	oldLasts := []*model.LastCommodity{}
	if err = i.repoLast.Query(ctx, store.NewFilter().In(field.CommodityHash, hashs), &oldLasts); err != nil {
		logx.Error("get last commodity failed", "hashs", hashs, "error", err)
		return nil, err
	}
	for _, l := range oldLasts {
		for _, c := range info.Commodity {
			if l.CommodityHash == c.CommodityHash {
				l.UsedCount += c.CommodityCount
				if l.CommodityCount < l.UsedCount {
					logx.Error("last commodity count not enough", "last_commodity", l, "error", err)
					return nil, errx.NewErrParam("", "商品【%s】库存不足", l.CommodityName)
				}
				break
			}
		}
	}
	// 修改批次商品数量
	stmts = append(stmts, utils.Converts(oldProduces, func(p *model.FullCommodity) *store.SessionStatement {
		return &store.SessionStatement{
			Repo: i.repoFull,
			Process: func(ctx context.Context, tab store.Tabler) error {
				return tab.Update(ctx,
					store.NewFilter().Eq(field.ID, p.ID),
					&coModel.UpdateCommodityCount{
						UpdatedAt:      nowTime,
						CommodityCount: p.CommodityCount,
						UsedCount:      p.UsedCount,
						LeftCount:      p.CommodityCount - p.UsedCount,
					})
			},
		}
	})...)
	// 修改聚合商品数量
	stmts = append(stmts, utils.Converts(oldLasts, func(l *model.LastCommodity) *store.SessionStatement {
		return &store.SessionStatement{
			Repo: i.repoLast,
			Process: func(ctx context.Context, tab store.Tabler) error {
				return tab.Update(ctx,
					store.NewFilter().Eq(field.ID, l.ID),
					&coModel.UpdateCommodityCount{
						UpdatedAt:      nowTime,
						CommodityCount: l.CommodityCount,
						UsedCount:      l.UsedCount,
						LeftCount:      l.CommodityCount - l.UsedCount,
					})
			},
		}
	})...)
	return
}
