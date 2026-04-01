package constants

type TaskStatus int

const (
	Todo TaskStatus = iota
	InProgress
	Done
)

var taskName = map[TaskStatus]string{
	Todo:       "Todo",
	InProgress: "In-Progress",
	Done:       "Done",
}

func (ss TaskStatus) String() string {
	return taskName[ss]
}
