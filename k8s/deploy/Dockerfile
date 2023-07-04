# 从官方的node镜像开始
FROM node:13-alpine

# 创建目录 /app
RUN mkdir /app
# 设置工作目录
WORKDIR /app
# 在 /app 内安装 mysql 依赖
RUN npm i mysql