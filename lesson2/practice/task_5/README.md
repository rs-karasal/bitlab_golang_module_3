# Задача 5: Мини-API приветствий

## Цель
Закрепить работу с:
- роутером `http.ServeMux`;
- query-параметрами (`r.URL.Query()`);
- телом запроса (`r.Body`) и JSON;
- статус-кодами и заголовками ответа (`Content-Type`);
- обработкой базовых ошибок (`http.Error`).

## Условие
Реализуй HTTP-сервис с двумя ручками:

1. `GET /hello`
   - Принимает необязательный query-параметр `name`.
   - Если параметр `name` **не передан** или пустой — используй значение по умолчанию: `"Bitlab"`.
   - Возвращает JSON-ответ вида:

     ```json
     {
       "message": "Hello, Bitlab"
     }
     ```

   - Примеры:
     - `GET /hello` → `{ "message": "Hello, Bitlab" }`
     - `GET /hello?name=Alice` → `{ "message": "Hello, Alice" }`

2. `POST /hello`
   - Принимает JSON в теле запроса вида:

     ```json
     {
       "name": "Alice"
     }
     ```

   - Если:
     - тело невалидный JSON, или
     - поле `name` отсутствует, или
     - поле `name` пустое,

     то нужно вернуть:

     - статус `400 Bad Request`;
     - текст ошибки в теле (например, `"invalid json"` или `"name is required"`).

   - Если всё ок — вернуть:

     - статус `200 OK`;
     - заголовок `Content-Type: application/json`;
     - JSON-ответ вида:

     ```json
     {
       "message": "Hello, Alice"
     }
     ```

## Требования к реализации

1. Использовать стандартную библиотеку `net/http` и роутер `http.NewServeMux()`.
2. Разделить обработку по методам:
   - внутри одного пути `/hello` использовать `switch r.Method`;
   - для `GET` вызывать отдельную функцию-хендлер, например, `helloGetHandler`;
   - для `POST` вызывать `helloPostHandler`;
   - для остальных методов возвращать `405 Method Not Allowed`.
3. Для JSON-ответов:
   - обязательно установить заголовок `Content-Type: application/json`;
   - использовать `json.NewEncoder(w).Encode(...)` для кодирования структуры в JSON.
4. Для ошибок использовать функцию `http.Error`:
   - `http.Error(w, "invalid json", http.StatusBadRequest)`;
   - `http.Error(w, "name is required", http.StatusBadRequest)`;
   - `http.Error(w, "method not allowed", http.StatusMethodNotAllowed)`.

## Подсказка по структурам

Создай структуры для запроса и ответа:

```go
type HelloRequest struct {
    Name string `json:"name"`
}

type HelloResponse struct {
    Message string `json:"message"`
}
```

## Пример структуры кода

Файл `main.go` может выглядеть примерно так (структура, не окончательное решение):

```go
package main

import (
    "encoding/json"
    "net/http"
)

type HelloRequest struct {
    Name string `json:"name"`
}

type HelloResponse struct {
    Message string `json:"message"`
}

func helloGetHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "Bitlab"
    }

    resp := HelloResponse{Message: "Hello, " + name}

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    _ = json.NewEncoder(w).Encode(resp)
}

func helloPostHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req HelloRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "invalid json", http.StatusBadRequest)
        return
    }

    if req.Name == "" {
        http.Error(w, "name is required", http.StatusBadRequest)
        return
    }

    resp := HelloResponse{Message: "Hello, " + req.Name}

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    _ = json.NewEncoder(w).Encode(resp)
}

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            helloGetHandler(w, r)
        case http.MethodPost:
            helloPostHandler(w, r)
        default:
            http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        }
    })

    http.ListenAndServe(":8080", mux)
}
```

> ⚠️ Важно: этот код дан как пример структуры. Попробуй сначала реализовать задачу самостоятельно, а потом свериться с примером.

## Как проверить решение

1. Запусти сервер:

   ```bash
   go run ./...
   ```

2. Проверь `GET /hello` в браузере или через `curl`:

   ```bash
   curl "http://localhost:8080/hello"
   curl "http://localhost:8080/hello?name=Alice"
   ```

3. Проверь `POST /hello` через `curl` или Postman:

   ```bash
   curl -X POST "http://localhost:8080/hello" \
        -H "Content-Type: application/json" \
        -d '{"name": "Alice"}'
   ```

4. Попробуй отправить невалидный JSON или пустое имя и убедись, что сервер возвращает `400 Bad Request` и текст ошибки.