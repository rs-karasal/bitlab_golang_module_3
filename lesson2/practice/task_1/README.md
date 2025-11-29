# Задача 1: Hello, Query! (разогрев)

## Цель

Закрепить работу с:
- query-параметрами (`r.URL.Query()`);
- простейшим текстовым ответом от сервера;
- базовым хендлером и роутером `http.ServeMux`.

## Условие

Реализуй HTTP-сервер, который обрабатывает запросы к одному эндпоинту:

- `GET /hello`

Поведение хендлера:

1. Сервер читает query-параметр `name`.
2. Если `name` **не передан** или пустой — вернуть строку:
   
   ```text
   Hello, Bitlab!
   ```

3. Если `name` передан (например, `name=Alice`) — вернуть строку:

   ```text
   Hello, Alice!
   ```

4. Статус-код всегда `200 OK`.
5. Тип контента можно оставить по умолчанию (Go сам поставит `text/plain; charset=utf-8`).

## Примеры запросов

### 1. Без параметра `name`

**Запрос:**

```http
GET /hello HTTP/1.1
Host: localhost:8080
```

**Ответ:**

```text
Hello, Bitlab!
```

### 2. С параметром `name`

**Запрос:**

```http
GET /hello?name=Bob HTTP/1.1
Host: localhost:8080
```

**Ответ:**

```text
Hello, Bob!
```

## Требования к реализации

1. Использовать стандартный роутер:
   
   ```go
   mux := http.NewServeMux()
   ```

2. Зарегистрировать хендлер по пути `/hello`:
   
   ```go
   mux.HandleFunc("/hello", helloHandler)
   ```

3. В функции `helloHandler`:
   - получить значение `name` через `r.URL.Query().Get("name")`;
   - если `name` пустой — заменить его на `"Bitlab"`;
   - сформировать ответ строкой `"Hello, " + name + "!"`;
   - отправить ответ через `fmt.Fprint` или `fmt.Fprintf`.

## Шаблон кода

```go
package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: прочитать query-параметр name и вернуть корректный ответ
}

func main() {
    mux := http.NewServeMux()

    // TODO: зарегистрировать хендлер /hello
    // mux.HandleFunc("/hello", helloHandler)

    http.ListenAndServe(":8080", mux)
}
```

## Как проверить решение

1. Запусти сервер:

   ```bash
   go run ./...
   ```

2. Открой в браузере:

   - `http://localhost:8080/hello`
   - `http://localhost:8080/hello?name=Alice`

3. Или используй `curl`:

   ```bash
   curl "http://localhost:8080/hello"
   curl "http://localhost:8080/hello?name=Alice"
   ```