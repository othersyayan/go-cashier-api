# Go Cashier API

Go Cashier API is a backend service developed in Go (Golang) for managing a
simple cashier system. It uses **PostgreSQL (via Supabase)** for persistence and
creates a RESTful API to manage Products and Categories.

This project uses a layered architecture, separating repositories, services, and
handlers.

## 🚀 Features

- **Layered Architecture**: Repository, Service, and Handler pattern.
- **PostgreSQL Database**: Persistent storage using `lib/pq` driver (configured
  for Supabase).
- **RESTful API**: Standard HTTP methods for CRUD operations.
- **Resource Management**:
  - **Products**: Manage inventory with price and stock.
  - **Categories**: Organize products into categories.
- **UUIDs**: Uses UUID string identifiers for all resources.
- **Validation**: Ensures integrity (e.g. required Category ID for products).

## 📂 Project Structure

```
go-cashier-api/
├── database/         # Database connection logic
├── handlers/         # HTTP Handlers (Controllers)
├── models/           # Domain data structures
├── repositories/     # Database access layer
├── services/         # Business logic layer
└── main.go           # Entry point and routing
```

## 🛠️ Configuration

The application requires environment variables for configuration. Create a
`.env` file in the root directory:

```env
PORT=8080
DATABASE_URL=postgres://user:password@host:port/dbname?sslmode=require
```

## 🚀 Run the Application

1.  **Clone the repository**
2.  **Install dependencies**:
    ```bash
    go mod tidy
    ```
3.  **Run the server**:
    ```bash
    go run main.go
    ```
4.  The server will start at `http://localhost:8080`.

## 🔗 API Endpoints

### Products

| Method   | Endpoint             | Description                                   |
| :------- | :------------------- | :-------------------------------------------- |
| `GET`    | `/api/products`      | Get all products (includes Category Name)     |
| `POST`   | `/api/products`      | Create a new product (requires `category_id`) |
| `GET`    | `/api/products/{id}` | Get a product by ID                           |
| `PUT`    | `/api/products/{id}` | Update a product                              |
| `DELETE` | `/api/products/{id}` | Delete a product                              |

### Categories

| Method   | Endpoint               | Description           |
| :------- | :--------------------- | :-------------------- |
| `GET`    | `/api/categories`      | Get all categories    |
| `POST`   | `/api/categories`      | Create a new category |
| `GET`    | `/api/categories/{id}` | Get a category by ID  |
| `PUT`    | `/api/categories/{id}` | Update a category     |
| `DELETE` | `/api/categories/{id}` | Delete a category     |

### Health Check

- `GET /health`: Check if the API is running.

## 📚 Tech Stack

- **Go**: Core language.
- **PostgreSQL**: Database.
- **lib/pq**: Postgres driver.
- **Viper**: Configuration management.
- **Standard Library**: `net/http`, `encoding/json` for API.
