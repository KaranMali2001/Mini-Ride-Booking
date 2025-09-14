# Mini Ride Service

A microservices-based ride-sharing platform built with Go, featuring booking and driver services with PostgreSQL and RabbitMQ.

## Architecture

This project consists of multiple microservices:

- **Booking Service** (`booking_svc/`) - Handles ride booking requests
- **Driver Service** (`driver_svc/`) - Manages driver operations and availability
- **Migration Service** (`migrate/`) - Database schema migrations
- **Common** (`common/`) - Shared utilities and models

## Tech Stack

- **Language**: Go 1.24
- **Database**: PostgreSQL
- **Message Queue**: RabbitMQ
- **Containerization**: Docker & Docker Compose

## Prerequisites

- Docker and Docker Compose
- Go 1.24+ (for local development)

## Quick Start

1. Clone the repository:

   ```bash
   git clone github.com/KaranMali2001/Mini-Ride-Booking.git
   cd Mini-Ride-Booking
   ```

2. Start the services:

   ```bash
   cd deploy
   docker-compose up -d
   ```

3. The services will be available at:
   - Booking Service: http://localhost:8080
   - Driver Service: http://localhost:8081
   - PostgreSQL: localhost:5432
   - RabbitMQ Management: http://localhost:15672 (guest/guest)

## Development

This project uses Go workspaces for multi-module development:

```bash
# Install dependencies
go mod download

# Run a specific service locally
cd booking_svc
go run ./cmd

cd driver_svc
go run ./cmd
```

## Services

### Booking Service (Port 8080)

Handles ride booking operations including:

- Creating booking requests
- Managing booking status
- Customer notifications

### Driver Service (Port 8081)

Manages driver operations including:

- Driver registration and authentication
- Location tracking
- Ride assignment

### Database Migrations

Automatic database schema setup and migrations are handled by the migrate service.

## Configuration

Environment variables are managed through `.env` files:

- `deploy/.env` - Docker container configuration
- `booking_svc/.env` - Booking service local config
- `driver_svc/.env` - Driver service local config
- `migrate/.env` - Migration service config

## Docker Setup

The project uses a multi-service Docker setup with:

- **PostgreSQL** container (`ride_sharing_pg`) for data persistence
- **RabbitMQ** container for message queuing
- Individual service containers built from source

All services are orchestrated through `deploy/docker-compose.yml` with proper dependency management and health checks.
