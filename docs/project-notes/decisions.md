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

**Clarification (2026-02-26):** Moving this logic to the service layer (per ADR-008) was considered and rejected. Password hashing is a persistence invariant — it must hold regardless of which service or consumer writes a `User`. Keeping it in the model hook makes it unconditional and impossible to accidentally bypass.

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
- `Order.PaymentStatus` field was removed during implementation — replaced with `Payments []Payment` association. Payment state is read by querying the `payments` table directly.
- No separate `PaymentMethod` model for MVP — gateway tokens (e.g., Stripe `pm_...`) stored as a string field on `Payment`.
- Refunds handled via status + `RefundedAmount` on the existing `Payment` row (not separate rows) for MVP simplicity.
- Actual card data never stored — delegated entirely to the payment gateway (PCI compliance).

---
## ADR-008 — Thin DTOs with service-layer mapping and business logic

**Date:** 2026-02-26
**Status:** Active — DTOs done; service layer in progress 2026-02-27

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

### Service Layer Design (2026-02-27)

**Structure:** One sub-package per domain, mirroring the DTO layout.
```
api/internal/services/
├── user/user_service.go
├── address/address_service.go
├── product/product_service.go
├── category/category_service.go
├── review/review_service.go
├── order/order_service.go
└── payment/payment_service.go
```
`OrderItem` has no dedicated service — managed within `OrderService`.

**Pattern:** Each file defines an interface (`XxxServiceI`) and a concrete struct (`XxxService`) that implements it. Constructor takes a repository interface and returns the service interface: `func NewXxxService(repo XxxRepositoryI) XxxServiceI`. Services never hold `*gorm.DB` directly — see ADR-009.

**DTO import aliasing** — service package, repo package, and DTO package all share the same domain name (e.g. all `package user`). Alias at the import site: `userdto "commerce/api/internal/dto/user"`, `userrepo "commerce/internal/shared/repositories/user"`.

**Interface signatures:**

```go
// UserService
GetById(id uint) (*userdto.User, error)
GetByEmail(email string) (*userdto.User, error)
Create(dto *userdto.User) (*userdto.User, error)
Update(id uint, dto *userdto.User) (*userdto.User, error)
Delete(id uint) error
Authenticate(email, password string) (*userdto.User, error)

// AddressService
GetById(id uint) (*addressdto.Address, error)
GetByUserId(userId uint) ([]addressdto.Address, error)
Create(dto *addressdto.Address) (*addressdto.Address, error)
Update(id uint, dto *addressdto.Address) (*addressdto.Address, error)
Delete(id uint) error
SetDefault(userId uint, addressId uint) error

// ProductService
GetById(id uint) (*productdto.Product, error)
GetAll() ([]productdto.Product, error)
Create(dto *productdto.Product) (*productdto.Product, error)
Update(id uint, dto *productdto.Product) (*productdto.Product, error)
Delete(id uint) error

// CategoryService
GetById(id uint) (*categorydto.Category, error)
GetAll() ([]categorydto.Category, error)
Create(dto *categorydto.Category) (*categorydto.Category, error)
Update(id uint, dto *categorydto.Category) (*categorydto.Category, error)
Delete(id uint) error

// ReviewService
GetById(id uint) (*reviewdto.Review, error)
GetByProductId(productId uint) ([]reviewdto.Review, error)
Create(dto *reviewdto.Review) (*reviewdto.Review, error)
Update(id uint, dto *reviewdto.Review) (*reviewdto.Review, error)
Delete(id uint) error

// OrderService
GetById(id uint) (*orderdto.Order, error)
GetByUserId(userId uint) ([]orderdto.Order, error)
Create(dto *orderdto.Order) (*orderdto.Order, error)  // must create OrderItems in same transaction
UpdateStatus(id uint, status string) (*orderdto.Order, error)
Delete(id uint) error

// PaymentService
GetById(id uint) (*paymentdto.Payment, error)
GetByOrderId(orderId uint) ([]paymentdto.Payment, error)
Create(dto *paymentdto.Payment) (*paymentdto.Payment, error)
UpdateStatus(id uint, status string) (*paymentdto.Payment, error)
Delete(id uint) error
```

**Notable implementation notes (service layer):**
- `UserService.Authenticate` — fetch by email, call `model.CheckPassword(password)`, return error if false.
- `AddressService.SetDefault` — call `repo.ClearDefault(userId)` then `repo.SetDefault(addressId)`.
- `OrderService.Create` — open a `db.Transaction(...)` and pass it down to create `Order` + all `OrderItems` atomically.
- `OrderService.UpdateStatus` / `PaymentService.UpdateStatus` — validate input string against model enum constants before calling repo.

---

## ADR-009 — Repository pattern for data access

**Date:** 2026-02-27
**Status:** Active — pending implementation

A repository layer is introduced between services and GORM. Services never hold `*gorm.DB` directly; they depend on repository interfaces.

**Layering:**
```
Handler → Service → Repository → GORM → DB
           (why)      (how)
```

**Location:** `internal/shared/repositories/` — sits alongside models in the shared module. GORM is already a dependency there, and repos are model-specific with no API concerns.

**Structure:** One sub-package per domain, same pattern as models and DTOs.
```
internal/shared/repositories/
├── user/user_repository.go
├── address/address_repository.go
├── product/product_repository.go
├── category/category_repository.go
├── review/review_repository.go
├── order/order_repository.go
└── payment/payment_repository.go
```
`OrderItem` has no dedicated repo — managed within `order/`.

**Pattern:** Each file defines an interface (`XxxRepositoryI`) and a concrete struct (`XxxRepository`). Constructor takes `*gorm.DB` and returns the interface: `func NewXxxRepository(db *gorm.DB) XxxRepositoryI`.

**Method naming:** `Find...` for reads, `Create`, `Update`, `SoftDelete` for writes.

**Soft-delete** — repos own the soft-delete logic so services don't need to know about it:
- All `Find...` methods filter: `.Where("deleted_date = ?", time.Time{})`
- `SoftDelete` sets: `.Update("deleted_date", time.Now())`

**Why repositories in `internal/shared/` and not `api/internal/`:**
- GORM is already a dependency of `internal/shared` — no new dependency introduced.
- Repos are model-specific (no API concerns) — they belong near models, not near handlers.
- Future consumers (e.g. a worker module) can reuse repos without importing the `api` module.

**Why not embed queries directly in services:**
- Services become testable without a real DB — inject a mock repo instead.
- Query logic is centralized; soft-delete filtering isn't scattered across services.
- Swapping GORM for another persistence mechanism only touches the repo layer.
---