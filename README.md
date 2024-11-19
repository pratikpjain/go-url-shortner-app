# go-url-shortner-app
URL SHORTNER APP

*Folder Structure*
```
myapp/
├── .env                # Environment variables file
├── go.mod              # Go module definition
├── main.go             # Application entry point
├── internal/           # Core application logic
│   ├── config/         # Configuration and environment setup
│   │   └── config.go
│   ├── database/       # Database initialization and connections
│   │   └── db.go
│   ├── models/         # Database models
│   │   └── user.go
│   ├── repositories/   # Database query logic
│   │   └── user_repo.go
│   ├── services/       # Business logic
│   │   └── user_service.go
│   ├── routes/         # API route definitions
│   │   └── user_routes.go
│   └── dtos/           # Data Transfer Objects
│       └── user_dto.go
├── pkg/                # Shared utilities and helpers
│   └── logger.go
└── README.md           # Project documentation
```