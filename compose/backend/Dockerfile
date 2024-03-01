FROM golang:1.21 AS builder
COPY ./backend /build
RUN cd /build && go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -installsuffix cgo -o /backend .

FROM alpine
COPY --from=builder /backend /app/backend
COPY ./backend/config /app/config

ENTRYPOINT ["/app/backend"]