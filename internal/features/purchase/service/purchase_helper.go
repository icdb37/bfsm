package service

import (
	"context"

	"github.com/icdb37/bfsm/internal/features/purchase/model"
	"github.com/icdb37/bfsm/internal/infra/logx"
	coModel "github.com/icdb37/bfsm/internal/model"
)

// saveInventory 保存库存
func (p *purchaseImpl) saveInventory(ctx context.Context, info *model.EntirePurchase) error {
	for _, c := range info.Commodities {
		bc := &coModel.ProduceBatch{
			ID:        info.ID,
			Commodity: c.Commodities,
		}
		if err := p.inventory.Produce(ctx, bc); err != nil {
			logx.Error("produce inventory failed", "error", err)
			return err
		}
	}
	return nil
}
