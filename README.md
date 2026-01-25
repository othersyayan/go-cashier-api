# Go Cashier API

Go Cashier API is a backend service developed in Go (Golang) for managing a
simple cashier system. It creates a RESTful API to manage Products and
Categories.

This project demonstrates the use of **Clean Architecture** in Go, separating
the code into specialized layers for better maintainability and scalability.

## 🚀 Features

- **Clean Architecture**: Structured into Entity, Repository, Usecase, and
  Delivery layers.
- **In-Memory Storage**: Uses in-memory data structures for data persistence
  (easy to run, no database required).
- **RESTful API**: Standard HTTP methods for CRUD operations.
- **Resource Management**: Full CRUD support for **Products** and
  **Categories**.

## 📂 Project Structure

```
go-cashier-api/
├── internal/
│   ├── entity/       # Domain models (structs)
│   ├── repository/   # Data access layer (in-memory implementation)
│   ├── usecase/      # Business logic layer
│   └── delivery/     # HTTP handlers and routing
└── main.go           # Entry point (Dependency Injection & Server Start)
```

## 🛠️ Installation & Run

1.  **Clone the repository** (if applicable) or go to the project directory.

2.  **Run the server**:

    ```bash
    go run main.go
    ```

3.  The server will start at `http://localhost:8080`.

## 🔗 API Endpoints

### Products

| Method   | Endpoint             | Description          | Body (JSON)                                            |
| :------- | :------------------- | :------------------- | :----------------------------------------------------- |
| `GET`    | `/api/products`      | Get all products     | -                                                      |
| `POST`   | `/api/products`      | Create a new product | `{"name": "...", "price": 1000, "stock": 10}`          |
| `GET`    | `/api/products/{id}` | Get a product by ID  | -                                                      |
| `PUT`    | `/api/products/{id}` | Update a product     | `{"id": 1, "name": "...", "price": 1000, "stock": 10}` |
| `DELETE` | `/api/products/{id}` | Delete a product     | -                                                      |

### Categories

| Method   | Endpoint               | Description           | Body (JSON)                                      |
| :------- | :--------------------- | :-------------------- | :----------------------------------------------- |
| `GET`    | `/api/categories`      | Get all categories    | -                                                |
| `POST`   | `/api/categories`      | Create a new category | `{"name": "...", "description": "..."}`          |
| `GET`    | `/api/categories/{id}` | Get a category by ID  | -                                                |
| `PUT`    | `/api/categories/{id}` | Update a category     | `{"id": 1, "name": "...", "description": "..."}` |
| `DELETE` | `/api/categories/{id}` | Delete a category     | -                                                |

### Health Check

- `GET /health`: Check if the API is running.

## 📚 Tech Stack

- **Go**: Core language.
- **Standard Library**: `net/http`, `encoding/json` for building the API.
