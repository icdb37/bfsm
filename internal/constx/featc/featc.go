package featc

import "strings"

const (
	// CommodityCommodity 商品管理-商品信息
	CommodityCommodity = "commodity.commodity"
	// CommodityTemplate 商品管理-商品模板
	CommodityTemplate = "commodity.template"

	// CompanyCompany - 企业管理-企业信息
	CompanyCompany = "company.company"
	// CompanyCommodity - 企业管理-企业商品
	CompanyCommodity = "company.commodity"

	// PurchasePurchase - 采购订单管理-采购订单信息
	PurchasePurchase = "purchase.purchase"

	// InventoryInventory - 库存管理-库存信息
	InventoryInventory = "inventory.inventory"
	// InventoryProduce - 库存管理-采购订单库存信息
	InventoryProduce = "inventory.produce"
	// InventoryConsume - 库存管理-销售订单库存信息
	InventoryConsume = "inventory.consume"

	// User - 用户模块
	User = "user"
)

// GetTableName 获取表名
func GetTableName(feature string) string {
	return strings.ReplaceAll(feature, ".", "_")
}
