# Architectural Decision Records

## ADR-001 — Go workspace with three modules

**Date:** (pre-existing)
**Status:** Closed — verified 2026-02-25

Split into `api`, `utils`, and `internal/shared` modules under a single `go.work` workspace. All external dependencies (GORM, PostgreSQL driver, bcrypt) are pinned to `internal/shared` only. `api` and `utils` consume `internal/shared` as a local dependency.

---

## ADR-002 — GORM + PostgreSQL for persistence

**Date:** (pre-existing)
**Status:** Closed — verified 2026-02-25

GORM is the ORM. All models embed a `Base` struct (`internal/shared/models/base.go`) providing:
- `Id uint` — auto-increment primary key (not UUID)
- `CreatedDate time.Time` — auto-set on create
- `UpdatedDate time.Time` — auto-set on update
- `DeletedDate time.Time` — indexed, but typed as `time.Time` not `gorm.DeletedAt`

> **Note:** `DeletedDate` uses `time.Time`, not `gorm.DeletedAt`. GORM's automatic soft-delete filtering requires `gorm.DeletedAt`. Current implementation does NOT auto-filter deleted records unless queries are written manually.

All tables live in the `commerce` PostgreSQL schema.

---

## ADR-003 — Embedded config with env var fallback

**Date:** (pre-existing)
**Status:** Closed — verified 2026-02-25

`utils/main.go` embeds `configs/config.json` at compile time via `//go:embed` and passes the raw bytes to `managers.NewDbConfig([]byte)`. If JSON parsing fails, `NewDbConfig` falls back to environment variables (`DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `DB_SSLMODE`, `DB_SCHEMA`) and returns `nil` error.

Config file: `utils/configs/config.json` — gitignored (contains credentials). `utils/configs/config.example` is committed as a reference.

> **Note:** Embed responsibility lives in `utils/main.go`, not `config_manager.go`. `NewDbConfig` accepts `[]byte` and has no knowledge of the filesystem.

---

## ADR-004 — HTTP framework not yet chosen

**Date:** (pre-existing)
**Status:** Pending

`api/internal/handlers/` and `api/internal/services/` exist but are empty. No HTTP framework has been selected. This is the next major architectural decision to make.

---

## ADR-005 — bcrypt password hashing via GORM hooks

**Date:** (pre-existing)
**Status:** Active

`User` model uses `BeforeCreate` and `BeforeUpdate` GORM hooks to automatically hash the `Password` field with bcrypt. A `CheckPassword()` method is provided for verification. Hashing is transparent to callers.

---

## ADR-006 — Shell script installation with compile-time config embedding

**Date:** 2026-02-26
**Status:** Active

`utils/install.sh` is the chosen installation mechanism for the migration binary. It builds the binary with `configs/config.json` embedded at compile time and installs it to `$GOPATH/bin/commerce-tools/` alongside a copy of the `configs/` directory.

**Workflow:** edit `config.json` → run `install.sh` → binary is built with credentials baked in → migrations run immediately.

**Why this over alternatives:**
- `go install` alone can't embed a local config file at the user's `$GOBIN` without a build step
- Runtime `--config` flag was considered but rejected as unnecessary complexity for this use case
- Custom install dir (`commerce-tools/`) keeps the binary isolated from other Go tools in `$GOBIN`

**Tradeoff:** targeting a different database requires editing `config.json` and re-running `install.sh` (rebuild required). This is acceptable given the tool's purpose as a one-time migration runner, not a frequently reconfigured service.

**Fix (2026-02-26):** `cp configs` corrected to `cp -r configs` — directory copy was silently failing without the `-r` flag.

---

## ADR-007 — Payment model as a separate table with audit trail

**Date:** 2026-02-26
**Status:** Active — implemented and migrated 2026-02-26

Rather than extending `Order` with more payment fields, `Payment` is its own table with a many-to-one relationship to `Order`. This preserves the full history of payment attempts (retries, refunds) rather than overwriting a single status.

**Fields:** `OrderId` (FK), `Amount`, `Currency`, `Status`, `Gateway`, `GatewayTransactionId`, `GatewayResponse`, `PaymentMethod`, `PaidAt` (nullable).

**Status enum:** `pending | authorized | captured | failed | refunded | partially_refunded`

**Key decisions:**
- `Order.PaymentStatus` is kept as a denormalized convenience flag ("is this order paid?") — the `Payment` table is the source of truth for *how* and *when*.
- No separate `PaymentMethod` model for MVP — gateway tokens (e.g., Stripe `pm_...`) stored as a string field on `Payment`.
- Refunds handled via status + `RefundedAmount` on the existing `Payment` row (not separate rows) for MVP simplicity.
- Actual card data never stored — delegated entirely to the payment gateway (PCI compliance).

---
## ADR-008 — Thin DTOs with service-layer mapping and business logic

**Date:** 2026-02-26
**Status:** Active — pending implementation

API payloads are represented as DTOs (request/response structs) living in `api/internal/dto/`. DTOs are plain data containers — json tags, validation tags, and mapping methods only. Business logic lives exclusively in `api/internal/services/`.

**Layer responsibilities:**

| Concern | Layer |
|---|---|
| JSON shape / validation tags | DTO (`api/internal/dto/`) |
| Mapping DTO ↔ model | DTO methods (`ToModel()` / `FromModel()`) |
| Business rules (e.g. order must exist, not already paid) | Service |
| Password hashing, GORM hook behaviour | Model |
| DB persistence | Service (via GORM) |

**Mapping convention:** `ToModel()` as a method on request DTOs; standalone `FromModel()` functions for response DTOs.

**Why not business logic in DTOs:**
- GORM hooks on models (e.g. `User.BeforeCreate` bcrypt) already own some business logic — duplicating concerns in DTOs creates conflicts.
- DTOs live in `api/`; if logic lives there it can't be reused by other consumers (CLI, workers) without creating cross-module coupling.

**Why not logic in models:**
- Models are shared across all consumers via `internal/shared` — embedding API-specific rules there pollutes the shared library.
---