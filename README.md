# GreatProject - Todo List API

REST API для управления задачами с поддержкой PostgreSQL, написанный на Go.

## 🚀 Возможности

- ✅ **REST API** с полной OpenAPI 3 спецификацией
- ✅ **CRUD операции** для задач
- ✅ **PostgreSQL** с типобезопасными запросами (sqlc)
- ✅ **Docker & Docker Compose** для простого развертывания
- ✅ **Автогенерация кода** из OpenAPI спеки
- ✅ **Логирование** с разными уровнями
- ✅ **Валидация данных** и обработка ошибок
- ✅ **Юнит и интеграционные тесты**
- ✅ **Swagger UI** для интерактивной документации

## 📋 Требования

- **Go 1.21+**
- **PostgreSQL 13+** (или Docker)
- **Docker & Docker Compose** (для простого запуска)

## 🚀 Быстрый старт

### 1. Клонирование и установка
```bash
git clone https://github.com/ТвойUsername/GreatProject.git
cd GreatProject
make dev-setup  # Полная настройка для разработки
```

### 2. Запуск приложения
```bash
make run  # Запуск сервера
```

### 3. Просмотр API документации
Открой http://localhost:8080/docs для интерактивной документации Swagger UI.

## 🐳 Docker (рекомендуемый способ)

### Запуск с Docker Compose
```bash
# Запуск PostgreSQL
make docker-up

# Применение миграций
make migrate-up

# Запуск приложения
make run
```

### Остановка
```bash
make docker-down
```

## 📂 Структура проекта

```
GreatProject/
├── api/                    # OpenAPI спецификация
│   ├── openapi.yml        # Основная спецификация
│   └── README.md          # Документация API
├── internal/              # Внутренний код приложения
│   └── generated/         # Автогенерированный код
├── Business/              # Бизнес-логика
│   └── TaskService.go
├── Data/
│   ├── Entity/            # Модели данных
│   └── Repository/        # Репозитории
├── Infrastructure/
│   ├── Factory/           # Фабрики
│   ├── Logger/            # Логирование
│   └── Configs/           # Конфигурация
├── Web/                   # HTTP handlers (будущее)
├── scripts/               # Скрипты генерации
├── Tests/                 # Тесты
├── docker-compose.yml     # Docker конфигурация
├── Makefile              # Команды разработки
└── main.go               # Точка входа
```

## 🔧 Команды разработки

```bash
make help              # Показать все команды
make install           # Установить зависимости
make generate          # Сгенерировать код из OpenAPI
make build             # Собрать приложение
make test              # Запустить тесты
make test-coverage     # Тесты с покрытием
make run               # Запустить приложение
make docker-up         # Запустить PostgreSQL
make docker-down       # Остановить контейнеры
make clean             # Очистить временные файлы
```

## 📚 API Документация

### Основные endpoints

| Метод | Путь | Описание |
|-------|------|----------|
| GET | `/tasks` | Получить все задачи |
| POST | `/tasks` | Создать новую задачу |
| GET | `/tasks/{id}` | Получить задачу по ID |
| PUT | `/tasks/{id}` | Обновить задачу |
| DELETE | `/tasks/{id}` | Удалить задачу |
| PATCH | `/tasks/{id}/complete` | Отметить выполненной |
| GET | `/tasks/completed` | Выполненные задачи |
| GET | `/tasks/pending` | Невыполненные задачи |
| GET | `/health` | Проверка здоровья |

### Примеры запросов

#### Создание задачи
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Купить молоко",
    "description": "В магазине на углу"
  }'
```

#### Получение всех задач
```bash
curl http://localhost:8080/tasks
```

## 🧪 Тестирование

```bash
# Все тесты
make test

# Тесты с покрытием
make test-coverage

# Только unit тесты
go test ./Tests/...

# Только интеграционные тесты
go test ./Tests/... -tags=integration
```

## 📝 Конфигурация

### Переменные окружения

| Переменная | По умолчанию | Описание |
|------------|--------------|----------|
| `DB_HOST` | localhost | Хост PostgreSQL |
| `DB_PORT` | 5434 | Порт PostgreSQL |
| `DB_USER` | postgres | Пользователь БД |
| `DB_PASSWORD` | postgres | Пароль БД |
| `DB_NAME` | tasks_db | Имя базы данных |
| `SERVER_PORT` | 8080 | Порт HTTP сервера |
| `LOG_LEVEL` | INFO | Уровень логирования |

### Пример .env файла
```env
DB_HOST=localhost
DB_PORT=5434
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=tasks_db
SERVER_PORT=8080
LOG_LEVEL=DEBUG
```

## 🔄 Разработка

### 1. Изменение API
1. Отредактируй `api/openapi.yml`
2. Сгенерируй код: `make generate`
3. Обнови handlers в `internal/handlers/`
4. Запусти тесты: `make test`

### 2. Добавление новых endpoints
1. Добавь в `api/openapi.yml`
2. Сгенерируй код: `make generate`
3. Реализуй handler
4. Добавь тесты

### 3. Работа с БД
1. Добавь SQL запросы в `sql/queries/`
2. Сгенерируй код: `make generate-sqlc`
3. Используй в репозитории

## 📄 Лицензия

MIT License - см. [LICENSE](LICENSE) файл для деталей.

## 🤝 Вклад в проект

1. Форкни репозиторий
2. Создай ветку для фичи (`git checkout -b feature/amazing-feature`)
3. Сделай коммит (`git commit -m 'Add amazing feature'`)
4. Запушь в ветку (`git push origin feature/amazing-feature`)
5. Открой Pull Request
