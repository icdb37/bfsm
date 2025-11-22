package service

import (
	"context"

	"github.com/icdb37/bfsm/internal/constx/featc"
	"github.com/icdb37/bfsm/internal/features/purchase/model"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
	coModel "github.com/icdb37/bfsm/internal/model"
	coService "github.com/icdb37/bfsm/internal/service"
	"github.com/icdb37/bfsm/internal/wire"
)

// BatchServer - 采购订单服务接口
type BatchServer interface {
	Search(ctx context.Context, req *coModel.SearchRequest[model.QueryPurchase]) (resp *coModel.SearchResponse[model.SimplePurchase], err error)
	Create(ctx context.Context, info *model.PurchaseBatch) error
	Update(ctx context.Context, info *model.PurchaseBatch) error
	UpdateStatus(ctx context.Context, req *coModel.UpdateStatus) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*model.PurchaseBatch, error)
}

// GoodsServer - 商品服务接口
type GoodsServer interface {
	Search(ctx context.Context, req *coModel.SearchRequest[model.QueryPurchaseGoods]) (resp *coModel.SearchResponse[model.PurchaseGoods], err error)
	Get(ctx context.Context, id string) (resp *model.PurchaseGoods, err error)
}

func Provide() {
	wire.ProvideName(featc.PurchaseBatch, func() BatchServer {
		repoBatch, err := store.NewTable(&model.PurchaseBatch{})
		if err != nil {
			logx.Fatal("create purchase batch repo failed", "error", err)
		}
		repoGoods, err := store.NewTable(&model.PurchaseGoods{})
		if err != nil {
			logx.Fatal("create purchase goods repo failed", "error", err)
		}
		inventory := wire.ResolveName[coService.InventorySaver](featc.InventorySave)
		bill := wire.ResolveName[coService.BillSaver](featc.BillSave)
		return &batchImpl{repoBatch: repoBatch, repoGoods: repoGoods, inventory: inventory, bill: bill}
	})
	wire.ProvideName(featc.PurchaseGoods, func() GoodsServer {
		repo, err := store.NewTable(&model.PurchaseGoods{})
		if err != nil {
			logx.Fatal("create purchase goods repo failed", "error", err)
		}
		return &goodsImpl{repo: repo}
	})
}
