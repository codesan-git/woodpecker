# ─── Stage 1: Build ──────────────────────────────────
FROM golang:1.22-alpine AS builder

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

# Cache dependency download dulu
COPY go.mod go.sum* ./
RUN go mod download 2>/dev/null || true

# Copy source code & build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /app/server .

# ─── Stage 2: Runtime minimal ────────────────────────
FROM scratch

# Timezone & SSL certs
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Binary
COPY --from=builder /app/server /server

# Static files (frontend)
COPY static/ /static/

EXPOSE 8080

ENV PORT=8080
ENV APP_VERSION=1.0.0

ENTRYPOINT ["/server"]
