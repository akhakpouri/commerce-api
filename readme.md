# Commerce API

Go workspace for an e-commerce backend, organized into multiple modules with a shared data layer (GORM + PostgreSQL).

## Current Status

- ✅ Go workspace (`go.work`) with 3 modules
- ✅ Shared database package with auto-migrations
- ✅ Data models: User, Address, Product, Category, ProductCategory, Review, Order, OrderItem
- ✅ `utils` embeds DB config from `utils/configs/config.json` at compile time, with env var fallback
- ✅ `utils/install.sh` — builds and installs the migration binary with custom config to `$GOPATH/bin/commerce-tools/`
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
│   ├── install.sh             # builds & installs binary with custom config
│   ├── configs/
│   │   ├── config.json        # gitignored — local credentials
│   │   └── config.example     # committed reference template
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
```

## Go / Dependencies

- Go: `1.26`
- ORM: `gorm.io/gorm v1.31.1`
- DB Driver: `gorm.io/driver/postgres v1.6.0`

## Prerequisites

- Go 1.26+
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

## Utils Configuration

`utils` embeds `configs/config.json` into the binary at compile time via `//go:embed`. If the file is missing or fails to parse, it falls back to environment variables and continues without error.

Copy the example to get started locally:

```bash
cp utils/configs/config.example utils/configs/config.json
```

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

### Environment Variable Fallback

| Variable      | Purpose           |
|---------------|-------------------|
| `DB_HOST`     | Database host     |
| `DB_PORT`     | Database port     |
| `DB_USER`     | Database user     |
| `DB_PASSWORD` | Database password |
| `DB_NAME`     | Database name     |
| `DB_SSLMODE`  | SSL mode          |
| `DB_SCHEMA`   | Schema name       |

## Installing `commerce-migrate`

`utils/install.sh` builds the migration binary with your local `config.json` baked in and installs it to `$GOPATH/bin/commerce-tools/`.

**Prerequisites:** `$GOPATH` must be set.

**Steps:**

1. Copy and edit the config template with your target database credentials:

```bash
cp utils/configs/config.example utils/configs/config.json
vim utils/configs/config.json
```

2. Run the install script from the `utils/` directory:

```bash
(cd utils && bash install.sh)
```

This will:
- Create `$GOPATH/bin/commerce-tools/` if it doesn't exist
- Copy `configs/` alongside the binary (for reference)
- Build the binary with `config.json` embedded at compile time
- Install it to `$GOPATH/bin/commerce-tools/utils`
- Execute the binary immediately to run migrations

**To run migrations again after install:**

```bash
$GOPATH/bin/commerce-tools/utils
```

> The database config is embedded at compile time. To target a different database, edit `config.json` and re-run `install.sh`.

---

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
- `utils`: loads DB config, then runs GORM auto-migrations

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

Run tidy inside each module:

```bash
(cd api && go mod tidy)
(cd utils && go mod tidy)
(cd internal/shared && go mod tidy)
go work sync
```

## Notes

- API routes/endpoints are not implemented yet (HTTP framework not yet chosen).
- `utils/configs/config.json` is gitignored. Use `config.example` as a template.
- `DeletedDate` on all models uses `time.Time`, not `gorm.DeletedAt` — soft-deleted records are not auto-filtered by GORM.