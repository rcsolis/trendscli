FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/ ./...

FROM alpine:3.14 as release

WORKDIR /app

COPY --from=builder /app/bin/ ./
RUN chmod +x ./trendscli

ENTRYPOINT ["./trendscli"]
