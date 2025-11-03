package model

import (
	"github.com/icdb37/bfsm/internal/constx/featc"
	coModel "github.com/icdb37/bfsm/internal/model"
)

// ProduceBatch 库存采购
type ProduceBatch coModel.EntireBatch

// TableName 数据库表名
func (e *ProduceBatch) TableName() string {
	return featc.GetTableName(featc.InventoryProduce)
}

// GetFeature 特征
func (e *ProduceBatch) GetFeature() string {
	return featc.InventoryProduce
}

// ConsumeBatch 库存销售
type ConsumeBatch coModel.EntireBatch

// TableName 数据库表名
func (e *ConsumeBatch) TableName() string {
	return featc.GetTableName(featc.InventoryConsume)
}

// GetFeature 特征
func (e *ConsumeBatch) GetFeature() string {
	return featc.InventoryProduce
}
