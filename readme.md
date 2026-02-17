# Commerce API

Go workspace for an e-commerce backend, organized into multiple modules with a shared data layer (GORM + PostgreSQL).

## Current Status

- ✅ Go workspace (`go.work`) with 3 modules
- ✅ Shared database package with auto-migrations
- ✅ Data models: User, Address, Product, Category, ProductCategory, Review, Order, OrderItem
- ✅ `utils` loads DB config from JSON file (`utils/configs/config.json`)
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
│   ├── configs/
│   │   ├── config.json
│   │   └── config copy.example
│   ├── internal/
│   │   └── managers/
│   │       └── config_manager.go
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

Current connection template (in `internal/shared/database/main.go`):

```go
connection := fmt.Sprintf(
	"host=%s user=%s dbname=%s port=%d password=%s sslmode=%s search_path=commerce",
	cfg.Host, cfg.User, cfg.DbName, cfg.Port, cfg.Password, cfg.SSLMode,
)
```

## Utils Configuration (JSON)

`utils` reads DB config from `configs/config.json` (relative to the `utils` module directory).

Expected JSON shape:

```json
{
	"host": "localhost",
	"port": 5432,
	"user": "commerce",
	"password": "commerce@123",
	"dbname": "commerce",
	"sslmode": "disable",
	"schema": "commerce"
}
```

If file loading fails, `utils` attempts to read environment variables (`DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `DB_SSLMODE`, `DB_SCHEMA`), but currently returns the file error and exits.

## Running

Run each executable from its own module directory:

```bash
# Run API executable
(cd api && go run .)

# Run utils executable
(cd utils && go run .)
```

Current behavior:

- `api`: prints `hello, world!`
- `utils`: prints `hello, world!`, loads DB config, then runs migrations

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
- `utils` requires `utils/configs/config.json` when run locally.
- `utils/configs/config copy.example` can be used as a template file.