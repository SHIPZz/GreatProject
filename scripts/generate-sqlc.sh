#!/bin/bash

# Скрипт для генерации Go кода из SQL запросов

set -e

echo "🗄️ Генерация Go кода из SQL запросов..."

# Проверяем, что sqlc установлен
if ! command -v sqlc &> /dev/null; then
    echo "❌ sqlc не установлен"
    echo "Установи: go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest"
    exit 1
fi

# Проверяем конфигурацию
if [ ! -f "sqlc.yaml" ]; then
    echo "❌ Файл sqlc.yaml не найден"
    exit 1
fi

echo "📝 Проверяем SQL запросы..."
sqlc compile

echo "📝 Генерируем Go код..."
sqlc generate

echo "✅ Генерация завершена!"
echo "📁 Сгенерированные файлы:"
echo "   - internal/database/db.go (интерфейсы и типы)"
echo "   - internal/database/models.go (модели данных)"
echo "   - internal/database/tasks.sql.go (функции для работы с БД)"

echo ""
echo "🚀 Следующие шаги:"
echo "1. Обнови репозиторий для использования sqlc"
echo "2. Замени sqlx на сгенерированные функции"
echo "3. Запусти тесты: go test ./..."

