.PHONY: run build docs clean test

# Run the application
run:
	go run main.go

# Build the application
build:
	go build -o bin/swaggo-poc main.go

# Generate Swagger documentation
docs:
	swag init

# Clean build artifacts
clean:
	rm -rf bin/
	rm -rf docs/

# Install dependencies
deps:
	go mod tidy
	go mod download

# Install swag CLI if not present
install-swag:
	go install github.com/swaggo/swag/cmd/swag@latest

# Run with live reload (requires air)
dev:
	air

# Format code
fmt:
	go fmt ./...

# Vet code
vet:
	go vet ./...

# Run tests
test:
	go test ./...

# Full setup for new developers
setup: install-swag deps docs
	@echo "Setup complete! Run 'make run' to start the server."
