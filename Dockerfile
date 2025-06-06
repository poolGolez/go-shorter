# Step 1: Build the Go app
FROM golang:1.23-alpine AS builder

# Install git (for go modules) and build tools
RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main ./cmd

# Step 2: Create a lightweight image
FROM alpine:latest

# Install necessary certificates
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your app runs on (default Gin port)
EXPOSE 8080

# Run the executable
CMD ["./main"]
