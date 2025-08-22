## 🚀 Getting Started

### Prerequisites

- Go 1.20+
- Docker & Docker Compose

### ⚙️ Configuration

Edit `config/config.yaml` as needed.  
See `config/example.config.yaml` for reference.

### 🛠️ Build & Run

#### Local

```bash
make run
```

#### Docker

```bash
docker-compose up --build
```

## 📂 Project Structure

```
├── Dockerfile
├── Makefile
├── cmd/app/main.go
├── config
│
├── docker-compose.yaml # Docker Compose configuration
├── docs
│ ├── docs.go # Swagger docs integration
│ ├── swagger.json
│ └── swagger.yaml # Swagger API definition
├── internal
│ ├── delivery/http # HTTP layer (router, handler, request/response DTOs)
│ ├── entities # Business entities (domain models)
│ ├── repository # Data access layer (database operations)
│ └── usecases # Business logic/service layer
├── mock
│ └── inventory_items.sql
└── pkg
└── database/database.go
```

## 📖 API Documentation

Swagger is available at:

```
http://localhost:8080/swagger/index.html
```

## 🗄️ Database

PostgreSQL is used as the database.  
Connection details can be configured inside `config.yaml`.

## 🧪 Mock Data

You can copy SQL script from `mock/inventory_items.sql`. to create table
