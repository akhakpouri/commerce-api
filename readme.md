# Commerce API

A well-structured REST API for e-commerce functionality built with Go.

## Project Structure

```
commerce-api/
├── cmd/
│   └── server/
│       └── main.go           # Application entry point
├── internal/
│   ├── handlers/             # HTTP request handlers
│   ├── services/             # Business logic layer
│   ├── models/               # Data structures and models
│   └── middleware/           # HTTP middleware
├── pkg/                      # Reusable packages
├── go.mod                    # Go module definition
├── go.sum                    # Go module checksums
└── README.md                 # This file
```

## Prerequisites

- Go 1.21 or later
- golangci-lint (for linting)

## Installation

```bash
# Clone the repository
git clone <repository-url>
cd commerce-api

# Install dependencies
go mod download

# Verify dependencies
go mod verify
```

## Building

```bash
# Build the application
go build -o commerce-api ./cmd/server

# Run the application
./commerce-api

# Or run directly
go run ./cmd/server
```

## Testing

```bash
# Run all tests
go test ./...

# Run tests with race detection
go test -race ./...

# View test coverage
go test -cover ./...

# Generate detailed coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Code Quality & Linting

### Setup golangci-lint

```bash
# Install golangci-lint (macOS)
brew install golangci-lint

# Verify installation
golangci-lint --version
```

### Lint the Project

```bash
# Run all linters
golangci-lint run ./...

# Run with verbose output
golangci-lint run -v ./...

# Fix issues automatically
golangci-lint run --fix ./...

# Check only new issues
golangci-lint run --new
```

### Format Code

```bash
# Format all Go files
go fmt ./...

# Verify formatting (without changes)
go fmt -n ./...
```

### Static Analysis

```bash
# Run Go vet
go vet ./...

# Check for vulnerabilities
govulncheck ./...
```

## Pre-Commit Verification

Before committing code, run this comprehensive check:

```bash
go mod tidy && \
go fmt ./... && \
go vet ./... && \
go test -race ./... && \
golangci-lint run ./...
```

Or create a shell script:

```bash
#!/bin/bash
set -e

echo "Running go mod tidy..."
go mod tidy

echo "Running go fmt..."
go fmt ./...

echo "Running go vet..."
go vet ./...

echo "Running tests with race detection..."
go test -race ./...

echo "Running golangci-lint..."
golangci-lint run ./...

echo "✅ All checks passed!"
```

## Dependency Management

```bash
# Check for available updates
go list -u -m all

# Upgrade a specific dependency
go get -u github.com/package/name

# Upgrade all direct dependencies
go get -u ./...

# Upgrade only patch versions (safer)
go get -u=patch ./...

# Clean up dependencies
go mod tidy
```

## Common Issues

### "return value of '...' is not checked"

This warning means you're calling a function that returns an error, but not checking it:

```go
// ❌ Wrong
database.AutoMigrate(&User{})

// ✅ Correct
if err := database.AutoMigrate(&User{}); err != nil {
    log.Fatalf("migration failed: %v", err)
}
```

### Import cycles

Check for circular dependencies:

```bash
go mod graph | grep -E "A -> B.*B -> A"
```

## Development Workflow

1. Make changes to code
2. Run `go fmt ./...` to format
3. Run `go test -race ./...` to test
4. Run `golangci-lint run ./...` to check code quality
5. Run `go mod tidy` to clean up dependencies
6. Commit changes

## API Endpoints

*Documentation for API endpoints coming soon*

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes and ensure all tests pass
4. Run linters and fix any issues
5. Submit a pull request
