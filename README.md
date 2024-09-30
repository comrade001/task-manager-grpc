
# Task Manager gRPC Service

This repository contains a basic Task Manager service implemented in Go using gRPC and PostgreSQL. The purpose of this project is to showcase CRUD operations for managing tasks in a simple and scalable backend service.

## Overview

The Task Manager service provides the following functionalities:

- **CreateTask**: Create a new task with a title and description.
- **GetTasks**: Retrieve a list of all existing tasks.
- **UpdateTaskStatus**: Update the status of a specific task (pending, in progress, completed).
- **DeleteTask**: Delete a task by its ID.

The project uses `Protobuf` to define the service and message structures, and `gRPC` to facilitate communication between the server and the client. PostgreSQL is used as the database for storing tasks.

## Project Structure

The project is organized into the following folders:

```
TaskManager/
├── cmd/                    # Contains the main entry point of the application
│   └── main.go
├── db/                     # Database-related files (scripts, migrations)
├── internal/               # Business logic and server implementation
├── pkg/                    # Generated protobuf files and shared utilities
│   ├── task_manager.pb.go
│   └── task_manager_grpc.pb.go
├── proto/                  # Protobuf definitions
│   └── task_manager.proto
└── README.md               # Project documentation
```

## Technologies Used

- **Go**: Backend programming language for implementing the gRPC server.
- **gRPC**: Framework for communication between client and server using HTTP/2 and Protobuf.
- **PostgreSQL**: Relational database for storing tasks.
- **Protobuf**: Serialization format used for defining gRPC messages and services.

## Setup and Installation

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.16 or higher)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Protobuf Compiler (`protoc`)](https://grpc.io/docs/protoc-installation/)

### Steps

1. **Clone the repository**:
   ```bash
   git clone https://github.com/your-username/task-manager-grpc.git
   cd task-manager-grpc
   ```

2. **Install Go dependencies**:
   ```bash
   go mod tidy
   ```

3. **Setup the database**:
   - Create a PostgreSQL database named `task_manager`.
   - Create a user `task_user` with password `task_password`.
   - Grant necessary permissions to `task_user` on the `task_manager` database.

4. **Run database migrations**:
   - (Optional) You can create a `db/migrations` folder and add SQL scripts to initialize the `tasks` table.

5. **Compile the `.proto` file**:
   ```bash
   protoc --go_out=./pkg --go-grpc_out=./pkg proto/task_manager.proto
   ```

6. **Run the gRPC server**:
   ```bash
   go run cmd/main.go
   ```

7. **Test the service**:
   - Use a gRPC client like [Postman](https://learning.postman.com/docs/sending-requests/grpc/) or `grpcurl` to test the service methods (`CreateTask`, `GetTasks`, etc.).

## Example Requests

### Create a New Task

```proto
rpc CreateTask(CreateTaskRequest) returns (CreateTaskResponse);
```

**Request**:
```json
{
  "title": "Finish Go project",
  "description": "Complete the task manager project and update README.md"
}
```

**Response**:
```json
{
  "id": 1,
  "success": true,
  "message": "Task created successfully"
}
```

### Get All Tasks

```proto
rpc GetTasks(GetTasksRequest) returns (GetTasksResponse);
```

**Request**:
```json
{}
```

**Response**:
```json
{
  "tasks": [
    {
      "id": 1,
      "title": "Finish Go project",
      "description": "Complete the task manager project and update README.md",
      "status": "pending",
      "created_at": "2024-09-30T10:15:30Z"
    }
  ]
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Feel free to open issues or submit pull requests if you have any suggestions or improvements. Contributions are welcome!

---

**Author**: Your Name (your-email@example.com)  
**LinkedIn**: [Your LinkedIn Profile](https://linkedin.com/in/your-profile)  
