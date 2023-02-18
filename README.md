├── cmd
│   └── main.go
├── internal
│   ├── domain
│   │   ├── user
│   │   │   ├── repository.go
│   │   │   ├── service.go
│   │   │   ├── factory.go
│   │   │   ├── contracts.go
│   │   │   ├── repository_test.go
│   │   │   ├── service_test.go
│   │   │   ├── factory_test.go
│   │   │   └── mock.go
│   ├── infra
│   │   ├── database
│   │   │   ├── mysql
│   │   │   │   ├── connect.go
│   │   │   │   ├── contracts.go
│   │   │   │   └── mock.go
│   │   ├── openapi
│   │   │   ├── openapi.go
│   │   │   ├── openapi_test.go
│   │   │   ├── contracts.go
│   │   │   └── mock.go
│   ├── helpers
│   │   ├── errors
│   │   │   ├── errors.go
│   │   │   └── errors_test.go
│   │   ├── logger
│   │   │   ├── logger.go
│   │   │   └── logger_test.go
│   │   ├── config
│   │   │   ├── environment.go
│   │   │   └── environment_test.go

mkdir -p cmd/internal/domain/user
mkdir -p cmd/internal/infra/database/mysql
mkdir -p cmd/internal/infra/openapi
mkdir -p cmd/internal/helpers/errors
mkdir -p cmd/internal/helpers/logger
mkdir -p cmd/internal/helpers/config
