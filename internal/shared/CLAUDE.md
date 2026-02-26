# CLAUDE.md

This file provides guidance to Claude Code with respect to the `internal/shared` directory.

## Overview & Purpose
Shared library used by both `api` and `utils` modules. Contains all GORM models and the database connection/migration logic. All external dependencies (GORM, PostgreSQL driver, bcrypt) are pinned here.

## Packages

### `database`
- `Migrate(cfg DbConfig)` — opens a GORM+PostgreSQL connection and runs `AutoMigrate` on all registered models
- `setup.go` — the only place to register new models for migration
- `DbConfig` — connection struct; all fields including `Schema` are dynamic (no hardcoded values)

### `models`
Eight domain models, all embedding `Base`:

- `Base` — `Id uint` (auto-increment PK), `CreatedDate`, `UpdatedDate`, `DeletedDate` (all `time.Time`)
- **Important:** `DeletedDate` is `time.Time`, NOT `gorm.DeletedAt` — GORM does not auto-filter soft-deleted records
- Every model implements `TableName() string` to explicitly set the table name
- Full relationship diagram: see `docs/project-notes/facts.md`

**Notable model behaviour:**
- `User` — `BeforeCreate`/`BeforeUpdate` hooks auto-bcrypt the `Password` field; `CheckPassword(string) bool` for verification
- `Category` — self-referential via `ParentId *uint` (nullable); supports unlimited-depth tree
- `Order` — uses string enum types `OrderStatus` and `PaymentStatus` defined in `order.go`

## Adding Dependencies
```bash
cd internal/shared
go get gorm.io/gorm gorm.io/driver/postgres
```