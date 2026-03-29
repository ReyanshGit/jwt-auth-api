# JWT Auth API — Go + Gin + PostgreSQL

REST API with JWT Authentication built using Go.

## Tech Stack
- Go (Golang)
- Gin Framework
- PostgreSQL + GORM
- JWT Authentication
- bcrypt Password Hashing
- Docker

## Endpoints
| Method | Route | Description |
|--------|-------|-------------|
| POST | /register | User Register |
| POST | /login | User Login + Token |
| GET | /api/profile | Protected Route |

## Run Locally
```bash
go run .
```