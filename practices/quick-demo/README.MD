# GoFrame Quick Demo（快速入门示例）BY ZFW

一个基于 GoFrame 框架的快速入门示例项目，适合 ZFW 同学在新手时期学习使用。

## 功能特性

- ✅ Hello World 接口
- ✅ 用户管理
  - 创建用户
  - 根据 ID 获取用户
  - 获取用户列表

## 技术栈

- **框架**: [GoFrame](https://goframe.org) v2
- **数据库**: MySQL
- **API 文档**: Swagger UI

## 项目结构

```
quick-demo/
├── api/                    # API 定义层
│   ├── hello/             # Hello 模块
│   └── user/             # User 模块
├── internal/
│   ├── cmd/               # 程序入口
│   ├── controller/        # 控制器层
│   ├── dao/               # 数据访问层
│   ├── model/             # 模型层
│   │   ├── do/            # 数据对象
│   │   └── entity/        # 实体对象
│   └── logic/             # 业务逻辑层
├── manifest/
│   ├── config/            # 配置文件
│   ├── deploy/            # K8s 部署配置
│   └── docker/            # Docker 相关
└── main.go
```

## 快速开始

### 环境要求

- Go 1.18+
- MySQL 8.0+

### 1. 克隆项目

```bash
cd /Users/bytedance/GolandProjects/goframe_shop/goframe_demo/practices/quick-demo
```

### 2. 配置数据库

修改 `manifest/config/config.yaml`：

```yaml
database:
  default:
    link: "mysql:root:123abcABC@tcp(127.0.0.1:3306)/test"
```

### 3. 初始化数据库

创建 `test` 数据库，并导入以下表结构：

```sql
CREATE TABLE user (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(10) NOT NULL,
    status TINYINT NOT NULL DEFAULT 0 COMMENT '0:OK 1:Disabled',
    age INT UNSIGNED NOT NULL
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

### Hello 模块

| 方法  | 路径       | 描述          | 鉴权 |
| --- | -------- | ----------- | -- |
| GET | `/hello` | Hello World | ❌  |

### User 模块

| 方法   | 路径           | 描述     | 鉴权 |
| ---- | ------------ | ------ | -- |
| POST | `/user`      | 创建用户   | ❌  |
| GET  | `/user/{id}` | 获取用户信息 | ❌  |
| GET  | `/user/list` | 获取用户列表 | ❌  |

## 使用示例

### 1. Hello World

```bash
curl http://127.0.0.1:8000/hello
```

返回：

```
Hello World! zfw666666111
```

### 2. 创建用户

```bash
curl -X POST http://127.0.0.1:8000/v1/user \
  -H "Content-Type: application/json" \
  -d '{
    "name": "张三",
    "age": 25
  }'
```

返回：

```json
{
  "code": 0,
  "message": "",
  "data": {
    "user": {
      "id": 1,
      "name": "张三",
      "status": 0,
      "age": 25
    }
  }
}
```

### 3. 获取用户

```bash
curl http://127.0.0.1:8000/v1/user/1
```

### 4. 获取用户列表

```bash
curl http://127.0.0.1:8000/v1/user/list
```

## 配置说明

配置文件：`manifest/config/config.yaml`

```yaml
server:
  address:     ":8000"
  openapiPath: "/api.json"
  swaggerPath: "/swagger"

logger:
  level : "all"
  stdout: true

database:
  default:
    link: "mysql:root:123abcABC@tcp(127.0.0.1:3306)/test"
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

## 学习资源

- 官方文档：<https://goframe.org>
- 快速入门：<https://goframe.org/quick>

## 许可证

MIT License
