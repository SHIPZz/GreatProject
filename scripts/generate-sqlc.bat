@echo off
REM Скрипт для генерации Go кода из SQL запросов (Windows)

echo 🗄️ Генерация Go кода из SQL запросов...

REM Проверяем, что sqlc установлен
where sqlc >nul 2>nul
if %errorlevel% neq 0 (
    echo ❌ sqlc не установлен
    echo Установи: go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
    pause
    exit /b 1
)

REM Проверяем конфигурацию
if not exist "sqlc.yaml" (
    echo ❌ Файл sqlc.yaml не найден
    pause
    exit /b 1
)

echo 📝 Проверяем SQL запросы...
sqlc compile

if %errorlevel% neq 0 (
    echo ❌ Ошибка в SQL запросах
    pause
    exit /b 1
)

echo 📝 Генерируем Go код...
sqlc generate

if %errorlevel% neq 0 (
    echo ❌ Ошибка генерации
    pause
    exit /b 1
)

echo ✅ Генерация завершена!
echo 📁 Сгенерированные файлы:
echo    - internal\database\db.go (интерфейсы и типы)
echo    - internal\database\models.go (модели данных)
echo    - internal\database\tasks.sql.go (функции для работы с БД)

echo.
echo 🚀 Следующие шаги:
echo 1. Обнови репозиторий для использования sqlc
echo 2. Замени sqlx на сгенерированные функции
echo 3. Запусти тесты: go test ./...

pause
