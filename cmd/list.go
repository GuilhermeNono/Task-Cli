package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"nono.guilherme/task-cli/model/constants"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista todas as tarefas.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(c *cobra.Command, args []string) {

		var t = Tasks.OrderById()

		if len(args) != 0 {
			toConst, err := constants.ToConst(args[0])

			if err == nil {
				t = *t.ByStatus(toConst)
			}

		}

		for _, item := range t {
			fmt.Printf("%d | %s | %s | %s | %s\n", item.Id, item.Description, item.Status, item.CreatedAt, item.UpdatedAt)
		}
	},
}

func init() {
	RootCmd.AddCommand(ListCmd)
}
