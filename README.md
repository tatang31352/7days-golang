# 7天用Go从零实现系列

7天能写什么呢? 类似gin的web框架？类似groupcache的分布式缓存？或者一个简单的Python解释器？希望这个仓库能给你答案。

推荐先阅读[Go语言简明教程](https://geektutu.com/post/quick-golang.html),一篇文章了解Go的基本语法,并发编程,依赖管理等内容。

### 7天用Go从零实现Web框架-Gee

Gee是一个模仿[gin](https://github.com/gin-gonic/gin)实现等Web框架,[Go Gin简明教程](https://geektutu.com/post/quick-go-gin.html)可以快速入门。

- 第一天：前置知识（http.Handler接口｜Code
- 第二天：上下文设计（Context）｜Code
- 第三天：Tire路由（Router）｜Code
- 第四天：分组控制（Group）｜Code
- 第五天：中间件（Middleware) | Code
- 第六天：HTML模版（Template）| Code
- 第七天：错误恢复（Panic Recover）｜Code

### 7天用Go从零实现分布式缓存GeeCache

- 第一天：LRU缓存淘汰策略｜Code
- 第二天：单机并发缓存｜Code
- 第三天：HTTP服务端｜Code
- 第四天：一致性哈希（Hash）| Code
- 第五天：分布式节点｜Code
- 第六天：防止缓存击穿｜Code
- 第七天：使用Protobuf通信｜Code

### 7天用Go从零实现ORM框架GeeORM

GeeORM是一个模仿[grom](https://github.com/jinzhu/gorm)和[xrom](https://github.com/go-xorm/xorm)的ORM框架

grom准备推出完全重写的v2版本(目前还在开发中)，相对grom-v1来说，xrom的设计更容易理解，所以geeorm接口设计上主要参考了xorm，
一些细节实现上参考了gorm。

- 第一天：database/sql基础｜Code
- 第二天：对象表结构映射｜Code
- 第三天：记录新增和查询｜Code
- 第四天：链式操作与更新删除｜Code
- 第五天：实现钩子(Hooks) | Code
- 第六天：支持事务（Transaction）｜Code
- 第七天：数据库迁移：数据库迁移（Migrate）｜Code

### WebAssembly使用示例

具体的实践过程记录在[Go WebAssembly简明教程](https://geektutu.com/post/quick-go-wasm.html)。

- 示例一：Hello World｜Code
- 示例二：注册函数｜Code
- 示例三：操作DOM｜Code
- 示例四：回调函数｜Code




