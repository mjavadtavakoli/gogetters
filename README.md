# GoGetters API

A clean, simple, and production-ready **RESTful API** built with:

* **Golang**
* **Gin Framework**
* **PostgreSQL** (via GORM)
* **Layered Architecture** (Handler → Service → Repository)

This project is designed as a practice environment for learning backend development with Go while following real-world architecture patterns.

---

## 🚀 Features

* **CRUD operations** for:

  * Motorcycles
  * Books
  * Laptops
  * Coffee
* **Clean Architecture** with clear separation of concerns
* **Service-layer validations** (business logic)
* **Repository-layer error handling** (including duplicate key detection)
* **PostgreSQL Integration** using GORM
* **Environment-based configuration**
* **Graceful shutdown** support
* **Developer-friendly structure** for extension and scaling

---

## 📂 Project Structure

```
.
├── cmd
│   └── server
│       └── main.go         # Server entrypoint
├── internal
│   ├── motorcycle          # Motorcycle domain module
│   │   ├── handler.go      # HTTP handlers (Gin)
│   │   ├── service.go      # Business logic & validations
│   │   ├── repository.go   # Database queries (GORM)
│   │   └── ...
│   ├── book                # Other domain examples
│   └── models              # Shared database models
├── pkg                     # Utility packages
├── go.mod
├── go.sum
└── README.md
```

---

## 🛠 Requirements

* **Go 1.22+**
* **PostgreSQL 14+**
* **Git**

---

## ⚙️ Setup & Installation

1. Clone the repo:

```bash
git clone https://github.com/yourusername/gogetters.git
cd gogetters
```

2. Install dependencies:

```bash
go mod tidy
```


---

## 📡 API Endpoints (Example: Motorcycles)

### ➕ Create Motorcycle

```
POST /motorcycles
```

#### Body:

```json
{
  "brand": "Yamaha",
  "fueltype": "gasoline",
  "price": 15000,
  "totalspeed": 80
}
```

### 📄 Get All Motorcycles

```
GET /motorcycles
```

### ✏️ Update Motorcycle

```
PUT /motorcycles/:id
```

### ❌ Delete Motorcycle

```
DELETE /motorcycles/:id
```

---

## 🧠 Business Logic (Service Layer)

Examples of validations:

* Price cannot be below **10000**
* Brand must be at least **3 characters**
* Totalspeed cannot exceed **90**
* Duplicate brand names are not allowed

---

## 🧱 Repository Error Handling

Detects PostgreSQL errors like:

* Duplicate key violation (`SQLSTATE 23505`)

And returns clean, readable errors.

---

## 🤝 Contributing

Pull requests are welcome! This repo is mainly for learning but follows clean and scalable patterns.

---

## 📜 License

MIT License.
