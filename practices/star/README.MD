# Star（星辰英语本）BY ZFW

一个基于 GoFrame 框架开发的单词本 API 服务，集成了飞书功能。

## 功能特性

### 用户功能

- ✅ 用户注册
- ✅ 用户登录（JWT Token）
- ✅ 获取用户信息
- ✅ 获取用户列表
- ✅ 删除用户

### 单词功能

- ✅ 创建单个单词
- ✅ 批量创建单词
- ✅ 分页查询单词列表
- ✅ 根据 ID 查询单个单词
- ✅ 更新单词（PUT）
- ✅ 删除单词
- ✅ 设置单词熟练度（PATCH，1-5 级）

### 飞书集成功能 ✨

- ✅ **导出单词到飞书多维表格** - 批量导出单词数据到多维表格
- ✅ **发送学习提醒** - 通过飞书消息发送学习提醒

## 技术栈

- **框架**: [GoFrame](https://goframe.org) v2
- **数据库**: MySQL
- **认证**: JWT Token
- **API 文档**: Swagger UI
- **飞书集成**: 飞书开放平台 API（多维表格、消息发送）

## 飞书集成实现说明 ✨

### 整体架构

```
star/
├── api/lark/               # API 定义层
│   └── v1/
│       └── lark.go         # 定义 API 请求/响应结构
├── internal/
│   ├── controller/lark/    # 控制器层
│   │   └── lark_v1_export_words.go
│   └── logic/lark/        # 业务逻辑层
│       └── lark.go        # 核心业务逻辑
└── utility/lark/          # API 工具类
    └── client.go         # 飞书 API 客户端封装
```

### 模块职责

#### 1. API 定义层 (`api/lark/v1/lark.go`)
- 定义请求结构：`ExportWordsReq`、`SendReminderReq`
- 定义响应结构：`ExportWordsRes`、`SendReminderRes`
- 使用 GoFrame 的 API 定义规范

#### 2. 控制器层 (`internal/controller/lark/lark_v1_export_words.go`)
- 接收 HTTP 请求
- 从 JWT Token 中获取用户 ID
- 调用业务逻辑层
- 返回响应结果

#### 3. 业务逻辑层 (`internal/logic/lark/lark.go`)
- **导出单词逻辑**
  - 从数据库查询用户的单词列表
  - 组织数据格式（表头 + 数据行）
  - 调用工具类写入多维表格
  - 返回表格链接和导出数量
  
- **发送提醒逻辑**
  - 查询用户待复习的单词
  - 构建提醒消息内容
  - 调用工具类发送飞书消息

#### 4. 工具类层 (`utility/lark/client.go`)
- **`getTenantAccessToken`** - 获取飞书租户访问令牌
- **`AppendSpreadsheetRows`** - 追加数据到多维表格
  - 先调用 API 获取表列表，匹配表名获取表 ID
  - 构建记录数据（fields 格式）
  - 调用 `batch_create` 接口批量写入
- **`SendMessage`** - 发送飞书消息
  - 构建消息内容（text 格式）
  - 调用飞书消息发送 API

### 关键代码结构

#### 配置文件 (`manifest/config/config.yaml`)
```yaml
lark:
  appId: "cli_xxxxxxxxxxxx"
  appSecret: "xxxxxxxxxxxxx"
  defaultReceiveId: "ou_xxxxxxxxxx"
  baseToken: "Cx6ubIEy3apIFVsssQdleh4tgeb"
```

#### 核心流程示例（导出单词）
1. 用户请求导出 → Controller 接收
2. 从 Token 提取 userId → Logic 层查询数据库
3. 组织数据（单词、定义、例句...）→ 调用 Client.AppendSpreadsheetRows
4. Client 获取表列表 → 匹配表名获取表 ID → 调用飞书 API
5. 返回结果给用户

### 飞书 API 调用要点
- **多维表格 API**：使用 `/open-apis/bitable/v1/apps/{app_token}/tables/{table_id}/records/batch_create`
- **消息发送 API**：使用 `/open-apis/im/v1/messages?receive_id_type=open_id`
- **认证方式**：使用租户访问令牌（tenant_access_token）
- **错误处理**：统一捕获错误，记录日志，返回友好提示

## 项目结构

```
star/
├── api/                    # API 定义层
│   ├── account/
│   ├── users/
│   ├── words/
│   └── lark/              # 飞书 API 定义
├── internal/
│   ├── cmd/               # 程序入口
│   ├── controller/        # 控制器层
│   ├── logic/             # 业务逻辑层
│   │   └── lark/         # 飞书业务逻辑
│   ├── dao/               # 数据访问层
│   ├── model/             # 模型层
│   │   ├── do/            # 数据对象
│   │   └── entity/        # 实体对象
│   └── middleware/        # 中间件（JWT 认证）
├── manifest/
│   └── config/            # 配置文件
├── utility/
│   └── lark/              # 飞书 API 工具类
└── main.go
```

## 快速开始

### 环境要求

- Go 1.18+
- MySQL 8.0+

### 1. 克隆项目

```bash
cd /Users/bytedance/GolandProjects/goframe_shop/goframe_demo/practices/star
```

### 2. 配置数据库

修改 `manifest/config/config.yaml`：

```yaml
database:
  default:
    link: "mysql:root:123abcABC@tcp(127.0.0.1:3306)/star?loc=Local"
```

### 3. 初始化数据库

创建 `star` 数据库，并导入以下表结构：

```sql
CREATE TABLE users (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(100) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME
);

CREATE TABLE words (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    uid INT UNSIGNED NOT NULL COMMENT '用户 ID',
    word VARCHAR(100) NOT NULL,
    definition VARCHAR(300) NOT NULL,
    example_sentence VARCHAR(300) NOT NULL,
    chinese_translation VARCHAR(300) NOT NULL,
    pronunciation VARCHAR(100) NOT NULL,
    proficiency_level TINYINT NOT NULL DEFAULT 1 COMMENT '1-5',
    created_at DATETIME,
    updated_at DATETIME,
    INDEX idx_uid (uid)
);
```

### 4. 安装依赖

```bash
go mod download
```

### 5. 运行项目

```bash
gf run main.go
# 或者
go run main.go
```

### 6. 访问 API 文档

浏览器打开：<http://127.0.0.1:8000/swagger>

## API 接口

### 用户模块

| 方法     | 路径                   | 描述         | 鉴权 |
| ------ | -------------------- | ---------- | -- |
| POST   | `/v1/users/register` | 用户注册       | ❌  |
| POST   | `/v1/users/login`    | 用户登录       | ❌  |
| GET    | `/v1/users/{id}`     | 获取用户信息     | ✅  |
| GET    | `/v1/users/list`     | 获取用户列表     | ✅  |
| DELETE | `/v1/users/{id}`     | 删除用户       | ✅  |
| GET    | `/v1/account/info`   | 获取当前登录用户信息 | ✅  |

### 单词模块

| 方法     | 路径                      | 描述       | 鉴权 |
| ------ | ----------------------- | -------- | -- |
| POST   | `/v1/word`              | 创建单词     | ✅  |
| POST   | `/v1/word/create/batch` | 批量创建单词   | ✅  |
| GET    | `/v1/words`             | 分页获取单词列表 | ✅  |
| GET    | `/v1/word/{id}`         | 获取单个单词   | ✅  |
| PUT    | `/v1/word/{id}`         | 更新单词     | ✅  |
| DELETE | `/v1/words/{id}`        | 删除单词     | ✅  |
| PATCH  | `/v1/words/{id}/level`  | 设置单词熟练度  | ✅  |

## 使用示例

### 1. 用户注册

```bash
curl -X POST http://127.0.0.1:8000/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "demo",
    "password": "123456",
    "email": "demo@example.com"
  }'
```

### 2. 用户登录

```bash
curl -X POST http://127.0.0.1:8000/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "demo",
    "password": "123456"
  }'
```

返回 Token：

```json
{
  "code": 0,
  "message": "",
  "data": {
    "token": "eyJhbGc...xxx"
  }
}
```

### 3. 创建单词

```bash
curl -X POST http://127.0.0.1:8000/v1/word \
  -H "Content-Type: application/json" \
  -H "Authorization: eyJhbGc...xxx" \
  -d '{
    "word": "hello",
    "definition": "Used as a greeting",
    "example_sentence": "Hello, nice to meet you!",
    "chinese_translation": "你好",
    "pronunciation": "/həˈləʊ/",
    "proficiency_level": 5
  }'
```

### 4. 获取单词列表

```bash
curl -X GET "http://127.0.0.1:8000/v1/words?page=1&size=10" \
  -H "Authorization: eyJhbGc...xxx"
```

### 5. 设置单词熟练度

```bash
curl -X PATCH http://127.0.0.1:8000/v1/words/1/level \
  -H "Content-Type: application/json" \
  -H "Authorization: eyJhbGc...xxx" \
  -d '{
    "level": 3
  }'
```

## 配置说明

配置文件：`manifest/config/config.yaml`

```yaml
server:
  address: ":8000"         # 服务端口
  openapiPath: "/api.json" # OpenAPI 文档
  swaggerPath: "/swagger"  # Swagger UI

database:
  default:
    link: "mysql:root:123abcABC@tcp(127.0.0.1:3306)/star?loc=Local"
    debug: true

# 飞书配置 ✨
lark:
  appId: "cli_a940a8f1f439dbca"       # 飞书应用 ID
  appSecret: "cGP2Tz38zr4bL81Musgi1d7YktlMuY5w"  # 飞书应用密钥
  defaultReceiveId: "ou_873d43281881570cdc0757ba3b977775"  # 默认接收人（飞书 Open ID）
  baseToken: "Cx6ubIEy3apIFVsssQdleh4tgeb"  # 默认多维表格 Token
```

## 开发说明

### 重新生成 DAO

```bash
gf gen dao
```

### 重新生成 Controller

```bash
gf gen ctrl
```

## 许可证

MIT License
