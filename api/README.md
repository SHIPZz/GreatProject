# OpenAPI 3 Specification

Этот каталог содержит OpenAPI 3 спецификацию для Todo List API.

## 📁 Файлы

- `openapi.yml` - Основная спецификация API в формате OpenAPI 3.0.0
- `README.md` - Документация по использованию

## 🚀 Быстрый старт

### 1. Просмотр документации

Открой файл `openapi.yml` в любом редакторе с поддержкой YAML или используй онлайн-редакторы:

- **Swagger Editor**: https://editor.swagger.io/
- **Redoc**: https://redocly.github.io/redoc/

### 2. Генерация Go кода

```bash
# Установи oapi-codegen
go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest

# Сгенерируй код
make generate
```

### 3. Локальный сервер документации

```bash
# Установи swagger-ui
npm install -g swagger-ui-serve

# Запусти локальный сервер
swagger-ui-serve api/openapi.yml
```

Открой http://localhost:3000 для просмотра интерактивной документации.

## 📋 Описание API

### Endpoints

| Метод | Путь | Описание |
|-------|------|----------|
| GET | `/tasks` | Получить все задачи |
| POST | `/tasks` | Создать новую задачу |
| GET | `/tasks/{id}` | Получить задачу по ID |
| PUT | `/tasks/{id}` | Обновить задачу |
| DELETE | `/tasks/{id}` | Удалить задачу |
| PATCH | `/tasks/{id}/complete` | Отметить задачу выполненной |
| GET | `/tasks/completed` | Получить выполненные задачи |
| GET | `/tasks/pending` | Получить невыполненные задачи |
| GET | `/health` | Проверка здоровья сервиса |

### Модели данных

#### Task
```yaml
Task:
  type: object
  properties:
    id: integer          # Уникальный ID
    name: string         # Название (1-255 символов)
    description: string  # Описание (до 1000 символов)
    completed: boolean   # Статус выполнения
    created_at: string   # Дата создания (ISO 8601)
    updated_at: string   # Дата обновления (ISO 8601)
```

#### CreateTaskRequest
```yaml
CreateTaskRequest:
  type: object
  properties:
    name: string         # Название (обязательно)
    description: string  # Описание (обязательно)
```

#### UpdateTaskRequest
```yaml
UpdateTaskRequest:
  type: object
  properties:
    name: string         # Название (обязательно)
    description: string  # Описание (обязательно)
    completed: boolean   # Статус (обязательно)
```

## 🔧 Генерация кода

### Типы данных
```bash
oapi-codegen -package generated -generate types api/openapi.yml > internal/generated/types.go
```

### Интерфейсы сервера
```bash
oapi-codegen -package generated -generate server api/openapi.yml > internal/generated/server.go
```

### HTTP клиент
```bash
oapi-codegen -package generated -generate client api/openapi.yml > internal/generated/client.go
```

## 📝 Примеры использования

### Создание задачи
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Купить молоко",
    "description": "В магазине на углу"
  }'
```

### Получение всех задач
```bash
curl http://localhost:8080/tasks
```

### Отметка задачи выполненной
```bash
curl -X PATCH http://localhost:8080/tasks/1/complete
```

## 🧪 Тестирование API

### Swagger UI
После запуска сервера открой http://localhost:8080/docs для интерактивного тестирования.

### Postman
Импортируй `openapi.yml` в Postman для автоматической генерации коллекции запросов.

### Insomnia
Импортируй `openapi.yml` в Insomnia для тестирования API.

## 🔄 Обновление спецификации

1. Отредактируй `openapi.yml`
2. Проверь синтаксис в Swagger Editor
3. Сгенерируй новый код: `make generate`
4. Обнови реализацию handlers
5. Запусти тесты: `make test`

## 📚 Полезные ссылки

- [OpenAPI Specification](https://swagger.io/specification/)
- [oapi-codegen](https://github.com/deepmap/oapi-codegen)
- [Swagger Editor](https://editor.swagger.io/)
- [OpenAPI Generator](https://openapi-generator.tech/)
