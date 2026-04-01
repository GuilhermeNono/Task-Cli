package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista todas as tarefas.",
	Run: func(c *cobra.Command, args []string) {

		for _, item := range Tasks.OrderById() {
			fmt.Printf("%d | %s | %s | %s | %s\n", item.Id, item.Description, item.Status, item.CreatedAt, item.UpdatedAt)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
