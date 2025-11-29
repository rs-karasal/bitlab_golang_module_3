# Задача 3: Пользователь по ID (path-параметр)

## Цель

Потренировать:

- работу с путём запроса `r.URL.Path`;
- ручной разбор path-параметра (ID, зашитого в URL);
- обработку ошибок и возврат корректных статус-кодов.

## Условие

Реализуй HTTP-эндпоинт:

- `GET /users/{id}`

Где `{id}` — **положительное целое число**.

Хендлер должен:

1. Прочитать путь запроса: `r.URL.Path`.
2. Убедиться, что путь имеет формат `/users/{id}`.
3. Достать `id` из пути.
4. Попробовать сконвертировать `id` в число (`int`).
5. Выполнить проверки и вернуть корректный ответ.

### Требуемое поведение

1. Если путь **не** в формате `/users/{id}` (например, `/users`, `/users/10/extra`) — вернуть:
   - статус: `400 Bad Request`;
   - тело ответа:

   ```text
   invalid path
   ```

2. Если `id` нельзя сконвертировать в число (например, `/users/abc`) — вернуть:
   - статус: `400 Bad Request`;
   - тело ответа:

   ```text
   invalid user id
   ```

3. Если `id <= 0` (например, `/users/0` или `/users/-5`) — вернуть:
   - статус: `400 Bad Request`;
   - тело ответа:

   ```text
   id must be positive
   ```

4. Если всё корректно — вернуть:
   - статус: `200 OK`;
   - тело ответа:

   ```text
   user id = {id}
   ```

   Например, для запроса `/users/10` ответ должен быть:

   ```text
   user id = 10
   ```

## Примеры запросов и ответов

### 1. Корректный запрос

**Запрос:**

```http
GET /users/10 HTTP/1.1
Host: localhost:8080
```

**Ответ:**

```text
user id = 10
```

---

### 2. Неверный путь

**Запрос:**

```http
GET /users HTTP/1.1
Host: localhost:8080
```

**Ответ:**

- статус: `400 Bad Request`
- тело:

```text
invalid path
```

---

### 3. Некорректный id

**Запрос:**

```http
GET /users/abc HTTP/1.1
Host: localhost:8080
```

**Ответ:**

- статус: `400 Bad Request`
- тело:

```text
invalid user id
```

---

### 4. Неположительный id

**Запрос:**

```http
GET /users/0 HTTP/1.1
Host: localhost:8080
```

**Ответ:**

- статус: `400 Bad Request`
- тело:

```text
id must be positive
```

## Требования к реализации

1. Использовать стандартный роутер:

   ```go
   mux := http.NewServeMux()
   ```

2. Зарегистрировать хендлеры для путей `/users` и `/users/`:

   ```go
   mux.HandleFunc("/users", userHandler)
   mux.HandleFunc("/users/", userHandler)
   ```

   > Так путь `/users` тоже будет обрабатываться в `userHandler`, и ты сможешь вернуть `400 invalid path` согласно условию.

3. В функции `userHandler`:

   - прочитать путь: `path := r.URL.Path`;
   - разбить путь на части с помощью `strings.Split(path, "/")`;
   - проверить, что путь состоит из трёх частей: `["", "users", "{id}"]`;
   - достать строковый `id` как `parts[2]`;
   - сконвертировать `id` в `int` через `strconv.Atoi`;
   - проверить, что `id > 0`;
   - при ошибках использовать `http.Error` с нужным текстом и кодом:
     - `http.Error(w, "invalid path", http.StatusBadRequest)`
     - `http.Error(w, "invalid user id", http.StatusBadRequest)`
     - `http.Error(w, "id must be positive", http.StatusBadRequest)`
   - при успехе — вывести `user id = {id}` через `fmt.Fprintf`.

## Шаблон кода

```go
// lesson2/practice/task_3/main.go
package main

import (
    "fmt"
    "net/http"
    "strconv"
    "strings"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: разобрать путь, достать id, обработать ошибки
}

func main() {
    mux := http.NewServeMux()

    // Обрабатываем и /users, и /users/{id}
    mux.HandleFunc("/users", userHandler)
    mux.HandleFunc("/users/", userHandler)

    http.ListenAndServe(":8080", mux)
}
```

## Как проверить решение

1. Запусти сервер:

   ```bash
   go run ./...
   ```

2. Отправь несколько запросов через `curl`:

   ```bash
   # Корректный запрос
   curl "http://localhost:8080/users/10"

   # Неверный путь (нет id)
   curl "http://localhost:8080/users"

   # Некорректный id
   curl "http://localhost:8080/users/abc"

   # Неположительный id
   curl "http://localhost:8080/users/0"
   ```

3. Убедись, что статус-коды и тексты ответов совпадают с описанием задачи.