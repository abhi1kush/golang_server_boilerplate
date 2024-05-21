# Stage 1: Build the Go binary
FROM golang:1.21-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN go build -o myapp ./cmd/server

# Stage 2: Create a lightweight image with the Go binary
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/myapp .

# Ensure configs directory exists
RUN mkdir /configs

# Copy the configuration files from the builder stage
COPY --from=builder /app/configs /configs

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the binary
CMD ["./myapp"]
