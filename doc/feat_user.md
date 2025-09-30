# 用户管理

## 简述
对于用户信息的管理，支持：新增、删除、修改、查询
### 基本信息

**用户信息**

| 字段 | 类型 | 描述 |
| --- | --- | --- |
| id | int | id |
| alias | string | 别名，默认为姓名 |
| desc | string | 描述 |
| phone | string | 手机号 |
| password | string | 密码 |
| created | time | 创建时间 |
| updated | time | 更新时间 |
| ic | IdentityCard | 身份证 |
| bc | BankCard | 银行卡 |
| tags | array | 标签 |
| contacts | array | 紧急联系人 |

**身份证信息** 

| 字段 | 类型 | 描述 |
| --- | --- | --- |
| no | string | 身份证号 |
| name | string | 姓名 |
| sex | uint8 | 性别：0未知、1男、2女 |
| birth | date | 出生日期 |
| address | string | 地址 |
| validBeg | date | 有效期开始时间 |
| validEnd | date | 有效期结束时间 |

**银行卡信息**

| 字段 | 类型 | 描述 |
| --- | --- | --- |
| no | string | 银行卡号 |
| name | string | 银行卡持有人 |
| bank | string | 银行名称 |
| validBeg | date | 有效期开始时间 |
| validEnd | date | 有效期结束时间 |


**紧急联系人**

| 字段 | 类型 | 描述 |
| --- | --- | --- |
| name | string | 紧急联系人姓名 |
| phone | string | 紧急联系人手机号 |
| desc | string | 紧急联系人关系 |

示例
```json
{
    "id": "1",
    "alias": "张三",
    "desc": "zs简介",
    "phone": "13800000000",
    "password": "123456",
    "created": "2023-01-01T00:00:00Z",
    "updated": "2023-01-01T00:00:00Z",
    "ic": {
        "no": "44030519900101001X",
        "name": "张三",
        "sex": 1,
        "birth": "1990-01-01T00:00:00Z",
        "address": "中国 北京",
        "validBeg": "2023-01-01T00:00:00Z", 
        "validEnd": "2025-01-01T00:00:00Z"
    },
    "bc": {
        "no": "6222021234567890123",
        "name": "张三",
        "bank": "中国建设银行",
        "validBeg": "2023-01-01T00:00:00Z",
        "validEnd": "2025-01-01T00:00:00Z"
    },
    "tags": ["A"],
    "contacts": [
        {
            "name": "李四",
            "phone": "13600000000",
            "desc": "四弟"
        }
    ]
}
```

## 接口

`internal/features/builder` 实现用户相关操作的业务逻辑。

### 搜索
POST /api/v1/users/search

**请求**
```json
{
    "page": 0,
    "size": 10,
    "sorts": ["-updated"],
    "query": {
        "alias": "模糊搜索",
        "desc": "模糊搜索",
        "phone": "模糊搜索",
        "ic": {
            "no": "模糊搜索",
            "name": "模糊搜索",
            "sex": 0,
            "birth": "2023-01-01T00:00:00Z",
            "address": "模糊搜索"
        },
        "bc": {
            "no": "模糊搜索",
            "name": "模糊搜索",
            "bank": "模糊搜索"
        },
        "tags": "模糊搜索",
        "contacts": {
            "name": "模糊搜索",
            "phone": "模糊搜索"
        },
        "created": {
            "beg": "2023-01-01T00:00:00Z",
            "end": "2023-01-01T00:00:00Z"
        },
        "updated": {
            "beg": "2023-01-01T00:00:00Z",
            "end": "2023-01-01T00:00:00Z"
        }
    }
}
```

**响应**
```json
{
    "id": "1",
    "alias": "张三",
    "desc": "zs简介",
    "phone": "13800000000",
    "created": "2023-01-01T00:00:00Z",
    "updated": "2023-01-01T00:00:00Z",
    "ic": {
        "no": "44030519900101001X",
        "name": "张三",
        "sex": 1,
        "address": "中国 北京"
    },
    "tags": ["A"]
}
```

### 详情
GET /api/v1/users/:id

**响应**
```json
{
    "id": "1",
    "alias": "张三",
    "desc": "zs简介",
    "phone": "13800000000",
    "password": "123456",
    "created": "2023-01-01T00:00:00Z",
    "updated": "2023-01-01T00:00:00Z",
    "ic": {
        "no": "44030519900101001X",
        "name": "张三",
        "sex": 1,
        "birth": "1990-01-01T00:00:00Z",
        "address": "中国 北京",
        "validBeg": "2023-01-01T00:00:00Z", 
        "validEnd": "2025-01-01T00:00:00Z"
    },
    "bc": {
        "no": "6222021234567890123",
        "name": "张三",
        "bank": "中国建设银行",
        "validBeg": "2023-01-01T00:00:00Z",
        "validEnd": "2025-01-01T00:00:00Z"
    },
    "tags": ["A"],
    "contacts": [
        {
            "name": "李四",
            "phone": "13600000000",
            "desc": "四弟"
        }
    ]
}
```

### 创建
POST /api/v1/users

**请求**
```json
{
    "alias": "张三",
    "desc": "zs简介",
    "phone": "13800000000",
    "password": "123456",
    "created": "2023-01-01T00:00:00Z",
    "updated": "2023-01-01T00:00:00Z",
    "ic": {
        "no": "44030519900101001X",
        "name": "张三",
        "sex": 1,
        "birth": "1990-01-01T00:00:00Z",
        "address": "中国 北京",
        "validBeg": "2023-01-01T00:00:00Z", 
        "validEnd": "2025-01-01T00:00:00Z"
    },
    "bc": {
        "no": "6222021234567890123",
        "name": "张三",
        "bank": "中国建设银行",
        "validBeg": "2023-01-01T00:00:00Z",
        "validEnd": "2025-01-01T00:00:00Z"
    },
    "tags": ["A"],
    "contacts": [
        {
            "name": "李四",
            "phone": "13600000000",
            "desc": "四弟"
        }
    ]
}
```

**响应**

```json
{
    "id": 1
}
```

### 修改
PUT /api/v1/users/:id

**请求**
```json
{
    "alias": "张三",
    "desc": "zs简介",
    "phone": "13800000000",
    "password": "123456",
    "created": "2023-01-01T00:00:00Z",
    "updated": "2023-01-01T00:00:00Z",
    "ic": {
        "no": "44030519900101001X",
        "name": "张三",
        "sex": 1,
        "birth": "1990-01-01T00:00:00Z",
        "address": "中国 北京",
        "validBeg": "2023-01-01T00:00:00Z", 
        "validEnd": "2025-01-01T00:00:00Z"
    },
    "bc": {
        "no": "6222021234567890123",
        "name": "张三",
        "bank": "中国建设银行",
        "validBeg": "2023-01-01T00:00:00Z",
        "validEnd": "2025-01-01T00:00:00Z"
    },
    "tags": ["A"],
    "contacts": [
        {
            "name": "李四",
            "phone": "13600000000",
            "desc": "四弟"
        }
    ]
}
```

**响应**

```json
{
    "id": 1
}
```

### 删除
DELETE /api/v1/users/:id

**响应**

```json
{
    "id": 1,
    "alias": "张三"
}
```

