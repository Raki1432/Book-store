# Book Store API

This repository contains the code for a simple Book Store API built with Go. It provides endpoints to create, read, update, and delete book records in a database. The project is a great starting point for learning how to build RESTful APIs using Go.

## Features

- Add new books (POST)
- View all books (GET)
- View a single book by ID (GET)
- Update book details (PUT)
- Delete a book (DELETE)

## Technologies Used

- Language: Go (Golang)
- Framework: Gorilla Mux
- Database: MySQL
- ORM: GORM

## Prerequisites

Before setting up the project, ensure you have the following installed on your system:

1. [Go](https://golang.org/dl/) (v1.16 or later)
2. [MySQL](https://dev.mysql.com/downloads/)
3. Git

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/Rakesh-honawad/Book-store.git
cd Book-store
```

### 2. Configure the Database

1. Create a MySQL database named `testdb`.
2. Update the database credentials in `pkg/config/config.go`:
   ```go
   d, err := gorm.Open("mysql", "<username>:<password>@/<database>?charset=utf8&parseTime=True&loc=Local")
   ```
   Replace `<username>`, `<password>`, and `<database>` with your MySQL credentials.

### 3. Install Dependencies

Run the following command to install the required Go modules:

```bash
go mod tidy
```

### 4. Run the Application

Start the server:

```bash
go run cmd/main.go
```

The server will start at `http://localhost:9010`.

### 5. Test the Endpoints

You can test the API endpoints using tools like [Postman](https://www.postman.com/) or `curl`.

#### Sample Endpoints:

- **Create a Book**

  ```bash
  POST http://localhost:9010/book/
  Body (JSON):
  {
    "name": "Book Title",
    "author": "Author Name",
    "publication": "Publication Name"
  }
  ```

- **Get All Books**

  ```bash
  GET http://localhost:9010/book/
  ```

- **Get a Book by ID**

  ```bash
  GET http://localhost:9010/book/{id}
  ```

- **Update a Book**

  ```bash
  PUT http://localhost:9010/book/{id}
  Body (JSON):
  {
    "name": "Updated Title",
    "author": "Updated Author",
    "publication": "Updated Publication"
  }
  ```

- **Delete a Book**

  ```bash
  DELETE http://localhost:9010/book/{id}
  ```

### 6. Push to GitHub

After making any changes, push your updates to GitHub:

```bash
git add .
git commit -m "Your commit message"
git push origin main
```

## Folder Structure

```plaintext
Book-store/
├── cmd/               # Main application entry point
├── pkg/               # Application packages
│   ├── config/        # Database configuration
│   ├── controllers/   # API controllers
│   ├── models/        # Data models
│   ├── routes/        # API routes
│   └── utils/         # Utility functions
├── go.mod             # Go module file
├── go.sum             # Dependency checksum file
└── README.md          # Project documentation
```
## Contact

For any questions or feedback, feel free to contact [Rakesh Honawad](mailto\:rakeshhonawadrk1432@gmail.com).

