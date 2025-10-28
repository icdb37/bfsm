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
  "cost": 123400, //总共费用，分
  "companies": [
    {
      "company": {
        "id": "",
        "name": ""
      },
      "status": "PENDING", //状态：PENDING, CONFIRMED, COMPLETED, CANCELLED
      "cost": 123400, //总共费用，分
      "extra_costs": [
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
