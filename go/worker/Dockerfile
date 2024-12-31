# Step 1: Build the Go worker binary
FROM golang:1.22.6-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go worker
RUN go build -o worker .

# Step 2: Create the production container
FROM alpine:3.18
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/worker .
COPY --from=builder /app/.env .

# Expose the worker port if needed
EXPOSE 8081

# Run the Go worker
CMD ["./worker"]