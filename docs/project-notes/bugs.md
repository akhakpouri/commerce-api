# Bug Log

## BUG-001 — `getPortFromEnv()` returns scan count instead of port number

**File:** `utils/internal/managers/config_manager.go`
**Discovered:** 2026-02-25
**Status:** Fixed

### Description
`fmt.Sscanf` returns the number of successfully scanned items, not the scanned value. The code was scanning into the `port` string variable (instead of an `int`) and returning `p` (always `1` on success), so `DB_PORT` env var was effectively ignored — port always resolved to `1`.

### Buggy Code
```go
p, err := fmt.Sscanf(port, "%d", &port) // scans into string, p = count
return p                                  // returns 1, not the port
```

### Fix
Replaced with `strconv.Atoi` which directly parses a string to int:
```go
p, err := strconv.Atoi(port)
return p
```

---

## BUG-002 — Missing `//go:embed` directive; `embed.FS` always empty

**File:** `utils/internal/managers/config_manager.go` (later refactored to `utils/main.go`)
**Discovered:** 2026-02-25
**Status:** Fixed

### Description
`var content embed.FS` was declared with no `//go:embed` directive. Go requires the directive on the line immediately preceding the variable. Without it the FS is empty and `content.ReadFile(...)` always returns "file does not exist".

### Fix
Added `//go:embed configs/config.json` directly above `var content embed.FS` in `utils/main.go` (embed responsibility moved there during refactor).

---

## BUG-003 — `//go:embed` directive bound to blank identifier `_`

**File:** `utils/main.go`
**Discovered:** 2026-02-25
**Status:** Fixed

### Description
```go
var content embed.FS        // no directive — always empty

//go:embed configs/config.json
var _ embed.FS              // directive discarded; _ is never read
```
The directive was attached to the wrong variable. `content` remained empty; `ReadFile` returned `*errors.errorString {s: "file does not exist"}`.

### Fix
```go
//go:embed configs/config.json
var content embed.FS
```

---

## BUG-004 — Wrong package name in `product/product.go`

**File:** `api/internal/dto/product/product.go`
**Discovered:** 2026-02-26
**Status:** Fixed

### Description
File declared `package dto` instead of `package product`. All other DTO sub-packages use their directory name as the package name. This causes an import conflict when callers import the package as `product`.

### Fix
Changed line 1 from `package dto` to `package product`.

---

## BUG-005 — Nil pointer dereference on `*time.Time` in `payment.FromModel`

**File:** `api/internal/dto/payment/payment.go`
**Discovered:** 2026-02-26
**Status:** Fixed

### Description
`payment.PaidAt.Format(...)` was called directly on a `*time.Time` field without a nil check. `PaidAt` is nullable — calling `.Format()` on a nil pointer panics at runtime.

### Fix
Wrapped in a nil check:
```go
PaidAt: func() string {
    if payment.PaidAt != nil {
        return payment.PaidAt.Format("01/02/2006 15:04:05")
    }
    return ""
}(),
```

---

## BUG-006 — Invalid Go time token `pm` in format string

**File:** `api/internal/dto/payment/payment.go`
**Discovered:** 2026-02-26
**Status:** Fixed

### Description
Format string `"01/02/2006 15:04pm"` was used. `pm` is not a valid Go time token — it is output as the literal string "pm". When combined with `15` (24-hour clock), this produces incorrect output like `"02/26/2026 14:04pm"`.

### Fix
Changed to `"01/02/2006 15:04:05"` (standard 24-hour format).

---

## BUG-007 — Format/parse layout mismatch causing silent nil on `PaidAt`

**File:** `api/internal/dto/payment/payment.go`
**Discovered:** 2026-02-26
**Status:** Fixed

### Description
`FromModel` formatted `PaidAt` with `"01/02/2006 15:04:05"` but `getTimeString` (used in `ToModel`) still used the old layout `"01/02/2006 15:04pm"`. `time.Parse` silently returns an error on mismatch, causing `getTimeString` to always return `nil` — `PaidAt` was never round-tripped correctly.

### Fix
Aligned both layouts to `"01/02/2006 15:04:05"`.

---

## BUG-008 — Duplicate `PaymentStatus` type in `order.go` and `payment.go`

**File:** `internal/shared/models/order.go`, `internal/shared/models/payment.go`
**Discovered:** 2026-02-26
**Status:** Fixed

### Description
`PaymentStatus` type and its constants were defined in both files. Go does not allow duplicate type definitions in the same package — compile error.

### Fix
Removed `PaymentStatus` from `order.go`. It now lives exclusively in `payment.go`. `Order` model references it from there. Also removed `Order.PaymentStatus` field; replaced with `Payments []Payment` association.

---

## BUG-009 — `GatewayTransactionId` marked `not null; unique` on `Payment`

**File:** `internal/shared/models/payment.go`
**Discovered:** 2026-02-26
**Status:** Fixed

### Description
`GatewayTransactionId` was tagged `gorm:"not null;unique"`. Failed and pending payments may not have a gateway transaction ID yet — `not null` would prevent inserting these rows.

### Fix
Removed `not null` constraint. Field is now nullable and unique only when populated.
