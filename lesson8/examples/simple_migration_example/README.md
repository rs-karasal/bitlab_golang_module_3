# Goose Migration Example

Пример использования библиотеки [goose](https://github.com/pressly/goose) для управления миграциями базы данных в Go.

## Структура проекта

```
simple_migration_example/
├── cmd/
│   └── migrate/
│       └── main.go          # Точка входа для запуска миграций
├── migrations/
│   ├── 001_create_users_table.sql
│   ├── 002_create_products_table.sql
│   └── 003_add_description_to_products.sql
├── go.mod
└── README.md
```

## Требования

- Go 1.21+
- PostgreSQL 14+
- Docker (опционально)

## Настройка базы данных

### Вариант 1: Docker

```bash
docker run --name migration-example-postgres \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=migration_example \
  -p 5432:5432 \
  -d postgres:16
```

### Вариант 2: Локальный PostgreSQL

```bash
psql -U postgres -c "CREATE USER \"user\" WITH PASSWORD 'password';"
psql -U postgres -c "CREATE DATABASE migration_example OWNER \"user\";"
```

## Установка зависимостей

```bash
cd simple_migration_example
go mod tidy
```

## Запуск миграций

### Применить все миграции

```bash
go run ./cmd/migrate -command=up
```

Вывод:
```
OK   001_create_users_table.sql
OK   002_create_products_table.sql
OK   003_add_description_to_products.sql
goose: successfully migrated database to version: 3
```

### Применить одну миграцию

```bash
go run ./cmd/migrate -command=up-one
```

### Откатить последнюю миграцию

```bash
go run ./cmd/migrate -command=down
```

### Откатить все миграции

```bash
go run ./cmd/migrate -command=reset
```

### Посмотреть статус миграций

```bash
go run ./cmd/migrate -command=status
```

Вывод:
```
    Applied At                  Migration
    =======================================
    Wed Jan 10 12:00:00 2024 -- 001_create_users_table.sql
    Wed Jan 10 12:00:01 2024 -- 002_create_products_table.sql
    Wed Jan 10 12:00:02 2024 -- 003_add_description_to_products.sql
```

### Посмотреть текущую версию

```bash
go run ./cmd/migrate -command=version
```

## Использование переменной окружения

Можно задать строку подключения через `DATABASE_URL`:

```bash
export DATABASE_URL="postgres://user:password@localhost:5432/migration_example?sslmode=disable"
go run ./cmd/migrate -command=up
```

## Формат файлов миграций

Goose использует специальные комментарии для разделения UP и DOWN миграций:

```sql
-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS users;
```

### Правила именования

- Файлы должны начинаться с номера версии: `001_`, `002_`, и т.д.
- Расширение `.sql` для SQL-миграций
- Goose также поддерживает Go-миграции (`.go` файлы)

## Проверка результата

После применения миграций можно проверить таблицы в БД:

```bash
psql -U user -d migration_example -c "\dt"
```

Вывод:
```
              List of relations
 Schema |       Name        | Type  | Owner
--------+-------------------+-------+-------
 public | goose_db_version  | table | user
 public | products          | table | user
 public | users             | table | user
```

Таблица `goose_db_version` создаётся автоматически и хранит историю применённых миграций.

## Команды goose CLI (альтернатива)

Можно также использовать CLI-инструмент goose напрямую:

```bash
# Установка
go install github.com/pressly/goose/v3/cmd/goose@latest

# Использование
goose -dir migrations postgres "postgres://user:password@localhost:5432/migration_example?sslmode=disable" up
goose -dir migrations postgres "postgres://user:password@localhost:5432/migration_example?sslmode=disable" status
```
