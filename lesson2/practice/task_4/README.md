# Задача 4: Создать пользователя (JSON в теле запроса)

## Цель

Попрактиковаться с:

- чтением тела запроса `r.Body`;
- разбором JSON через пакет `encoding/json`;
- базовой валидацией данных;
- выставлением корректных статус-кодов (`201`, `400`, `405`).

## Условие

Реализуй HTTP-эндпоинт:

- `POST /users`

Тело запроса (JSON):

```json
{
  "name": "Alice",
  "email": "alice@example.com"
}
```

Хендлер должен:

1. Проверить HTTP-метод.
2. Прочитать и распарсить JSON из тела запроса в Go-структуру.
3. Проверить, что обязательные поля не пустые.
4. Вернуть корректный статус-код и текстовый ответ.

### Требуемое поведение

1. Если метод запроса **не POST** — вернуть:
   - статус: `405 Method Not Allowed`;
   - тело ответа (например):

   ```text
   method not allowed
   ```

2. Декодировать тело запроса в структуру:

   ```go
   type CreateUserRequest struct {
       Name  string `json:"name"`
       Email string `json:"email"`
   }
   ```

3. Если JSON **невалидный** (ошибка при декодировании) — вернуть:
   - статус: `400 Bad Request`;
   - тело:

   ```text
   invalid json
   ```

4. Если поля `name` или `email` пустые — вернуть:
   - статус: `400 Bad Request`;
   - тело:

   ```text
   name and email are required
   ```

5. Если всё корректно — вернуть:
   - статус: `201 Created`;
   - заголовок `Content-Type: text/plain` (можно оставить дефолтный от Go);
   - тело:

   ```text
   user {name} with email {email} created
   ```

   Например:

   ```text
   user Alice with email alice@example.com created
   ```

## Пример запроса и ответа

**Запрос:**

```http
POST /users HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{"name": "Alice", "email": "alice@example.com"}
```

**Ответ:**

```http
HTTP/1.1 201 Created
Content-Type: text/plain; charset=utf-8

user Alice with email alice@example.com created
```

## Требования к реализации

1. Использовать стандартный роутер:

   ```go
   mux := http.NewServeMux()
   ```

2. Зарегистрировать хендлер по пути `/users`:

   ```go
   mux.HandleFunc("/users", createUserHandler)
   ```

3. В функции `createUserHandler`:

   - проверить метод запроса:

     ```go
     if r.Method != http.MethodPost {
         http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
         return
     }
     ```

   - создать переменную типа `CreateUserRequest`;
   - декодировать JSON из `r.Body` с помощью:

     ```go
     if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
         http.Error(w, "invalid json", http.StatusBadRequest)
         return
     }
     ```

   - проверить, что `req.Name` и `req.Email` не пустые;
   - при ошибке — вернуть `400 Bad Request` с текстом `name and email are required`;
   - при успехе — выставить статус `201 Created` и вывести текст:

     ```go
     w.WriteHeader(http.StatusCreated)
     fmt.Fprintf(w, "user %s with email %s created", req.Name, req.Email)
     ```

## Шаблон кода

```go
// lesson2/practice/task_4/main.go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type CreateUserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: реализовать всю логику из условия
}

func main() {
    mux := http.NewServeMux()

    // POST /users
    mux.HandleFunc("/users", createUserHandler)

    http.ListenAndServe(":8080", mux)
}
```

## Как проверить решение

1. Запусти сервер:

   ```bash
   go run ./...
   ```

2. Отправь корректный запрос:

   ```bash
   curl -X POST "http://localhost:8080/users" \
        -H "Content-Type: application/json" \
        -d '{"name":"Alice","email":"alice@example.com"}'
   ```

3. Попробуй отправить:

   - запрос с методом `GET` на `/users`;
   - некорректный JSON (например, обрезать кавычку);
   - JSON с пустым `name` или `email`.

4. Убедись, что статус-коды и тексты ответов совпадают с описанием задачи.