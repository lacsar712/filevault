# 官方 Go 镜像，保留完整工具链（评测用）
FROM golang:1.22

WORKDIR /app

# 本项目无外部依赖：无 go.sum，跳过 go mod download
COPY go.mod ./
COPY . .

# 预编译一次，把编译缓存留在镜像里
RUN go build ./...

CMD ["bash"]
