# Task Manager API (Clean Architecture Refactored)

A RESTful API for managing tasks, refactored using Clean Architecture principles for better structure, scalability, and maintainability.

---

---

##  What Changed (Compared to Old Code)

| **Old Code**                             | **Refactored Code (Clean Arch)**                             |
| ---------------------------------------- | ------------------------------------------------------------ |
| Flat structure (`controllers`, `data`)   | Organized into **layers** (Delivery, Domain, Usecases, etc.) |
| Business logic in controllers            | Business logic moved to **Usecases**                         |
| MongoDB code in `data` package           | Database code moved to **Repositories**                      |
| Middleware mixed with JWT logic          | JWT logic split into **Infrastructure**                      |
| Tight coupling (Gin, MongoDB everywhere) | Decoupled layers with clear responsibilities                 |

---

##  Folder Responsibilities

* **Delivery/** – Handles HTTP layer (controllers, routes).
* **Domain/** – Pure business models (Task, User).
* **Usecases/** – Business rules and application logic.
* **Repositories/** – Interfaces for database operations (MongoDB).
* **Infrastructure/** – External services (JWT, password hashing, middleware).

---

## How to Run


 Run the app:

   ```bash
   cd Delivery
   go run main.go
   ```
 API available at:

   ```
   http://localhost:8080
   ```