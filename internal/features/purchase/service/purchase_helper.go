package service

import (
	"context"

	"github.com/icdb37/bfsm/internal/features/purchase/model"
	"github.com/icdb37/bfsm/internal/infra/logx"
	coModel "github.com/icdb37/bfsm/internal/model"
)

// saveInventory 保存库存
func (p *purchaseImpl) saveInventory(ctx context.Context, info *model.PurchaseBatch) error {
	bg := &coModel.BatchGoods{
		RefBatch: info.GetBatch(),
	}
	for _, c := range info.Companies {
		for _, g := range c.Goods {
			d := &coModel.RefGoods{
				Goods:      *g,
				RefCompany: c.Company,
			}
			bg.Datas = append(bg.Datas, d)
		}
	}
	if err := p.inventory.Save(ctx, bg); err != nil {
		logx.Error("save purchase inventory failed", "error", err)
		return err
	}
	return nil
}
