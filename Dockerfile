# Build Stage: Use an official Go image to build the binary
FROM golang:1.23-alpine AS builder

# Set environment variables to ensure a Linux-compatible binary
ENV GOOS=linux
ENV GOARCH=amd64  
#Change to `arm64` if targeting ARM architecture
#Change to `amd64` if targeting AMD architecture
ENV CGO_ENABLED=0

# Set the working directory
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the application source code
COPY . .

# Build the application binary
RUN go build -o main cmd/app/main.go

################################################################################################

# Final Stage: Minimal image using `scratch`
FROM scratch

WORKDIR /app

# Copy certificates and the binary from the builder stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/main .

COPY ./config/config.yaml /app/config/config.yaml

ENV CONFIG_FILE_PATH=/app/config/config.yaml

# Expose the application port (if applicable)
EXPOSE 8080

# Run the application
CMD ["./main"]
