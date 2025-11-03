package service

import (
	"context"

	coModel "github.com/icdb37/bfsm/internal/model"
)

// InventoryProducer - 仓库生产服务接口
type InventoryProducer interface {
	Produce(ctx context.Context, info *coModel.EntireBatch) error
}

// InventoryConsumer - 仓库消费服务接口
type InventoryConsumer interface {
	Consume(ctx context.Context, info *coModel.EntireBatch) error
}
