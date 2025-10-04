#!/bin/bash

# Скрипт для генерации Go кода из OpenAPI спеки

set -e

echo "🔧 Генерация Go кода из OpenAPI спеки..."

# Проверяем, что oapi-codegen установлен
if ! command -v oapi-codegen &> /dev/null; then
    echo "❌ oapi-codegen не установлен"
    echo "Установи: go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest"
    exit 1
fi

# Создаем папку для сгенерированного кода
mkdir -p internal/generated

echo "📝 Генерируем типы данных..."
oapi-codegen \
    -package generated \
    -generate types \
    api/openapi.yml > internal/generated/types.go

echo "📝 Генерируем интерфейсы сервера..."
oapi-codegen \
    -package generated \
    -generate server \
    api/openapi.yml > internal/generated/server.go

echo "📝 Генерируем клиент..."
oapi-codegen \
    -package generated \
    -generate client \
    api/openapi.yml > internal/generated/client.go

echo "✅ Генерация завершена!"
echo "📁 Сгенерированные файлы:"
echo "   - internal/generated/types.go (модели данных)"
echo "   - internal/generated/server.go (интерфейсы сервера)"
echo "   - internal/generated/client.go (HTTP клиент)"

echo ""
echo "🚀 Следующие шаги:"
echo "1. Реализуй интерфейсы в internal/handlers/"
echo "2. Настрой роутинг в main.go"
echo "3. Запусти сервер: go run main.go"
