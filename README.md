# Go REST API Server

From my time as a Hashicorp intern where the I was mentored in go by the talented [Johan Brandhorst-Satzkorn](https://github.com/johanbrandhorst)

This began using the standard library then as an evolution to teach myself popular libraries, i integrated [Gorilla Mux](https://github.com/gorilla/mux)

A production-grade RESTful API server built with Go, implementing modern backend development practices and features. This project represents my first venture into building a production-ready Go server, incorporating essential features for robust API development.

## 🌟 Key Features

- **MVC Architecture**: Clean separation of concerns with Models, Controllers, and a structured data store
- **PostgreSQL Integration**: Robust database integration using a PostgreSQL store
- **RESTful Endpoints**: Well-structured API endpoints for:
  - User management (CRUD operations)
  - Todo management
- **Authentication**: User authentication system with login functionality
- **Error Handling**: Centralized error handling with proper HTTP status codes
- **Middleware Support**: Using Gorilla Mux for routing and middleware
- **Clean Project Structure**:
  ```
  ├── controllers/    # Request handlers and business logic
  ├── models/        # Data models
  ├── server/        # Server configuration and setup
  ├── store/         # Database interactions
  ├── todos/         # Todo-related functionality
  ├── users/         # User management functionality
  └── utils/         # Shared utilities and helpers
  ```

## 🛠️ Technologies Used

- Go (Golang)
- Gorilla Mux (Router)
- PostgreSQL
- JWT for authentication

## 🚀 Getting Started

1. Ensure you have Go installed and PostgreSQL running
2. Clone the repository
3. Set up your PostgreSQL database
4. Update the database connection string in `main.go`
5. Run the server:
   ```bash
   go run main.go
   ```
   The server will start on `localhost:5000`

## 📡 API Endpoints

### Users
- `POST /login` - User authentication
- `GET /users` - Get all users
- `GET /users/{id}` - Get user by ID
- `POST /users` - Create new user
- `DELETE /users/{id}` - Delete user

### Todos
- `GET /todos` - Get all todos
- `GET /todos/{id}` - Get todo by ID

## 💡 Learning Highlights

This project served as an excellent learning experience in:
- Implementing clean architecture in Go
- Setting up proper error handling and response structures
- Working with PostgreSQL in a Go environment
- Building secure authentication systems
- Structuring a production-ready REST API

## 🔒 Security Features

- Secure password handling
- JWT-based authentication
- Proper error handling to prevent information leakage
