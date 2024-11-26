# CoreBaseGo

**CoreBaseGo** is a boilerplate Go project designed with **clean architecture (hexagonal architecture)** principles. The goal is to provide a scalable and maintainable structure for Go applications with clear separations of concerns, making it easy to extend and adapt to changing requirements.

---

## Features
- **Clean Architecture**: Separation of concerns into domain, application, infrastructure, and interface layers.
- **Database Integration**: Support for database operations through the `gorm` ORM.
- **RESTful API**: Predefined routes and controllers for APIs.
- **Modular Design**: Organized into feature-specific modules for scalability.
- **Extensibility**: Designed to easily integrate new external services, routes, and features.

---

## Folder Structure Overview
The project follows a modular and layered architecture:

```
1- cmd/: Entry points for the application (e.g., API server, workers).
2- configs/: Configuration templates.
3- internal/: Application code (business logic, use cases, and adapters).
    - domain/: Core business logic and rules.
    - application/: Application workflows and orchestrations.
    - infrastructure/: Technical implementations like database and HTTP server.
    - interfaces/: Adapters for APIs and external interactions.

4- pkg/: Shared reusable packages (if needed).
5- scripts/: Build and deployment scripts.
6- test/: Unit and integration tests.
```

---

## Folder Structure
The project follows a modular and layered architecture:
```
.
├── cmd/                # Application entry points (main files)
│   ├── api/            # API server entry point (main.go or similar to start the server)
│   └── worker/         # Background workers entry point (not implemented yet)
├── configs/            # Configuration files (YAML, JSON, or environment templates)
├── internal/           # Application-specific code (not exposed as public Go packages)
│   ├── domain/         # Core business logic and entities
│   │   ├── sampleFeature/  
│   │   │   ├── entity/ 
│   │   │   │    └── sampleFeatureEntity.go # Defines the SampleFeature entity
│   │   │   └── service/ 
│   │   │       └── sampleFeatureService.go # Business rules for SampleFeature (validation, logic)
│   │   └── common/     # Shared domain entities or logic (if required in future)
│   │
│   ├── application/    # Use cases and application-level workflows
│   │   ├── sampleFeature/   # Application logic for SampleFeature workflows
│   │   │   └── errors/      # Custom application-specific errors (not yet implemented)
│   │   └── common/          # Shared application logic (not yet implemented)
│   │
│   ├── infrastructure/ # External dependencies and technical details
│   │   ├── persistence/ # Database-related code (repositories and migrations)
│   │   │   ├── Migrations/ # Database schema migrations (not yet implemented)
│   │   │   │   └── migrations.go # Handles database schema changes
│   │   │   ├── repositories/ 
│   │   │   │   └── sampleFeatureRepo.go # Repository for SampleFeature database operations
│   │   │   └── database.go  # Database connection and setup
│   │   ├── http/        # HTTP server setup and middleware
│   │   │   ├── Middlewares/ # Middleware for logging, authentication, and rate limiting
│   │   │   │   ├── common/ # Shared middlewares
│   │   │   │   └── sampleFeature/ # Middlewares specific to SampleFeature (if needed)
│   │   │   └── server.go  # HTTP server setup
│   │   │   └── routes.go  # Route initialization and grouping
│   │   ├── logging/     # Logging configuration and utilities
│   │   │   └── log.go   # Logger setup (e.g., structured logging)
│   │   └── config/      # Configuration loading and management
│   │       └── base.go  # Base configuration loader
│   │
│   ├── interfaces/      # Adapters for user-facing APIs (REST, gRPC, CLI)
│   │   ├── rest/        # REST API entry points
│   │   │   ├── sampleFeature/ # HTTP routes for SampleFeature
│   │   │   │   ├── controller/  # Handlers for SampleFeature endpoints
│   │   │   │   └── sampleFeatureRoutes.go  # Route definitions for SampleFeature
│   │   ├── grpc/        # gRPC endpoints (not yet implemented)
│   │   └── cli/         # CLI commands (if required, not yet implemented)
│   │
│   └── utils/           # Utility functions (e.g., helpers for JSON, dates, and string parsing)
│
├── pkg/                 # Shared reusable packages (for sharing across projects, if needed)
├── scripts/             # Deployment and build scripts (not yet implemented)
├── storage/             # static file should be here
├── test/                # Test cases
│   ├── unit/            # Unit tests
│   └── integration/     # Integration tests
├── go.mod               # Go module definition (dependencies)
└── go.sum               # Go module checksum file
   

```
## WorkFellow For Sample Feature
what happened in this system for example:
```
Routes (sampleFeatureRoute.Routes)
├── Endpoint: GET /api/v1/test
│   ├── sampleFeatureController.List
│   │   ├── sampleFeatureApplication.ListSampleFeature
│   │   │   ├── persistence.GetInstance (DB Connection)
│   │   │   └── sampleFeatureRepo.GetSampleFeature (Fetch Data)
│   │   │       └── Gorm: db.Find (Retrieve Records from Database)
│   │   └── rest.JSONOutput (Return JSON Response)
│   └── Response: List of SampleFeature records or error
│
├── Endpoint: POST /api/v1/test
│   ├── sampleFeatureController.Store
│   │   ├── c.ShouldBindJSON (Parse Input from Request Body)
│   │   ├── sampleFeatureApplication.CreateSampleFeature
│   │   │   ├── service.ValidateSampleFeatureInput (Validate Name)
│   │   │   ├── service.CreateSampleFeature (Create SampleFeature Entity)
│   │   │   ├── persistence.GetInstance (DB Connection)
│   │   │   └── sampleFeatureRepo.StoreSampleFeature (Save Data)
│   │   │       └── Gorm: db.Create (Insert Record into Database)
│   │   └── rest.JSONOutput (Return JSON Response)
│   └── Response: Created SampleFeature or error


```


---

## Getting Started

Follow these steps to set up and run the **CoreBaseGo** project:

### Prerequisites
- **Go 1.20+** installed on your system.
- A configured database (compatible with `gorm`, e.g., MySQL, PostgreSQL).
- [Git](https://git-scm.com/) installed for cloning the repository.

---

### Steps

1. **Clone the Repository**  
   Clone the project repository from GitHub:
   ```bash
   git clone https://github.com/username/CoreBaseGo.git
   cd CoreBaseGo

2. **Install Dependencies**  
   Install all required dependencies for the project:
   ```bash
   go mod tidy

3. **Configure the Application**
   Copy the example configuration file to create your own configuration file:
   ```bash
   cp configs/.env.example configs/.env
   make generate-key
   

4. **Run Database Migrations**  
   Ensure your database schema is up-to-date:
   ```bash
   make migrate



5. **Run the API Server**  
   Start the application by running the API server:
   ```bash
   make run

---

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