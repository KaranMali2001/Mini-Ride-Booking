# Stage 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.work and go.work.sum from repo root
COPY go.work go.work.sum ./

# Copy all modules
COPY common/ ./common/
COPY booking_svc/ ./booking_svc/
COPY driver_svc/ ./driver_svc/
COPY migrate/ ./migrate/

RUN go mod download

# Run tests for booking service
RUN go test ./booking_svc/internal/tests/... -v
RUN echo "Tests passed"

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./booking_svc/cmd/booking_svc/*.go

# Stage 2: Run
FROM alpine

WORKDIR /app

COPY --from=builder /app/main .
COPY booking_svc/.env .
EXPOSE 8080

CMD ["./main"]
