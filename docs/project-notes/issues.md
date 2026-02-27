# Work Log

## Issue — Repository + Service layer implementation (ADR-008, ADR-009)

**Date:** 2026-02-27
**Status:** In progress
**Branch:** `main`

Implementing the repository layer (ADR-009) and service layer (ADR-008). See both ADRs in `decisions.md` for full interface signatures and implementation notes.

**Repository layer** (`internal/shared/repositories/`) — owns GORM queries, returns models:
- [ ] `repositories/user/user_repository.go`
- [ ] `repositories/address/address_repository.go`
- [ ] `repositories/product/product_repository.go`
- [ ] `repositories/category/category_repository.go`
- [ ] `repositories/review/review_repository.go`
- [ ] `repositories/order/order_repository.go`
- [ ] `repositories/payment/payment_repository.go`

**Service layer** (`api/internal/services/`) — owns business logic, returns DTOs:
- [ ] `services/user/user_service.go`
- [ ] `services/address/address_service.go`
- [ ] `services/product/product_service.go`
- [ ] `services/category/category_service.go`
- [ ] `services/review/review_service.go`
- [ ] `services/order/order_service.go`
- [ ] `services/payment/payment_service.go`

---

## Issue #38 — Payment model implementation

**Date:** 2026-02-26
**Status:** Done
**Branch:** `feature/issue-9`
**GitHub Issue:** #9

Designing and implementing the `Payment` entity as per ADR-007. Model lives in `internal/shared/models/payment.go` and must be registered in `internal/shared/database/setup.go`.

**Scope:**
- [x] `Payment` model with all fields from ADR-007
- [x] Register model for GORM AutoMigrate
- [x] Update `Order` model if needed (e.g., `Payments []Payment` association)

See ADR-007 in `decisions.md` for full field list and rationale.

---

## Issue #37 — ADR-003 embed fix

**Status:** Done
**Branch:** `feature/issue-37`

Resolved three bugs related to the `//go:embed` config setup (see BUG-002, BUG-003):

1. `config_manager.go` had `var content embed.FS` with no `//go:embed` directive — FS was always empty.
2. Embed responsibility was refactored: `NewDbConfig` now accepts `[]byte`; file reading and embedding moved to `utils/main.go`.
3. In `main.go`, the `//go:embed` directive was attached to `var _ embed.FS` (blank identifier) instead of `var content embed.FS` — fixed by moving the directive to the correct variable.
4. Fixed fallback logic: env var path now returns `nil` error so the caller can proceed.
5. Restored `utils/configs/config.json` as the canonical config location; updated `.gitignore` to match.

---

## Issue #34 — (merged)

**Branch:** `feature/issue-33`
**Merged commit:** `82a534f`
**Status:** Done

---

## Issue #33 — (merged)

**Status:** Done
**Notes:** Readme update included (`5c69c89`), config file removed (`109803b`).
