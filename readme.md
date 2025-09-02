# Go Test SM (Anang Novriadi)

A simple API using [Go Fiber](https://gofiber.io/), GORM, and JWT.

---

## Features
- User registration with name, email & password
- Login with JWT authentication
- Protected routes using middleware
- MySQL as database


## Installation

```bash
git clone https://github.com/anangnovriadi/go-test-sm.git
cd go-test-sm
```

## How to Run

```bash
go mod tidy
cp .env.example .env
go run main.go
```

## Example API

Login
```curl
curl -X POST http://localhost:3010/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Anang Novriadi",
    "email": "anangnov99@gmail.com",
    "password": "12345678"
  }'
```

Register
```curl
curl -X POST http://localhost:3010/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "anangnov99@gmail.com",
    "password": "12345678"
  }'
```

Get User
```curl
curl -X GET http://localhost:3010/api/user \
  -H "Authorization: Bearer <your_token>"
```

## Run Tests
```bash
go test -v ./...
```