package main

import (
	"encoding/json"
	"errors"
	"os"

	"nono.guilherme/task-cli/cmd"
	"nono.guilherme/task-cli/model"
)

const saveFile = "tasks.json"

func main() {
	tasks, err := loadTasks()

	if err != nil {
		panic(err)
	}

	cmd.Execute(tasks)

	if err := saveTasks(cmd.Tasks); err != nil {
		panic(err)
	}
}

func loadTasks() ([]model.Task, error) {
	data, errF := os.ReadFile(saveFile)

	if errF != nil {
		if errors.Is(errF, os.ErrNotExist) {
			return []model.Task{}, nil
		}

		return nil, errF
	}

	if len(data) == 0 {
		return []model.Task{}, nil
	}

	var tasks []model.Task

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func saveTasks(tasks []model.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile(saveFile, data, 0o644)
}
