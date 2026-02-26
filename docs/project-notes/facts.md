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

### Base (embedded by all models)

| Field         | Type        | Notes                                      |
|---------------|-------------|--------------------------------------------|
| `Id`          | `uint`      | Primary key, auto-increment                |
| `CreatedDate` | `time.Time` | Auto-set on insert                         |
| `UpdatedDate` | `time.Time` | Auto-set on update                         |
| `DeletedDate` | `time.Time` | Soft-delete marker — **not** `gorm.DeletedAt`; GORM does NOT auto-filter deleted records |

---

### Entity Relationship Diagram

```
┌─────────────────────────────────────────────────────────────────┐
│                            User                                 │
│  Id · FirstName · LastName · Email · Password(bcrypt)           │
└────────┬───────────────────────┬──────────────────┬────────────┘
         │ 1:many                │ 1:many           │ 1:many
         ▼                       ▼                  ▼
  ┌─────────────┐         ┌────────────┐     ┌───────────┐
  │   Address   │◄────────│   Order    │     │  Review   │
  │─────────────│ shipping│────────────│     │───────────│
  │ UserId (FK) │ billing │ UserId(FK) │     │ UserId(FK)│
  │ Street      │         │ ShipAddr   │     │ ProductId │
  │ City        │         │ BillAddr   │     │ Rating    │
  │ State       │         │ OrderNum   │     │ Title     │
  │ PostalCode  │         │ Status     │     │ Comment   │
  │ Country     │         │ Payment    │     └─────┬─────┘
  │ IsDefault   │         │ SubTotal   │           │ many:1
  └─────────────┘         │ Tax·Total  │           ▼
                          └─────┬──────┘   ┌───────────────┐
                                │ 1:many   │    Product    │
                                ▼          │───────────────│
                         ┌────────────┐    │ Name · Price  │
                         │ OrderItem  │    │ Description   │
                         │────────────│    │ Sku (unique)  │
                         │ OrderId(FK)│    │ Stock         │
                         │ ProductId ─┼───►│ IsActive      │
                         │ Quantity   │    │ IsFeatured    │
                         │ UnitPrice  │    └───────┬───────┘
                         │ TaxAmount  │            │ 1:many (via junction)
                         └────────────┘            ▼
                                          ┌──────────────────┐
                                          │  ProductCategory │
                                          │──────────────────│
                                          │ ProductId (FK)   │
                                          │ CategoryId (FK)  │
                                          └────────┬─────────┘
                                                   │ many:1
                                                   ▼
                                          ┌──────────────────┐
                                          │    Category      │
                                          │──────────────────│
                                          │ Name · Slug      │
                                          │ Description      │
                                          │ ParentId (*uint) │◄─┐
                                          │ IsActive         │  │ self-ref
                                          │ Children []      │──┘ (tree)
                                          └──────────────────┘
```

---

### Relationship Summary

| From         | To              | Type          | FK field(s)                              |
|--------------|-----------------|---------------|------------------------------------------|
| `Address`    | `User`          | many:1        | `Address.UserId`                         |
| `Order`      | `User`          | many:1        | `Order.UserId`                           |
| `Order`      | `Address`       | many:1 (×2)   | `Order.ShippingAddressId`, `Order.BillingAddressId` |
| `Order`      | `OrderItem`     | 1:many        | `OrderItem.OrderId`                      |
| `OrderItem`  | `Product`       | many:1        | `OrderItem.ProductId`                    |
| `Review`     | `User`          | many:1        | `Review.UserId`                          |
| `Review`     | `Product`       | many:1        | `Review.ProductId`                       |
| `Product`    | `Category`      | many:many     | via `ProductCategory` junction           |
| `Category`   | `Category`      | self-ref tree | `Category.ParentId` (`*uint`, nullable)  |

---

### Model Notes

**User** (`users`)
- `BeforeCreate` hook: bcrypt-hashes `Password`; rejects empty password
- `BeforeUpdate` hook: re-hashes only if `Password` field changed
- `CheckPassword(string) bool` — bcrypt comparison helper
- `FullName() string` — concatenates `FirstName + LastName`

**Category** (`categories`)
- `ParentId *uint` is nullable — `nil` means root category
- Self-referential `Children []Category` enables an unlimited-depth tree

**Order** (`orders`)
- `Status OrderStatus` — enum: `pending`, `shipped`, `delivered`, `cancelled`
- `PaymentStatus PaymentStatus` — enum: `pending`, `completed`, `failed`, `refunded`
- References `Address` twice (shipping + billing) via explicit FK fields

**ProductCategory** (`product_categories`)
- Pure junction table; carries its own `Base` (Id + timestamps)

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
