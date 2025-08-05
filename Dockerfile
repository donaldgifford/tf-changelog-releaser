# Multi-stage Dockerfile for tf-changelog CLI
FROM golang:1.24-alpine AS builder

# Install necessary packages
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tf-changelog .

# Final stage - minimal image
FROM alpine:latest

# Install necessary packages for git operations and CA certificates
RUN apk --no-cache add ca-certificates git openssh-client curl

# Install terraform-docs (optional, can be configured via action input)
RUN curl -sSLo terraform-docs.tar.gz "https://github.com/terraform-docs/terraform-docs/releases/download/v0.17.0/terraform-docs-v0.17.0-linux-amd64.tar.gz" && \
    tar -xzf terraform-docs.tar.gz && \
    mv terraform-docs /usr/local/bin/ && \
    rm terraform-docs.tar.gz

# Create non-root user
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

WORKDIR /workspace

# Copy the binary from builder stage
COPY --from=builder /app/tf-changelog /usr/local/bin/tf-changelog

# Make sure binary is executable
RUN chmod +x /usr/local/bin/tf-changelog

# Switch to non-root user
USER appuser

# Set entrypoint
ENTRYPOINT ["/usr/local/bin/tf-changelog"]
CMD ["--help"]