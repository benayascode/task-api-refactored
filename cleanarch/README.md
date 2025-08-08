# Task Management API - Clean Architecture

This project demonstrates the implementation of Clean Architecture principles in a Go-based Task Management API.

## Architecture Overview

The application is organized into distinct layers following Clean Architecture principles:

### 1. Domain Layer (`Domain/`)
- **Purpose**: Contains the core business entities and interfaces
- **Components**:
  - `domain.go`: Defines Task and User entities, along with repository and use case interfaces
  - Business rules and domain errors

### 2. Use Cases Layer (`Usecases/`)
- **Purpose**: Implements application-specific business logic
- **Components**:
  - `task_usecases.go`: Handles task-related business operations
  - `user_usecases.go`: Handles user-related business operations
- **Responsibilities**: Orchestrates interactions between different layers and enforces business rules

### 3. Repository Layer (`Repositories/`)
- **Purpose**: Abstracts data access logic
- **Components**:
  - `task_repository.go`: Implements task data access operations
  - `user_repository.go`: Implements user data access operations
- **Responsibilities**: Provides data persistence abstraction

### 4. Infrastructure Layer (`Infrastructure/`)
- **Purpose**: Implements external dependencies and services
- **Components**:
  - `jwt_service.go`: JWT token generation and validation
  - `password_service.go`: Password hashing and comparison
  - `auth_middleware.go`: Authentication and authorization middleware
- **Responsibilities**: Handles external concerns like security and authentication

### 5. Delivery Layer (`Delivery/`)
- **Purpose**: Handles incoming requests and responses
- **Components**:
  - `main.go`: Application entry point with dependency injection
  - `controllers/controller.go`: HTTP request handlers
  - `routers/router.go`: Route configuration
- **Responsibilities**: Manages HTTP requests and responses

## Key Features

### Clean Architecture Benefits
1. **Separation of Concerns**: Each layer has a specific responsibility
2. **Dependency Inversion**: High-level modules don't depend on low-level modules
3. **Testability**: Easy to unit test each layer independently
4. **Maintainability**: Clear boundaries make the codebase easier to maintain
5. **Scalability**: New features can be added without affecting existing code

### Business Logic
- Task management (CRUD operations)
- User registration and authentication
- Role-based access control (Admin/User)
- JWT-based authentication
- Password hashing for security

### API Endpoints

#### Public Endpoints
- `POST /register` - Register a new user
- `POST /login` - Authenticate user and get JWT token

#### Protected Endpoints (Require JWT)
- `GET /tasks` - Get all tasks
- `GET /tasks/:id` - Get task by ID

#### Admin Only Endpoints
- `POST /tasks` - Create a new task
- `PUT /tasks/:id` - Update a task
- `DELETE /tasks/:id` - Delete a task
- `POST /promote/:username` - Promote user to admin

## Getting Started

### Prerequisites
- Go 1.21 or higher
- MongoDB running on localhost:27017

### Installation
1. Navigate to the project directory:
   ```bash
   cd task-manager
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run Delivery/main.go
   ```

The server will start on port 8080.

## Testing the API

### 1. Register a User
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser", "password": "password123"}'
```

### 2. Login
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser", "password": "password123"}'
```

### 3. Create a Task (Admin only)
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -H "Authorization: token YOUR_JWT_TOKEN" \
  -d '{"title": "Test Task", "description": "Test Description", "status": "pending"}'
```

### 4. Get All Tasks
```bash
curl -X GET http://localhost:8080/tasks \
  -H "Authorization: token YOUR_JWT_TOKEN"
```

## Architecture Principles Applied

1. **Dependency Rule**: Dependencies point inward. Domain has no dependencies, while Delivery depends on all other layers.

2. **Interface Segregation**: Each layer defines interfaces that are implemented by the layer below it.

3. **Single Responsibility**: Each component has a single, well-defined responsibility.

4. **Open/Closed Principle**: The system is open for extension but closed for modification.

5. **Dependency Injection**: Dependencies are injected rather than created within components.

## Future Enhancements

- Add comprehensive unit tests for each layer
- Implement logging and monitoring
- Add database migrations
- Implement caching layer
- Add API documentation with Swagger
- Implement rate limiting
- Add input validation middleware

## Contributing

When contributing to this project, please follow the Clean Architecture principles and maintain the separation of concerns between layers.
