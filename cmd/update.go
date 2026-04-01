package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Atualiza a descrição da tarefa.",
	Args:  cobra.MinimumNArgs(2),
	Run: func(c *cobra.Command, args []string) {

		if len(args) == 1 {
			return
		}

		if len(args) == 0 {
			fmt.Println("Por favor, informe um Id valido.")
			return
		}

		id, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Println("Informe o Id da tarefa em formato numérico.")
			return
		}

		if id == 0 {
			fmt.Println("Por favor, informe um Id valido.")
			return
		}

		updatedItem := Tasks[id]
		updatedItem.Description = args[1]
		updatedItem.UpdatedAt = time.Now()
		Tasks[id] = updatedItem
	},
}

func init() {
	RootCmd.AddCommand(updateCmd)
}
