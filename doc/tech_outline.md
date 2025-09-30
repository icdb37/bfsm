# 技术简述

整体采用 RESTful API 设计，前端采用 Vue3 + ElementPlus 实现，后端采用 golang 实现。

## 后端技术
- golang 后端语言
- sqlite3 数据库
- echo 框架
- xorm 数据库 ORM 框架
- jwt 认证框架
- bcrypt 密码加密框架
- validator 校验框架
- zap 日志框架

采用三层架构设计，分别为：
- 控制器层：处理 HTTP 请求，调用服务层处理业务逻辑，返回 HTTP 响应。
- 服务层：处理业务逻辑，调用数据访问层操作数据库。
- 数据访问层：操作数据库，执行 CRUD 操作。