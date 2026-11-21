# Task Management REST API

A task manager built with **Go, Gin and PostgreSQL**. It exposes a REST API for tasks and serves a small static frontend.

## Features

- Create, read, update, partially update and delete tasks
- Request validation with Gin binding rules
- Layered structure: routes → handlers → repository → database
- Configuration through environment variables

## Tech stack

Go · Gin · PostgreSQL (`lib/pq`) · godotenv · HTML/CSS/JavaScript

## API

| Method | Path | Description |
|---|---|---|
| `GET` | `/` | Frontend |
| `GET` | `/tasks` | List tasks |
| `GET` | `/tasks/:id` | Get one task |
| `POST` | `/tasks` | Create a task |
| `PUT` | `/tasks/:id` | Replace a task |
| `PATCH` | `/tasks/:id` | Update selected fields |
| `DELETE` | `/tasks/:id` | Delete a task |

Example request body for `POST /tasks`:

```json
{
  "title": "Write documentation",
  "description": "Add a README",
  "priority": 2
}
```

`title` needs at least 3 characters and `priority` is required. `description` is optional. `PATCH` accepts a priority from 1 to 5.

## Project layout

```text
Todo list/
├── main.go
├── database/     PostgreSQL connection
├── handlers/     HTTP handlers
├── models/       Task and request models
├── repository/   SQL queries
├── routes/       Route definitions
└── public/       Static frontend
```

## Run locally

Requires Go 1.25+ and PostgreSQL.

1. Create the table. The repository has no migration, so this schema is written to match its queries:

   ```sql
   CREATE TABLE tasks (
       id          SERIAL PRIMARY KEY,
       title       TEXT NOT NULL,
       description TEXT,
       completed   BOOLEAN NOT NULL DEFAULT false,
       priority    INT NOT NULL
   );
   ```

2. Create `Todo list/.env`:

   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_user
   DB_PASSWORD=your_password
   DB_NAME=your_database
   ```

3. Start the server on port 8080:

   ```bash
   cd "Todo list"
   go run .
   ```
