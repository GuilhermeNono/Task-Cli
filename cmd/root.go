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

var Tasks = make(model.TaskMap)
var NextId = 1

func Execute(loadedTasks []model.Task) {

	if len(loadedTasks) != 0 {
		sort.Sort(model.TaskSlice(loadedTasks))
		NextId = loadedTasks[len(loadedTasks)-1].Id + 1
		for _, item := range loadedTasks {
			Tasks[item.Id] = item
		}
	}

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
