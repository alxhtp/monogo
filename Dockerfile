# ----------- Builder Stage -----------
# Use official Go image for building the binary
FROM golang:1.25-alpine AS builder

# Install git (required for go mod) and air for dev
RUN apk add --no-cache git curl
RUN go install github.com/air-verse/air@latest

WORKDIR /app

# Cache go mod deps
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Install swag CLI tool
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Generate Swagger documentation
RUN swag init -g cmd/main.go --output docs --parseDependency


# Build a statically linked binary for production
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/monogo -ldflags="-s -w" ./cmd/main.go

# ----------- Development Stage -----------
FROM golang:1.25-alpine AS dev
WORKDIR /app

# Install air and git
RUN apk add --no-cache git curl
RUN go install github.com/air-verse/air@latest

COPY --from=builder /app /app

# Expose port for the app
EXPOSE 8080

# Use air for live reload in dev
CMD ["air", "-c", ".air.toml"]

# ----------- Production Stage -----------
FROM gcr.io/distroless/static-debian12 AS prod

# Create non-root user
USER nonroot:nonroot

WORKDIR /app

# Copy the statically built binary from builder
COPY --from=builder /app/bin/monogo /app/monogo

# Copy Swagger docs
COPY --from=builder /app/docs /app/docs

# Expose port for the app
EXPOSE 8080

# Run the binary
ENTRYPOINT ["/app/monogo"]
