# ---- Builder ----
FROM golang:1.24 AS builder

WORKDIR /app

# Install dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o themis cmd/themis/main.go

# ---- Runtime ----
FROM gcr.io/distroless/base

WORKDIR /app

COPY --from=builder /app/themis /app/themis

EXPOSE 8080

ENTRYPOINT ["/app/themis"]
