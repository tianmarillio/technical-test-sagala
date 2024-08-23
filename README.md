# Technical Test Sagala

## Setup

> Go & PostgreSQL are required on your system

- Copy .env file from .env.example manually or run command:

```
cp .env.example .env
```

- Fill out PORT and database configs

- Install dependencies

```
go mod tidy
```

> For simplicity purpose, built-in GORM auto migrate is used for database migration. In production or real cases, it's recommended to use SQL or migration tools like golang-migrate or goose to avoid data loss and increase overall migration control

- Run development server

```
go run ./src/main.go
```

## Testing

> SQLite (in-memory database) is required on your system

Provided end-to-end (e2e) API testing by running command:

```
go test ./...
```

## API Documentation & Important Notes

- API Documentation:

- Valid "status" formats:
  - waiting_list (default)
  - in_progress
  - done

- Examples of valid "due_date" formats:
  - "2017"
  - "2017-10"
  - "1999-12-12 12"
  - "1999-12-12 12:20"
  - "1999-12-12 12:20:21"