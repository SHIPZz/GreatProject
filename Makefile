.PHONY: generate build run test clean docker-up docker-down

# Генерация кода
generate:
	@echo "🔄 Generating OpenAPI Go code..."
	./scripts/generate.sh
	@echo "🔄 Generating SQLC Go code..."
	./scripts/generate-sqlc.sh

# Сборка приложения
build:
	@echo "🔨 Building application..."
	go build -o bin/server cmd/server/main.go

# Запуск сервера
run:
	@echo "🚀 Starting server..."
	go run cmd/server/main.go

# Тесты
test:
	@echo "🧪 Running tests..."
	go test -v ./...

# Очистка
clean:
	@echo "🧹 Cleaning up..."
	go clean
	rm -rf bin/

# Docker Compose
docker-up:
	@echo "🐳 Starting Docker containers..."
	docker-compose up -d

docker-down:
	@echo "🛑 Stopping Docker containers..."
	docker-compose down

# Полная настройка
setup: generate build
	@echo "✅ Setup complete!"

# Разработка
dev: docker-up
	@echo "🚀 Starting development server..."
	go run cmd/server/main.go