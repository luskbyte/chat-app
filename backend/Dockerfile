# Build stage
FROM golang:1.21-alpine AS builder

# Install git (needed for some Go dependencies)
RUN apk add --no-cache git

WORKDIR /app

# Copy all source code
COPY . .

# Tidy up dependencies (creates go.sum based on imports)
RUN go mod tidy

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o chat-app-backend cmd/server/main.go

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/chat-app-backend .

EXPOSE 8080

CMD ["./chat-app-backend"]

