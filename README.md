# 🧑‍💻 User Age API

A RESTful API built with **Go** to manage users with their name and date of birth. The API dynamically calculates and returns a user's age when fetching user details.

Built with clean architecture principles — **Handler → Service → Repository** layering, type-safe SQL via **SQLC**, input validation with **go-playground/validator**, and structured logging with **Uber Zap**.

---

## 🔧 Tech Stack

| Layer              | Technology                                                                 |
| ------------------ | -------------------------------------------------------------------------- |
| **Language**       | [Go 1.25](https://golang.org/)                                            |
| **HTTP Framework** | [GoFiber v2](https://gofiber.io/)                                         |
| **Database**       | [PostgreSQL 16](https://www.postgresql.org/)                               |
| **SQL Code Gen**   | [SQLC](https://sqlc.dev/)                                                  |
| **Validation**     | [go-playground/validator v10](https://github.com/go-playground/validator)  |
| **Logging**        | [Uber Zap](https://github.com/uber-go/zap)                                |
| **Containerization** | [Docker Compose](https://docs.docker.com/compose/)                      |

---

## 📁 Project Structure

```
user-age-api/
├── cmd/server/              # Application entrypoint
│   └── main.go
├── config/                  # Environment config loader
│   └── config.go
├── db/
│   ├── migrations/          # SQL migration files
│   │   └── 001_create_users.sql
│   └── sqlc/                # SQLC config + generated Go code
│       ├── db.go
│       ├── models.go
│       ├── query.sql
│       └── query.sql.go
├── internal/
│   ├── handler/             # HTTP handlers (request parsing, response)
│   │   └── user_handler.go
│   ├── service/             # Business logic (age calculation, transforms)
│   │   └── user_service.go
│   ├── repository/          # Data access layer (wraps SQLC queries)
│   │   └── user_repository.go
│   ├── models/              # Request & Response DTOs
│   │   └── user.go
│   ├── routes/              # Fiber route registration
│   │   └── routes.go
│   ├── middleware/          # Custom middleware
│   └── logger/              # Zap logger initialization
│       └── logger.go
├── docker-compose.yml       # PostgreSQL container
├── sqlc.yaml                # SQLC configuration
├── .env.example             # Environment variable template
├── .gitignore
├── go.mod
└── go.sum
```

---

## 🗂️ Database Schema

**`users`** table:

| Field  | Type     | Constraints  |
| ------ | -------- | ------------ |
| `id`   | SERIAL   | PRIMARY KEY  |
| `name` | TEXT     | NOT NULL     |
| `dob`  | DATE     | NOT NULL     |

> The `age` field is **not stored** in the database. It is computed dynamically using Go's `time` package every time a user is fetched.

---

## 🚀 Getting Started

### Prerequisites

- [Go 1.25+](https://golang.org/dl/)
- [Docker & Docker Compose](https://docs.docker.com/get-docker/)
- [SQLC](https://docs.sqlc.dev/en/latest/overview/install.html) *(only needed if regenerating query code)*

### 1. Clone the Repository

```bash
git clone https://github.com/ksploitx/user-age-api.git
cd user-age-api
```

### 2. Configure Environment Variables

```bash
cp .env.example .env
```

Edit `.env` with your database credentials:

```env
DB_HOST=localhost
DB_PORT=5433
DB_USER=admin
DB_PASSWORD=secret
DB_NAME=userdb
APP_PORT=3000
```

### 3. Start PostgreSQL

```bash
docker compose up -d
```

This starts a PostgreSQL 16 container on port **5433**.

### 4. Run the Database Migration

```bash
psql -h localhost -p 5433 -U admin -d userdb -f db/migrations/001_create_users.sql
```

> Enter password `secret` when prompted (or set `PGPASSWORD=secret` beforehand).

### 5. Install Dependencies & Run the Server

```bash
go mod download
go run cmd/server/main.go
```

The server starts at **`http://localhost:3000`**.

---

## 📡 API Endpoints

### Health Check

```
GET /health
```

**Response:** `200 OK`
```json
{
  "status": "ok"
}
```

---

### Create User

```
POST /users/
```

**Request Body:**
```json
{
  "name": "Alice",
  "dob": "1990-05-10"
}
```

**Response:** `201 Created`
```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10"
}
```

---

### Get User by ID

```
GET /users/:id
```

**Response:** `200 OK`
```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 36
}
```

> `age` is dynamically calculated from `dob` using Go's `time` package.

---

### List All Users

```
GET /users/
```

**Response:** `200 OK`
```json
[
  {
    "id": 1,
    "name": "Alice",
    "dob": "1990-05-10",
    "age": 36
  }
]
```

---

### Update User

```
PUT /users/:id
```

**Request Body:**
```json
{
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

**Response:** `200 OK`
```json
{
  "id": 1,
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

---

### Delete User

```
DELETE /users/:id
```

**Response:** `204 No Content`

---

## 🔄 Architecture

```
HTTP Request
     │
     ▼
┌──────────┐    ┌──────────┐    ┌──────────────┐    ┌──────┐    ┌────────────┐
│  Handler  │───▶│ Service  │───▶│  Repository  │───▶│ SQLC │───▶│ PostgreSQL │
└──────────┘    └──────────┘    └──────────────┘    └──────┘    └────────────┘
     │               │
     │               ├── Parses DOB strings
  Validates          ├── Calculates age
  input with         └── Transforms DB models → API responses
  go-playground/
  validator
```

- **Handler** — Parses HTTP input, validates request bodies with `go-playground/validator`, returns structured JSON responses with appropriate HTTP status codes.
- **Service** — Business logic layer. Parses date-of-birth strings, calculates age dynamically, and transforms database models into API response DTOs.
- **Repository** — Thin wrapper over SQLC-generated queries. Provides a clean interface between the service layer and database access.
- **SQLC** — Generates type-safe Go code from raw SQL queries, ensuring compile-time safety for all database operations.

---

## ✅ Error Handling

| Scenario               | HTTP Status              | Response                                    |
| ---------------------- | ------------------------ | ------------------------------------------- |
| Invalid request body   | `400 Bad Request`        | `{"error": "invalid request body"}`         |
| Validation failure     | `400 Bad Request`        | `{"error": "<validation details>"}`         |
| Invalid ID parameter   | `400 Bad Request`        | `{"error": "invalid id"}`                   |
| User not found         | `404 Not Found`          | `{"error": "user not found"}`               |
| Internal server error  | `500 Internal Server Error` | `{"error": "could not create user"}`     |

---

## 🧪 Example cURL Commands

```bash
# Health check
curl http://localhost:3000/health

# Create a user
curl -X POST http://localhost:3000/users/ \
  -H "Content-Type: application/json" \
  -d '{"name": "Alice", "dob": "1990-05-10"}'

# Get user by ID
curl http://localhost:3000/users/1

# List all users
curl http://localhost:3000/users/

# Update a user
curl -X PUT http://localhost:3000/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Alice Updated", "dob": "1991-03-15"}'

# Delete a user
curl -X DELETE http://localhost:3000/users/1
```

---

## 📦 Key Design Decisions

- **Age is calculated, not stored** — `dob` is persisted in the database as a `DATE`. The `age` field is computed on every GET request using Go's `time` package, so it's always accurate and never stale.
- **Separate response DTOs** — `UserResponse` (for create/update) omits `age`, while `UserDetailResponse` (for get/list) includes it, matching the API specification exactly.
- **SQLC over ORM** — Raw SQL with SQLC provides full control over queries with compile-time type safety, avoiding the overhead and magic of a traditional ORM.
- **Structured logging** — All key actions (user created, errors) are logged with Uber Zap in structured JSON format for production-grade observability.

---

## 📜 License

MIT

---