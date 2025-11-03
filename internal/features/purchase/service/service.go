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

// PurchaseServer - 采购订单服务接口
type PurchaseServer interface {
	Search(ctx context.Context, req *coModel.SearchRequest[model.QueryPurchase]) (resp *coModel.SearchResponse[model.EntirePurchase], err error)
	Create(ctx context.Context, info *model.EntirePurchase) error
	Update(ctx context.Context, info *model.EntirePurchase) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*model.EntirePurchase, error)
}

func Provide() {
	wire.ProvideName(featc.CommodityCommodity, func() PurchaseServer {
		repo, err := store.NewTable(&model.EntirePurchase{})
		if err != nil {
			logx.Fatal("create purchase repo failed", "error", err)
		}
		inventory := wire.ResolveName[coService.InventoryProducer](featc.InventoryProduce)
		return &purchaseImpl{repo: repo, inventory: inventory}
	})
}
