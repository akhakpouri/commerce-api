# Project Facts & Configuration

## Database

| Key      | Value           |
|----------|-----------------|
| Engine   | PostgreSQL 13+  |
| Database | `commerce`      |
| Schema   | `commerce`      |
| User     | `commerce`      |
| Password | `commerce@123`  |
| Host     | `localhost`     |
| Port     | `5432`          |
| SSL Mode | `disable`       |

Config file: `utils/configs/config.json` (gitignored; `config.example` committed as reference)
Embedded via `//go:embed configs/config.json` in `utils/main.go` and passed as `[]byte` to `managers.NewDbConfig`.

### Setup SQL
```sql
CREATE DATABASE commerce;
CREATE USER commerce WITH ENCRYPTED PASSWORD 'commerce@123';
GRANT ALL PRIVILEGES ON DATABASE commerce TO commerce;
\c commerce
CREATE SCHEMA commerce AUTHORIZATION commerce;
```

---

## Domain Models (all in `internal/shared/models`)

`User`, `Address`, `Product`, `Category`, `ProductCategory` (junction), `Review`, `Order`, `OrderItem`

All embed `Base`: `uint` PK (`Id`), `CreatedDate`, `UpdatedDate`, `DeletedDate` (time.Time, not gorm.DeletedAt).

---

## Environment Variables (DB config fallback)

`DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `DB_SSLMODE`, `DB_SCHEMA`

---

## Module Paths

| Module            | Go module name            |
|-------------------|---------------------------|
| `api`             | `commerce/api`            |
| `utils`           | `commerce/utils`          |
| `internal/shared` | `commerce/internal/shared`|

---

## Linter

Tool: `golangci-lint` — config at `.golangci.yml` (workspace root)
Enabled rules: `errcheck`, `ineffassign`, `unused`, `govet`, `staticcheck`
