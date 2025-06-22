.PHONY: build test clean install run serve help

# Default target
all: build

# Build the CLI application
build:
	@echo "Building str2go CLI..."
	go build -o bin/str2go cmd/str2go/main.go

# Install the CLI application
install:
	@echo "Installing str2go CLI..."
	go install ./cmd/str2go

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	rm -f coverage.out coverage.html

# Run the CLI with example
run: build
	@echo "Running str2go with example..."
	./bin/str2go convert "42" --type int

# Start the web server
serve: build
	@echo "Starting str2go web server..."
	./bin/str2go serve --port 8080

# Build and run tests
check: test build

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Lint code
lint:
	@echo "Linting code..."
	golangci-lint run

# Generate documentation
docs:
	@echo "Generating documentation..."
	godoc -http=:6060

# Create release build
release: clean
	@echo "Creating release builds..."
	GOOS=linux GOARCH=amd64 go build -o bin/str2go-linux-amd64 cmd/str2go/main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/str2go-darwin-amd64 cmd/str2go/main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/str2go-darwin-arm64 cmd/str2go/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/str2go-windows-amd64.exe cmd/str2go/main.go

# Show help
help:
	@echo "Available targets:"
	@echo "  build        - Build the CLI application"
	@echo "  install      - Install the CLI application"
	@echo "  test         - Run tests"
	@echo "  test-coverage- Run tests with coverage report"
	@echo "  clean        - Clean build artifacts"
	@echo "  run          - Run the CLI with example"
	@echo "  serve        - Start the web server"
	@echo "  check        - Run tests and build"
	@echo "  fmt          - Format code"
	@echo "  lint         - Lint code"
	@echo "  docs         - Generate documentation"
	@echo "  release      - Create release builds for multiple platforms"
	@echo "  help         - Show this help message" 