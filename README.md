# Job Search Backend

Учебный backend-проект сервиса поиска работы, написанный на Go.

## Возможности

- Регистрация пользователей
- Авторизация (JWT)
- Создание вакансий
- Создание резюме
- Создание откликов
- Добавление вакансий в избранное
- PostgreSQL
- Docker Compose

## Структура проекта

```
cmd/
internal/
    auth/
    database/
    dto/
    handler/
    middleware/
    models/
    repository/
    routes/
    service/
```

## Используемые технологии

- Go
- PostgreSQL
- pgx
- JWT
- Docker

## Запуск проекта

```bash
go run ./cmd
```