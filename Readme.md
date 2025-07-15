# URL Shortener Service

A scalable URL Shortener service written in Go, using PostgreSQL as the backend store. It exposes REST APIs to create short links, redirect to original URLs, and fetch stats. It includes a background worker to clean expired links and logs every important event.

---

## 🏗 Folder Structure
```
├── main.go                    # Entry point
├── internal
│   ├── config                 # DB setup
│   │   └── config.go
│   ├── handlers               # HTTP handlers
│   │   ├── handlers.go
│   │   └── urlHandlers.go
│   ├── models                 # Struct definitions
│   │   └── url.go
│   ├── repository             # DB operations
│   │   ├── initDbRepo.go
│   │   └── urlRepo.go
│   ├── services               # Business logic
│   │   └── url_service.go
│   ├── utils                  # Logger, helpers
│   │   └── logger.go
│   └── worker                 # Cleanup routines
│       └── cleaner.go
├── .env                       # Environment variables
├── go.mod                     # Module definition
├── go.sum
├── README.md                  # Project documentation
```

---

## ⚙️ Setup Instructions

### 1. Clone the repository
```bash
git clone https://github.com/yourname/urlshortener.git
cd urlshortener
```

### 2. Set up `.env`
Create a `.env` file:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=yourpassword
DB_NAME=shortenerdb
```

### 3. Run PostgreSQL
Make sure PostgreSQL is running and a database is created:
```sql
CREATE DATABASE urlshortner_qmkr;
```

### 4. Initialize and run
```bash
go mod tidy
go run main.go
```

---

## 🧪 API Endpoints

### POST `/shorten`
Create a short URL.
```json
Request: { "url": "https://example.com" }
Response: { "code": "aZxP9k" }
```

### GET `/{code}`
Redirect to the original URL.

### GET `/stats/{code}`
Get stats like hit count and expiry for a short URL.

---

## 🔄 Background Job
A Go routine periodically runs and deletes expired URLs from the DB every hour.