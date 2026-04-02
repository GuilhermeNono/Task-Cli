package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"nono.guilherme/task-cli/model/constants"
)

var status string

func init() {

	var updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Atualiza a descrição da tarefa.",
		Args:  cobra.MinimumNArgs(1),
		Run:   command,
	}

	updateCmd.Flags().StringVarP(&status, "status", "s", "", "Alterar o estado atual da tarefa.")

	RootCmd.AddCommand(updateCmd)
}

func command(_ *cobra.Command, args []string) {
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

	if status == constants.InProgress.String() || status == constants.Done.String() {
		updatedItem.Status = status
	}

	if len(args) == 1 {
		updatedItem.Description = Tasks[id].Description
	} else {
		updatedItem.Description = args[1]
	}

	updatedItem.UpdatedAt = time.Now()
	Tasks[id] = updatedItem
}
