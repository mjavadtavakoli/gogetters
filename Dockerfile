# Step 1: Build stage
FROM golang:1.24.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# Step 2: Run stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .
COPY internal ./internal

EXPOSE 8080

CMD ["./server"]
