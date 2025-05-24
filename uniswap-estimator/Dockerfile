FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/uniswap-estimator ./cmd/main.go

FROM alpine:latest

COPY --from=builder /app/uniswap-estimator /usr/local/bin/uniswap-estimator

ENV PORT=1337

EXPOSE 1337

CMD ["uniswap-estimator"]
