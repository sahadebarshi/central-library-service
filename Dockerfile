# Stage 1: Build the Go binary
# -----------------------------
FROM golang:1.24.5-alpine AS builder

# Set the working directory inside the container
WORKDIR /go/src/rest/cls

# Install git (some Go modules may require it)
RUN apk add --no-cache git

# Copy go.mod and go.sum separately to leverage Docker cache
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the rest of your application code
COPY . .

#   Build the Go binary
# - CGO_ENABLED=0: builds a static binary (you can set to 1 if you need C libs)
# - GOOS=linux / GOARCH=amd64: build for Linux AMD64
# - -ldflags="-s -w": strip debug info to reduce size
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o cls

# -----------------------------------
# Stage 2: Minimal Ubuntu runtime
# -----------------------------------
FROM ubuntu:22.04

# Create a non-root user
RUN useradd -m -s /bin/bash clsuser

# Set the working directory
WORKDIR /clsapp

# Copy the compiled binary from the builder
COPY --from=builder /go/src/rest/cls/cls .

#Ensure executable permissions
RUN chmod +x /clsapp/cls

# Switch to non-root user
USER clsuser

# Environment variables (can be overridden in Helm or kubectl)
ENV PORT=8443
ENV HEALTH_PORT=8083

# Expose ports for mTLS and health checks
EXPOSE 8443
EXPOSE 8083

# Start the application
ENTRYPOINT ["/clsapp/cls"]