# Echo-Gorm API Project

This project is a simple REST API built using Go with the Echo framework and Gorm ORM. It provides CRUD operations for managing products, categories, and carts, using SQLite as the database.

## Features

- CRUD operations for Products, Categories, and Carts.
- Uses Gorm ORM with relationships.
- Supports preloading related entities.
- Scoped queries for better organization.

## Prerequisites

Ensure you have the following installed:

- [Go](https://go.dev/dl/) (latest version)
- [GCC compiler](https://gcc.gnu.org/) (required for SQLite support)
- [SQLite](https://www.sqlite.org/) (optional for database management)

## Setup Instructions

```sh
# Clone the repository
git clone git@github.com:mazurmilosz000/e-biznes.git
cd .\go\

# Install dependencies
go mod tidy

# Run the project (Windows PowerShell)
$env:CGO_ENABLED=1; go run main.go

# Run the project (Linux/macOS)
export CGO_ENABLED=1 && go run main.go
```
The server should start on http://localhost:8080.

## License

This project is open-source under the MIT License.