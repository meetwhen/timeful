# Module-download-only stage. CI test images (compose.test.yaml) stop here:
# server-route-test and server-test mount the source and compile live, so the
# full server binary build below would be discarded work for them.
FROM golang:1.26.4-alpine3.24 AS testdeps

WORKDIR /app

# Cache module downloads separately from application source changes.
COPY go.mod go.sum ./
RUN go mod download

# Build stage
FROM testdeps AS builder

WORKDIR /app

# Copy the remaining build inputs and compile the server binary.
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -o /out/server .

# Goose runs schema changes in its own one-shot Compose service.
FROM golang:1.26.4-alpine3.24 AS goose-builder
RUN CGO_ENABLED=0 go install github.com/pressly/goose/v3/cmd/goose@v3.26.0

FROM alpine:3.24 AS migrator
RUN apk add --no-cache postgresql18-client
COPY --from=goose-builder /go/bin/goose /usr/local/bin/goose
COPY migrations /migrations
CMD ["goose"]

# Runtime stage
FROM alpine:3.24

WORKDIR /app

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Create non-root user and writable runtime directories
RUN adduser -D -g '' appuser \
	&& mkdir -p /app/logs \
	&& chown -R appuser:appuser /app

# Copy binary from builder
COPY --from=builder /out/server ./server

USER appuser

EXPOSE 3002

CMD ["./server"]
