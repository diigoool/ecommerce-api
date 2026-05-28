# Ecommerce API

RESTful Ecommerce API built with Golang, Chi Router, PostgreSQL, JWT Authentication, Docker, Swagger Documentation, and CI/CD.

---

## Features

* JWT Authentication & Authorization
* Role-based Access (Admin/User)
* Product Management
* Shopping Cart
* Checkout & Orders
* Pagination & Search
* Request Validation
* Swagger API Documentation
* Docker Support
* PostgreSQL Database
* Unit Test & Integration Test
* CI/CD with GitHub Actions
* Railway Deployment

---

## Tech Stack

* Golang
* Chi Router
* PostgreSQL
* GORM
* JWT
* Docker
* Swagger
* GitHub Actions
* Railway

---

## Live Demo

* API Base URL:
  [Railway Deployment](https://ecommerce-api-production-8669.up.railway.app?utm_source=chatgpt.com)

* Swagger Documentation:
  [Swagger Docs](https://ecommerce-api-production-8669.up.railway.app/swagger/index.html?utm_source=chatgpt.com)

---

## Project Structure

```bash
.
├── config/
├── dto/
├── handlers/
├── mapper/
├── middlewares/
├── models/
├── repositories/
├── routes/
├── seeders/
├── services/
├── tests/
├── utils/
├── Dockerfile
├── docker-compose.yml
├── .github/workflows/
└── main.go
```

---

## Installation

### Clone Repository

```bash
git clone https://github.com/yourusername/ecommerce-api.git
cd ecommerce-api
```

---

## Environment Variables

Create `.env` file:

```env
PORT=3000

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=ecommerce_db

JWT_SECRET=supersecretkey
```

---

## Run Locally

```bash
go mod tidy
go run main.go
```

Server running at:

```bash
http://localhost:3000
```

---

## Run With Docker

```bash
docker compose up --build
```

---

## Swagger Documentation

Generate swagger docs:

```bash
swag init
```

Swagger URL:

```bash
http://localhost:3000/swagger/index.html
```

---

## Authentication

Use JWT token in Authorization header:

```bash
Authorization: Bearer your_token
```

---

## API Endpoints

### Auth

| Method | Endpoint             | Description   |
| ------ | -------------------- | ------------- |
| POST   | `/api/auth/register` | Register user |
| POST   | `/api/auth/login`    | Login user    |

### Products

| Method | Endpoint                   | Description            |
| ------ | -------------------------- | ---------------------- |
| GET    | `/api/product`             | Get all products       |
| GET    | `/api/product/{id}`        | Get product by id      |
| POST   | `/api/product/create`      | Create product (Admin) |
| PATCH  | `/api/product/update/{id}` | Update product (Admin) |
| DELETE | `/api/product/delete/{id}` | Delete product (Admin) |

### Cart

| Method | Endpoint         | Description |
| ------ | ---------------- | ----------- |
| POST   | `/api/cart`      | Add to cart |
| GET    | `/api/cart`      | Get cart    |
| DELETE | `/api/cart/{id}` | Remove item |

### Orders

| Method | Endpoint        | Description     |
| ------ | --------------- | --------------- |
| POST   | `/api/checkout` | Checkout        |
| GET    | `/api/orders`   | Get user orders |

---

## Running Tests

### Unit Test

```bash
go test ./services/...
```

### All Tests

```bash
go test ./...
```

---

## CI/CD

This project uses GitHub Actions for automated testing workflow.

---

## Future Improvements

* Refresh Token
* Redis Cache
* Payment Gateway
* Upload Product Image
* Rate Limiter
* Email Verification
* Microservices Architecture

---

## Author

Dicky Kurniawan

* LinkedIn
* GitHub
