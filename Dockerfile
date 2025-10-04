# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git make

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build application (generated code should already exist)
RUN echo "Building Go application..." && \
    go build -v -o /app/server cmd/server/main.go && \
    echo "Build successful, checking binary..." && \
    ls -la /app/server && \
    test -f /app/server && echo "Binary exists!"

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy SQL files first
COPY --from=builder /app/sql ./sql

# Copy binary from builder
COPY --from=builder /app/server /usr/local/bin/server

# Verify binary was copied and make it executable
RUN echo "Verifying binary in runtime stage..." && \
    ls -la /app && \
    ls -la /usr/local/bin && \
    test -f /usr/local/bin/server && echo "Binary found!" || echo "Binary NOT found!" && \
    chmod +x /usr/local/bin/server && \
    /usr/local/bin/server --version || echo "Binary test complete"

# Expose port
EXPOSE 8080

# Run application
ENTRYPOINT ["/usr/local/bin/server"]
