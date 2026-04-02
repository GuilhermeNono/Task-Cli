# ✅ Task CLI

A lightweight task tracker built with Go and Cobra.

This project is a finished command-line application for creating, updating, listing, and deleting tasks from the terminal. Tasks are stored locally in a JSON file so the data persists between runs.

## ✨ Highlights

- 🚀 Simple terminal workflow for daily task tracking
- 💾 Local persistence in `tasks.json`
- 🧭 Cobra-based command structure
- 📅 Automatic timestamps for creation and updates
- 🔎 Optional status filtering when listing tasks

## 🛠️ Tech Stack

- [Go](https://go.dev/)
- [Cobra](https://github.com/spf13/cobra) for CLI commands
- Standard library packages for file handling and JSON serialization

## 📦 Project Structure

```text
task-cli/
├── cmd/              # CLI commands
├── model/            # Task model, sorting, and status helpers
├── main.go           # Load tasks, execute CLI, save tasks
└── tasks.json        # Local data file created at runtime
```

## 🚀 Getting Started

### Run with Go

```bash
go run . version
go run . add "Buy groceries"
go run . list
```

### Build the binary

```bash
go build -o task-cli
./task-cli version
./task-cli add "Buy groceries"
./task-cli list
```

## 🧭 Available Commands

### `add <description>`

Creates a new task with a unique ID, the `Todo` status, and timestamps.

```bash
task-cli add "Read a book"
```

**Example:**
Adding a new migration task to the tracker.
![Add Task Example](https://i.imgur.com/b9pgKaP.png)

### `list [status]`

Lists all tasks in ascending ID order.

You can also filter by status:

- `Todo`
- `InProgress`
- `Done`

```bash
task-cli list
task-cli list Todo
task-cli list Done
```

**Example:**
Retrieving the full list of tracked tasks.
![List Tasks Example](https://i.imgur.com/ot7iDWt.png)

Filtering the output to show specific statuses (e.g., `Todo` and `Done`).
![Filter Tasks Example](https://i.imgur.com/oA6Xr5e.png)

### `update <id> [description] --status <status>`

Updates a task description and/or status.

```bash
task-cli update 1 "Read two chapters" --status Done
task-cli update 2 --status InProgress
```

> Status values are case-sensitive and follow the CLI’s internal status names.

**Example:**
Updating the description and moving the status to `Done` in a single command, followed by a list verification.
![Update Task Example](https://i.imgur.com/WMocgBj.png)

### `delete <id>`

Removes a task by ID.

```bash
task-cli delete 1
```

**Example:**
Deleting a specific task by its ID and verifying the empty list.
![Delete Task Example](https://i.imgur.com/M49TAdE.png)

### `version`

Prints the current CLI version.

```bash
task-cli version
```

## 🗂️ Task Storage

Tasks are persisted in a `tasks.json` file in the **current working directory**.

That means:

- if you run the app inside the project folder, `tasks.json` is created there
- if you run it from another folder, the JSON file is created in that folder instead

Each task uses this model:

- `id`: unique numeric identifier
- `description`: task text
- `status`: current task state
- `createdAt`: creation timestamp
- `updatedAt`: last modification timestamp

Example `tasks.json` entry:

```json
[
  {
    "id": 1,
    "description": "Buy groceries",
    "status": "Todo",
    "createdAt": "2026-04-01T10:00:00Z",
    "updatedAt": "2026-04-01T10:00:00Z"
  }
]
```

## 🔄 How It Works

1. The app loads existing tasks from `tasks.json`.
2. Commands operate on the in-memory task collection.
3. After the command finishes, the tasks are written back to disk.
4. IDs are automatically incremented based on the highest existing ID.

## 🤖 Continuous Integration

A GitHub Actions workflow is configured in `.github/workflows/go.yml`.

- Runs on `push` and `pull_request` to `master`
- Uses Ubuntu latest
- Builds the project
- Runs the test suite

## 📋 Notes

- The project uses the MIT License.
- `tasks.json` is meant to hold local development data and is created automatically when needed.
- Status filters are case-sensitive.

---

Built as a hands-on CLI project to practice filesystem access, argument parsing, sorting, and structured Go application design. 💡

This project is part of a coding exercise from [Roadmap.SH | Task-Tracker](https://roadmap.sh/projects/task-tracker)
