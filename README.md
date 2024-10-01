
# Task Manager gRPC Service

This is a gRPC-based Task Manager service built with Go, PostgreSQL, and Protobuf. It provides functionalities to create, read, update, and delete tasks in a PostgreSQL database using a gRPC interface.

## Prerequisites

1. [Go](https://golang.org/doc/install) (version 1.17 or higher).
2. [PostgreSQL](https://www.postgresql.org/download/) installed and running.
3. `psql` configured in the `PATH` environment variable for command-line operations.
4. `protoc` (Protocol Buffers compiler) installed for generating Go files from `.proto` definitions.

## Getting Started

### Step 1: Clone the repository

```bash
git clone https://github.com/your-username/task-manager.git
cd task-manager
```

### Step 2: Set up the PostgreSQL database

1. Connect to PostgreSQL as an admin user (e.g., `postgres`):

   ```bash
   psql -U postgres
   ```

2. Create the database `task_manager`:

   ```sql
   CREATE DATABASE task_manager;
   ```

3. Create the user `task_user` with a password:

   ```sql
   CREATE USER task_user WITH PASSWORD 'task_password';
   ```

4. Grant all privileges on the `task_manager` database to `task_user`:

   ```sql
   GRANT ALL PRIVILEGES ON DATABASE task_manager TO task_user;
   ```

5. Grant usage and create permissions on the `public` schema to `task_user`:

   ```sql
   GRANT USAGE ON SCHEMA public TO task_user;
   GRANT CREATE ON SCHEMA public TO task_user;
   ```

### Step 3: Create the `tasks` table

1. Connect to the `task_manager` database as `task_user`:

   ```bash
   psql -U task_user -d task_manager
   ```

2. Create the `tasks` table:

   ```sql
   CREATE TABLE tasks (
       id SERIAL PRIMARY KEY,
       title VARCHAR(255) NOT NULL,
       description TEXT,
       status VARCHAR(50) NOT NULL DEFAULT 'pending',
       created_at TIMESTAMP NOT NULL DEFAULT NOW()
   );
   ```

### Step 4: Generate Go files from the `.proto` file

Make sure you're in the `proto` directory:

```bash
cd proto
```

Run the following command to generate the Go files:

```bash
protoc --go_out=../pkg --go-grpc_out=../pkg task_manager.proto
```

### Step 5: Update the configuration in `main.go`

Check that the database connection string in `main.go` is correct:

```go
connStr := "user=task_user password=task_password dbname=task_manager sslmode=disable"
```

### Step 6: Run the gRPC server

Start the server by running the following command:

```bash
go run cmd/main.go
```

You should see a message like:

```
gRPC server listening on port 50051...
```

### Step 7: Test the service with Postman or `grpcurl`

1. Import the `task_manager.proto` file in Postman and configure the connection to `localhost:50051`.
2. Use the `CreateTask` method with the following payload:

   ```json
   {
     "title": "New Task",
     "description": "This is a test task from Postman"
   }
   ```

3. You should receive a response like:

   ```json
   {
     "id": 1,
     "success": true,
     "message": "Task created successfully"
   }
   ```

## Additional Information

This project uses gRPC for communication and PostgreSQL for data storage. For more information on how to extend the service or modify it, check the [official gRPC documentation](https://grpc.io/docs/) and [PostgreSQL documentation](https://www.postgresql.org/docs/).
