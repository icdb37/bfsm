package service

import (
	"context"

	coModel "github.com/icdb37/bfsm/internal/model"
)

// InventorySaver - 仓库库存保存接口
type InventorySaver interface {
	Save(ctx context.Context, info *coModel.BatchGoods) error
}
