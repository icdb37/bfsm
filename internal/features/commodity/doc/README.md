# 商品功能

## 数据类型

### 商品

商品是货物抽象的描述，基本信息如下：
有效期：
* 0 ：永久有效
* n ：n年有效

```json
{
    "id": 1,
    "name": "商品名称",
    "desc": "商品描述",
    "spec": "商品规格",
    "size": "商品尺寸",
    "validity": 1,
    "price": 100,
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
```

### 商品模板

某一公司下的商品模板，定义了商品的基本属性，例如：

```json
{
    "id": "",
    "company_id": "",
    "company_name": "",
    "commodities": [
        {
            "name": "商品名称",
            "desc": "商品描述",
            "spec": "商品规格",
            "size": "商品尺寸",
            "validity": 1,
            "price": 100,
            "count": 100,
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
```