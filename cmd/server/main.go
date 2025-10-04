package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"GreatProject/internal/db"
	"GreatProject/internal/generated"
	"GreatProject/internal/handlers"
	"GreatProject/internal/repository"
	"GreatProject/internal/service"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Подключение к PostgreSQL
	conn, err := pgx.Connect(context.Background(), getDatabaseURL())

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer conn.Close(context.Background())

	// Создаем Queries для работы с БД
	queries := db.New(conn)

	// Создаем слои приложения (Repository → Service → Handler)
	taskRepo := repository.NewTaskRepository(queries)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	// Создаем Echo сервер
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Регистрируем роуты
	generated.RegisterHandlers(e, taskHandler)

	// Health check endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
			"time":   time.Now().Format(time.RFC3339),
		})
	})

	// Запуск сервера
	port := getPort()

	// Graceful shutdown
	go func() {
		if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	fmt.Printf("🚀 Server started on port %s\n", port)
	fmt.Println("📚 API Documentation: http://localhost:" + port + "/swagger")
	fmt.Println("🔍 Health check: http://localhost:" + port + "/health")

	// Ждем сигнал для graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("🛑 Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	fmt.Println("✅ Server stopped")
}

func getDatabaseURL() string {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "Kekos228@")
	dbname := getEnv("DB_NAME", "tasks_db")
	sslmode := getEnv("DB_SSLMODE", "disable")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, password, host, port, dbname, sslmode)
}

func getPort() string {
	return getEnv("PORT", "8080")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
