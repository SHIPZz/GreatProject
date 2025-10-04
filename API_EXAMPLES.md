# Примеры использования API

## Базовый URL
```
http://localhost:8080
```

## 1. Health Check
```bash
curl http://localhost:8080/health
```

## 2. Получить все задачи
```bash
curl http://localhost:8080/tasks
```

## 3. Создать новую задачу
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"Купить хлеб\",\"description\":\"В ближайшем магазине\"}"
```

PowerShell:
```powershell
$body = @{ name = "Купить хлеб"; description = "В ближайшем магазине" } | ConvertTo-Json
Invoke-RestMethod -Uri "http://localhost:8080/tasks" -Method Post -Body $body -ContentType "application/json"
```

## 4. Получить задачу по ID
```bash
curl http://localhost:8080/tasks/1
```

## 5. Обновить задачу
```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"Купить молоко и хлеб\",\"description\":\"В супермаркете\",\"completed\":false}"
```

PowerShell:
```powershell
$body = @{ name = "Купить молоко и хлеб"; description = "В супермаркете"; completed = $false } | ConvertTo-Json
Invoke-RestMethod -Uri "http://localhost:8080/tasks/1" -Method Put -Body $body -ContentType "application/json"
```

## 6. Отметить задачу выполненной
```bash
curl -X PATCH http://localhost:8080/tasks/1/complete
```

PowerShell:
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/tasks/1/complete" -Method Patch
```

## 7. Получить только выполненные задачи
```bash
curl http://localhost:8080/tasks/completed
```

## 8. Получить только невыполненные задачи
```bash
curl http://localhost:8080/tasks/pending
```

## 9. Удалить задачу
```bash
curl -X DELETE http://localhost:8080/tasks/1
```

PowerShell:
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/tasks/1" -Method Delete
```

## Swagger UI (API Documentation)

Открой в браузере:
```
http://localhost:8080/swagger
```

## Postman Collection

Импортируй OpenAPI спецификацию в Postman:
```
File: api/openapi.yml
```

## Параметры запроса (Query Parameters)

### Получить задачи с фильтрами:
```bash
# Первые 10 задач
curl "http://localhost:8080/tasks?limit=10"

# Пропустить первые 5 и взять следующие 10
curl "http://localhost:8080/tasks?limit=10&offset=5"

# Только выполненные
curl "http://localhost:8080/tasks?completed=true"

# Только невыполненные
curl "http://localhost:8080/tasks?completed=false"
```

## Примеры ответов

### Успешное создание задачи (201):
```json
{
  "id": 1,
  "name": "Купить хлеб",
  "description": "В ближайшем магазине",
  "completed": false,
  "created_at": "2025-10-04T12:00:00Z",
  "updated_at": "2025-10-04T12:00:00Z"
}
```

### Список задач (200):
```json
{
  "tasks": [
    {
      "id": 1,
      "name": "Купить хлеб",
      "description": "В ближайшем магазине",
      "completed": false,
      "created_at": "2025-10-04T12:00:00Z",
      "updated_at": "2025-10-04T12:00:00Z"
    }
  ],
  "total": 1,
  "limit": 50,
  "offset": 0
}
```

### Ошибка валидации (400):
```json
{
  "error": "Validation failed",
  "message": "Name is required",
  "code": "VALIDATION_ERROR"
}
```

### Задача не найдена (404):
```json
{
  "error": "Not Found",
  "message": "Task with ID 999 not found",
  "code": "TASK_NOT_FOUND"
}
```

