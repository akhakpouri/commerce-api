# Commerce API

A well-structured REST API for e-commerce functionality built with Go, using GORM and PostgreSQL.

## Project Structure

```
commerce-api/
├── main.go                   # Application entry point
├── internal/
│   ├── database/             # Database connection and migrations
│   │   ├── main.go          # Database initialization
│   │   └── setup.go         # Migration setup
│   ├── handlers/             # HTTP request handlers (coming soon)
│   ├── services/             # Business logic layer (coming soon)
│   └── models/              # Data models
│       ├── address.go       # Address model
│       ├── base.go          # Base model with common fields
│       ├── category.go      # Category model
│       ├── product.go       # Product model
│       ├── product-category.go  # Product-Category relationship
│       ├── review.go        # Review model
│       └── user.go          # User model
├── pkg/                      # Reusable packages
├── .golangci.yml            # Linter configuration
├── go.mod                    # Go module definition
├── go.sum                    # Go module checksums
└── README.md                 # This file
```

## Tech Stack

- **Go**: 1.25.0
- **Database**: PostgreSQL
- **ORM**: GORM v1.31.1
- **Database Driver**: PostgreSQL driver for GORM

## Prerequisites

- Go 1.25 or later
- PostgreSQL 13 or later
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

## Database Setup

### 1. Install PostgreSQL

```bash
# macOS
brew install postgresql@15
brew services start postgresql@15

# Verify installation
psql --version
```

### 2. Create Database

```bash
# Connect to PostgreSQL
psql postgres

# Create database and user
CREATE DATABASE commerce;
CREATE USER commerce WITH ENCRYPTED PASSWORD 'commerce@123';
GRANT ALL PRIVILEGES ON DATABASE commerce TO commerce;

# Create schema
\c commerce
CREATE SCHEMA commerce;
GRANT ALL ON SCHEMA commerce TO commerce;

# Exit PostgreSQL
\q
```

### 3. Update Database Configuration

The database connection string is currently in `internal/database/main.go`:

```go
connection := "host=localhost user=commerce dbname=commerce port=5432 password=commerce@123 sslmode=disable search_path=commerce"
```

**Note**: For production, use environment variables for database credentials instead of hardcoding them.

## Building and Running

```bash
# Build the application
go build -o commerce-api

# Run the application (will auto-migrate database)
./commerce-api

# Or run directly without building
go run main.go
```

The application will:
1. Connect to the PostgreSQL database
2. Run automatic migrations for all models
3. Start the server

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

## Data Models

The API includes the following data models:

- **User**: Customer/user account information
- **Address**: Shipping and billing addresses
- **Product**: Product catalog
- **Category**: Product categories
- **ProductCategory**: Many-to-many relationship between products and categories
- **Review**: Product reviews and ratings

All models are automatically migrated to the database on application startup.

## API Endpoints

*API endpoints are currently under development. The project includes:*

- ✅ Database connection and migrations
- ✅ Data models (User, Product, Category, Review, Address)
- 🔄 HTTP handlers (in progress)
- 🔄 Business logic services (in progress)
- 🔄 RESTful API endpoints (planned)

## Environment Variables

For production deployments, use environment variables for configuration:

```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=commerce
export DB_PASSWORD=your_secure_password
export DB_NAME=commerce
export DB_SCHEMA=commerce
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes and ensure all tests pass
4. Run linters and fix any issues
5. Submit a pull request
