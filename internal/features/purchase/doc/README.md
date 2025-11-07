# 采购功能

## 数据结构

* 企业商品表（表名：company_commodity）
  * id 唯一标识
  * company_id 企业标识
  * company_name 企业名称
  * commodity_id 商品标识
  * commodity_name 商品名称
  * commodity_spec 商品规格
  * commodity_size 商品尺寸
  * commodity_validity 商品有效期（天）
  * commodity_price 商品单价（分）
* 采购表（表名：purchase_batch）
  * id 采购标识
  * name 采购名称
  * desc 采购描述
  * created_at 采购创建时间
  * status 采购状态
  * total_amount 采购总金额（分）
  * clear_amount 已结算金额（分）
* 采购商品表（表名：purchase_commodity）
  * purchase_id 采购标识
  * company_id 企业标识
  * commodity_id 商品标识
  * count 商品数量
  * price 商品单价（分）

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
      "commodity": [
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


```sql

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