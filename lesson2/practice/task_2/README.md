# Задача 2: Калькулятор в URL

## Цель

Потренировать:
- чтение query-параметров (`r.URL.Query()`);
- конвертацию строк в числа с помощью `strconv.Atoi`;
- обработку ошибок и возврат корректных статус-кодов (`400`, `200`).

## Условие

Реализуй HTTP-эндпоинт:

- `GET /sum?a=...&b=...`

Хендлер должен:

1. Считать query-параметры `a` и `b`.
2. Если хотя бы один параметр **не передан** или пустой:
   - вернуть статус `400 Bad Request`;
   - вернуть тело ответа:
     
     ```text
     a and b are required
     ```
3. Попробовать сконвертировать значения `a` и `b` в целые числа (`int`) с помощью `strconv.Atoi`.
   - Если конвертация **не удалась** (например, `a=abc`):
     - вернуть статус `400 Bad Request`;
     - тело ответа:

     ```text
     a and b must be integers
     ```

4. Если всё корректно — посчитать сумму `a + b` и вернуть:
   - статус `200 OK`;
   - тело ответа в формате:

     ```text
     sum = {результат}
     ```

   Например: `sum = 12`.

## Примеры запросов и ответов

### 1. Корректный запрос

**Запрос:**

```http
GET /sum?a=5&b=7 HTTP/1.1
Host: localhost:8080
```

**Ответ:**

```text
sum = 12
```

### 2. Не хватает параметра

**Запрос:**

```http
GET /sum?a=5 HTTP/1.1
Host: localhost:8080
```

**Ответ:**

- статус: `400 Bad Request`
- тело:

```text
a and b are required
```

### 3. Некорректный формат числа

**Запрос:**

```http
GET /sum?a=abc&b=2 HTTP/1.1
Host: localhost:8080
```

**Ответ:**

- статус: `400 Bad Request`
- тело:

```text
a and b must be integers
```

## Требования к реализации

1. Использовать стандартный роутер `http.NewServeMux()`.
2. Зарегистрировать хендлер по пути `/sum`:

   ```go
   mux.HandleFunc("/sum", sumHandler)
   ```

3. В функции `sumHandler`:
   - получить параметры `a` и `b` через `r.URL.Query().Get("a")` и `r.URL.Query().Get("b")`;
   - проверить, что оба параметра не пустые;
   - сконвертировать значения в `int` через `strconv.Atoi`;
   - при ошибках использовать `http.Error(w, ..., http.StatusBadRequest)`;
   - при успехе посчитать сумму и вывести результат через `fmt.Fprintf`.

## Шаблон кода

```go
package main

import (
    "fmt"
    "net/http"
    "strconv"
)

func sumHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: реализовать логику из условия
}

func main() {
    mux := http.NewServeMux()

    // TODO: зарегистрировать хендлер /sum
    // mux.HandleFunc("/sum", sumHandler)

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
   # Успешный запрос
   curl "http://localhost:8080/sum?a=5&b=7"

   # Не хватает параметра
   curl "http://localhost:8080/sum?a=5"

   # Некорректное число
   curl "http://localhost:8080/sum?a=abc&b=2"
   ```

3. Убедись, что статус-коды и тексты ответов совпадают с условием задачи.