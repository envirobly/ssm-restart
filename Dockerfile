FROM golang:alpine AS builder

WORKDIR /app

# Copy the source code
COPY ssm_restart.go .

# Build the Go binary
RUN go build -o ssm_restart_agent ssm_restart.go

# Stage 2: Create a small final image
FROM alpine:latest

# Install necessary runtime dependencies
# RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/ssm_restart_agent .

EXPOSE 9009

CMD ["./ssm_restart_agent"]
