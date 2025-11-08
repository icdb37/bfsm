package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/icdb37/bfsm/internal/constx/enum"
	"github.com/icdb37/bfsm/internal/constx/field"
	"github.com/icdb37/bfsm/internal/features/purchase/model"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
	coModel "github.com/icdb37/bfsm/internal/model"
)

func (b *batchImpl) savePurchase(ctx context.Context, status *coModel.UpdateStatus, info *model.PurchaseBatch) error {
	stmts := []*store.SessionStatement{
		b.saveBatchStmt(status),
		b.saveGoodsStmt(info),
	}
	if err := store.Transaction(ctx, stmts...); err != nil {
		return err
	}
	if err := b.saveInventory(ctx, info); err != nil {
		return err
	}
	if err := b.saveBill(ctx, info); err != nil {
		return err
	}
	return nil
}

func (b *batchImpl) saveBatchStmt(info *coModel.UpdateStatus) *store.SessionStatement {
	info.UpdatedAt = time.Now()
	return &store.SessionStatement{
		Repo: b.repoBatch,
		Process: func(ctx context.Context, tab store.Tabler) error {
			return tab.Update(ctx,
				store.NewFilter().Eq(field.ID, info.ID),
				info)
		},
	}
}
func (b *batchImpl) saveGoodsStmt(info *model.PurchaseBatch) *store.SessionStatement {
	gs := []any{}
	for _, c := range info.Companies {
		for _, g := range c.Goods {
			rg := &coModel.RefGoods{
				Goods:      *g,
				RefCompany: c.Company,
			}
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
	return &store.SessionStatement{
		Repo: b.repoGoods,
		Process: func(ctx context.Context, tab store.Tabler) error {
			return tab.Insert(ctx, gs...)
		},
	}
}

// saveInventory 保存库存
func (b *batchImpl) saveInventory(ctx context.Context, info *model.PurchaseBatch) error {
	bg := &coModel.BatchGoods{
		RefBatch: info.GetBatch(),
	}
	for _, c := range info.Companies {
		for _, g := range c.Goods {
			rg := &coModel.RefGoods{
				Goods:      *g,
				RefCompany: c.Company,
			}
			bg.Datas = append(bg.Datas, rg)
		}
	}
	if err := b.inventory.Save(ctx, bg); err != nil {
		logx.Error("save purchase inventory failed", "error", err)
		return err
	}
	return nil
}

// saveBill 保存账单
func (b *batchImpl) saveBill(ctx context.Context, info *model.PurchaseBatch) error {
	bd := &coModel.BatchDeal{
		Category: enum.DealCategoryExpense,
		RefBatch: info.GetBatch(),
	}
	for _, c := range info.Companies {
		d := &coModel.RefDeal{
			RefCompany:  c.Company,
			AmountTotal: c.AmountTotal,
			AmountClear: c.AmountClear,
			DealDesc:    c.Desc,
		}
		bd.Datas = append(bd.Datas, d)
	}
	if len(info.Extras) > 0 {
		d := &coModel.RefDeal{
			RefCompany: coModel.ExtraCompany,
			DealDesc:   info.Desc,
		}
		for _, e := range info.Extras {
			d.AmountTotal += e.Amount
		}
		bd.Datas = append(bd.Datas, d)
	}
	if err := b.bill.Save(ctx, bd); err != nil {
		logx.Error("save batch bill failed", "error", err)
		return err
	}
	return nil
}
