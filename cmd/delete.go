package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove uma tarefa",
	Args:  cobra.MinimumNArgs(1),
	Run: func(c *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Printf("Não foi possivel excluir a tarefa de codigo %s\n", args[0])
			return
		}

		if Tasks[id].Id == 0 {
			return
		}

		delete(Tasks, id)
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
