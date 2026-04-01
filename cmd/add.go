package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"nono.guilherme/task-cli/model"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adiciona uma nova task.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(c *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Falha ao criar uma nova task. Por favor, informe uma descrição para a task.")
			return
		}

		Tasks = append(Tasks, model.Task{
			Id:          NextId,
			Description: args[0],
			Status:      "todo",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})

		fmt.Printf("Task added successfully (ID: %d)\n", NextId)
	},
}

func init() {
	RootCmd.AddCommand(AddCmd)
}
