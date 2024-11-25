

```
.
├── cmd/                # Application entry points (main files)
│   ├── api/            # API server entry
│   └── worker/         # Background workers (not implemented yet)
├── configs/            # Configuration files (YAML, JSON, or ENV templates)
├── internal/           # Application code (not exposed as a public package)
│   ├── domain/         # Core business logic (rules and behaviors, domain models, and validators)
│   │   ├── importer/   # Excel importer domain
│   │   │   ├── entity/ # Import-related domain models/entities
│   │   │   └── service/ # Excel import business logic
│   │   └── common/     # Shared domain entities (if needed)
│   ├── application/    # Application services (coordinates the use cases and orchestrates tasks like workflows)
│   │   ├── importer/   # Excel import workflows (not yet implemented)
│   │   └── common/     # Shared application logic (not yet implemented)
│   │   └── errors/     # Application-level error handling (not yet implemented)
│   ├── infrastructure/ # External dependencies and frameworks (handles technical details)
│   │   ├── persistence/ # Database-related code (repositories)
│   │   │   ├── Migrations/ # Database migrations (not yet implemented)
│   │   ├── http/        # HTTP server setup and handlers
│   │   │   ├── Middlewares/ # Middlewares (e.g., logging, auth, rate limiting)
│   │   │   ├── server.go  # HTTP server setup (not yet implemented)
│   │   ├── logging/     # Logging setup (not yet implemented)
│   │   └── config/      # Configuration and environment loading
│   │       └── base.go  # Base configuration setup (not yet implemented)
│   ├── interfaces/      # Adapters for APIs, CLI, and external systems
│   │   ├── rest/        # REST API endpoints
│   │   │   ├── sampleFeature/ # Importer-specific HTTP routes
│   │   │   │   ├── controller/  # Handlers for importer endpoints
│   │   │   │   └── sampleFeatureRoutes.go  # Importer-specific HTTP route definitions (not yet implemented)
│   │   ├── grpc/        # gRPC endpoints (not yet implemented)
│   │   └── cli/         # CLI tools (if required, not yet implemented)
│   └── utils/           # Utilities and helper functions (not yet implemented)
├── pkg/                 # Shared reusable packages (if needed for other projects, not yet implemented)
├── scripts/             # Deployment and build scripts (not yet implemented)
├── test/                # Test cases
│   ├── unit/            # Unit tests (not yet implemented)
│   └── integration/     # Integration tests (not yet implemented)
├── go.mod               # Go module definition
└── go.sum               # Go module checksum
                         # Go module checksum

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