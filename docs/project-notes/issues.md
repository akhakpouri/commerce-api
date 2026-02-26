# Work Log

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
