# Commerce API

Go workspace for an e-commerce backend, organized into multiple modules with a shared data layer (GORM + PostgreSQL).

## Current Status

- ✅ Go workspace (`go.work`) with 3 modules
- ✅ Shared database package with auto-migrations
- ✅ Data models: User, Address, Product, Category, ProductCategory, Review, Order, OrderItem
- 🔄 API handlers and services scaffolding are present but not implemented yet

## Workspace Structure

```text
commerce-api/
├── go.work
├── go.work.sum
├── .golangci.yml
├── readme.md
├── api/                       # API executable module
│   ├── go.mod
│   ├── main.go
│   └── internal/
│       ├── handlers/          # (currently empty)
│       └── services/          # (currently empty)
├── utils/                     # Utility executable module
│   ├── go.mod
│   └── main.go
├── internal/
│   └── shared/                # Shared module used by executables
│       ├── go.mod
│       ├── database/
│       │   ├── main.go        # DB connection + migration trigger
│       │   └── setup.go       # AutoMigrate model registration
│       └── models/
│           ├── address.go
│           ├── base.go
│           ├── category.go
│           ├── order.go
│           ├── order_item.go
│           ├── product.go
│           ├── product-category.go
│           ├── review.go
│           └── user.go
└── pkg/
```

## Go / Dependencies

- Go: `1.25.7`
- ORM: `gorm.io/gorm v1.31.1`
- DB Driver: `gorm.io/driver/postgres v1.6.0`

## Prerequisites

- Go 1.25+
- PostgreSQL 13+
- `golangci-lint` (optional but recommended)

## Database Setup

```sql
CREATE DATABASE commerce;
CREATE USER commerce WITH ENCRYPTED PASSWORD 'commerce@123';
GRANT ALL PRIVILEGES ON DATABASE commerce TO commerce;

\c commerce
CREATE SCHEMA commerce;
GRANT ALL ON SCHEMA commerce TO commerce;
```

Current connection string (in `internal/shared/database/main.go`):

```go
connection := "host=localhost user=commerce dbname=commerce port=5432 password=commerce@123 sslmode=disable search_path=commerce"
```

## Running

From repository root:

```bash
# Run API executable
go run ./api

# Run utils executable
go run ./utils
```

Both executables currently print a message and run shared database migrations.

## Build

From each module:

```bash
(cd api && go build -o ../bin/api .)
(cd utils && go build -o ../bin/utils .)
```

## Linting / Vet / Tests

Run per module:

```bash
(cd api && go test ./...)
(cd utils && go test ./...)
(cd internal/shared && go test ./...)

(cd api && go vet ./...)
(cd utils && go vet ./...)
(cd internal/shared && go vet ./...)

(cd api && golangci-lint run ./...)
(cd utils && golangci-lint run ./...)
(cd internal/shared && golangci-lint run ./...)
```

## Module Maintenance

Because this repository uses a Go workspace (no root `go.mod`), run tidy inside each module:

```bash
(cd api && go mod tidy)
(cd utils && go mod tidy)
(cd internal/shared && go mod tidy)
go work sync
```

## Notes

- API routes/endpoints are not implemented yet.
- Database credentials are currently hardcoded for local development.
- For production, move DB configuration to environment variables.