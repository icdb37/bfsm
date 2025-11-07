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

func (i *inventoryImpl) b2pCommodity(info *coModel.BatchGoods) func(p *coModel.RefGoods) *model.FullGoods {
	nowTime := time.Now()
	return func(p *coModel.RefGoods) *model.FullGoods {
		fc := &model.FullGoods{
			CreatedAt: nowTime,
			UpdatedAt: nowTime,
			RefBatch:  info.RefBatch,
			RefGoods:  *p,
			LeftCount: p.Count,
		}
		fc.ID = uuid.NewString()
		fc.Normalize()
		return fc
	}
}
func (i *inventoryImpl) getProduceHash(c *model.FullGoods) string {
	return c.Hash
}

func (i *inventoryImpl) saveProduceStatement(ctx context.Context, info *coModel.BatchGoods) (stmts []*store.SessionStatement, err error) {
	nowTime := time.Now()
	// newFulls 批次商品扁平化
	newFulls := utils.Converts(info.Datas, i.b2pCommodity(info))
	hashs := utils.Converts(newFulls, i.getProduceHash)
	oldLasts := []*model.LastCommodity{}
	if err := i.repoLast.Query(ctx, store.NewFilter().In(field.Hash, hashs), &oldLasts); err != nil {
		logx.Error("get last commodity failed", "hashs", hashs, "error", err)
		return nil, err
	}
	// id2Lasts 商品HASH -> 聚合商品
	id2Lasts := utils.Convertm(oldLasts, func(l *model.LastCommodity) string { return l.Hash })
	newLasts := map[string]*model.LastCommodity{}
	// 将批次商品转换为聚合商品，修改老商品，新增新商品
	for _, c := range newFulls {
		if l, ok := id2Lasts[c.Hash]; ok {
			l.UpdatedAt = nowTime
			l.Count += c.Count
			continue
		}
		l, ok := newLasts[c.Hash]
		if !ok {
			l = &model.LastCommodity{
				ID:        uuid.NewString(),
				CreatedAt: nowTime,
				UpdatedAt: nowTime,
				Commodity: c.Commodity,
				LeftCount: c.Count,
			}
			newLasts[c.Hash] = l
		}
		l.Count = c.Count
	}
	// 老聚合商品更新商品总量
	stmts = append(stmts, utils.Converts(oldLasts, func(l *model.LastCommodity) *store.SessionStatement {
		return &store.SessionStatement{
			Repo: i.repoLast,
			Process: func(ctx context.Context, tab store.Tabler) error {
				return tab.Update(ctx,
					store.NewFilter().Eq(field.Hash, l.Hash),
					&coModel.GoodsCount{
						UpdatedAt: nowTime,
						Count:     l.Count,
						UsedCount: l.UsedCount,
						LeftCount: l.Count - l.UsedCount,
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
			return tab.Insert(ctx, utils.Converts(newFulls, func(p *model.FullGoods) any { return p })...)
		},
	})
	return
}

func (i *inventoryImpl) saveConsumeStatement(ctx context.Context, info *coModel.BatchGoods) (stmts []*store.SessionStatement, err error) {
	nowTime := time.Now()
	goodsIDs := utils.Converts(info.Datas, func(c *coModel.RefGoods) string { return c.ID })
	// 修改批次库存
	oldFulls := []*model.FullGoods{}
	if err = i.repoFull.Query(ctx, store.NewFilter().In(field.ID, goodsIDs), &oldFulls); err != nil {
		logx.Error("get last commodity failed", "ref_ids", goodsIDs, "error", err)
		return nil, err
	}
	var hashs []string
	for _, c := range info.Datas {
		var p *model.FullGoods
		for _, tp := range oldFulls {
			if tp.ID == c.ID {
				p = tp
				break
			}
		}
		if p == nil {
			logx.Error("get consume commodity failed", "ref_goods_id", c.ID, "error", err)
			return nil, errx.NewErrParam("", "引用商品【%s】不存在", c.ID)
		}
		p.UsedCount += c.Count
		if p.UsedCount > p.Count {
			logx.Error("consume commodity count not enough", "left_count", p.LeftCount, "consume_count", c.Count, "ref_goods_id", c.ID, "error", err)
			return nil, errx.NewErrParam("", "引用商品【%s】库存不足，批次【%s】库存剩余 %d", p.Name, p.BatchDesc, p.LeftCount)
		}
		c.Hash = p.Hash
		hashs = append(hashs, p.Hash)
	}
	oldLasts := []*model.LastCommodity{}
	if err = i.repoLast.Query(ctx, store.NewFilter().In(field.Hash, hashs), &oldLasts); err != nil {
		logx.Error("get last commodity failed", "hashs", hashs, "error", err)
		return nil, err
	}
	for _, l := range oldLasts {
		for _, c := range info.Datas {
			if l.Hash == c.Hash {
				l.UsedCount += c.Count
				if l.Count < l.UsedCount {
					logx.Error("last commodity count not enough", "last_commodity", l, "error", err)
					return nil, errx.NewErrParam("", "商品【%s】库存不足", l.Name)
				}
				break
			}
		}
	}
	// 修改批次商品数量
	stmts = append(stmts, utils.Converts(oldFulls, func(p *model.FullGoods) *store.SessionStatement {
		return &store.SessionStatement{
			Repo: i.repoFull,
			Process: func(ctx context.Context, tab store.Tabler) error {
				return tab.Update(ctx,
					store.NewFilter().Eq(field.ID, p.ID),
					&coModel.GoodsCount{
						UpdatedAt: nowTime,
						Count:     p.Count,
						UsedCount: p.UsedCount,
						LeftCount: p.Count - p.UsedCount,
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
					&coModel.GoodsCount{
						UpdatedAt: nowTime,
						Count:     l.Count,
						UsedCount: l.UsedCount,
						LeftCount: l.Count - l.UsedCount,
					})
			},
		}
	})...)
	return
}
