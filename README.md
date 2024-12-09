
# Todo API with Go, Fiber, and MongoDB

This is a simple RESTful API for task (todo) management built with Go, Fiber (web framework), and MongoDB (NoSQL database). This project is an example for study and learning purposes about backend, RESTful APIs, and MongoDB integration using Go.

## Features
*  **GET /api/todos:**  List all tasks.
*  **POST /api/todos:**  Create a new task.
*  **PATCH /api/todos/:id**:  Update a task (mark as completed).
*  **DELETE /api/todos/:id**:  Delete a task.

## Technologies
* **GO** (Golang)
* **Fiber** (Web framework for Go)
* **MongoDB** (NoSQL database)
* **dotenv** (Environment variable loading)

## Project Structure
The project is organized as follows:
```
go-react-tasks/
├── config/                # Configuration for environment and variables
│   └── config.go          # Load environment variables
├── controllers/           # Controllers for routing (request logic)
│   └── todoController.go  # Functions controlling the API routes
├── models/                # Data models (structures)
│   └── todo.go            # Task (todo) structure
├── repository/            # Logic for database access
│   └── todoRepository.go  # Functions to interact with MongoDB
├── routes/                # Routing configuration
│   └── routes.go          # Definition of the application's routes
├── .env                   # Environment variables
├── main.go                # Main file that starts the application
├── go.mod                 # Dependencies file
├── go.sum                 # Dependencies file
```
## Prerequisites
Before running the application, make sure you have:
* **Go** (version 1.18 or later) intalled. [Install Go](https://go.dev/doc/install)
* **MongoDB** running locally or an account on [MongoDB Atlas.](https://www.mongodb.com/atlas)
* **Git** to clone the repository.

## License
This project is for study purposes only. You can use it as a base to learn more about Go, RESTful APIs, and MongoDB, but do not use this code in production without proper improvements and adjustments.
