# Dockerfile for Go API

# Step 1: Build the Go API binary
FROM golang:1.22.6-alpine AS builder
WORKDIR /app

# Copy the go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go binary using the main.go entry point
RUN go build -o ./bin/ecom ./go/api//main.go

# Step 2: Create the production container
FROM alpine:3.18
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/bin/ecom .
COPY --from=builder /app/.env .

# Expose the API port
EXPOSE 8080

# Run the Go API server
CMD ["./ecom"]
