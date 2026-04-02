package constants

import "fmt"

type TaskStatus int

const (
	None TaskStatus = iota
	Todo
	InProgress
	Done
)

var taskName = map[TaskStatus]string{
	None:       "None",
	Todo:       "Todo",
	InProgress: "In-Progress",
	Done:       "Done",
}

func (ss TaskStatus) String() string {
	return taskName[ss]
}

func ToConst(s string) (TaskStatus, error) {
	switch s {
	case "Done":
		return Done, nil
	case "InProgress":
		return InProgress, nil
	case "Todo":
		return Todo, nil
	case "None":
		return None, nil
	default:
		return None,
			fmt.Errorf(fmt.Sprintln("Não foi possivel encontrar uma constante válida."))
	}
}
