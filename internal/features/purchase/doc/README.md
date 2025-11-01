# 采购功能

## 数据结构

### 采购订单

```json
{
  "id": "采购标识",
  "name": "采购名称",
  "desc": "采购描述",
  "created_at": "2025-10-28T00:00:00Z",
  "status": "PENDING", //状态：PENDING, CONFIRMED, COMPLETED, CANCELLED
  "total_amount": 123400, //总共费用，分
  "clear_amount": 123400, //已结算费用，分
  "companies": [
    {
      "company": {
        "id": "",
        "name": ""
      },
      "status": "PENDING", //状态：PENDING, CONFIRMED, COMPLETED, CANCELLED
      "cost": 123400, //总共费用，分
      "extras": [
        {
            "name": "额外费用名称",
            "desc": "额外费用描述",
            "cost": 23400 //额外费用，分
        }
      ],
      "commodities": [
        {
          "id": 1,
          "name": "商品名称",
          "desc": "商品描述",
          "spec": "商品规格",
          "size": "商品尺寸",
          "validity": 1,
          "price": 100,
          "count": 1000,
          "attrs": [
            {
              "name": "材质",
              "value": "纯棉"
            },
            {
              "name": "颜色",
              "value": "red"
            }
          ]
        }
      ]
    }
  ]
}
```

## 功能列表

### 基本功能

- 创建采购订单
- 查询采购订单
- 更新采购订单状态
- 删除采购订单
- 采购订单状态流转
- 采购商品入库

### 高级功能

- 采购订单商品导出（word模板化）
  - A4纸张
  - 三联清单