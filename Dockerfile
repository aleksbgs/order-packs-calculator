# Stage 1: Build
FROM golang:1.23.4-alpine AS builder

WORKDIR /app

# Install OS dependencies
RUN apk add --no-cache git

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./

# Download all Go dependencies specified in go.mod
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -o order-packs-calculator ./cmd/api

# Verify the binary exists and is executable
RUN test -f /app/order-packs-calculator || (echo "Binary not found after build!" && exit 1)
RUN chmod +x /app/order-packs-calculator
RUN ls -l /app/order-packs-calculator

# Stage 2: Runtime
FROM alpine:latest

WORKDIR /root/

# Copy the built binary and resources
COPY --from=builder /app/order-packs-calculator .
COPY --from=builder /app/web ./web
COPY --from=builder /app/config.yaml .

# Verify the binary exists in the runtime stage
RUN test -f /root/order-packs-calculator || (echo "Binary not found in runtime stage!" && exit 1)
RUN ls -l /root/order-packs-calculator

# Ensure the binary has executable permissions
RUN chmod +x /root/order-packs-calculator

EXPOSE 3000

CMD ["./order-packs-calculator"]