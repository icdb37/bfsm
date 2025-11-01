package service

import (
	"context"

	"github.com/icdb37/bfsm/internal/constx/featc"
	"github.com/icdb37/bfsm/internal/features/inventory/model"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
	coModel "github.com/icdb37/bfsm/internal/model"
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
	Search(ctx context.Context, req *coModel.SearchRequest[model.QueryCommodity]) (resp *coModel.SearchResponse[model.EntireCommodity], err error)
	Get(ctx context.Context, id string) (resp []*model.EntireCommodity, err error)
}

func Provide() {
	wire.ProvideName(featc.InventoryConsume, func() InventoryConsumer {
		repo, err := store.NewTable(&model.ConsumeBatch{})
		if err != nil {
			logx.Fatal("create repo failed", "feature", featc.InventoryConsume, "error", err)
		}
		return &consumeImpl{repo: repo}
	})
	wire.ProvideName(featc.InventoryProduce, func() InventoryProducer {
		repo, err := store.NewTable(&model.ProduceBatch{})
		if err != nil {
			logx.Fatal("create repo failed", "feature", featc.InventoryProduce, "error", err)
		}
		return &produceImpl{repo: repo}
	})
	wire.ProvideName(featc.InventoryInventory, func() InventoryInventory {
		repo, err := store.NewTable(&model.EntireCommodity{})
		if err != nil {
			logx.Fatal("create repo failed", "feature", featc.InventoryInventory, "error", err)
		}
		return &inventoryImpl{repo: repo}
	})
}
