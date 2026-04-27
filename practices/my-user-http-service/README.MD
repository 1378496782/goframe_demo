# 我的用户 HTTP 服务

## 介绍

基于 GoFrame 框架构建的用户管理 HTTP 服务。该示例包含：

- 用户注册、登录和登出功能
- 基于会话的认证机制
- HTTP 中间件
- MySQL 数据库集成
- 自动生成的 OpenAPI 文档
- Service 层业务逻辑
- 现代化的前端界面（渐变背景、动画效果、响应式设计）

## 目录结构

```text
.
├── api/                    # API 定义
│   └── user/             # 用户服务 API
│       ├── user.go         # 用于API完整性实现的检测接口(自动生成)
│       └── v1/          # API定义，v1版本
├── hack/                 # 开发工具
│   └── config.yaml       # CLI 工具配置
├── internal/             # 内部包
│   ├── cmd/              # 命令定义
│   │   └── cmd.go        # 主命令与服务器设置
│   ├── consts/           # 常量
│   ├── controller/      # HTTP 控制器
│   │   └── user/       # 用户控制器
│   ├── dao/              # 数据访问对象(自动生成)
│   ├── logic/            # 业务逻辑(已迁移到 service)
│   ├── model/            # 数据模型
│   │   ├── do/            # 领域对象(自动生成)
│   │   └── entity/        # 数据库实体(自动生成)
│   ├── service/        # 业务逻辑
│   │   ├── bizctx/     # 业务上下文服务
│   │   ├── middleware/   # 中间件服务
│   │   ├── session/       # 会话服务
│   │   └── user/          # 用户服务
│   └── packed/         # 打包文件
├── manifest/           # 部署清单
│   ├── config/         # 配置文件
│   │   └── config.yaml    # 应用配置
│   ├── deploy/        # 部署文件
│   ├── docker/        # Docker 文件
│   └── sql/           # SQL 脚本
├── resource/           # 前端资源
│   ├── public/       # 静态资源
│   │   ├── html/      # HTML 页面
│   │   │   ├── index.html      # 首页
│   │   │   ├── login.html      # 登录页
│   │   │   ├── signup.html     # 注册页
│   │   │   ├── list.html       # 用户列表页
│   │   │   └── profile.html    # 个人资料页
│   │   └── resource/   # 资源文件
│   │       ├── css/       # CSS 样式
│   │       ├── js/        # JavaScript
│   │       └── image/     # 图片
│   └── template/     # 模板文件
├── utility/           # 工具类/工具函数
├── main.go          # 应用入口
├── go.mod           # Go 模块文件
├── go.sum           # Go 模块依赖锁定文件
└── Makefile         # 构建自动化
```

## 功能特性

### 用户管理

- 用户注册（`SignUp`：支持昵称自动用护照号
- 用户登录（`SignIn`）
- 用户登出（`SignOut`）
- 用户列表（`List`）
- 用户查询（`Profile`）
- 登录状态检查（`IsSignedIn`）

### 前端界面优化

- 现代化渐变背景（紫蓝配色）
- 响应式设计，支持各种屏幕尺寸
- 流畅的动画效果（页面滑入、按钮悬停、消息提示）
- 密码显示/隐藏切换功能
- 实时密码强度检测
- 加载状态按钮，防止重复提交
- 键盘快捷键支持（Ctrl+Enter 提交）
- 美观的表单设计，自适应屏幕宽度50%
- 错误/成功消息提示（带抖动/弹出动画）

### 认证与授权

- 基于会话的用户认证
- 自定义业务上下文注入
- 基于中间件的授权机制
- 受保护路由

### 接口文档

- 自动生成的OpenAPI规范
- 自动生成的Swagger UI位于/swagger

### 数据库集成

- MySQL 数据库连接
- DAO 模式数据访问
- 自动生成 DAO、DO和Entity代码
- 事务支持

## 环境要求

### 本地开发
- Go 1.23或更高版本
- MySQL 5.7或更高版本

### Docker 部署
- Docker
- Docker Compose (可选)

# 前置准备

### 配置 MySQL 数据库

使用 Docker 运行 MySQL 数据库：

```bash
docker run -d \
  --name mysql-user-service \
  -p 3306:3306 \
  -e MYSQL_ROOT_PASSWORD=12345678 \
  -e MYSQL_DATABASE=test \
  mysql:8.0
```

如果本地有 MySQL，可以改用端口冲突，可以：

```bash
docker run -d \
  --name mysql-user-service \
  -p 3307:3306 \
  -e MYSQL_ROOT_PASSWORD=12345678 \
  -e MYSQL_DATABASE=test \
  mysql:8.0
```

并修改 `config.yaml` 的端口为 3307。

### 初始化数据库

从旁边的例子项目复制 SQL 脚本或直接在项目执行 SQL：

```sql
CREATE TABLE `user`(
    `id`        int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'User ID',
    `passport`  varchar(45) NOT NULL COMMENT 'User Passport',
    `nickname`  varchar(45) NOT NULL COMMENT 'User Nickname',
    `password`  varchar(45) NOT NULL COMMENT 'User Password',
    `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
    `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
    `delete_at` datetime DEFAULT NULL COMMENT 'Deleted Time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_passport` (`passport`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

## 配置说明

在 `manifest/config/config.yaml` 中配置：

```yaml
server:
  address: ":8000"
  openapiPath: "/api.json"
  swaggerPath: "/swagger"

database:
  default:
    link: "mysql:root:12345678@tcp(127.0.0.1:3306)/test?loc=Asia/Shanghai&parseTime=true"
    timezone: "Asia/Shanghai"
```

## 界面预览

### 首页
![首页](./docs/images/home.png)

### 登录页
![登录页](./docs/images/login.png)

### 注册页
![注册页](./docs/images/signup.png)

### 用户列表页
![用户列表](./docs/images/list.png)

### 个人资料页
![个人资料](./docs/images/profile.png)

## 部署指南

### 方式一：使用一键部署脚本（推荐）

项目提供了 `deploy.sh` 一键部署脚本，最简单的部署方式：

```bash
# 1. 进入项目目录
cd goframe_demo/practices/my-user-http-service

# 2. 构建并启动服务
./deploy.sh start

# 3. 查看服务状态
./deploy.sh status

# 4. 查看日志
./deploy.sh logs
```

**访问地址：** http://localhost:8000/html/index.html

**常用命令：**
```bash
./deploy.sh start    # 启动服务
./deploy.sh stop     # 停止服务
./deploy.sh restart  # 重启服务
./deploy.sh status   # 查看状态
./deploy.sh logs     # 查看日志
./deploy.sh clean    # 清理所有资源
```

### 方式二：Docker Compose 部署

```bash
# 1. 构建 Linux 二进制文件
GOOS=linux GOARCH=amd64 go build -o temp/linux_amd64/main main.go

# 2. 构建 Docker 镜像
docker build -t my-user-service:latest -f manifest/docker/Dockerfile .

# 3. 启动服务
docker-compose up -d

# 4. 查看服务状态
docker-compose ps
```

### 方式三：Kubernetes 部署

项目已包含 Kubernetes 部署配置（位于 `manifest/deploy/kustomize/`）：

```bash
# 使用 Kustomize 部署
kubectl apply -k manifest/deploy/kustomize/overlays/develop
```

### 部署文件说明

| 文件 | 说明 |
|------|------|
| `docker-compose.yml` | Docker Compose 配置文件 |
| `deploy.sh` | 一键部署脚本 |
| `manifest/docker/Dockerfile` | Docker 镜像构建文件 |
| `manifest/config/config-docker.yaml` | Docker 环境专用配置 |
| `manifest/sql/init.sql` | 数据库初始化脚本 |

### Docker 环境变量

| 变量名 | 默认值 | 说明 |
|--------|--------|------|
| GF_SERVER_ADDRESS | :8000 | 服务监听地址 |
| GF_DATABASE_HOST | mysql:3306 | 数据库地址 |
| GF_DATABASE_NAME | test | 数据库名称 |
| GF_DATABASE_USER | root | 数据库用户名 |
| GF_DATABASE_PASS | 12345678 | 数据库密码 |

## 部署原理解析

### 🏗️ 整体架构

项目采用 **Docker Compose** 容器化部署，包含两个主要服务：

```
┌───────────────────────────────────────────────┐
│              User's Browser                   │
│    (访问 localhost:8000)                     │
└────────────────┬──────────────────────────────┘
                 │
                 ▼
┌───────────────────────────────────────────────┐
│      my-user-service (Go应用容器)             │
│         - 监听 8000 端口                      │
│         - 处理 HTTP 请求                      │
└────────────────┬──────────────────────────────┘
                 │
                 ▼
┌───────────────────────────────────────────────┐
│   user-service-mysql (MySQL数据库容器)        │
│         - 监听 3306 端口 (容器内)             │
│         - 映射到宿主 3307 端口               │
└───────────────────────────────────────────────┘
```

### 📦 关键部署文件详解

#### 1️⃣ Dockerfile - 构建应用镜像

```dockerfile
# 基础镜像：轻量级 Alpine Linux
FROM alpine:3.18

# 时区配置
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai

# 设置工作目录
ENV WORKDIR /app
WORKDIR $WORKDIR

# 复制资源文件
ADD resource /app/resource
ADD manifest /app/manifest
ADD ./temp/linux_amd64/main /app/main
RUN chmod +x /app/main

# 使用 Docker 专用配置文件
RUN mv /app/manifest/config/config-docker.yaml /app/manifest/config/config.yaml

# 启动应用
CMD ./main
```

**工作原理：**
- 基于 Alpine 3.18 构建（仅 5MB）
- 复制编译好的 Go 二进制文件（Linux/amd64）
- 复制前端资源文件
- 替换配置文件为 Docker 环境专用版本

#### 2️⃣ docker-compose.yml - 服务编排

```yaml
services:
  # MySQL 数据库服务
  mysql:
    image: mysql:8.0
    container_name: user-service-mysql
    environment:
      MYSQL_ROOT_PASSWORD: 12345678
      MYSQL_DATABASE: test
    ports:
      - "3307:3306"  # 容器内 3306 → 宿主 3307
    volumes:
      - mysql-data:/var/lib/mysql  # 数据持久化
      - ./manifest/sql:/docker-entrypoint-initdb.d  # 自动初始化
    healthcheck:
      test: ["CMD", "mysqladmin", "ping", "-h", "localhost", "-uroot", "-p12345678"]
      interval: 10s
      timeout: 5s
      retries: 5

  # Go 应用服务
  app:
    image: my-user-service:latest
    container_name: my-user-service
    ports:
      - "8000:8000"  # 容器内 8000 → 宿主 8000
    depends_on:
      mysql:
        condition: service_healthy  # 等 MySQL 健康再启动
```

**工作原理：**
- **服务依赖：** `app` 依赖 `mysql`，并等待 MySQL 健康检查通过
- **网络：** 两个容器在同一个 Docker 网络，可以通过服务名互相访问
- **数据持久化：** MySQL 数据存储在 Docker Volume，容器删除数据不丢失
- **端口映射：** 外部访问 8000 → 容器 8000，本地 MySQL 工具访问 3307 → 容器 3306

#### 3️⃣ config-docker.yaml - Docker 环境配置

```yaml
database:
  default:
    type: "mysql"
    host: "mysql"  # 关键！使用 Docker 服务名而非 IP
    port: "3306"
    user: "root"
    pass: "12345678"
    name: "test"
```

**核心原理：**
- Docker Compose 会自动创建一个内部 DNS
- 容器可以通过服务名（`mysql`）互相访问
- 不需要知道具体的 IP 地址

#### 4️⃣ deploy.sh - 部署脚本

**核心流程：**

```bash
# 1. 构建 Linux 二进制文件
GOOS=linux GOARCH=amd64 go build -o temp/linux_amd64/main main.go

# 2. 构建 Docker 镜像
docker build -t my-user-service:latest -f manifest/docker/Dockerfile .

# 3. 启动服务
docker-compose up -d
```

### 🔄 部署完整流程

```
┌───────────────────────────────────────────────────────────┐
│                    1. 本地编译                             │
│  GOOS=linux GOARCH=amd64 go build → Linux 二进制文件       │
└───────────────────────┬───────────────────────────────────┘
                        ▼
┌───────────────────────────────────────────────────────────┐
│                    2. 构建镜像                             │
│  docker build → 打包二进制 + 资源 + 配置 → my-user-service │
└───────────────────────┬───────────────────────────────────┘
                        ▼
┌───────────────────────────────────────────────────────────┐
│                    3. 启动容器                             │
│  docker-compose up → 启动 MySQL 和 App 容器               │
└───────────────────────┬───────────────────────────────────┘
                        ▼
┌───────────────────────────────────────────────────────────┐
│                    4. 服务就绪                             │
│  - MySQL 初始化数据库表                                    │
│  - App 连接 MySQL (通过服务名)                           │
│  - 监听 8000 端口                                         │
└───────────────────────────────────────────────────────────┘
```

### 🎯 关键技术点

#### 1️⃣ 跨平台编译
```bash
GOOS=linux GOARCH=amd64 go build
```
- 在 Mac 上编译出 Linux 可执行文件
- 这样二进制就能在 Alpine Linux 容器中运行

#### 2️⃣ Docker 网络
- 同一 Compose 项目的容器在同一网络
- 通过服务名（`mysql`）互相访问
- 自动服务发现

#### 3️⃣ 数据持久化
```yaml
volumes:
  - mysql-data:/var/lib/mysql
```
- 使用 Docker Volume 存储 MySQL 数据
- 容器重启/删除数据不丢失

#### 4️⃣ 健康检查
```yaml
healthcheck:
  test: ["CMD", "mysqladmin", "ping", ...]
```
- 确保 MySQL 真正就绪后才启动应用
- 避免连接错误

### 🚀 为什么这样部署？

| 优势 | 说明 |
|------|------|
| **环境一致性** | 开发/测试/生产环境完全一致 |
| **依赖隔离** | MySQL 和应用分别容器化，互不影响 |
| **快速部署** | 一条命令启动所有服务 |
| **易于扩展** | 可以轻松添加 Redis、Nginx 等服务 |
| **便于迁移** | 可以在任何 Docker 环境运行 |

### 💡 总结

这个部署方案的核心思想是：
1. **编译一次** - 构建出 Linux 二进制
2. **打包镜像** - 把所有依赖打包进 Docker
3. **编排启动** - 用 Docker Compose 管理服务关系
4. **服务发现** - 通过服务名互相通信

这样无论你在本地、云服务器、还是 Kubernetes 上，都可以用同样的方式运行！

## 使用说明

### 运行服务器

启动 HTTP 服务器：

```bash
cd goframe_demo/practices/my-user-http-service

# 使用 gf run
gf run main.go

# 或直接 go run
go run main.go
```

服务器将在端口 8000 上启动，提供以下接口：

- Swagger UI：<http://localhost:8000/swagger>
- OpenAPI 规范：<http://localhost:8000/api.json>
- 首页：<http://localhost:8000/html/index.html>
- 登录页：<http://localhost:8000/html/login.html>
- 注册页：<http://localhost:8000/html/signup.html>
- 用户列表：<http://localhost:8000/html/list.html>

### 测试服务

#### 使用 Swagger UI

浏览器打开 [http://localhost:8000/swagger](http://localhost:8000/swagger。)

#### 使用 cURL

1. **注册**：

```bash
curl -X POST http://localhost:8000/user/sign-up \
  -H "Content-Type: application/json" \
  -d '{
    "Passport": "user001",
    "Password": "123456",
    "Password2": "123456",
    "Nickname": "测试用户"
  }'
```

1. **登录**：

```bash
curl -X POST http://localhost:8000/user/sign-in \
  -H "Content-Type: application/json" \
  -d '{
    "Passport": "user001",
    "Password": "123456"
  }' \
  -c cookies.txt
```

1. **检查登录状态**：

```bash
curl -X POST http://localhost:8000/user/is-signed-in \
  -b cookies.txt
```

1. **获取用户列表**：

```bash
curl -X GET http://localhost:8000/user/list
```

1. **登出**：

```bash
curl -X POST http://localhost:8000/user/sign-out \
  -b cookies.txt
```

## 接口文档

### 公开接口

#### SignUp

**POST** `/user/sign-up`

注册新用户。

**请求体：**

```json
{
  "Passport": "string (6-16 字符，必填)",
  "Password": "string (6-16 字符，必填)",
  "Password2": "string (必须与 Password 相同，必填)",
  "Nickname": "string (可选)"
}
```

**响应：**

```json
{
  "code": 0,
  "message": "",
  "data": {}
}
```

#### SignIn

**POST** `/user/sign-in`

用户认证并创建会话。

**请求体：**

```json
{
  "Passport": "string (必填)",
  "Password": "string (必填)"
}
```

#### List

**GET** `/user/list`

获取用户列表。

**响应：**

```json
{
  "code": 0,
  "message": "",
  "data": {
    "UserList": [...]
  }
}
```

## 实现细节

### Controller 层

`internal/controller/user/` 处理 HTTP 请求：

- 接收并验证请求参数
- 调用 Service 层处理业务逻辑
- 返回标准化的 JSON 响应

### Service 层

`internal/service/` 包含业务逻辑：

#### 用户服务

- 用户创建带验证
- 认证和会话管理
- 用户资料查询
- 事务支持

#### 中间件服务

- 自定义业务数据的上下文注入
- 受保护路由的认证中间件
- CORS 跨域支持

#### 会话服务

- 用户会话管理
- 会话存储和检索

## 注意事项

- 启动应用前请确保 MySQL 已运行
- 默认数据库凭据为 root:12345678 (生产环境请修改)
- 服务默认使用端口 8000
- 如果端口冲突可以配置 config.yaml 修改地址

