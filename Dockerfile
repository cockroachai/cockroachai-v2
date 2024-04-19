# 使用官方 Golang Alpine 镜像作为基础镜像
FROM golang:alpine as builder

# 安装 ca-certificates，这样你的应用就可以访问带有 SSL 证书的站点了
RUN apk update && apk add --no-cache ca-certificates tzdata

# 设置工作目录
WORKDIR /app

# 将 go.mod 和 go.sum 文件复制到工作目录
COPY go.mod go.sum ./

# 下载所有依赖项
RUN go mod download

# 将源代码复制到工作目录
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 使用 scratch 作为基础镜像
FROM scratch

# 设置工作目录
WORKDIR /app

# 从 builder 镜像中复制 /etc/ssl/certs 到当前镜像中，这样你的应用就可以访问带有 SSL 证书的站点了
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# 从 builder 镜像中复制 /usr/share/zoneinfo 到当前镜像中，这样你的应用就可以支持环境变量指定的时区了
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# 从 builder 镜像中复制应用到当前镜像中
COPY --from=builder /app/main /app/main

# 从 builder 镜像中复制目录到当前镜像中
COPY --from=builder /app/resource /app/resource

# 指定环境变量 TZ，你可以在运行 Docker 容器时通过 -e 参数来覆盖这个值
ENV TZ=Asia/Shanghai

# 指定容器启动时运行的命令
ENTRYPOINT ["/app/main"]
