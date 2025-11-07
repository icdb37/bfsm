package service

import (
	"context"

	"github.com/icdb37/bfsm/internal/constx/featc"
	"github.com/icdb37/bfsm/internal/features/inventory/model"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
	coModel "github.com/icdb37/bfsm/internal/model"
	coService "github.com/icdb37/bfsm/internal/service"
	"github.com/icdb37/bfsm/internal/wire"
)

// InventoryInventory 库存管理
type InventoryInventory interface {
	SearchLast(ctx context.Context, req *coModel.SearchRequest[model.QueryLastCommodity]) (resp *coModel.SearchResponse[model.LastCommodity], err error)
	SearchFull(ctx context.Context, req *coModel.SearchRequest[model.QueryFullGoods]) (resp *coModel.SearchResponse[model.FullGoods], err error)
	UpdateFull(ctx context.Context, info *model.FullGoods) error
	UpdateLast(ctx context.Context, newLast *model.LastCommodity) error
}

func Provide() {
	repoFull, err := store.NewTable(&model.FullGoods{})
	if err != nil {
		logx.Fatal("create repo failed", "feature", featc.InventoryGoodsFull, "error", err)
	}
	repoLast, err := store.NewTable(&model.LastCommodity{})
	if err != nil {
		logx.Fatal("create repo failed", "feature", featc.InventoryGoodsLast, "error", err)
	}
	i := &inventoryImpl{repoFull: repoFull, repoLast: repoLast}
	wire.ProvideName(featc.InventoryInventory, func() InventoryInventory {
		return i
	})
	wire.ProvideName(featc.InventorySave, func() coService.InventorySaver {
		return i
	})
}
