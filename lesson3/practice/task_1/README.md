# Задача 1: Слоистое приложение: заметки (получить и добавить)

## Цель

Закрепить:

- базовую структуру проекта (`cmd`, `internal`);
- разделение кода на слои: `handlers → service → repository → entities`;
- умение обрабатывать **несколько операций** над сущностью (получить список и добавить запись);
- роль `main.go` как точки входа и «сборщика» приложения.

## Условие

Нужно создать слоистое приложение c одной сущностью `Note` (заметка) и двумя операциями:

1. **Получить список заметок**
   - `GET /notes`
   - В ответе сервер должен вернуть JSON-массив заметок:
     ```json
     [
       {
         "id": 1,
         "text": "First note"
       },
       {
         "id": 2,
         "text": "Second note"
       }
     ]
     ```
   - Статус-код: `200 OK`.
   - Если заметок пока нет — вернуть пустой массив `[]` и статус `200 OK`.

2. **Добавить новую заметку**
   - `POST /notes`
   - В теле запроса приходит JSON с текстом заметки:
     ```json
     {
       "text": "Buy milk"
     }
     ```
   - В ответе сервер должен вернуть JSON с созданной заметкой и сгенерированным `id`:
     ```json
     {
       "id": 1,
       "text": "Buy milk"
     }
     ```
   - Статус-код: `201 Created`.
   - Если поле `text` пустое — вернуть `400 Bad Request` и сообщение об ошибке.

Пока **без базы данных**: данные о заметках храним в in-memory репозитории на основе `map[int]Note`.

Приложение должно быть разложено по слоям:

- `cmd/api` — точка входа (`main.go`);
- `internal/entities` — доменная сущность `Note`;
- `internal/repository` — репозиторий заметок (in-memory);
- `internal/service` — сервис заметок (бизнес-логика);
- `internal/handlers` — HTTP-хендлеры (обработчики `/notes`).

## Требования к структуре проекта

Структура папок должна быть примерно такой:

```text
lesson3/
  practice/
    task_1/
      go.mod
      cmd/
        api/
          main.go
      internal/
        entities/
          note.go
        repository/
          note_repo.go
        service/
          note_service.go
        handlers/
          note_handler.go
```

Имена файлов могут отличаться, но важно выдержать идею слоёв.

## Требования к реализации

### 1. Сущность `Note` (entities)

В `internal/entities/note.go` опиши доменную сущность:

```go
package entities

type Note struct {
    ID   int    `json:"id"`
    Text string `json:"text"`
}
```

### 2. Репозиторий (repository)

В `internal/repository/note_repo.go` создай репозиторий заметок.

Ниже приведён шаблон для реализации. Это только каркас, студенту нужно самостоятельно дописать реализацию по условиям задачи.

```go
package repository

import "bitlab_golang_module_3/lesson3/practice/task_1/internal/entities"

type NoteRepository struct {
    // TODO: добавить необходимые поля (например, map[int]entities.Note и счётчик nextID)
}

func NewNoteRepository() *NoteRepository {
    // TODO: инициализировать структуру репозитория (карта заметок, начальное значение nextID)
    return &NoteRepository{}
}

func (r *NoteRepository) Create(text string) (entities.Note, error) {
    // TODO: создать новую заметку, присвоить ей ID, сохранить в хранилище и вернуть
    return entities.Note{}, nil
}

func (r *NoteRepository) List() []entities.Note {
    // TODO: вернуть все заметки из хранилища
    return nil
}
```

### 3. Сервис (service)

В `internal/service/note_service.go` опиши сервис, который использует репозиторий.

Ниже приведён шаблон для реализации. Это только каркас, студенту нужно самостоятельно дописать реализацию по условиям задачи.

```go
package service

import "bitlab_golang_module_3/lesson3/practice/task_1/internal/entities"

type NoteRepository interface {
    Create(text string) (entities.Note, error)
    List() []entities.Note
}

type NoteService struct {
    repo NoteRepository
}

func NewNoteService(r NoteRepository) *NoteService {
    // TODO: сохранить репозиторий в структуре сервиса
    return &NoteService{}
}

func (s *NoteService) CreateNote(text string) (entities.Note, error) {
    // TODO: добавить при необходимости валидацию текста и вызвать метод репозитория
    return entities.Note{}, nil
}

func (s *NoteService) ListNotes() []entities.Note {
    // TODO: вернуть список заметок из репозитория
    return nil
}
```

### 4. Хендлер (handlers)

В `internal/handlers/note_handler.go` опиши HTTP-слой.

Ниже приведён шаблон для реализации. Это только каркас, студенту нужно самостоятельно дописать реализацию по условиям задачи.

```go
package handlers

import (
    "net/http"

    "bitlab_golang_module_3/lesson3/practice/task_1/internal/entities"
)

type NoteService interface {
    CreateNote(text string) (entities.Note, error)
    ListNotes() []entities.Note
}

type NoteHandler struct {
    service NoteService
}

func NewNoteHandler(s NoteService) *NoteHandler {
    // TODO: сохранить сервис в структуре хендлера
    return &NoteHandler{}
}

type createNoteRequest struct {
    Text string `json:"text"`
}

func (h *NoteHandler) NotesHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        // TODO: реализовать обработку GET /notes (вызвать h.handleListNotes)
    case http.MethodPost:
        // TODO: реализовать обработку POST /notes (вызвать h.handleCreateNote)
    default:
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    }
}

// TODO: реализовать вспомогательные методы handleListNotes и handleCreateNote
```

### 5. Точка входа (cmd/api/main.go)

В `cmd/api/main.go` собери все слои.

Ниже приведён шаблон для реализации. Это только каркас, студенту нужно самостоятельно дописать реализацию по условиям задачи.

```go
package main

import (
    "log"
    "net/http"

    "bitlab_golang_module_3/lesson3/practice/task_1/internal/handlers"
    "bitlab_golang_module_3/lesson3/practice/task_1/internal/repository"
    "bitlab_golang_module_3/lesson3/practice/task_1/internal/service"
)

func main() {
    // TODO: создать репозиторий заметок
    noteRepo := repository.NewNoteRepository()

    // TODO: создать сервис заметок и передать в него репозиторий
    noteService := service.NewNoteService(noteRepo)

    // TODO: создать хендлер и передать в него сервис
    noteHandler := handlers.NewNoteHandler(noteService)

    // TODO: создать роутер и зарегистрировать путь "/notes"
    mux := http.NewServeMux()
    // mux.HandleFunc("/notes", noteHandler.NotesHandler)

    log.Println("starting server on :8080")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatal(err)
    }
}
```