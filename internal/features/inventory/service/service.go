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

// InventoryProducer 库存生产者
type InventoryProducer interface {
	Search(ctx context.Context, req *coModel.SearchRequest[coModel.QueryBatch]) (resp *coModel.SearchResponse[coModel.EntireBatch], err error)
	Create(ctx context.Context, info *coModel.EntireBatch) error
	Update(ctx context.Context, info *coModel.EntireBatch) error
	Delete(ctx context.Context, id string) error
}

// InventoryConsumer 库存消费者
type InventoryConsumer interface {
	Search(ctx context.Context, req *coModel.SearchRequest[coModel.QueryBatch]) (resp *coModel.SearchResponse[coModel.EntireBatch], err error)
	Create(ctx context.Context, info *coModel.EntireBatch) error
	Update(ctx context.Context, info *coModel.EntireBatch) error
	Delete(ctx context.Context, id string) error
}

// InventoryInventory 库存管理
type InventoryInventory interface {
	Search(ctx context.Context, req *coModel.SearchRequest[model.QueryCommodity]) (resp *coModel.SearchResponse[model.LastCommodity], err error)
	Get(ctx context.Context, id string) (resp []*model.FullCommodity, err error)
}

// wire.ProvideName(featc.InventoryConsume, func() InventoryConsumer {
// 	repo, err := store.NewTable(&model.ConsumeBatch{})
// 	if err != nil {
// 		logx.Fatal("create repo failed", "feature", featc.InventoryConsume, "error", err)
// 	}
// 	return &consumeImpl{repo: repo}
// })
// wire.ProvideName(featc.InventoryProduce, func() InventoryProducer {
// 	repo, err := store.NewTable(&model.ProduceBatch{})
// 	if err != nil {
// 		logx.Fatal("create repo failed", "feature", featc.InventoryProduce, "error", err)
// 	}
// 	return &produceImpl{repo: repo}
// })

func Provide() {
	repoFull, err := store.NewTable(&model.FullCommodity{})
	if err != nil {
		logx.Fatal("create repo failed", "feature", featc.InventoryInventory, "error", err)
	}
	repoLast, err := store.NewTable(&model.LastCommodity{})
	if err != nil {
		logx.Fatal("create repo failed", "feature", featc.InventoryInventory, "error", err)
	}
	i := &inventoryImpl{repoFull: repoFull, repoLast: repoLast}
	wire.ProvideName(featc.InventoryInventory, func() InventoryInventory {
		return i
	})
	wire.ProvideName(featc.InventoryConsume, func() coService.InventoryConsumer {
		return i
	})
	wire.ProvideName(featc.InventoryProduce, func() coService.InventoryProducer {
		return i
	})
}
