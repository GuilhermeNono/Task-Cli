# ✅ Task CLI

A lightweight command-line Task Tracker built in Go.

This project is based on the challenge described in `exercise.md`: track tasks, persist them in a local JSON file, and manage task lifecycle from the terminal.

## 🚀 Goals

- Track what you need to do (`todo`)
- Track what is in progress (`in-progress`)
- Track what is finished (`done`)
- Persist everything in `tasks.json` in the current directory

## 🧩 Current Status

### Implemented

- `add` command to create a task
- `list` command to print tasks
- `version` command
- JSON persistence (`tasks.json`) using Go standard library

### Planned (from challenge requirements)

- `update <id> <description>`
- `delete <id>`
- `mark-in-progress <id>`
- `mark-done <id>`
- `list done`
- `list todo`
- `list in-progress`

## 🛠️ Tech Stack

- Go
- `cobra` for CLI command structure
- Native filesystem APIs (`os`, `encoding/json`) for persistence

## 📦 Project Structure

```text
task-cli/
├── cmd/            # CLI commands
├── model/          # Task model and sorting helpers
├── main.go         # Load tasks, execute CLI, save tasks
└── tasks.json      # Auto-created data file
```

## ▶️ Running the project

You can run directly with Go:

```bash
go run . version
go run . add "Buy groceries"
go run . list
```

Or build and run the binary:

```bash
go build -o task-cli
./task-cli version
./task-cli add "Buy groceries"
./task-cli list
```

## 🧪 Command Examples

```bash
# Add a new task
task-cli add "Buy groceries"

# List all tasks
task-cli list

# Show CLI version
task-cli version
```

> Note: In the current implementation, `list` prints the in-memory slice structure. Output formatting can be improved in next iterations.

## 🗂️ Task Data Model

Each task follows this schema:

- `id`: unique identifier
- `description`: short task text
- `status`: `todo` | `in-progress` | `done`
- `createdAt`: creation timestamp
- `updatedAt`: last update timestamp

Example `tasks.json` entry:

```json
[
  {
    "id": 1,
    "description": "Buy groceries",
    "status": "todo",
    "createdAt": "2026-04-01T10:00:00Z",
    "updatedAt": "2026-04-01T10:00:00Z"
  }
]
```

## ✅ Challenge Checklist

- [x] CLI entry point
- [x] Add task
- [x] Basic list task
- [x] JSON file persistence
- [ ] Update task
- [ ] Delete task
- [ ] Mark task as in progress
- [ ] Mark task as done
- [ ] List by status filters
- [ ] Better error handling and UX output

## 📌 Notes

- Data is stored in `tasks.json` in the current working directory.
- The file is created automatically when needed.
- The project follows a step-by-step implementation approach from the original challenge.

---

Built as a hands-on CLI practice project to strengthen filesystem, argument parsing, and Go application structure skills. 💡

