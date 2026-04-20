# GoFrame Demo（GoFrame 示例项目集合）BY ZFW

ZFW 同学基于 [GoFrame](https://goframe.org) 框架的学习和实践项目集合。

## 📂 项目列表

| 项目                                    | 描述              | 技术栈                            |
| ------------------------------------- | --------------- | ------------------------------ |
| [quick-demo](./practices/quick-demo/) | GoFrame 快速入门示例  | GoFrame、MySQL、Swagger          |
| [star](./practices/star/)             | 星辰英语本（单词管理 API） | GoFrame、MySQL、JWT、Swagger、并发处理 |

## 🚀 快速开始

### 前置要求

- Go 1.18+
- MySQL 8.0+
- gf 命令行工具（可选，推荐）

### 安装 gf 工具（可选）

```bash
go install github.com/gogf/gf/cmd/gf/v2@latest
```

## 📖 项目说明

### 1. quick-demo（快速入门）

适合新手学习 GoFrame 的基础项目，包含：

- Hello World 接口
- 用户 CRUD 操作
- Swagger API 文档

详细文档：[quick-demo/README.md](./practices/quick-demo/README.md)

### 2. star（星辰英语本）

完整的单词学习 API 项目，包含：

- 用户注册/登录（JWT 认证）
- 单词 CRUD
- 批量创建（并发处理）
- 单词熟练度管理
- 分页查询

详细文档：[star/README.md](./practices/star/README.md)

## 🛠️ 技术栈

- **框架**: GoFrame v2
- **数据库**: MySQL
- **认证**: JWT（star 项目）
- **API 文档**: Swagger UI
- **并发处理**: errgroup（star 项目）

## 📝 常用命令

### 运行项目

```bash
cd practices/quick-demo
gf run main.go
# 或者
go run main.go
```

### 代码生成

```bash
# 生成 DAO
gf gen dao

# 生成 Controller
gf gen ctrl
```

### 查看 API 文档

启动项目后访问：<http://127.0.0.1:8000/swagger>

## 📚 学习资源

- GoFrame 官方文档：<https://goframe.org>
- GoFrame 快速入门：<https://goframe.org/quick>

## 📄 许可证

MIT License
