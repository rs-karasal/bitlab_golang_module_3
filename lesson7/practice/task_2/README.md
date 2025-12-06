# Задача 2: Репозиторий `ProductRepository` на PostgreSQL

## Цель

Закрепить работу с PostgreSQL из Go и научиться переносить хранилище данных:

- с in-memory реализации (карты/слайсы);
- на реальный репозиторий поверх PostgreSQL.

После выполнения задачи ты:

- опишешь интерфейс репозитория для продуктов;
- реализуешь `ProductRepository` на основе `*sql.DB` и SQL-запросов;
- заменишь in-memory репозиторий на Postgres-реализацию в `main.go`;
- добавишь валидацию данных в сервисном слое;
- проверишь, что поведение сервиса осталось прежним, а данные хранятся в БД.

---

## Общая идея

В предыдущих занятиях мы уже делали:

- сущность `Product` (товар);
- сервис, который работает с продуктами;
- in-memory репозиторий (на `map` или `[]Product`).

Теперь нужно:

1. Подключиться к базе PostgreSQL (таблица `products` уже должна быть создана, либо ты создашь её по схеме).
2. Реализовать репозиторий `ProductRepository`, который вместо памяти использует таблицу `products`.
3. В `main.go` подставить новый репозиторий и убедиться, что API продолжает работать, но теперь данные живут в БД.

---

## Предполагаемая схема таблицы `products`

Используем такую схему (если у тебя другая — адаптируй код под неё):

```sql
CREATE TABLE products (
    id       SERIAL PRIMARY KEY,
    name     TEXT NOT NULL,
    price    NUMERIC(10, 2) NOT NULL,
    quantity INT NOT NULL
);
```

Поля в Go-сущности могут выглядеть так:

```go
// internal/entities/product.go
package entities

import "errors"

type Product struct {
    ID       int
    Name     string
    Price    float64
    Quantity int
}

// Validate проверяет корректность данных продукта
func (p Product) Validate() error {
    if p.Name == "" {
        return errors.New("name is required")
    }
    if p.Price <= 0 {
        return errors.New("price must be positive")
    }
    if p.Quantity < 0 {
        return errors.New("quantity cannot be negative")
    }
    return nil
}
```

> Обрати внимание: правила вроде `price > 0` или `quantity >= 0` мы здесь проверяем в коде (через метод `Validate()`), а не в схеме БД. В реальном проекте можно дублировать важные ограничения ещё и на уровне базы (через `CHECK` или другие механизмы), но в этой задаче нам важно потренироваться именно с валидацией в Go.

---

## Структура проекта (пример)

```text
lesson7/
  practice/
    task_2/
      go.mod
      cmd/
        api/
          main.go                 // точка входа, здесь выбираем нужный репозиторий
      internal/
        entities/
          product.go              // сущность Product
        repository/
          product_repo.go         // интерфейс репозитория (если нужен)
          product_repo_pg.go      // реализация на PostgreSQL
          product_repo_memory.go  // (опционально) старая in-memory реализация
        service/
          product_service.go      // бизнес-логика
        handlers/
          product_handler.go      // HTTP-хендлеры для /products
```

Точная структура может отличаться, но важно, чтобы был **слой repository**, который использует `*sql.DB`.

---

## Часть 1. Интерфейс `ProductRepository`

**Файл (пример):** `internal/service/product_service.go` (или отдельный файл в слое `service`)

Опиши интерфейс репозитория, который будет использовать сервис:

```go
package service

import "bitlab_golang_module_3/lesson7/practice/task_2/internal/entities"

type ProductRepository interface {
    Create(p entities.Product) (entities.Product, error)
    GetByID(id int) (entities.Product, error)
    List() ([]entities.Product, error)
}
```

> 💡 Интерфейс `ProductRepository` достаточно описать **один раз** в слое `service` и использовать его и в сервисе, и в реализациях репозитория. Не нужно дублировать определение интерфейса в других пакетах.

---

## Часть 2. Реализация на PostgreSQL

**Файл:** `internal/repository/product_repo_pg.go`

### 2.1. Структура репозитория

Создай структуру, которая хранит `*sql.DB`:

```go
package repository

import (
    "database/sql"

    "bitlab_golang_module_3/lesson7/practice/task_2/internal/entities"
)

type PostgresProductRepository struct {
    db *sql.DB
}

func NewPostgresProductRepository(db *sql.DB) *PostgresProductRepository {
    return &PostgresProductRepository{db: db}
}
```

Эта структура должна реализовать интерфейс `ProductRepository`.

### 2.2. Метод `Create`

Реализуй метод, который:

- вставляет продукт в таблицу `products`;
- использует `INSERT ... RETURNING id` для получения нового `id`;
- возвращает заполненную сущность `Product` с присвоенным `ID`.

Пример заготовки:

```go
func (r *PostgresProductRepository) Create(p entities.Product) (entities.Product, error) {
    const query = `
        INSERT INTO products (name, price, quantity)
        VALUES ($1, $2, $3)
        RETURNING id;
    `

    // TODO: выполнить запрос r.db.QueryRow(...) и считать возвращённый id в p.ID

    return p, nil
}
```

### 2.3. Метод `GetByID`

Реализуй метод, который:

- выполняет `SELECT` по `id`;
- возвращает найденный продукт;
- если продукт не найден — возвращает понятную ошибку (например, `sql.ErrNoRows` обернуть в свою ошибку).

Заготовка:

```go
func (r *PostgresProductRepository) GetByID(id int) (entities.Product, error) {
    const query = `
        SELECT id, name, price, quantity
        FROM products
        WHERE id = $1;
    `

    var p entities.Product

    // TODO: выполнить r.db.QueryRow(...) и считать результат в p

    // TODO: обработать случай, когда продукт не найден (sql.ErrNoRows)

    return p, nil
}
```

### 2.4. Метод `List`

Реализуй метод, который:

- выбирает все продукты (например, с `ORDER BY id`);
- обходит `rows.Next()` и собирает слайс `[]entities.Product`.

Заготовка:

```go
func (r *PostgresProductRepository) List() ([]entities.Product, error) {
    const query = `
        SELECT id, name, price, quantity
        FROM products
        ORDER BY id;
    `

    // TODO: выполнить r.db.Query(...)
    // TODO: не забыть rows.Close()
    // TODO: обойти все строки и добавить продукты в слайс

    var products []entities.Product

    return products, nil
}
```

Не забудь проверить `rows.Err()` после цикла.

---

## Часть 3. Сервисный слой с валидацией

**Файл:** `internal/service/product_service.go`

```go
package service

import "bitlab_golang_module_3/lesson7/practice/task_2/internal/entities"

// Интерфейс ProductRepository описан выше в этом же пакете.

type ProductService struct {
    repo ProductRepository
}

func NewProductService(r ProductRepository) *ProductService {
    return &ProductService{repo: r}
}

func (s *ProductService) CreateProduct(p entities.Product) (entities.Product, error) {
    // TODO: вызвать p.Validate()
    // TODO: если Validate вернёт ошибку — вернуть её из метода
    // TODO: если всё ок — вызвать s.repo.Create(p) и вернуть результат
    return entities.Product{}, nil
}

func (s *ProductService) GetProductByID(id int) (entities.Product, error) {
    return s.repo.GetByID(id)
}

func (s *ProductService) ListProducts() ([]entities.Product, error) {
    return s.repo.List()
}
```

> Подсказка: реализация `CreateProduct` должна сначала вызвать `p.Validate()`, а затем, если ошибок нет, вызвать `s.repo.Create(p)`.

---

## Часть 4. Подключение к PostgreSQL в `main.go`

**Файл:** `cmd/api/main.go`

В `main.go` тебе нужно:

1. Прочитать строку подключения к БД (DSN) — из конфига или переменных окружения.
2. Инициализировать `*sql.DB` через `sql.Open` и проверить соединение `db.Ping()`.
3. Создать `PostgresProductRepository` и передать его в сервис.

Пример заготовки:

```go
package main

import (
    "database/sql"
    "log"
    "net/http"
    "os"

    _ "github.com/jackc/pgx/v5/stdlib"

    "bitlab_golang_module_3/lesson7/practice/task_2/internal/handlers"
    "bitlab_golang_module_3/lesson7/practice/task_2/internal/repository"
    "bitlab_golang_module_3/lesson7/practice/task_2/internal/service"
)

func main() {
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        // TODO: либо задать значение по умолчанию, либо завершить с ошибкой
    }

    // Пример значения переменной окружения `DATABASE_URL`:

    /*
    export DATABASE_URL="postgres://postgres:postgres@localhost:5432/myshop?sslmode=disable"
    */

    db, err := sql.Open("pgx", dsn)
    if err != nil {
        log.Fatalf("failed to open db: %v", err)
    }
    defer db.Close()

    if err := db.Ping(); err != nil {
        log.Fatalf("failed to ping db: %v", err)
    }

    // Инициализируем репозиторий на PostgreSQL
    productRepo := repository.NewPostgresProductRepository(db)

    // Создаём сервис
    productService := service.NewProductService(productRepo)

    // Создаём хендлер
    productHandler := handlers.NewProductHandler(productService)

    mux := http.NewServeMux()
    mux.HandleFunc("/products", productHandler.HandleProducts)
    mux.HandleFunc("/products/", productHandler.HandleProduct) // для GET /products/{id}

    log.Println("starting server on :8080")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatal(err)
    }
}
```

> 💡 Если у тебя уже есть in-memory реализация, можешь временно держать обе и переключаться между ними в зависимости от конфигурации. Для этой задачи достаточно просто заменить in-memory на PostgreSQL-репозиторий.

---

## Часть 5. Заменить in-memory репозиторий на PostgreSQL

Если раньше в `main.go` было что-то вроде:

```go
productRepo := repository.NewInMemoryProductRepository()
```

то теперь нужно использовать:

```go
productRepo := repository.NewPostgresProductRepository(db)
```

Убедись, что:

- сервис и хендлеры **не изменились** — они по-прежнему зависят от интерфейса, а не от конкретной реализации;
- меняется только то место, где «склеивается» приложение (слой `main` / `cmd/api`).

---

## Как проверить решение

1. Убедись, что база и таблица `products` созданы:

   ```sql
   CREATE TABLE products (
       id       SERIAL PRIMARY KEY,
       name     TEXT NOT NULL,
       price    NUMERIC(10, 2) NOT NULL,
       quantity INT NOT NULL
   );
   ```

2. Запусти сервер:

   ```bash
   cd bitlab_golang_module_3/lesson7/practice/task_2
   go run ./cmd/api
   ```

3. Отправь запрос на создание товара:

   ```bash
   curl -X POST "http://localhost:8080/products" \
        -H "Content-Type: application/json" \
        -d '{"name":"Laptop","price":1200.5,"quantity":10}'
   ```

4. Проверь, что:

   - в ответе приходит JSON с `id`, `name`, `price`, `quantity`;
   - статус — `201 Created`;
   - если сделать:

     ```bash
     curl "http://localhost:8080/products"
     ```

     — ты видишь список продуктов, включая только что созданный.

5. Получи продукт по ID:

   ```bash
   curl "http://localhost:8080/products/1"
   ```

   Проверь, что возвращается JSON с данными продукта.

6. Проверь валидацию — отправь некорректные данные:

   ```bash
   curl -X POST "http://localhost:8080/products" \
        -H "Content-Type: application/json" \
        -d '{"name":"","price":0,"quantity":10}'
   ```

   Должна вернуться ошибка валидации.

7. Открой базу (через `psql` или GUI) и убедись, что запись реально появилась в таблице `products`.

8. ⭐ По желанию: добавь метод `Delete(id int) error` в интерфейс и реализуй его в репозитории, а затем повесь на него эндпоинт `DELETE /products/{id}`.

Если всё работает, и твой сервис без изменения хендлеров/сервисов начал использовать PostgreSQL — ты успешно перенёс репозиторий на БД 💪
