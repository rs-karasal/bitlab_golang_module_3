# Задача 2: Несколько эндпоинтов на одном сервере

## Цель

Закрепить базовые навыки работы с HTTP-сервером в Go:

- запуск одного сервера с несколькими ручками (эндпоинтами);
- регистрацию разных путей через `http.HandleFunc` / `http.NewServeMux`;
- проверку ответов через браузер и `curl`.

## Условие

Реализуй HTTP-сервер, который обрабатывает **два разных URL**:

1. `GET /hello` — должен возвращать текст:

   ```text
   Hello!
   ```

2. `GET /goodbye` — должен возвращать текст:

   ```text
   Goodbye!
   ```

Сервер должен слушать порт `8080`.

## Требования к реализации

1. Использовать пакет `net/http`.
2. Создать роутер с помощью `http.NewServeMux()`.
3. Зарегистрировать два хендлера на одном сервере:
   - путь `/hello`;
   - путь `/goodbye`.
4. В хендлере для `/hello` вернуть текст `Hello!` с помощью `fmt.Fprint` или `fmt.Fprintf`.
5. В хендлере для `/goodbye` вернуть текст `Goodbye!`.
6. Запустить сервер на порту `8080` с помощью `http.ListenAndServe(":8080", mux)`.

## Шаблон кода

Создай файл `main.go` в папке `lesson1/practice/task_2` и используй следующий шаблон:

```go
package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: вернуть "Hello!"
}

func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: вернуть "Goodbye!"
}

func main() {
    mux := http.NewServeMux()

    // TODO: зарегистрировать хендлеры
    // mux.HandleFunc("/hello", helloHandler)
    // mux.HandleFunc("/goodbye", goodbyeHandler)

    // TODO: запустить сервер на порту 8080
    // http.ListenAndServe(":8080", mux)
}
```

## Как проверить решение

1. Запусти сервер из папки `lesson1/practice/task_2`:

   ```bash
   go run main.go
   ```

2. Открой эндпоинты в браузере:

   - `http://localhost:8080/hello`
   - `http://localhost:8080/goodbye`

3. Или используй `curl`:

   ```bash
   curl "http://localhost:8080/hello"
   curl "http://localhost:8080/goodbye"
   ```

4. Убедись, что ответы полностью совпадают с условием задачи.
