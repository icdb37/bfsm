package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/icdb37/bfsm/internal/features/purchase/model"
	"github.com/icdb37/bfsm/internal/infra/logx"
	coModel "github.com/icdb37/bfsm/internal/model"
)

// saveInventory 保存库存
func (p *batchImpl) saveInventory(ctx context.Context, info *model.PurchaseBatch) error {
	bg := &coModel.BatchGoods{
		RefBatch: info.GetBatch(),
	}
	gs := []any{}
	for _, c := range info.Companies {
		for _, g := range c.Goods {
			rg := &coModel.RefGoods{
				Goods:      *g,
				RefCompany: c.Company,
			}
			bg.Datas = append(bg.Datas, rg)
			pg := &model.PurchaseGoods{
				PurchaseID:   info.ID,
				PurchaseName: info.Name,
				RefGoods:     *rg,
				CreatedAt:    time.Now(),
			}
			pg.ID = uuid.NewString()
			gs = append(gs, pg)
		}
	}
	if err := p.inventory.Save(ctx, bg); err != nil {
		logx.Error("save purchase inventory failed", "error", err)
		return err
	}
	if err := p.repoGoods.Insert(ctx, gs...); err != nil {
		logx.Error("save purchase goods failed", "error", err)
		return err
	}
	return nil
}
