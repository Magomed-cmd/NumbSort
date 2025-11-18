# Numeric Sorter Service

Простой REST сервис на Go + Gin с разделением на handler/service/repository и хранилищем в PostgreSQL. Принимает число, сохраняет его и возвращает отсортированный список всех сохранённых чисел.

## Запуск локально

```bash
cp .env.example .env
go run ./cmd/server
```

Базу можно поднять через `docker-compose` (см. ниже). При старте сервис сам создаст таблицу `numbers`. В docker-compose БД слушает на хостовом порту `5433`, чтобы не конфликтовать с локальной установкой.

## Docker / Docker Compose

```bash
docker-compose up --build
```

- сервис: `http://localhost:8080`
- POST `/numbers` с телом `{"value": 3}` → ответ `{"numbers":[3]}` и т.д.
- GET `/health` для проверки статуса

## Тесты

```bash
go test ./...
```

## Структура

- `cmd/server` — точка входа
- `internal/handler` — HTTP слой (Gin)
- `internal/service` — бизнес-логика
- `internal/repository` — pgx доступ к БД + schema init
- `internal/routes` — маршрутизация и HTTP сервер
- `internal/config` — конфигурация из env
# numbsort
# numbsort
# NumbSort
