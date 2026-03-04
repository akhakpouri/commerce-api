# CLAUDE.md

This file provides guidance to Claude Code with respect to the `api/internal` directory.

## Overview & Purpose
Application layer for the HTTP API. Contains two packages: `dto` (request/response shapes) and `services` (business logic). No HTTP framework has been chosen yet — see ADR-004.

**Hard rule:** Services must never import or reference `gorm.io/gorm` directly. All DB access goes through repository interfaces injected at construction time.

---

## Packages

### `dto`
Thin data containers for API payloads. One sub-package per domain under `api/internal/dto/<name>/`.

**Responsibilities:**
- JSON shape (`json` tags)
- Input validation (`validate` tags)
- Mapping to/from models (`ToModel()` / `FromModel()`)

**No business logic in DTOs.** See ADR-008.

**Mapping convention:**
- `ToModel()` — method on request DTOs, returns a model
- `FromModel(...)` — standalone function on response DTOs, accepts a model

**Structure:**
```
dto/
├── user/user.go
├── address/address.go
├── product/product.go
├── category/category.go
├── review/review.go
├── order/order.go
├── order-item/order_item.go
└── payment/payment.go
```

---

### `services`
Business logic layer. One sub-package per domain under `api/internal/services/<name>/`.

**Responsibilities:**
- Enforce business rules (e.g. order must exist before payment, status validation)
- Orchestrate repository calls
- Map between models and DTOs
- Return DTOs to callers — never raw models

**Pattern:** Each file defines an interface (`XxxServiceI`) and a concrete struct (`XxxService`) that implements it.

**Constructor:** takes one or more repository interfaces, returns the service interface:
```go
func NewXxxService(repo xrepo.XxxRepositoryI) XxxServiceI {
    return &XxxService{repo: repo}
}
```

**Import aliasing** — service, repo, and DTO packages share the same domain name. Alias at the import site:
```go
import (
    userdto  "commerce/api/internal/dto/user"
    userrepo "commerce/internal/shared/repositories/user"
)
```

**Structure:**
```
services/
├── user/user_service.go
├── address/address_service.go
├── product/product_service.go
├── category/category_service.go
├── review/review_service.go
├── order/order_service.go
└── payment/payment_service.go
```

**Notable implementation rules (from ADR-008):**
- `UserService.Authenticate` — fetch by email, call `model.CheckPassword(password)`, return error if false
- `AddressService.SetDefault` — clear existing default for user, then set new one
- `OrderService.Create` — must create `Order` + all `OrderItems` atomically in a single transaction
- `OrderService.UpdateStatus` / `PaymentService.UpdateStatus` — validate input string against model enum constants before calling repo
- `OrderService` injects both `OrderRepositoryI` and `OrderItemRepositoryI`

---

## Key ADRs
| ADR | Title |
|-----|-------|
| ADR-004 | HTTP framework not yet chosen |
| ADR-008 | Thin DTOs with service-layer mapping and business logic |
| ADR-009 | Repository pattern for data access |

Full details in `docs/project-notes/decisions.md`.
