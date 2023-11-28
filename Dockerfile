FROM golang:1.21.4-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件并下载依赖项
COPY go.mod go.sum ./
RUN go mod download

# 复制项目文件
COPY ./cmd ./cmd
COPY ./internal ./internal

# 构建可执行文件
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags go_json -o main ./cmd/web/main.go

# 第二阶段：创建运行镜像
# 使用轻量级的 alpine 镜像
FROM alpine:latest

# 在 alpine 中安装 ca-certificates 以支持 SSL
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 从构建器阶段复制编译后的应用
COPY --from=builder /app/main .

# 暴露端口
EXPOSE 9000

# 设置容器启动命令
CMD ["./main"]