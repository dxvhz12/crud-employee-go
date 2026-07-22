# ---- Stage 1: build ----
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# ---- Stage 2: run ----
FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/server .
COPY views ./views

RUN mkdir -p public/uploads

EXPOSE 9090

CMD ["./server"]