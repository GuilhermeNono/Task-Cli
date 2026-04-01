package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/spf13/cobra"
	"nono.guilherme/task-cli/model"
)

var RootCmd = &cobra.Command{
	Use:   "task-cli",
	Short: "Meu app de task-cli",
}

var Tasks []model.Task
var NextId = 1

func Execute(loadedTasks []model.Task) {
	if len(loadedTasks) != 0 {
		Tasks = loadedTasks
		sort.Sort(model.ById(Tasks))
		NextId = Tasks[len(Tasks)-1].Id + 1
	}

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
