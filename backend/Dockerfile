# ---------- Build stage ----------
FROM golang:1.24.5-alpine AS build

WORKDIR /src

# Install ca-certificates (needed for HTTPS calls)
RUN apk add --no-cache ca-certificates tzdata

# Copy go.mod / go.sum first for efficient caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build the binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o corpchat-server ./cmd/server

# ---------- Runtime stage ----------
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata postgresql-client

WORKDIR /app

# Copy the statically-linked binary
COPY --from=build /src/corpchat-server .

# Expose port
EXPOSE 8080

# Run the server
CMD ["./corpchat-server"]