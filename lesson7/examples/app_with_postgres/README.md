# App with PostgreSQL

REST API приложение для управления продуктами с использованием PostgreSQL.

## Требования

- Go 1.21+
- PostgreSQL 14+
- Docker (опционально)

## Настройка базы данных

### Вариант 1: Используя Docker

```bash
docker run --name myshop-postgres \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=myshop \
  -p 5432:5432 \
  -d postgres:16
```

### Вариант 2: Локальная установка PostgreSQL

Подключитесь к PostgreSQL и создайте базу данных:

```bash
psql -U user -d myshop
```

## Создание таблиц

Выполните следующий SQL-скрипт для создания таблицы `products`:

```sql
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    quantity INTEGER NOT NULL DEFAULT 0
);
```

### Опционально: добавление тестовых данных

```sql
INSERT INTO products (name, price, quantity) VALUES
    ('Laptop', 999.99, 10),
    ('Mouse', 29.99, 50),
    ('Keyboard', 79.99, 30);
```

## Запуск приложения

```bash
go run cmd/app/main.go
```

## API Endpoints

| Метод | URL | Описание |
|-------|-----|----------|
| GET | /products | Получить список всех продуктов |
| POST | /products | Создать новый продукт |
| GET | /products/{id} | Получить продукт по ID |

### Примеры запросов

**Создание продукта:**

```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{"name": "Phone", "price": 599.99, "quantity": 25}'
```

**Получение списка продуктов:**

```bash
curl http://localhost:8080/products
```

**Получение продукта по ID:**

```bash
curl http://localhost:8080/products/1
```
