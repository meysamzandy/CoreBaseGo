# Define variables
GO = go
BINARY_NAME = goplan
BUILD_DIR = build
LINT = golint
TEST = $(GO) test
VET = $(GO) vet
FMT = $(GO) fmt
RUN = $(GO) run
MOD = $(GO) mod
VERSION = $(shell git describe --tags --always)

# Default target
all: build

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	$(GO) build -o $(BINARY_NAME) cmd/main.go

# Run the application
run:
	@echo "Running $(BINARY_NAME)..."
	$(RUN) cmd/api/main.go

generate-key:
	@echo "Running $(BINARY_NAME)..."
	$(RUN) scripts/keyGenerator.go

migrate:
	@echo "Running $(BINARY_NAME)..."
	$(RUN) scripts/initDatabase.go

# Test the application
test:
	@echo "Running tests..."
	$(TEST) ./...

# Format the code
fmt:
	@echo "Formatting code..."
	$(FMT) ./...

# Vet the code
vet:
	@echo "Vetting code..."
	$(VET) ./...

# Lint the code
lint:
	@echo "Linting code..."
	$(LINT) ./...

# Clean the build artifacts
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)

# Run Workers
worker:
	@echo "Running $(BINARY_NAME)..."
	$(RUN) cmd/worker/main.go

cli-run:
	@echo "Running CLI command..."
	@go run internal/interfaces/cli/main.go greet meysam

# Build and run the application
build-run: build
	@echo "Running $(BINARY_NAME)..."
	$(RUN) ./$(BINARY_NAME)

# Generate a release
release: build
	@echo "Creating release for version $(VERSION)..."
	@mkdir -p $(BUILD_DIR)
	@cp $(BINARY_NAME) $(BUILD_DIR)/$(BINARY_NAME)-$(VERSION)



# Help message
help:
	@echo "Makefile commands:"
	@echo "  make build        - Build the binary"
	@echo "  make run          - Run the application"
	@echo "  make generate-key - Create APP_KEY"
	@echo "  make migrate      - Init database migrations"
	@echo "  make test         - Run tests"
	@echo "  make fmt          - Format code"
	@echo "  make vet          - Vet code"
	@echo "  make lint         - Lint code"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make build-run    - Build and run the application"
	@echo "  make release      - Create a release"
	@echo "  make help         - Show this help message"

# Phony targets
.PHONY: all build run test fmt vet lint clean build-run release help