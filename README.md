# 🧠 User Management API - Go (Golang)

A simple RESTful API for user management built with Go, Gin, GORM, and MySQL.

---

## 🚀 Features

- 🔐 Register & Login (with hashed password)
- 🔑 JWT authentication middleware
- 📄 CRUD user
- 🧪 Input validation and error handling
- 🌱 .env support using `godotenv`

---

## 🛠️ Tech Stack

- **Language**: Go
- **Framework**: Gin Gonic
- **ORM**: GORM
- **Database**: MySQL
- **JWT**: `github.com/golang-jwt/jwt/v5`
- **Validation**: `go-playground/validator`
- **DotEnv**: `github.com/joho/godotenv`

---

## 📁 Project Structure

```text ├── config/ # .env loader ├── controllers/ # Login, register, user CRUD ├── database/ # Database setup & migration ├── helpers/ # Hash, JWT, validation helpers ├── middlewares/ # Auth middleware ├── models/ # User model ├── routes/ # API route definitions ├── structs/ # Request & response structs ├── main.go # Entry point ├── README.md # This file ```


---

## ⚙️ Getting Started

### Clone the repo
```bash
git clone https://github.com/anggakrnwn/user-management-api.git
cd user-management-api
