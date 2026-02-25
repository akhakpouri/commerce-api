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
