FROM golang:1.21 As builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -installsuffix cgo -o webapp .

# 第二阶段：使用官方的 scratch 镜像创建一个最小的运行环境
FROM scratch

# 从构建器中复制构建的应用程序
COPY --from=builder /app/webapp .

# 指定运行时的命令
CMD ["./webapp"]