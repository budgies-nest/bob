FROM golang:1.24.2-alpine AS builder
WORKDIR /app
COPY . .
#COPY go.mod .

RUN <<EOF
go mod download
go mod tidy 
go build -o mcp-http-server
EOF

FROM scratch
#FROM debian:12-slim

WORKDIR /app
COPY --from=builder /app/mcp-http-server .


CMD ["./mcp-http-server"]