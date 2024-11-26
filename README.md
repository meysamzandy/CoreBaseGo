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

```
request flow
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