package db

import (
	"GreatProject/internal/database"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

// New создает новый экземпляр Queries
func New(conn *pgx.Conn) *db.Queries {
	return db.New(conn)
}

// Connect подключается к PostgreSQL и создает схему
func Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), getDatabaseURL())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Проверяем подключение
	if err := conn.Ping(context.Background()); err != nil {
		conn.Close(context.Background())
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Создаем схему если нужно
	if err := createSchema(conn); err != nil {
		conn.Close(context.Background())
		return nil, fmt.Errorf("failed to create schema: %w", err)
	}

	return conn, nil
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

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func createSchema(conn *pgx.Conn) error {
	// Читаем SQL схему
	schemaSQL, err := os.ReadFile("sql/schema/001_initial.sql")
	if err != nil {
		return fmt.Errorf("failed to read schema file: %w", err)
	}

	// Выполняем SQL
	_, err = conn.Exec(context.Background(), string(schemaSQL))
	if err != nil {
		return fmt.Errorf("failed to execute schema: %w", err)
	}

	return nil
}
