#!/bin/bash

echo "====================================="
echo "用户管理系统 - 部署脚本"
echo "====================================="

# 颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 项目根目录
PROJECT_DIR=$(cd "$(dirname "$0")" && pwd)
cd "$PROJECT_DIR"

case "$1" in
  build)
    echo -e "${GREEN}[1/2] 构建 Linux 二进制文件...${NC}"
    mkdir -p temp/linux_amd64
    GOOS=linux GOARCH=amd64 go build -o temp/linux_amd64/main main.go
    if [ $? -eq 0 ]; then
      echo -e "${GREEN}✓ 二进制文件构建成功${NC}"
    else
      echo -e "${RED}✗ 二进制文件构建失败${NC}"
      exit 1
    fi

    echo -e "${GREEN}[2/2] 构建 Docker 镜像...${NC}"
    docker build -f manifest/docker/Dockerfile -t my-user-service:latest .
    if [ $? -eq 0 ]; then
      echo -e "${GREEN}✓ Docker 镜像构建成功${NC}"
      echo -e "${YELLOW}镜像名称: my-user-service:latest${NC}"
    else
      echo -e "${RED}✗ Docker 镜像构建失败${NC}"
      exit 1
    fi
    ;;

  start)
    echo -e "${GREEN}启动服务...${NC}"
    if [ ! -f temp/linux_amd64/main ]; then
      echo -e "${YELLOW}未找到二进制文件，先构建...${NC}"
      $0 build
    fi
    docker-compose up -d
    echo -e "${GREEN}✓ 服务已启动${NC}"
    echo -e "${YELLOW}访问地址: http://localhost:8000/html/index.html${NC}"
    ;;

  stop)
    echo -e "${GREEN}停止服务...${NC}"
    docker-compose down
    echo -e "${GREEN}✓ 服务已停止${NC}"
    ;;

  restart)
    echo -e "${GREEN}重启服务...${NC}"
    $0 stop
    $0 start
    ;;

  logs)
    echo -e "${GREEN}查看日志...${NC}"
    docker-compose logs -f
    ;;

  status)
    echo -e "${GREEN}服务状态...${NC}"
    docker-compose ps
    ;;

  clean)
    echo -e "${YELLOW}清理资源...${NC}"
    docker-compose down -v
    rm -rf temp
    docker rmi my-user-service:latest 2>/dev/null
    echo -e "${GREEN}✓ 清理完成${NC}"
    ;;

  k8s)
    echo -e "${GREEN}部署到 Kubernetes...${NC}"
    echo -e "${YELLOW}确保已配置 kubectl${NC}"
    kubectl apply -k manifest/deploy/kustomize/overlays/develop
    echo -e "${GREEN}✓ 已部署到 K8s${NC}"
    ;;

  *)
    echo "用法: $0 {build|start|stop|restart|logs|status|clean|k8s}"
    echo ""
    echo "命令说明:"
    echo "  build    - 构建 Docker 镜像"
    echo "  start    - 启动服务（MySQL + App）"
    echo "  stop     - 停止服务"
    echo "  restart  - 重启服务"
    echo "  logs     - 查看日志"
    echo "  status   - 查看状态"
    echo "  clean    - 清理所有资源"
    echo "  k8s      - 部署到 Kubernetes"
    exit 1
    ;;
esac

exit 0
