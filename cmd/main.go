package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	pb "task-manager/pkg/taskmanager" // Importa el paquete generado de gRPC

	_ "github.com/lib/pq" // Importar el driver de PostgreSQL
	"google.golang.org/grpc"
)

// Definir la estructura del servidor
type server struct {
	pb.UnimplementedTaskManagerServer
	db *sql.DB // Conexión a la base de datos
}

// Implementar el método CreateTask
func (s *server) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	// Insertar una nueva tarea en la base de datos
	query := "INSERT INTO tasks (title, description, status, created_at) VALUES ($1, $2, 'pending', NOW()) RETURNING id"
	var id int32
	err := s.db.QueryRowContext(ctx, query, req.Title, req.Description).Scan(&id)
	if err != nil {
		return nil, err
	}

	// Devolver el ID generado y un mensaje de éxito
	return &pb.CreateTaskResponse{
		Id:      id,
		Success: true,
		Message: "Task created successfully",
	}, nil
}

// GetTasks obtiene la lista de todas las tareas desde la base de datos.
func (s *server) GetTasks(ctx context.Context, req *pb.GetTasksRequest) (*pb.GetTasksResponse, error) {
	query := "SELECT id, title, description, status, created_at FROM tasks"
	rows, err := s.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*pb.Task
	for rows.Next() {
		var task pb.Task
		if err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return &pb.GetTasksResponse{
		Tasks: tasks,
	}, nil
}

func main() {
	// Conectar a la base de datos PostgreSQL
	connStr := "user=task_user password=task_password dbname=task_manager sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Iniciar el listener en el puerto 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// Crear un nuevo servidor gRPC
	grpcServer := grpc.NewServer()
	pb.RegisterTaskManagerServer(grpcServer, &server{db: db})

	fmt.Println("gRPC server listening on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
