# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.work go.work
COPY core-service/go.mod core-service/go.mod
COPY shared/go.mod shared/go.mod

# Download dependencies
RUN go work sync

# Copy source code
COPY core-service/ core-service/
COPY shared/ shared/

# Build the application
RUN cd core-service && go build -o main main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary
COPY --from=builder /app/core-service/main .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]
