FROM alpine:3.18

###############################################################################
#                                INSTALLATION
###############################################################################

# 设置时区
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai

ENV WORKDIR                 /app
WORKDIR $WORKDIR

# 复制资源文件
ADD resource                $WORKDIR/resource
ADD manifest                $WORKDIR/manifest
ADD ./temp/linux_amd64/main $WORKDIR/main
RUN chmod +x $WORKDIR/main

# 替换配置文件为 Docker 专用版本
RUN mv $WORKDIR/manifest/config/config-docker.yaml $WORKDIR/manifest/config/config.yaml

###############################################################################
#                                   START
###############################################################################
CMD ./main
