package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista todas as tarefas.",
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(Tasks)
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
