FROM golang:1.21 AS builder

WORKDIR /app

COPY . /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

FROM scratch

COPY --from=builder /app/myapp .

ENTRYPOINT ["./myapp"]