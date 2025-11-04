package model

import (
	"github.com/icdb37/bfsm/internal/constx/featc"
	coModel "github.com/icdb37/bfsm/internal/model"
)

// ProduceBatch 库存采购
type ProduceBatch coModel.ProduceBatch

// TableName 数据库表名
func (e *ProduceBatch) TableName() string {
	return featc.GetTableName(featc.InventoryProduce)
}

// GetFeature 特征
func (e *ProduceBatch) GetFeature() string {
	return featc.InventoryProduce
}
